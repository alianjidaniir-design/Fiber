package mySqlDS

import (
	"Fiber/API/2/apiSchema/studentsSchema"
	"context"
	"errors"
	"fmt"
	"time"

	studentDataModel "Fiber/API/2/models/student/dataModel"

	"github.com/go-sql-driver/mysql"
)

type UserDBDS struct {
	tablename string
	tableSQL  string
	db        DBExecutor
}

func myLocation() *time.Location {
	loc, err := time.LoadLocation("Asia/Tehran")
	if err != nil {
		return time.FixedZone("Asia/Tehran", 3*3600+30*60)
	}
	return loc
}

func isUnknownColmunErr(err error) bool {
	var mysqlErr *mysql.MySQLError
	return errors.As(err, &mysqlErr) && mysqlErr.Number == 1054
}

func NewTaskDBDSFromEnv() (*UserDBDS, bool, error) {
	cfg, err := LoadConfigFromEnv()
	if err != nil {
		return nil, false, err
	}
	if cfg.DSN == "" {
		return nil, false, nil
	}

	tableSQL, err := studentTableIdentifier(cfg.TaskTableName)
	if err != nil {
		return nil, false, err
	}

	db, err := open(cfg)
	if err != nil {
		return nil, false, err
	}

	if err := EnsureTaskTable(db, cfg.TaskTableName); err != nil {
		return nil, false, err
	}
	return &UserDBDS{
		tablename: cfg.TaskTableName,
		tableSQL:  tableSQL,
		db:        db,
	}, true, nil

}

func (db *UserDBDS) CreateStudent(ctx context.Context, req studentsSchema.CreateUserRequest) (studentDataModel.Students, error) {
	insertQuery := fmt.Sprintf("INSERT INTO %s (title , description) VALUES (?, ?)", db.tableSQL)
	insertResult, err := ds.db.ExecContent(ctx, insertQuery, req.FirstName, req.LastName, req.LastName)
	if err != nil {
		return studentDataModel.Students{}, err
	}
	insertedID, err := insertResult.LastInsertId()
	if err != nil {
		return studentDataModel.Students{}, err
	}
	return ds.readTaskByID(ctx, insertedID)
}

func joinCSV(parts []string) string {
	if len(parts) == 0 {
		return ""
	}
	joined := parts[0]
	for i := 1; i < len(parts); i++ {
		joined += ", " + parts[i]
	}
	return joined
}

func (ds *UserDBDS) TableName() string {
	return ds.tablename
}
