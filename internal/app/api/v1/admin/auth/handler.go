package auth

import (
	"net/http"

	"github.com/anonychun/ecorp/internal/api"
	"github.com/labstack/echo/v4"
)

func (h *Handler) Login(c echo.Context) error {
	req := &LoginRequest{
		IpAddress: c.RealIP(),
		UserAgent: c.Request().UserAgent(),
	}
	err := c.Bind(req)
	if err != nil {
		return err
	}

	res, err := h.usecase.Login(c.Request().Context(), *req)
	if err != nil {
		return err
	}

	c.SetCookie(&http.Cookie{
		Name:     "admin_session",
		Value:    res.Token,
		Path:     "/",
		HttpOnly: true,
	})

	return c.JSON(http.StatusOK, res)
}

func (h *Handler) Logout(c echo.Context) error {
	cookie, err := c.Cookie("admin_session")
	if err != nil {
		return err
	}

	req := &LogoutRequest{
		Token: cookie.Value,
	}

	err = h.usecase.Logout(c.Request().Context(), *req)
	if err != nil {
		return err
	}

	c.SetCookie(&http.Cookie{
		Name:     "admin_session",
		Value:    "",
		Path:     "/",
		HttpOnly: true,
		MaxAge:   -1,
	})

	return c.NoContent(http.StatusOK)
}

func (h *Handler) Me(c echo.Context) error {
	return api.NewResponse(c).SendOk()
}
