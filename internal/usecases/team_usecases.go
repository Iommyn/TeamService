package usecases

import (
	"TeamService/internal/entity"
	"TeamService/internal/repository"
	"context"
	"fmt"
	"github.com/rs/zerolog"
	"net/http"
)

type TeamServiceImpl struct {
	teamService TeamService

	teamRepo repository.TeamRepository
	logger   zerolog.Logger
}

func NewTeamService(teamRepo repository.TeamRepository, logger zerolog.Logger) TeamService {
	return &TeamServiceImpl{
		teamRepo: teamRepo,
		logger:   logger,
	}
}

func (svc *TeamServiceImpl) CreateTeam(ctx context.Context, team *entity.Team) error {
	if team.TeamName == "" {
		svc.logger.Error().Msg("Поле TeamName не может быть пустым")
		return fmt.Errorf("HTTP %d: %s", http.StatusBadRequest, "Поле TeamName не может быть пустым")
	}
	if team.TeamLogo == "" {
		svc.logger.Error().Msg("Поле TeamLogo не может быть пустым")
		return fmt.Errorf("HTTP %d: %s", http.StatusBadRequest, "Поле TeamLogo не может быть пустым")
	}
	if team.Description == "" {
		svc.logger.Error().Msg("Поле Description не может быть пустым")
		return fmt.Errorf("HTTP %d: %s", http.StatusBadRequest, "Поле Description не может быть пустым")
	}

	return svc.teamRepo.CreateTeam(ctx, team)
}

func (svc *TeamServiceImpl) GetTeamByID(ctx context.Context, id int64) (*entity.Team, error) {
	return svc.teamRepo.GetTeamByID(ctx, id)
}

func (svc *TeamServiceImpl) DeleteTeam(ctx context.Context, id int64) error {
	return svc.teamRepo.DeleteTeam(ctx, id)
}

func (svc *TeamServiceImpl) GetAllTeams(ctx context.Context) ([]*entity.Team, error) {
	teams, err := svc.teamRepo.GetAllTeams(ctx)

	if err != nil {
		return nil, err
	}

	if len(teams) == 0 {
		return nil, fmt.Errorf("404")
	}

	return teams, nil
}

func (svc *TeamServiceImpl) UpdateTeam(ctx context.Context, team *entity.Team) error {
	existingTeam, err := svc.teamRepo.GetTeamByID(ctx, team.ID)
	if err != nil {
		return err
	}

	existingTeam.TeamName = team.TeamName
	existingTeam.TeamLogo = team.TeamLogo
	existingTeam.Description = team.Description

	err = svc.teamRepo.UpdateTeam(ctx, existingTeam)
	if err != nil {
		return err
	}

	return nil
}
