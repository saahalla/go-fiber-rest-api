package models

// Todo struct
type Todo struct {
	ID                int    `json:"id"`
	Activity_Group_ID int    `json:"activity_group_id"`
	Title             string `json:"title"`
	Is_Active         bool   `json:"is_active"`
	Priority          string `json:"priority"`
	Create_At         string `json:"create_at"`
	Update_At         string `json:"update_at"`
	Delete_At         string `json:"delete_at"`
}
