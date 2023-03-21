package DataReps

const tableName = "data_reps"
const pid = "key"
const fieldValue = "value"

type DataReps struct {
	Key   string `gorm:"primaryKey;column:key;autoIncrement;not null;default:'';" json:"key"` //
	Value string `gorm:"column:value;type:varchar(255);not null;default:'';" json:"value"`    //

}

// func (itself *DataReps) BeforeSave(tx *gorm.DB) (err error) {}
// func (itself *DataReps) BeforeCreate(tx *gorm.DB) (err error) {}
// func (itself *DataReps) AfterCreate(tx *gorm.DB) (err error) {}
// func (itself *DataReps) BeforeUpdate(tx *gorm.DB) (err error) {}
// func (itself *DataReps) AfterUpdate(tx *gorm.DB) (err error) {}
// func (itself *DataReps) AfterSave(tx *gorm.DB) (err error) {}
// func (itself *DataReps) BeforeDelete(tx *gorm.DB) (err error) {}
// func (itself *DataReps) AfterDelete(tx *gorm.DB) (err error) {}
// func (itself *DataReps) AfterFind(tx *gorm.DB) (err error) {}

func (DataReps) TableName() string {
	return tableName
}
