package models

type Employee struct {
	ID          uint64  `gorm:"primaryKey" json:"id"`
	FirstName   string  `json:"firstName"`
	LastName    string  `json:"lastName"`
	Salary      float64 `json:"salary"`
	Description string  `json:"description"`
}
