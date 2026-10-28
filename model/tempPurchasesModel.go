package model

type TempPurchases struct {
	ID        uint16      `gorm:"primaryKey;autoIncrement"`
	BillingID uint16      `gorm:"not null;index"`
	Billing   TempBilling `gorm:"foreignKey:BillingID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Mobile    string      `gorm:"index"`
	MrpOrNet  int32
	Item      string
	Quantity  int32
	Discount  string
	SubTotal  float64
}