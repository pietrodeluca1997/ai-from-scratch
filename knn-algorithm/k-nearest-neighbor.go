/*
	O algoritmo KNN pode ser aplicado em uma variedade de problemas de classificação e regressão em diferentes domínios.
	Aqui estão alguns exemplos de casos de uso onde o KNN pode ser aplicado com eficácia:

	Classificação:
		Reconhecimento de Padrões: Identificação de padrões em imagens, como reconhecimento facial, reconhecimento de caracteres escritos à mão, classificação de objetos em imagens, etc.

		Diagnóstico Médico: Classificação de pacientes em grupos com base em sintomas, características clínicas ou resultados de testes médicos para diagnóstico de doenças.

		Filtragem de Spam: Classificação de e-mails como spam ou não spam com base em características como palavras-chave, remetente, estrutura do e-mail, etc.

		Análise de Sentimentos: Classificação de sentimentos em textos, como análise de sentimentos em mídias sociais, avaliação de opiniões de clientes, etc.

		Sistemas de Recomendação: Identificação de itens semelhantes ou, recomendação de produtos com base nos padrões de comportamento do usuário ou nas características dos itens.

	Regressão:
		Previsão de Preços: Previsão de preços de imóveis, ações, commodities, etc., com base em características como localização, tamanho, características, histórico de preços, etc.

		Previsão de Demanda: Previsão de demanda por produtos ou serviços com base em características como preço, promoções, dados históricos de vendas, etc.

		Previsão de Desempenho: Previsão de desempenho de alunos em exames com base em seu histórico acadêmico, tempo de estudo, etc.

		Análise de Tendências: Identificação de padrões temporais em séries temporais, como previsão de temperatura, vendas sazonais, tráfego na web, etc.

		Personalização de Serviços: Personalização de serviços com base no comportamento do usuário, como previsão de preferências de música, recomendações de filmes, etc.

*/

package knn_algorithm

import (
	datasets_customer "datasets"
	"entities"
	"math"
)

/*
Classifica um cliente usando o algoritmo KNN (K-Vizinhos Mais Próximos).
Este método calcula as distâncias entre o cliente a ser classificado e todos os clientes no conjunto de dados.
O cliente é classificado com base nas classes dos k vizinhos mais próximos.

Parâmetros:
- customerToClassify: O cliente a ser classificado.
- k: O número de vizinhos mais próximos a serem considerados na classificação.
- dataset: O conjunto de dados contendo os clientes e suas respectivas classes.
*/
func Classify(customerToClassify entities.Customer, k int, dataset *datasets_customer.CustomerDataset) string {
	distances := make([]float64, len(dataset.Customers))

	for i, point := range dataset.Customers {
		distances[i] = euclideanDistance(customerToClassify, point)
	}

	sortedIndexes := make([]int, len(dataset.Customers))

	for i := range distances {
		sortedIndexes[i] = i
	}

	for i := range distances {
		for j := i + 1; j < len(distances); j++ {
			if distances[i] > distances[j] {
				distances[i], distances[j] = distances[j], distances[i]
				sortedIndexes[i], sortedIndexes[j] = sortedIndexes[j], sortedIndexes[i]
			}
		}
	}

	classifierNearestOccurrences := make(map[datasets_customer.BinaryClassification]int)

	for i := 0; i < k; i++ {
		classifierNearestOccurrences[dataset.Classifications[sortedIndexes[i]]]++
	}

	maxCount := 0
	var predictedClass datasets_customer.BinaryClassification
	for class, count := range classifierNearestOccurrences {
		if count > maxCount {
			maxCount = count
			predictedClass = class
		}
	}

	return datasets_customer.GetClassificationDescription(predictedClass)
}

/*
Essa função retorna a distância euclidiana entre os dois pontos fornecidos como argumentos.
Essa distância é usada para medir a proximidade entre os clientes no espaço de características.
*/
func euclideanDistance(p1, p2 entities.Customer) float64 {
	return math.Sqrt(
		math.Pow(float64(p1.Age-p2.Age), 2) +
			math.Pow(p1.MonthlyIncome-p2.MonthlyIncome, 2) +
			math.Pow(p1.AverageIncome-p2.AverageIncome, 2),
	)
}
