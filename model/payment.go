package model

type Payment struct {
	ID uint `gorm:"primary_key; auto_increment" json:"id"`
	Amount float64 `json:"amount"`
	Paid bool `gorm:"not null" json:"paid"`
	ResID uint `sql:"type:integer REFERENCES reservations(id)" json:"res_id"`
	Reservations []Reservation `gorm:"foreign_key:ID; association_foreignkey:ResID"`
}

func (payment *Payment) Create() (error) {
	return db.Create(payment).Error
}

func UpdatePayment(payment *Payment) (error) {
	return db.Model(payment).Updates(payment).Error
}

func QueryPayment(query string) ([]*Payment, error) {
	var err error
	payments := make([]*Payment, 0)
	err = db.Table("payments").Where(query).Find(&payments).Error
	return payments, err
}

func GetPayments(id uint) ([]*Payment, error) {
	var err error
	payments := make([]*Payment, 0)
	if id == 0 {
		err = db.Table("payments").Find(&payments).Error
	} else {
		err = db.Table("payments").Where("id = ?", id).Find(&payments).Error
	}
	return payments, err
}