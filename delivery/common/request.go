package common

type ActivityCreate struct {
	Title string `json:"title" validate:"required"`
	Email string `json:"email" validate:"required,email"`
}

type ActivityUpdate struct {
	Title string `json:"title" validate:"required"`
	Email string `json:"email" validate:"required,email"`
}

type TodoCreate struct {
	ActivityGroupId int    `json:"activity_group_id" validate:"required"`
	Title           string `json:"title" validate:"required"`
	IsActive        bool   `json:"is_active"`
}

type TodoUpdate struct {
	Title    string `json:"title" validate:"required"`
	IsActive bool   `json:"is_active"`
	Priority string `json:"priority" validate:"required"`
}
