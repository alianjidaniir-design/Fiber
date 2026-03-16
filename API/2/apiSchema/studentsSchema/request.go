package studentsSchema

type CreateUserRequest struct {
	StudentCode string `json:"studentCode" msgpack:"studentCode" valid:"required , max=128"`
	FirstName   string `json:"firstName" msgpack:"firstName" valid:"required"`
	LastName    string `json:"lastName" msgpack:"lastName" valid:"required"`
}
