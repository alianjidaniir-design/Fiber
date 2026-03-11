package dataModel

type Studentss struct {
	ID        int64  `gorm:"id;primaryKey" json:"id"`
	StudentID string `gorm:"studentId" json:"studentId"`
	FirstName string `gorm:"firstName" json:"firstName"`
	LastName  string `gorm:"lastName" json:"lastName"`
	CreatedAt string `gorm:"createdAt" json:"createdAt"`
	UpdatedAt string `gorm:"updatedAt" json:"updatedAt"`
	DeletedAt string `gorm:"deletedAt" json:"deletedAt"`
}
