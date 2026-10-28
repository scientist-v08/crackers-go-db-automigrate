package model

type TempBilling struct {
	ID           uint16 `gorm:"primaryKey;autoIncrement"`
	User         string
	Mobile       string
	Comments     string
	GrandTotal   float64
	FinalizedAmt float64
}