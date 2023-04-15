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
	ActivityGroupId int    `json:"activity_group_id"`
	Title           string `json:"title"`
	IsActive        bool   `json:"is_active"`
}

type TodoUpdate struct {
	Title    string `json:"title"`
	IsActive bool   `json:"is_active"`
	Priority string `json:"priority"`
}
