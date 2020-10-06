package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/grassiricardo/web/model"
)

func main() {
	cliente := &http.Client{
		Timeout: time.Second * 30,
	}

	usuario := model.Usuario{}
	usuario.ID = 1
	usuario.Nome = "Grassi"

	conteudoEnviar, err := json.Marshal(usuario)
	if err != nil {
		fmt.Println("[main] Erro ao gerar o usuário do programa",
			err.Error())
		return
	}

	request, err := http.NewRequest("POST", "http://requestbin.net/r/1ixphll1", bytes.NewBuffer(conteudoEnviar))
	if err != nil {
		fmt.Println("[main] Erro ao enviar um novo usuario. ERRO",
			err.Error())
		return
	}
	request.SetBasicAuth("fizz", "buzz")
	request.Header.Set("content-type", "application/json; charset=utf-8")
	resposta, err := cliente.Do(request)
	if resposta.StatusCode == 200 {
		corpo, err := ioutil.ReadAll(resposta.Body)
		if err != nil {
			fmt.Println("[main] Erro ao abrir enviar um usuario. ERRO",
				err.Error())
			return
		}
		fmt.Println(string(corpo))
	}

	// resposta, err := cliente.Get("https://www.google.com.br")
	// if err != nil {
	// 	fmt.Println("[main] Erro ao abrir a pagina do Google. ERRO",
	// 		err.Error())
	// 	return
	// }
	// defer resposta.Body.Close()

	// if resposta.StatusCode == 200 {
	// 	corpo, err := ioutil.ReadAll(resposta.Body)
	// 	if err != nil {
	// 		fmt.Println("[main] Erro ao abrir a pagina do Google. ERRO",
	// 			err.Error())
	// 		return
	// 	}
	// 	fmt.Println(string(corpo))
	// }

	// request, err := http.NewRequest("GET", "https://www.google.com.br", nil)
	// if err != nil {
	// 	fmt.Println("[main] Erro ao criar uma request Google. ERRO",
	// 		err.Error())
	// 	return
	// }
	// request.SetBasicAuth("teste", "teste")
	// resposta, err = cliente.Do(request)
	// if resposta.StatusCode == 200 {
	// 	corpo, err := ioutil.ReadAll(resposta.Body)
	// 	if err != nil {
	// 		fmt.Println("[main] Erro ao abrir a pagina do Google. ERRO",
	// 			err.Error())
	// 		return
	// 	}
	// 	fmt.Println(string(corpo))
	// }
}
