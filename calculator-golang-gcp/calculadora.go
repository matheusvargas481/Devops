package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"errors"

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

func sum(PrimeiroNumero float64, SegundoNumero float64) (float64, error) {
	Resultado := PrimeiroNumero + SegundoNumero
	return Resultado, nil
}
func sub(PrimeiroNumero float64, SegundoNumero float64) (float64, error) {
	Resultado := PrimeiroNumero - SegundoNumero
	return Resultado, nil
}
func mul(PrimeiroNumero float64, SegundoNumero float64) (float64, error) {
	Resultado := PrimeiroNumero * SegundoNumero
	return Resultado, nil
}

func div(PrimeiroNumero float64, SegundoNumero float64) (float64, error) {
	if SegundoNumero == 0 {
		return 0.0, errors.New("Não existe divisão por zero !")
	}
	Resultado := PrimeiroNumero / SegundoNumero
	return Resultado, nil
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
	var Resultado float64
	var erroCalcular error 
	switch Operacao {
	case "sum":
		Resultado, erroCalcular = sum(PrimeiroNumero, SegundoNumero)
	case "sub":
		Resultado, erroCalcular = sub(PrimeiroNumero, SegundoNumero)
	case "mul":
		Resultado, erroCalcular = mul(PrimeiroNumero, SegundoNumero)
	case "div":
		Resultado, erroCalcular = div(PrimeiroNumero, SegundoNumero)
	default:
		fmt.Fprint(response, "Operação inválida, digite sum[somar], sub[subtrair], mul[multiplicar] ou div[dividir] no /calc/[AQUI]/valor1/valor2 para continuar")
	}
	if erroCalcular == nil {
		response.Write(novoCalculo(PrimeiroNumero, Operacao, SegundoNumero, Resultado))
	}else{
		response.WriteHeader(http.StatusExpectationFailed)
		response.Write([]byte(erroCalcular.Error()))
	}
}

func main() {
	router := mux.NewRouter()
	api := router.PathPrefix("/calc").Subrouter()
	api.HandleFunc("/{operacao}/{num1}/{num2}", calcular)
	api.HandleFunc("/historico", historicoOperacao)
	log.Println("server port 8080")
	http.ListenAndServe(":8080", router)
}