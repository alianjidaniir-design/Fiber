package dataModel

type Students struct {
	ID          int64  `gorm:"column:id;primaryKey" json:"id" msgpack:"id"`
	StudentCode string `gorm:"column:studentId" json:"studentCode" msgpack:"studentCode"`
	FirstName   string `gorm:"column:firstName" json:"firstName" msgpack:"firstName"`
	LastName    string `gorm:"column:lastName" json:"lastName" msgpack:"lastName"`
	CreatedAt   string `gorm:"column:createdAt" json:"createdAt" msgpack:"createdAt"`
	UpdatedAt   string `gorm:"column:updatedAt" json:"updatedAt" msgpack:"updatedAt"`
	DeletedAt   string `gorm:"column:deletedAt" json:"deletedAt" msgpack:"deletedAt"`
}
