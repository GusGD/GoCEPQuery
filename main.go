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
	return fmt.Sprintf("CEP: %s, Logradouro: %s, Complemento: %s, Bairro:%s, Localidade: %s, UF: %s", e.Cep, e.Logradouro, e.Complemento, e.Bairro, e.Localidade, e.Uf)
}

type Address struct {
	Cep          string `json:"cep"`
	State        string `json:"state"`
	City         string `json:"city"`
	Neighborhood string `json:"neighborhood"`
	Street       string `json:"street"`
}

func (a *Address) String() string {
	return fmt.Sprintf("CEP: %s, State: %s, City: %s, Neighborhood: %s, Street: %s", a.Cep, a.State, a.City, a.Neighborhood, a.Street)
}

func main() {
	var cep string
	fmt.Print("Digite o CEP: ")
	fmt.Scanln(&cep)

	chViaCep := make(chan interface{})
	chBrasilCep := make(chan interface{})
	timeout := time.After(time.Second)

	go func() {
		// time.Sleep(time.Second)
		endereco := &Endereco{}
		requestAPI("https://viacep.com.br/ws/"+cep+"/json/", chViaCep, &endereco)
	}()

	go func() {
		// time.Sleep(time.Second)
		address := &Address{}
		requestAPI("https://brasilapi.com.br/api/cep/v1/"+cep, chBrasilCep, address)
	}()

	select {
	case res := <-chViaCep:
		endereco := res.(*Endereco)
		fmt.Println("Resultado via CEP:", endereco.String())
	case res := <-chBrasilCep:
		address := res.(*Address)
		fmt.Println("Resultado Brasil API:", address.String())
	case <-timeout:
		fmt.Println("Timeout")
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
