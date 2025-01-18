package controllers

import (
	"os"
	"testing"
)

func TestLoadRules(t *testing.T) {
	rulesFile := "test_rules.json"
	rulesData := `[{
		"id": "rule1",
		"description": "10% off for premium customers",
		"condition": {"customer_type": "premium"},
		"discount_percentage": 10
	}]`

	// Create a temporary rules file
	if err := os.WriteFile(rulesFile, []byte(rulesData), 0644); err != nil {
		t.Fatalf("Failed to create test rules file: %v", err)
	}
	// Remove the temporary rules file
	defer os.Remove(rulesFile)

	// Load rules
	if err := LoadRules(rulesFile); err != nil {
		t.Fatalf("Failed to load rules: %v", err)
	}

	// Verify rules are loaded correctly
	rulesLock.RLock()
	defer rulesLock.RUnlock()

	if len(rules) != 1 {
		t.Fatalf("Expected 1 rule, got %d", len(rules))
	}
	if rules[0].ID != "rule1" {
		t.Errorf("Expected rule ID 'rule1', got '%s'", rules[0].ID)
	}
	if rules[0].Condition["customer_type"] != "premium" {
		t.Errorf("Expected customer_type 'premium', got '%v'", rules[0].Condition["customer_type"])
	}
}