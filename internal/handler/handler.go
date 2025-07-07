package handler

import (
	"log"
	"net/http"

	"github.com/bhushan-aruto/justarquest/internal/entity"
	"github.com/bhushan-aruto/justarquest/internal/usecase" 
	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	UserUsecase *usecase.CreatesuersUsecase
}

func NewUserHandler(u *usecase.CreatesuersUsecase) *UserHandler {
	return &UserHandler{
		UserUsecase: u,
	}
}

func (h *UserHandler) RegisterUser(c echo.Context) error {
	var req entity.Jsutrequest

	if err := c.Bind(&req); err != nil {
		log.Println("Failed to bind JSON:", err)
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
	}

	err := h.UserUsecase.CreateUserUseacase(req.UserName, req.Email, req.Password)
	if err != nil {
		log.Println("Error creating user:", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Could not create user"})
	}

	return c.JSON(http.StatusCreated, map[string]string{"message": "User created successfully"})
}
