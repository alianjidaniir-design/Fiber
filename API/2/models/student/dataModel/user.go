package dataModel

type Students struct {
	ID        int64  `gorm:"column:id;primaryKey" json:"id"`
	StudentID string `gorm:"column:studentId" json:"studentId"`
	FirstName string `gorm:"column:firstName" json:"firstName"`
	LastName  string `gorm:"column:lastName" json:"lastName"`
	CreatedAt string `gorm:"column:createdAt" json:"createdAt"`
	UpdatedAt string `gorm:"column:updatedAt" json:"updatedAt"`
	DeletedAt string `gorm:"column:deletedAt" json:"deletedAt"`
}
