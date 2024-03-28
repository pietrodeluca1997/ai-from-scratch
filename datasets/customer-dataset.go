package datasets_customer

import "entities"

type BinaryClassification int

const (
	GoodPayer BinaryClassification = iota
	BadPayer
)

var classificationDescriptions = map[BinaryClassification]string{
	GoodPayer: "Good Payer",
	BadPayer:  "Bad Payer",
}

func GetClassificationDescription(classification BinaryClassification) string {
	return classificationDescriptions[classification]
}

type CustomerDataset struct {
	Customers       []entities.Customer
	Classifications []BinaryClassification
}

func Get() *CustomerDataset {
	return &CustomerDataset{
		Customers: []entities.Customer{
			{Age: 25, MonthlyIncome: 2500, AverageIncome: 1000},
			{Age: 20, MonthlyIncome: 2500, AverageIncome: -1000},
			{Age: 35, MonthlyIncome: 3500, AverageIncome: 2000},
			{Age: 35, MonthlyIncome: 1500, AverageIncome: -2000},
			{Age: 45, MonthlyIncome: 4500, AverageIncome: 3000},
			{Age: 55, MonthlyIncome: 5500, AverageIncome: 4000},
		},
		Classifications: []BinaryClassification{GoodPayer, BadPayer, GoodPayer, BadPayer, BadPayer, BadPayer},
	}
}
