package PhoneLocation

const tableName = "phone_location"
const pid = "id"
const fieldPref = "pref"
const fieldPhone = "phone"
const fieldProvince = "province"
const fieldCity = "city"
const fieldIsp = "isp"
const fieldPostCode = "post_code"
const fieldCityCode = "city_code"
const fieldAreaCode = "area_code"

type PhoneLocation struct {
	Id       uint64 `gorm:"primaryKey;column:id;autoIncrement;not null;" json:"id"`                     //
	Pref     string `gorm:"column:pref;index;type:varchar(255);not null;default:;" json:"pref"`         //
	Phone    string `gorm:"column:phone;index;type:varchar(255);not null;default:;" json:"phone"`       //
	Province string `gorm:"column:province;index;type:varchar(255);not null;default:;" json:"province"` //
	City     string `gorm:"column:city;index;type:varchar(255);not null;default:;" json:"city"`         //
	Isp      string `gorm:"column:isp;type:varchar(255);not null;default:0;" json:"isp"`                //
	PostCode string `gorm:"column:post_code;type:varchar(255);not null;default:;" json:"postCode"`      //
	CityCode string `gorm:"column:city_code;type:varchar(255);not null;default:;" json:"cityCode"`      //
	AreaCode string `gorm:"column:area_code;type:varchar(255);not null;default:;" json:"areaCode"`      //

}

// func (itself *PhoneLocation) BeforeSave(tx *gorm.DB) (err error) {}
// func (itself *PhoneLocation) BeforeCreate(tx *gorm.DB) (err error) {}
// func (itself *PhoneLocation) AfterCreate(tx *gorm.DB) (err error) {}
// func (itself *PhoneLocation) BeforeUpdate(tx *gorm.DB) (err error) {}
// func (itself *PhoneLocation) AfterUpdate(tx *gorm.DB) (err error) {}
// func (itself *PhoneLocation) AfterSave(tx *gorm.DB) (err error) {}
// func (itself *PhoneLocation) BeforeDelete(tx *gorm.DB) (err error) {}
// func (itself *PhoneLocation) AfterDelete(tx *gorm.DB) (err error) {}
// func (itself *PhoneLocation) AfterFind(tx *gorm.DB) (err error) {}

func (PhoneLocation) TableName() string {
	return tableName
}
