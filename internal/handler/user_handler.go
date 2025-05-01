package handler

import (
	"jagratama-backend/internal/dto"
	"jagratama-backend/internal/helpers"
	"jagratama-backend/internal/model"
	"jagratama-backend/internal/service"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	userService service.UserService
}

func NewUserHandler(userService service.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

func (h *UserHandler) Login(c echo.Context) error {
	ctx := c.Request().Context()
	login := &dto.LoginRequest{}

	if err := c.Bind(&login); err != nil {
		return helpers.SendResponseHTTP(c, http.StatusBadRequest, "Invalid payload", err.Error())
	}

	users, err := h.userService.Login(ctx, login.Email, login.Password)
	if err != nil {
		if err.Error() == "record not found" {
			return helpers.SendResponseHTTP(c, http.StatusNotFound, "User not found", nil)
		}
		if err.Error() == "invalid password" {
			return helpers.SendResponseHTTP(c, http.StatusUnauthorized, " email/ password", nil)
		}
		return helpers.SendResponseHTTP(c, http.StatusInternalServerError, "Failed to login", err.Error())
	}

	return helpers.SendResponseHTTP(c, http.StatusOK, "Successfully to logged in", users)
}

func (h *UserHandler) GetAllUsers(c echo.Context) error {
	ctx := c.Request().Context()
	users, err := h.userService.GetAllUsers(ctx)
	if err != nil {
		return helpers.SendResponseHTTP(c, http.StatusInternalServerError, "Failed to get all users", err.Error())
	}

	return helpers.SendResponseHTTP(c, http.StatusOK, "Successfully to get all users", users)
}

func (h *UserHandler) CreateUser(c echo.Context) error {
	ctx := c.Request().Context()
	user := &model.User{}

	if err := c.Bind(user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	newUser, err := h.userService.CreateUser(ctx, user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, newUser)
}

func (h *UserHandler) GetUserByID(c echo.Context) error {
	ctx := c.Request().Context()
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return helpers.SendResponseHTTP(c, http.StatusBadRequest, "Invalid user ID", nil)
	}

	user, err := h.userService.GetUserByID(ctx, id)
	if err != nil {
		return helpers.SendResponseHTTP(c, http.StatusInternalServerError, "Failed to get user", err.Error())
	}

	return helpers.SendResponseHTTP(c, http.StatusOK, "Successfully to get user", user)
}

func (h *UserHandler) UpdateUser(c echo.Context) error {
	ctx := c.Request().Context()
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return helpers.SendResponseHTTP(c, http.StatusBadRequest, "Invalid user ID", nil)
	}

	user := &model.User{}
	if err := c.Bind(user); err != nil {
		return helpers.SendResponseHTTP(c, http.StatusBadRequest, "Invalid request body", err.Error())
	}

	user.ID = uint(id)
	updatedUser, err := h.userService.UpdateUser(ctx, user)
	if err != nil {
		return helpers.SendResponseHTTP(c, http.StatusInternalServerError, "Failed to update user", err.Error())
	}

	return helpers.SendResponseHTTP(c, http.StatusOK, "Successfully to update user", updatedUser)
}

func (h *UserHandler) DeleteUser(c echo.Context) error {
	ctx := c.Request().Context()
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return helpers.SendResponseHTTP(c, http.StatusBadRequest, "Invalid user ID", nil)
	}

	err = h.userService.DeleteUser(ctx, id)
	if err != nil {
		return helpers.SendResponseHTTP(c, http.StatusInternalServerError, "Failed to delete user", err.Error())
	}

	return helpers.SendResponseHTTP(c, http.StatusOK, "Successfully to delete user", nil)
}

func (h *UserHandler) GetMe(c echo.Context) error {
	ctx := c.Request().Context()
	userID, ok := c.Get("userID").(int)
	if !ok {
		return helpers.SendResponseHTTP(c, http.StatusForbidden, "Unauthorized", nil)
	}

	userLogged, err := h.userService.GetMe(ctx, userID)
	if err != nil {
		return helpers.SendResponseHTTP(c, http.StatusInternalServerError, "Failed to get user", err.Error())
	}

	return helpers.SendResponseHTTP(c, http.StatusOK, "Successfully to get user", userLogged)
}

func (h *UserHandler) GetApproverReviewerUsers(c echo.Context) error {
	ctx := c.Request().Context()
	users, err := h.userService.GetApproverReviewerUsers(ctx)
	if err != nil {
		return helpers.SendResponseHTTP(c, http.StatusInternalServerError, "Failed to get approver/reviewer users", err.Error())
	}

	return helpers.SendResponseHTTP(c, http.StatusOK, "Successfully to get approver/reviewer users", users)
}

func (h *UserHandler) UpdateUserProfile(c echo.Context) error {
	ctx := c.Request().Context()

	profile := &dto.UpdateProfileRequest{}
	if err := c.Bind(profile); err != nil {
		return helpers.SendResponseHTTP(c, http.StatusBadRequest, "Invalid request body", err.Error())
	}

	userID, ok := c.Get("userID").(int)
	if !ok {
		return helpers.SendResponseHTTP(c, http.StatusForbidden, "Unauthorized", nil)
	}

	updatedUser, err := h.userService.UpdateUserProfile(ctx, profile, userID)
	if err != nil {
		return helpers.SendResponseHTTP(c, http.StatusInternalServerError, "Failed to update user", err.Error())
	}

	return helpers.SendResponseHTTP(c, http.StatusOK, "Successfully to update user", updatedUser)
}

func (h *UserHandler) RefreshToken(c echo.Context) error {
	ctx := c.Request().Context()

	refreshToken := struct {
		RefreshToken string `json:"refresh_token"`
	}{}
	if err := c.Bind(&refreshToken); err != nil {
		return helpers.SendResponseHTTP(c, http.StatusBadRequest, "Invalid payload", err.Error())
	}

	newToken, err := h.userService.RefreshToken(ctx, refreshToken.RefreshToken)
	if err != nil {
		return helpers.SendResponseHTTP(c, http.StatusInternalServerError, "Failed to refresh token", err.Error())
	}

	return helpers.SendResponseHTTP(c, http.StatusOK, "Successfully to refresh token", newToken)
}

func (h *UserHandler) Logout(c echo.Context) error {
	ctx := c.Request().Context()
	userID, ok := c.Get("userID").(int)
	if !ok {
		return helpers.SendResponseHTTP(c, http.StatusForbidden, "Unauthorized", nil)
	}

	err := h.userService.Logout(ctx, userID)
	if err != nil {
		return helpers.SendResponseHTTP(c, http.StatusInternalServerError, "Failed to logout", err.Error())
	}

	return helpers.SendResponseHTTP(c, http.StatusOK, "Successfully to logout", nil)
}
