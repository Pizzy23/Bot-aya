package db

import "time"

type Client struct {
	BaseModel
	Cellphone string `json:"cellphone"`
}

type DataClient struct {
	BaseModel
	ClientID    uint         `json:"client_id"`
	Investments []Investment `json:"investments" gorm:"foreignKey:ClientID;references:ID"`
	Slips       []Slip       `json:"slips" gorm:"foreignKey:ClientID;references:ID"`
	Balances    Balance      `json:"balances" gorm:"foreignKey:ClientID;references:ID"`
	Recharge    []Recharge   `json:"recharge" gorm:"foreignKey:ClientID;references:ID"`
	Debt        float64      `json:"debt"`
}

type Investment struct {
	BaseModel
	ClientID     uint    `json:"client_id"`
	Name         string  `json:"name"`
	Value        float64 `json:"value"`
	Income       float64 `json:"income"`
	TypeOfInvest string  `json:"type_of_invest"`
}

type Slip struct {
	BaseModel
	ClientID uint    `json:"client_id"`
	Name     string  `json:"name"`
	Value    float64 `json:"value"`
	BarCode  string  `json:"bar_code"`
}

type Balance struct {
	BaseModel
	ClientID     uint    `json:"client_id"`
	Balance      float64 `json:"balance"`
	TotalPayment float64 `json:"total_payment"`
}

type Recharge struct {
	BaseModel
	ClientID       uint    `json:"client_id"`
	RechargeType   string  `json:"recharge_Type"`
	RechargeNumber string  `json:"recharge_Number"`
	RechargeValue  float64 `json:"recharge_Value"`
}

type Navegation struct {
	BaseModel
	ClientID  uint `json:"client_id"`
	Payment   int  `json:"payment"`
	Recharge  int  `json:"recharge"`
	Invest    int  `json:"invest"`
	Treatment bool `json:"treatment"`
}

type BaseModel struct {
	ID        uint       `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at" gorm:"index"`
}
