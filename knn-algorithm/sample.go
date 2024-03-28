package knn_algorithm

import (
	datasets_customer "datasets"
	"entities"
	"fmt"
)

func Run() {
	dataset := datasets_customer.Get()

	newCustomer := entities.Customer{
		Age:           26,
		Name:          "Pietro",
		MonthlyIncome: 1000,
		AverageIncome: 500,
	}

	classification := Classify(newCustomer, 2, dataset)

	fmt.Println(classification)
}
