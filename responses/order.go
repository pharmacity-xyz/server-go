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

type OrderDetailsResponse struct {
	OrderDate                   time.Time
	TotalPrice                  float64
	OrderDetailsProductResponse []OrderDetailsProductResponse
}
