package database

import (
	"context"

	"github.com/lucassimon/desafio-client-server-api/internal/entity"
	"gorm.io/gorm"
)

type DollarExchangeDB struct {
	DB *gorm.DB
}

func NewDollarExchange(db *gorm.DB) *DollarExchangeDB {
	return &DollarExchangeDB{DB: db}
}

func (p *DollarExchangeDB) Create(ctx context.Context, usdbrl *entity.UsdBrl) error {
	return p.DB.WithContext(ctx).Create(usdbrl).Error
}
