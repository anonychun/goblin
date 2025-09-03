package admin

import "github.com/anonychun/ecorp/internal/entity"

type AdminResponse struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

func NewAdminResponse(admin *entity.Admin) *AdminResponse {
	return &AdminResponse{
		Id:   admin.Id.String(),
		Name: admin.Name,
	}
}

type FindByIdRequest struct {
	Id string `param:"id"`
}

type CreateRequest struct {
	Name string `json:"name"`
}

type UpdateRequest struct {
	Id   string `param:"id"`
	Name string `json:"name"`
}

type DeleteRequest struct {
	Id string `param:"id"`
}
