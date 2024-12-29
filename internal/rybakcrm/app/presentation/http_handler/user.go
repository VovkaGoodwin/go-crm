package http_handler

import "github.com/gin-gonic/gin"

type CreateUserRequest struct {
	Email           string `json:"email"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirm_password"`
	FirstName       string `json:"first_name"`
	LastName        string `json:"last_name"`
	StartWork       string `json:"start_work"`
	BirthDate       string `json:"birth_date"`
	IsSuperUser     bool   `json:"is_super_user"`
	DepartmentID    string `json:"department_id"`
	PositionID      string `json:"position_id"`
}

func (h *Handler) createUser(c *gin.Context) {
}
