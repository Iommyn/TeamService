package repository

import (
	"TeamService/internal/app/postgre"
	"TeamService/internal/app/redis"
	"TeamService/internal/entity"
	"TeamService/internal/models"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/lib/pq"
	"github.com/rs/zerolog"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"strconv"
	"time"
)

type teamRepository struct {
	db    *postgre.PostgreSQLService
	redis *redis.RedisService

	logger zerolog.Logger
}

func NewTeamRepository(db *postgre.PostgreSQLService, redis *redis.RedisService, logger zerolog.Logger) TeamRepository {
	return &teamRepository{db: db, redis: redis, logger: logger}
}

func (t teamRepository) CreateTeam(ctx context.Context, team *entity.Team) error {

	existingTeam, err := models.Teams(qm.Where("team_name = ?", team.TeamName)).One(ctx, t.db.GetMaster())
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		t.logger.Error().Err(err).Msg("Ошибка при проверке существования команды")
		return fmt.Errorf("500")
	}
	if existingTeam != nil {
		return fmt.Errorf("409")
	}

	newTeam := models.Team{
		TeamName:    team.TeamName,
		TeamLogo:    team.TeamLogo,
		Description: team.Description,
	}

	dbMaster := t.db.GetMaster()

	err = newTeam.Insert(ctx, dbMaster, boil.Infer())

	if err != nil {
		var pqErr *pq.Error
		if errors.As(err, &pqErr) {
			if pqErr.Code == "23505" {
				return fmt.Errorf("409")
			}
		}
		return fmt.Errorf("500")
	}

	return nil
}

func (t teamRepository) GetTeamByID(ctx context.Context, id int64) (*entity.Team, error) {
	dbReplica := t.db.GetReplica()

	teamModel, err := models.Teams(qm.Where("id = ?", id)).One(ctx, dbReplica)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("404")
		}
		return nil, fmt.Errorf("500")
	}

	team := &entity.Team{
		ID:          teamModel.ID,
		UserID:      strconv.FormatInt(teamModel.UserID, 10),
		TeamName:    teamModel.TeamName,
		TeamLogo:    teamModel.TeamLogo,
		Description: teamModel.Description,
		UpdatedAt:   teamModel.UpdatedAt.Format(time.RFC3339),
		CreatedAt:   teamModel.CreatedAt.Format(time.RFC3339),
	}
	return team, nil
}

func (t teamRepository) GetAllTeams(ctx context.Context) ([]*entity.Team, error) {
	dbReplica := t.db.GetReplica()

	teamModels, err := models.Teams().All(ctx, dbReplica)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("404")
		}

		return nil, fmt.Errorf("505")
	}

	var teams []*entity.Team

	for _, teamModel := range teamModels {
		team := &entity.Team{
			ID:          teamModel.ID,
			UserID:      strconv.FormatInt(teamModel.UserID, 10),
			TeamName:    teamModel.TeamName,
			TeamLogo:    teamModel.TeamLogo,
			Description: teamModel.Description,
			UpdatedAt:   teamModel.UpdatedAt.Format(time.RFC3339),
			CreatedAt:   teamModel.CreatedAt.Format(time.RFC3339),
		}
		teams = append(teams, team)
	}
	return teams, nil
}

func (t teamRepository) DeleteTeam(ctx context.Context, id int64) error {
	dbMaster := t.db.GetMaster()

	teamModels := &models.Team{
		ID: id,
	}

	rowsAffected, err := teamModels.Delete(ctx, dbMaster)
	if err != nil {
		var pqErr *pq.Error
		if errors.As(err, &pqErr) {
			if pqErr.Code == "23505" {
				return fmt.Errorf("409")
			}
		}
		return fmt.Errorf("500")
	}

	if rowsAffected == 0 {
		return fmt.Errorf("404")
	}

	return nil
}

func (t teamRepository) UpdateTeam(ctx context.Context, team *entity.Team) error {
	dbMaster := t.db.GetMaster()

	teamModel := &models.Team{
		ID:          team.ID,
		TeamName:    team.TeamName,
		TeamLogo:    team.TeamLogo,
		Description: team.Description,
		UpdatedAt:   time.Now(),
	}

	rowsAffected, err := teamModel.Update(ctx, dbMaster, boil.Infer())
	if err != nil {
		return fmt.Errorf("500")
	}

	if rowsAffected == 0 {
		return fmt.Errorf("404")
	}

	return nil
}
