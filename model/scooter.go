package model

type Scooter struct {
	ID uint `gorm:"primary_key; auto_increment" json:"id"`
	Latitude float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Reserved bool `gorm:"not null" json:"reserved"`
}

func (scooter *Scooter) Create() (error) {
	return db.Create(scooter).Error
}

func GetAvailableScooters() ([]*Scooter, error) {
	scooters := make([]*Scooter, 0)
	err := db.Table("scooters").Where("reserved = false").Find(&scooters).Error
	return scooters, err
}

func UpdateScooter(scooter *Scooter) (error) {
	return db.Model(scooter).Updates(scooter).Error
}

func SaveScooter(scooter *Scooter) (error) {
	return db.Save(scooter).Error
}

func GetScooters(id uint) ([]*Scooter, error) {
	var err error
	scooters := make([]*Scooter, 0)
	if id == 0 {
		err = db.Table("scooters").Find(&scooters).Error
	} else {
		err = db.Table("scooters").Where("id = ?", id).Find(&scooters).Error
	}
	return scooters, err
}
