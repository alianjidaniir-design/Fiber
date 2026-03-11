package datasourse

import (
	"Fiber/API/2/apiSchema/studentsSchema"
	"context"

	studentDataModel "github.com/alianjidaniir-design/Fiber/API/2/models/student/dataModel"
)

type StudentDBDS interface{
	CreateStudent(ctx context.Context, req studentsSchema.CreateUserRequest ) (studentDataModel., error)
}
