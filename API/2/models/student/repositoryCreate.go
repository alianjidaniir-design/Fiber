package student

import (
	"Fiber/API/2/apiSchema/commonSchema"
	"Fiber/API/2/apiSchema/studentsSchema"
	"Fiber/API/2/statics/constants/status"
	"context"
)

func (repo *Repository) Create(ctx context.Context, req commonSchema.BaseRequest[studentsSchema.CreateUserRequest]) (res studentsSchema.UserLoginResponse, errStr string, code int, err error) {
	if repo.initErr != nil {
		return studentsSchema.UserLoginResponse{}, "03", status.StatusNotImplemented, repo.initErr
	}
	createdUser, err := repo.db().CreateStudent(ctx, req.Body)
	if err != nil {
		return studentsSchema.UserLoginResponse{}, "04", status.StatusNotImplemented, err
	}
	return studentsSchema.UserLoginResponse{User: createdUser}, "", status.StatusOK, nil
}
