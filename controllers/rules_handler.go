package controllers

import (
	"encoding/json"
	"os"

	"github.com/Ris-Codes/discount-engine/models"
)

// Load the discount rules from the given file
func LoadRules(filePath string) error {
	data, err := os.ReadFile(filePath)
	if err != nil{
		return err
	}

	// Unmarshal the JSON data into the rules slice
	var loadedRules []models.DiscountRule
	if err := json.Unmarshal(data, &loadedRules); err != nil {
		return err
	}

	rulesLock.Lock()
	rules = loadedRules
	rulesLock.Unlock()
	return nil
}

// Check if the given order satisfies the conditions of the discount rule
func ValidateRule(order models.Order, rule models.DiscountRule) bool {
	for key, value := range rule.Condition {
		switch key {
		case "min_order_value":
			var minOrderValue float64
			switch v := value.(type) {
			case int:
				minOrderValue = float64(v)
			case float64:
				minOrderValue = v
			default:
				// Unsupported type, consider the rule invalid
				return false
			}
			if order.OrderTotal < minOrderValue {
				return false
			}

		case "customer_type":
			customerType, ok := value.(string)
			if !ok || order.CustomerType != customerType {
				return false
			}
		}
	}
	return true
}