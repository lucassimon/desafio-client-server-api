package database

import (
	"context"

	"github.com/lucassimon/desafio-client-server-api/internal/entity"
)

type DollarExchangeInterface interface {
	Create(ctx context.Context, usdbrl *entity.UsdBrl) error
}
