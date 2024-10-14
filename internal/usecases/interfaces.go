package usecases

import (
	"TeamService/internal/entity"
	"context"
)

type (
	TeamService interface {
		CreateTeam(ctx context.Context, team *entity.Team) error
		GetTeamByID(ctx context.Context, id int64) (*entity.Team, error)
		GetAllTeams(ctx context.Context) ([]*entity.Team, error)
		DeleteTeam(ctx context.Context, id int64) error
		UpdateTeam(ctx context.Context, team *entity.Team) error
	}
)
