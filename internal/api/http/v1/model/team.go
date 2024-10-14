package model

type TeamRequest struct {
	TeamName    string `json:"team_name"`
	TeamLogo    string `json:"team_logo"`
	Description string `json:"description"`
}
