package responses

import (
	"time"

	"github.com/google/uuid"
)

type OrderOverviewResponse struct {
	OrderId         uuid.UUID
	OrderData       time.Time
	TotalPrice      float64
	Product         string
	ProductImageUrl string
}
