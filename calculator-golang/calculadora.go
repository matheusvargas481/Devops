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
	PrimeiroNumero float64 `json:"PrimeiroNumero,omitempty"`
	Operacao       string  `json:"Operacao,omitempty"`
	SegundoNumero  float64 `json:"SegundoNumero,omitempty"`
	Resultado      float64 `json:"Resultado,omitempty"`
}

var calculo []*Calculadora

func historicoOperacao(response http.ResponseWriter, request *http.Request) {
	historicoJson, err := json.Marshal(calculo)
	if err == nil {
		fmt.Print(err)
	}
	response.Write(historicoJson)
}

func sum(response http.ResponseWriter, request *http.Request) {
	num1 := mux.Vars(request)["num1"]
	PrimeiroNumero, err := strconv.ParseFloat(num1, 64)
	if err == nil {
		fmt.Print(err)
	}
	fmt.Println("Primeiro Número:", PrimeiroNumero)

	Operacao := mux.Vars(request)["operacao"]
	fmt.Println("Operação:", Operacao)

	num2 := mux.Vars(request)["num2"]
	SegundoNumero, err := strconv.ParseFloat(num2, 64)
	if err == nil {
		fmt.Print(err)
	}
	fmt.Println("Segundo Número:", SegundoNumero)

	Resultado := PrimeiroNumero + SegundoNumero

	sum := new(Calculadora)
	sum.PrimeiroNumero = PrimeiroNumero
	sum.Operacao = Operacao
	sum.SegundoNumero = SegundoNumero
	sum.Resultado = Resultado

	sumJson, _ := json.Marshal(sum)
	calculo = append(calculo, sum)
	response.Write(sumJson)
}

func sub(response http.ResponseWriter, request *http.Request) {
	num1 := mux.Vars(request)["num1"]
	PrimeiroNumero, err := strconv.ParseFloat(num1, 64)
	if err == nil {
		fmt.Print(err)
	}
	fmt.Println("Primeiro Número:", PrimeiroNumero)

	Operacao := mux.Vars(request)["operacao"]
	fmt.Println("Operação:", Operacao)

	num2 := mux.Vars(request)["num2"]
	SegundoNumero, err := strconv.ParseFloat(num2, 64)
	if err == nil {
		fmt.Print(err)
	}
	fmt.Println("Segundo Número:", SegundoNumero)

	Resultado := (PrimeiroNumero - SegundoNumero)

	sub := new(Calculadora)
	sub.PrimeiroNumero = PrimeiroNumero
	sub.Operacao = Operacao
	sub.SegundoNumero = SegundoNumero
	sub.Resultado = Resultado

	subJson, _ := json.Marshal(sub)
	calculo = append(calculo, sub)
	response.Write(subJson)
}

func mul(response http.ResponseWriter, request *http.Request) {
	num1 := mux.Vars(request)["num1"]
	PrimeiroNumero, err := strconv.ParseFloat(num1, 64)
	if err == nil {
		fmt.Print(err)
	}
	fmt.Println("Primeiro Número:", PrimeiroNumero)

	Operacao := mux.Vars(request)["operacao"]
	fmt.Println("Operação:", Operacao)

	num2 := mux.Vars(request)["num2"]
	SegundoNumero, err := strconv.ParseFloat(num2, 64)
	if err == nil {
		fmt.Print(err)
	}
	fmt.Println("Segundo Número:", SegundoNumero)

	Resultado := PrimeiroNumero * SegundoNumero

	mul := new(Calculadora)
	mul.PrimeiroNumero = PrimeiroNumero
	mul.Operacao = Operacao
	mul.SegundoNumero = SegundoNumero
	mul.Resultado = Resultado

	mulJson, _ := json.Marshal(mul)
	calculo = append(calculo, mul)
	response.Write(mulJson)
}

func div(response http.ResponseWriter, request *http.Request) {
	num1 := mux.Vars(request)["num1"]
	PrimeiroNumero, err := strconv.ParseFloat(num1, 64)
	if err == nil {
		fmt.Print(err)
	}
	fmt.Println("Primeiro Número:", PrimeiroNumero)

	Operacao := mux.Vars(request)["operacao"]
	fmt.Println("Operação:", Operacao)

	num2 := mux.Vars(request)["num2"]
	SegundoNumero, err := strconv.ParseFloat(num2, 64)
	if err == nil {
		fmt.Print(err)
	}
	fmt.Println("Segundo Número:", SegundoNumero)

	Resultado := PrimeiroNumero / SegundoNumero

	div := new(Calculadora)
	div.PrimeiroNumero = PrimeiroNumero
	div.Operacao = Operacao
	div.SegundoNumero = SegundoNumero
	div.Resultado = Resultado

	divJson, _ := json.Marshal(div)
	calculo = append(calculo, div)
	response.Write(divJson)
}

func calcular(response http.ResponseWriter, request *http.Request) {
	Operacao := mux.Vars(request)["operacao"]
	fmt.Println("Operação:", Operacao)

	switch Operacao {
	case "sum":
		sum(response, request)
	case "sub":
		sub(response, request)
	case "mul":
		mul(response, request)
	case "div":
		div(response, request)
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
