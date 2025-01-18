package models

type DiscountRule struct {
	ID                 string         `json:"id"`
	Description        string         `json:"description"`
	Condition          map[string]any `json:"condition"`
	DiscountPercentage float64        `json:"discount_percentage,omitempty"`
	DiscountFixed      float64        `json:"discount_fixed,omitempty"`
	Priority           int            `json:"priority"`
}

type Order struct {
	OrderTotal   float64 `json:"order_total"`
	CustomerType string  `json:"customer_type"`
}

type DiscountResult struct {
	DiscountAmt  float64  `json:"discount_amount"`
	FinalTotal   float64  `json:"final_total"`
	AppliedRules []string `json:"applied_rules"`
}