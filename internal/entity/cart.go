package entity

import "time"

type UserCart struct {
	Id        string
	UserId    string
	ProductId string
	Quantity  int
	CreatedAt time.Time
	CreatedBy string
	UpdateAt  *time.Time
	UpdatedBy *string
}
