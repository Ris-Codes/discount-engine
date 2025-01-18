package controllers

import (
	"sync"
	"testing"

	"github.com/Ris-Codes/discount-engine/models"
)

func TestConcurrentRequests(t *testing.T) {
	rules = []models.DiscountRule{
		{
			ID:                 "rule_1",
			Description:        "10% off for orders over $100",
			Condition:          map[string]any{"min_order_value": 100},
			DiscountPercentage: 10,
		},
		{
			ID:            "rule_2",
			Description:   "$20 off for premium customers",
			Condition:     map[string]any{"customer_type": "premium"},
			DiscountFixed: 20,
		},
		{
			ID:                 "rule_3",
			Description:        "5% off for orders over $50",
			Condition:          map[string]any{"min_order_value": 50},
			DiscountPercentage: 5,
		},
		{
			ID:          "rule_4",
			Description: "$10 off for regular customers on orders over $75",
			Condition: map[string]any{
				"customer_type":   "regular",
				"min_order_value": 75,
			},
			DiscountFixed: 10,
		},
	}

	order := models.Order{
		OrderTotal:   85,
		CustomerType: "regular",
	}

	var wg sync.WaitGroup
	const numRequests = 100

	for i := 0; i < numRequests; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			result := CalculateDiscount(order)
			if result.DiscountAmt != 10 {
				t.Errorf("Expected discount amount 20, got %.2f", result.DiscountAmt)
			}
			if result.FinalTotal != 75 {
				t.Errorf("Expected final total 180, got %.2f", result.FinalTotal)
			}
		}()
	}

	wg.Wait()
}
