package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type Endereco struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
}

func (e *Endereco) String() string {
	return fmt.Sprintf("Cep: %s, Logradouro: %s, Complemento: %s, Bairro: %s, Localidade: %s, Uf: %s",
		e.Cep, e.Logradouro, e.Complemento, e.Bairro, e.Localidade, e.Uf)
}

type Address struct {
	Cep          string `json:"cep"`
	State        string `json:"state"`
	City         string `json:"city"`
	Neighborhood string `json:"neighborhood"`
	Street       string `json:"street"`
}

func (a *Address) String() string {
	return fmt.Sprintf("Cep: %s, Street: %s, Neighborhood: %s, City: %s, State: %s",
		a.Cep, a.Street, a.Neighborhood, a.City, a.State)
}

func main() {
	var cep string
	fmt.Print("Digite o CEP: ")
	fmt.Scanln(&cep)

	chViaCep := make(chan interface{})
	chBrasilCep := make(chan interface{})
	timeout := time.After(time.Second)

	go func() {
		// time.Sleep(time.Second * 2)
		endereco := &Endereco{}
		requestAPI("https://viacep.com.br/ws/"+cep+"/json/", chViaCep, endereco)
	}()

	go func() {
		// time.Sleep(time.Second * 2)
		address := &Address{}
		requestAPI("https://brasilapi.com.br/api/cep/v1/"+cep, chBrasilCep, address)
	}()

	select {
	case res := <-chViaCep:
		endereco := res.(*Endereco)
		fmt.Println("Resultado da API ViaCep:", endereco.String())
	case res := <-chBrasilCep:
		address := res.(*Address)
		fmt.Printf("Resultado da API BrasilApi: %+v\n", address)
	case <-timeout:
		fmt.Println("Erro de timeout: nenhuma resposta em 1 segundo")
	}
}

func requestAPI(url string, ch chan<- interface{}, result interface{}) {
	client := http.Client{
		Timeout: time.Second,
	}
	res, err := client.Get(url)

	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}
	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	json.Unmarshal(body, result)

	ch <- result
}
