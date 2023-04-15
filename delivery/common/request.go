package common

type ActivityCreate struct {
	Title string `json:"title"`
	Email string `json:"email"`
}

type ActivityUpdate struct {
	Title string `json:"title"`
	Email string `json:"email"`
}

type TodoCreate struct {
	Id              uint   `json:"id"`
	ActivityGroupId int    `json:"activity_group_id"`
	Title           string `json:"title"`
	IsActive        bool   `json:"is_active"`
	CreatedAt       string `json:"createdAt"`
	UpdatedAt       string `json:"updatedAt"`
}

type TodoUpdate struct {
	// Id        uint   `json:"id"`
	Title     string `json:"title"`
	IsActive  bool   `json:"is_active"`
	Priority  string `json:"priority"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}
