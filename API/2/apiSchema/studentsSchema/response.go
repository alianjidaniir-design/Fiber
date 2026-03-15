package studentsSchema

import "Fiber/API/2/models/student/dataModel"

type UserLoginResponse struct {
	User dataModel.Students `json:"user" `
}
