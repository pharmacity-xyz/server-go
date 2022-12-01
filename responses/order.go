package responses

import (
	"time"

	"github.com/google/uuid"
)

type OrderOverviewResponse struct {
	OrderId         uuid.UUID
	OrderDate       time.Time
	TotalPrice      float64
	Product         string
	ProductImageUrl string
}
