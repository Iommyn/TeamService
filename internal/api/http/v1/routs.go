package v1

import "github.com/gin-gonic/gin"

func RegisterRoutes(router *gin.Engine, teamHandler *TeamHandler) {
	router.Use(metricsMiddleware())

	TeamGroupV1 := router.Group("/v1")
	{
		TeamGroupV1.POST("/team-create", teamHandler.CreateTeam)
		//TeamGroupV1.POST("/team-useradd/:id", teamHandler.AddUserTeam)
		//TeamGroupV1.POST("/team-turadd/:id", teamHandler.AddTur)

		TeamGroupV1.GET("/teams", teamHandler.GetAllTeams)
		TeamGroupV1.GET("/team/:id", teamHandler.GetTeamByID)
		//TeamGroupV1.GET("/team-stats/", teamHandler.GetStatsAllTeams)
		//TeamGroupV1.GET("/team-stats/:id", teamHandler.GetStatsTeam)

		TeamGroupV1.PATCH("/team-edit/:id", teamHandler.EditTeam)
		//TeamGroupV1.PATCH("/team-roleedit/:id", teamHandler.EditRoleTeam)
		//TeamGroupV1.PATCH("/team-statsedit/:id", teamHandler.EditStatsTeam)
		//TeamGroupV1.PATCH("/team-modifyelo/:id", teamHandler.EditModifyELO)

		TeamGroupV1.DELETE("/team-del/:id", teamHandler.DeleteTeam)
		//TeamGroupV1.DELETE("/team-userdel/:id", teamHandler.DeleteUserTeam)
		//TeamGroupV1.DELETE("/team-userleave/:id", teamHandler.UserLeave)
		//TeamGroupV1.DELETE("/team-turdel/:id", teamHandler.DeleteTur)
	}
}
