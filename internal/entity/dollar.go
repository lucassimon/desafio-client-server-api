package entity

import (
	"github.com/lucassimon/desafio-client-server-api/internal/dto"
	"gorm.io/gorm"
)

type UsdBrl struct {
	Code       string `json:"code"`
	Codein     string `json:"codein"`
	Name       string `json:"name"`
	High       string `json:"high"`
	Low        string `json:"low"`
	VarBid     string `json:"varBid"`
	PctChange  string `json:"pctChange"`
	Bid        string `json:"bid"`
	Ask        string `json:"ask"`
	Timestamp  string `json:"timestamp"`
	CreateDate string `json:"create_date"`
	gorm.Model
}

func NewUsdBrl(dollar_real *dto.CreateDollarInput) *UsdBrl {
	return &UsdBrl{
		Code:       dollar_real.USDBRL.Code,
		Codein:     dollar_real.USDBRL.Codein,
		Name:       dollar_real.USDBRL.Name,
		High:       dollar_real.USDBRL.High,
		Low:        dollar_real.USDBRL.Low,
		VarBid:     dollar_real.USDBRL.VarBid,
		PctChange:  dollar_real.USDBRL.PctChange,
		Bid:        dollar_real.USDBRL.Bid,
		Ask:        dollar_real.USDBRL.Ask,
		Timestamp:  dollar_real.USDBRL.Timestamp,
		CreateDate: dollar_real.USDBRL.CreateDate,
	}
}
