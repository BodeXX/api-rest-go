package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func Home(w http.ResponseWriter, r *http.Request) {
	moedaOrigem := r.URL.Query().Get("moedaOrigem")
	moedaDestino := r.URL.Query().Get("moedaDestino")
	valorStr := r.URL.Query().Get("valor")

	valor, _ := strconv.ParseFloat(valorStr, 64)

	//Ultima atualização de taxas dia 24 Set 2024 as 14:00 Horário de Brasilia

	var taxas = map[string]map[string]float64{
		"USD": {
			"EUR": 0.90,   //1 USD = 0.90 EUR (EURO)
			"BRL": 5.46,   //1 USD = 5.46 BRL (REAL)
			"JPY": 143.69, //1 USD = 110.25 JPY (IENES)
			"CHF": 0.85,   //1 USD = 0.85 CHF (FRANCO SUÍÇO)
		},
		"EUR": {
			"USD": 1.12,   //1 EUR = 1.12 USD (DOLAR)
			"BRL": 6.09,   //1 EUR = 6.09 BRL (REAL)
			"JPY": 160.27, //1 EUR = 160.27 JPY (IENES)
			"CHF": 0.94,   //1 EUR = 0.94 CHF (FRANCO SUÍÇO)

		},
		"BRL": {
			"EUR": 0.16,  //1 BRL = 0.16 EUR (EURO)
			"USD": 0.18,  //1 BRL = 0.18 USD (DOLAR)
			"JPY": 26.32, //1 BRL = 26.32 JPY (IENES)
			"CHF": 0.15,  //1 BRL = 0.15 CHF (FRANCO SUÍÇO)
		},
		"JPY": {
			"USD": 0.0070, //1 JPY = 0.0070 USD (DOLAR)
			"EUR": 0.0063, //1 JPY = 0.0063 EUR (EURO)
			"BRL": 0.0381, //1 JPY = 0.047 BRL (REAL)
			"CHF": 0.0059, //1 JPY = 0.0059 CHF (FRANCO SUÍÇO)
		},
		"CHF": {
			"USD": 1.18,   //1 CHF = 1.18 USD (DOLAR)
			"EUR": 1.06,   //1 CHF = 1.06 EUR (EURO)
			"BRL": 6.46,   //1 CHF = 6.46 BRL (REAL)
			"JPY": 169.55, //1 CHF = 169.55 JPY (IENES)
		},
	}

	if taxaDestino, destinoExiste := taxas[moedaOrigem][moedaDestino]; destinoExiste {
		valorConvertido := valor * taxaDestino
		fmt.Fprintf(w, "Valor convertido: %.2f %s", valorConvertido, moedaDestino)
	} else {
		http.Error(w, "Moeda de destino inválida", http.StatusBadRequest)
	}
}

func HandleRequest() {
	http.HandleFunc("/", Home)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
	fmt.Printf("Iniciando o servidor Rest com Go")
	HandleRequest()
}
