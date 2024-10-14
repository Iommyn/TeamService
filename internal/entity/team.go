package entity

type Team struct {
	ID          int64  `json:"id"`
	UserID      string `json:"user_id"`
	TeamName    string `json:"team_name"`
	TeamLogo    string `json:"team_logo"`
	Description string `json:"description"`
	UpdatedAt   string `json:"updated_at"`
	CreatedAt   string `json:"created_at"`
}
