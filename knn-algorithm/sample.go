package knn_algorithm

import (
	datasets_customer "datasets"
	"entities"
	"fmt"
)

func Run() {
	dataset := datasets_customer.Get()

	newCustomer := entities.Customer{
		Age:           74,
		Name:          "Vera",
		MonthlyIncome: 10000,
		AverageIncome: 2000,
	}

	classification := Classify(newCustomer, 2, dataset)

	fmt.Println(classification)
}
