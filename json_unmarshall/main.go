package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/grassiricardo/json_unmarshall/model"
)

func main() {
	cliente := &http.Client{
		Timeout: time.Second * 30,
	}

	request, err := http.NewRequest("GET", "https://jsonplaceholder.typicode.com/posts/1", nil)
	if err != nil {
		fmt.Println("[main] Erro para pegar a request. ERRO",
			err.Error())
		return
	}

	resposta, err := cliente.Do(request)
	if resposta.StatusCode == 200 {
		corpo, err := ioutil.ReadAll(resposta.Body)
		if err != nil {
			fmt.Println("[main] Erro ao abrir enviar um usuario. ERRO",
				err.Error())
			return
		}
		fmt.Println(string(corpo))
		post := model.BlogPost{}
		err = json.Unmarshal(corpo, &post)
		if err != nil {
			fmt.Println("[main] Erro ao converter o retorno da API",
				err.Error())
			return
		}
		fmt.Printf("Post é: %+v\r\n", post)
	}
}
