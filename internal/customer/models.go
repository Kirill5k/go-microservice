package customer

type Entity struct {
	ID        string `gorm:"primaryKey"`
	FirstName string
	LastName  string
	Email     string `gorm:"uniqueIndex"`
	Phone     string
	Address   string
}

func (Entity) TableName() string {
	return "wisdom.customers"
}

type Customer struct {
	ID        string `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	Address   string `json:"address"`
}
