package main

import (
	"fmt"
	"net/http"

	"github.com/Ris-Codes/discount-engine/controllers"
)

func main() {
	// Load the configuration file with the discount rules
	if err := controllers.LoadRules("go_task_rules.json"); err != nil {
		panic(fmt.Sprintf("Failed to load rules: %v", err))
	}

	// Call the DiscountHandler function on the apply-discount endpoint
	http.HandleFunc("/apply-discount", controllers.DiscountHandler)

	// Run the server
	fmt.Println("Server running on port 8000")
	if err := http.ListenAndServe(":8000", nil); err != nil {
		panic(err)
	}
}
