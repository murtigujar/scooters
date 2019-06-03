package model

type Reservation struct {
	ID uint `gorm:"primary_key; auto_increment" json:"id"`
	ScooterID uint `sql:"type:integer REFERENCES scooters(id)" json:"scooter_id"`
	Scooters []Scooter `gorm:"foreign_key:ID; association_foreignkey:ScooterID"`
}

func (reservation *Reservation) Create() (error) {
	return db.Create(reservation).Error
}

func GetReservations(id uint) ([]*Reservation, error) {
	var err error
	reservations := make([]*Reservation, 0)
	if id == 0 {
		err = db.Table("reservations").Find(&reservations).Error
	} else {
		err = db.Table("reservations").Where("id = ?", id).Find(&reservations).Error
	}
	return reservations, err
}