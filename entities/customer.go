package entities

/*
	MonthlyIncome:
		Esta propriedade representa a renda mensal do cliente.
		Ela indica a quantidade de dinheiro que o cliente ganha em um determinado período, geralmente mensalmente.
		A renda mensal é uma métrica importante para entender a capacidade financeira do cliente e pode influenciar suas decisões de gastos e investimentos.

	AverageIncome:
		Esta propriedade representa o saldo médio na conta bancária do cliente.
		Reflete a média dos saldos diários ou mensais em uma conta bancária ao longo de um determinado período.
		O saldo médio na conta é uma medida útil para avaliar a estabilidade financeira do cliente e sua capacidade de manter fundos disponíveis para despesas ou investimentos.
*/

type Customer struct {
	Age           int
	Name          string
	MonthlyIncome float64
	AverageIncome float64
}
