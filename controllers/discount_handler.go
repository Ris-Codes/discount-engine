package controllers

import (
	"encoding/json"
	"net/http"
	"sync"

	"github.com/Ris-Codes/discount-engine/models"
)

var (
	rules     []models.DiscountRule
	rulesLock sync.RWMutex
)

func DiscountHandler(w http.ResponseWriter, r *http.Request) {
	// Accept POST requests
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST method is allowed", http.StatusBadRequest)
		return
	}

	// Decode the JSON data from the request
	var order models.Order
	err := json.NewDecoder(r.Body).Decode(&order)
	if err != nil {
		http.Error(w, "Invalid Request", http.StatusBadRequest)
		return
	}

	// Calculate the discount to be applied and store it as result
	result := CalculateDiscount(order)

	// Return the result in the JSON format
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

func CalculateDiscount(order models.Order) models.DiscountResult {
	// Lock the rules 
	rulesLock.RLock()
	// Unlock the lock after the function returns
	defer rulesLock.RUnlock()

	// Store the highest discount and applied rules
	var (
		highestDiscount float64
		appliedRules    []string
	)
	
	// Iterate over each rule and check if it is applicable
	for _, rule := range rules {
		if ValidateRule(order, rule) {
			var discount float64
			if rule.DiscountPercentage > 0 {
				discount = order.OrderTotal * rule.DiscountPercentage / 100
			} else if rule.DiscountFixed > 0 {
				discount = rule.DiscountFixed
			}

			// If the current discount is higher than the previous highest discount
			if discount > highestDiscount {
				highestDiscount = discount
				appliedRules = []string{rule.ID}
			} else if discount == highestDiscount {
				appliedRules = append(appliedRules, rule.ID)
			}
		}
	}

	return models.DiscountResult{
		DiscountAmt:  highestDiscount,
		FinalTotal:   order.OrderTotal - highestDiscount,
		AppliedRules: appliedRules,
	}
}
