package consts

import (
	"net/http"

	"github.com/anonychun/ecorp/internal/api"
)

var (
	ErrUnauthorized       = &api.Error{Status: http.StatusUnauthorized, Errors: "You are not allowed to perform this action"}
	ErrInvalidCredentials = &api.Error{Status: http.StatusUnauthorized, Errors: "Invalid email or password"}

	ErrAdminNotFound = &api.Error{Status: http.StatusNotFound, Errors: "Admin not found"}
)
