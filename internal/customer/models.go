package customer

type Entity struct {
	ID        string `gorm:"primaryKey"`
	FirstName string
	LastName  string
	Email     string
	Phone     string
	Address   string
}

func toDomain(e Entity) Customer {
	return Customer{
		ID:        e.ID,
		FirstName: e.FirstName,
		LastName:  e.LastName,
		Email:     e.Email,
		Phone:     e.Phone,
		Address:   e.Address,
	}
}

type Customer struct {
	ID        string `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	Address   string `json:"address"`
}
