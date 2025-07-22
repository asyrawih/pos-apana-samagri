package models

import "time"

type RecommendationType string

const (
	Upsell          RecommendationType = "upsell"
	CrossSell       RecommendationType = "cross_sell"
	Personalized    RecommendationType = "personalized"
	Trending        RecommendationType = "trending"
)

type Recommendation struct {
	ID                uint               `json:"id"`
	CustomerID        *uint              `json:"customer_id,omitempty"`
	ProductID         uint               `json:"product_id"`
	RecommendedType   RecommendationType `json:"recommended_type"`
	Score             float64            `json:"score"`
	CreatedAt         time.Time          `json:"created_at"`
}

type RecommendationRequest struct {
	CustomerID     *uint   `json:"customer_id,omitempty"`
	ProductIDs     []uint  `json:"product_ids,omitempty"`
	MaxSuggestions int     `json:"max_suggestions" binding:"min=1,max=10"`
}

type RecommendationResponse struct {
	CustomerID    *uint               `json:"customer_id,omitempty"`
	Recommendations []ProductRecommendation `json:"recommendations"`
	GeneratedAt    time.Time          `json:"generated_at"`
}

type ProductRecommendation struct {
	ProductID    uint     `json:"product_id"`
	Name         string   `json:"name"`
	Price        float64  `json:"price"`
	ImageURL     string   `json:"image_url,omitempty"`
	Reason       string   `json:"reason"`
	Score        float64  `json:"score"`
	Type         RecommendationType `json:"type"`
}

type RecommendationMetrics struct {
	TotalImpressions int     `json:"total_impressions"`
	TotalClicks      int     `json:"total_clicks"`
	ClickThroughRate float64 `json:"click_through_rate"`
	ConversionRate   float64 `json:"conversion_rate"`
}
