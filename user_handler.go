package example

import (
	"fmt"

	"github.com/go-openapi/runtime/middleware"
	"github.com/shellingford330/lambda-go/gen/models"
	"github.com/shellingford330/lambda-go/gen/restapi/operations"
)

func GetUsers(p operations.GetUsersParams) middleware.Responder {
	ctx := p.HTTPRequest.Context()
	users, err := scanUsers(ctx)
	if err != nil {
		return operations.NewGetUsersInternalServerError().WithPayload(&models.Error{
			Message: fmt.Sprintf("scan users error: %v", err),
		})
	}
	var resp models.Users
	for _, u := range users {
		u := u
		resp = append(resp, &models.User{
			UserID: &u.UserID,
			Name:   &u.UserName,
		})
	}
	return operations.NewGetUsersOK().WithPayload(resp)
}
