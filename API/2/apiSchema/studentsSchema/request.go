package studentsSchema

type CreateUserRequest struct {
	StudentCode string `json:"studentCode"  valid:"required , max=128"`
	FirstName   string `json:"firstName"  valid:"required"`
	LastName    string `json:"lastName" valid:"required"`
}
