package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Calculadora struct {
	PrimeiroNumero float64 `json:"PrimeiroNumero"`
	Operacao       string  `json:"Operacao,omitempty"`
	SegundoNumero  float64 `json:"SegundoNumero"`
	Resultado      float64 `json:"Resultado"`
}

var calculo []*Calculadora

func historicoOperacao(response http.ResponseWriter, request *http.Request) {
	historicoJson, err := json.Marshal(calculo)
	if err == nil {
		fmt.Print(err)
	}
	response.Write(historicoJson)
}

func sum(PrimeiroNumero float64, SegundoNumero float64) float64 {
	Resultado := PrimeiroNumero + SegundoNumero
	return Resultado
}
func sub(PrimeiroNumero float64, SegundoNumero float64) float64 {
	Resultado := PrimeiroNumero - SegundoNumero
	return Resultado
}
func mul(PrimeiroNumero float64, SegundoNumero float64) float64 {
	Resultado := PrimeiroNumero * SegundoNumero
	return Resultado
}

func div(response http.ResponseWriter, PrimeiroNumero float64, SegundoNumero float64) float64 {
	if SegundoNumero == 0 {
		response.WriteHeader(http.StatusExpectationFailed)
		response.Write([]byte("Não existe divisão por zero !"))
	} else {
		SegundoNumero = SegundoNumero
	}
	Resultado := PrimeiroNumero / SegundoNumero
	return Resultado
}

func novoCalculo(PrimeiroNumero float64, Operacao string, SegundoNumero float64, Resultado float64) []byte {
	retorno := new(Calculadora)
	retorno.PrimeiroNumero = PrimeiroNumero
	retorno.Operacao = Operacao
	retorno.SegundoNumero = SegundoNumero
	retorno.Resultado = Resultado
	retornoJson, _ := json.Marshal(retorno)
	calculo = append(calculo, retorno)
	return retornoJson
}

func calcular(response http.ResponseWriter, request *http.Request) {
	num1 := mux.Vars(request)["num1"]
	num2 := mux.Vars(request)["num2"]
	PrimeiroNumero, err1 := strconv.ParseFloat(num1, 64)
	SegundoNumero, err2 := strconv.ParseFloat(num2, 64)
	if err1 != nil || err2 != nil {
		response.WriteHeader(http.StatusExpectationFailed)
		response.Write([]byte("Parâmetro incorreto, informe um número !"))
		return
	}
	Operacao := mux.Vars(request)["operacao"]

	switch Operacao {
	case "sum":
		Resultado := sum(PrimeiroNumero, SegundoNumero)
		response.Write(novoCalculo(PrimeiroNumero, Operacao, SegundoNumero, Resultado))
	case "sub":
		Resultado := sub(PrimeiroNumero, SegundoNumero)
		response.Write(novoCalculo(PrimeiroNumero, Operacao, SegundoNumero, Resultado))
	case "mul":
		Resultado := mul(PrimeiroNumero, SegundoNumero)
		response.Write(novoCalculo(PrimeiroNumero, Operacao, SegundoNumero, Resultado))
	case "div":
		Resultado := div(response, PrimeiroNumero, SegundoNumero)
		response.Write(novoCalculo(PrimeiroNumero, Operacao, SegundoNumero, Resultado))
	default:
		fmt.Fprint(response, "Operação inválida, digite sum[somar], sub[subtrair], mul[multiplicar] ou div[dividir] no /calc/[AQUI]/valor1/valor2 para continuar")
	}
}

func main() {
	router := mux.NewRouter()
	api := router.PathPrefix("/calc").Subrouter()
	api.HandleFunc("/{operacao}/{num1}/{num2}", calcular)
	api.HandleFunc("/historico", historicoOperacao)
	log.Println("server port 9000")
	http.ListenAndServe(":9000", router)
}