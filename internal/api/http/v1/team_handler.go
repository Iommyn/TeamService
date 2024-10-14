package v1

import (
	"TeamService/internal/api/http/v1/model"
	"TeamService/internal/entity"
	"TeamService/internal/usecases"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"net/http"
	"strconv"
	"strings"
)

type TeamHandler struct {
	teamUsecase usecases.TeamService
	logger      zerolog.Logger
}

func NewTeamHandler(teamUsecase usecases.TeamService, logger zerolog.Logger) *TeamHandler {
	return &TeamHandler{teamUsecase: teamUsecase, logger: logger}
}

func (h *TeamHandler) CreateTeam(ctx *gin.Context) {
	var team model.TeamRequest
	if err := ctx.ShouldBindJSON(&team); err != nil {
		h.handleErrorResponse(ctx, http.StatusBadRequest, "Неверный формат данных", err)
		return
	}

	newTeam := &entity.Team{
		TeamName:    team.TeamName,
		TeamLogo:    team.TeamLogo,
		Description: team.Description,
	}

	err := h.teamUsecase.CreateTeam(ctx, newTeam)
	if err != nil {
		switch {
		case strings.Contains(err.Error(), "HTTP 400: Поле TeamName не может быть пустым"):
			h.handleErrorResponse(ctx, http.StatusBadRequest, "Поле TeamName не может быть пустым", err)
		case strings.Contains(err.Error(), "HTTP 400: Поле TeamLogo не может быть пустым"):
			h.handleErrorResponse(ctx, http.StatusBadRequest, "Поле TeamLogo не может быть пустым", err)
		case strings.Contains(err.Error(), "HTTP 400: Поле Description не может быть пустым"):
			h.handleErrorResponse(ctx, http.StatusBadRequest, "Поле Description не может быть пустым", err)
		case err.Error() == "409":
			h.handleErrorResponse(ctx, http.StatusConflict, "Команда уже существует", err)
			h.logger.Error().Msg("Команда уже существует")
		case err.Error() == "500":
			h.handleErrorResponse(ctx, http.StatusInternalServerError, "Ошибка на стороне сервера", err)
		}
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "Команда успешно сформировалась!"})
}

func (h *TeamHandler) GetTeamByID(ctx *gin.Context) {
	teamID := ctx.Param("id")
	id, err := strconv.ParseInt(teamID, 10, 64)
	if err != nil {
		h.handleErrorResponse(ctx, http.StatusBadRequest, "Неверный формат ID", err)
		return
	}

	team, err := h.teamUsecase.GetTeamByID(ctx, id)
	if err != nil {
		h.handleErrorResponse(ctx, http.StatusInternalServerError, "Ошибка при получении команды", err)
		return
	}

	ctx.JSON(http.StatusOK, team)
}

func (h *TeamHandler) GetAllTeams(ctx *gin.Context) {
	teams, err := h.teamUsecase.GetAllTeams(ctx)
	if err != nil {
		h.handleErrorResponse(ctx, http.StatusInternalServerError, "Ошибка при получении всех команд", err)
		return
	}

	ctx.JSON(http.StatusOK, teams)
}

func (h *TeamHandler) DeleteTeam(ctx *gin.Context) {
	teamID := ctx.Param("id")

	id, err := strconv.ParseInt(teamID, 10, 64)
	if err != nil {
		h.handleErrorResponse(ctx, http.StatusBadRequest, "Неверный формат ID", err)
		return
	}

	err = h.teamUsecase.DeleteTeam(ctx, id)
	if err != nil {
		switch {
		case err.Error() == "404":
			h.handleErrorResponse(ctx, http.StatusNotFound, "Команда не найдена", err)
		case err.Error() == "500":
			h.handleErrorResponse(ctx, http.StatusInternalServerError, "Ошибка на стороне сервера!", err)
		}
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Команда успешно удалена"})
}

func (h *TeamHandler) EditTeam(ctx *gin.Context) {
	teamID := ctx.Param("id")

	id, err := strconv.ParseInt(teamID, 10, 64)
	if err != nil {
		h.handleErrorResponse(ctx, http.StatusBadRequest, "Неверный формат ID", err)
		return
	}

	var team model.TeamRequest
	if err := ctx.ShouldBindJSON(&team); err != nil {
		h.logger.Error().Err(err).Msg("HTTP-Handler: Ошибка десериализации структуры TeamRequestUpdate")
		h.handleErrorResponse(ctx, http.StatusBadRequest, "Неверный формат данных", err)
		return
	}

	updatedTeam := &entity.Team{
		ID:          id,
		TeamName:    team.TeamName,
		TeamLogo:    team.TeamLogo,
		Description: team.Description,
	}

	err = h.teamUsecase.UpdateTeam(ctx, updatedTeam)
	if err != nil {
		switch {
		case err.Error() == "404":
			h.handleErrorResponse(ctx, http.StatusNotFound, "Команда не найдена", err)
		case err.Error() == "500":
			h.handleErrorResponse(ctx, http.StatusInternalServerError, "Ошибка на стороне сервера!", err)
		}
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Данные команды успешно обновлены"})

}

func (h *TeamHandler) handleErrorResponse(ctx *gin.Context, statusCode int, message string, err error) {
	ctx.JSON(statusCode, model.ErrorResponse{
		Status:  statusCode,
		Message: message,
		Error:   err.Error(),
	})
}
