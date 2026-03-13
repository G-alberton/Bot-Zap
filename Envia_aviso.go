/*BY Gustavo Borghezan Alberton
  Version 1.0
  Esse programa é responsavel por enviar as mensagens de informação de data de pagamento de boleto*/

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func EnviarAviso() {
	whatsapp := os.Args[1]
	nome := os.Args[2]
	valorPagar := os.Args[3]
	dataVencimento := os.Args[4]
	phoneNumberID := //ID do numero que utilizara na API (Deve ser numero virgem) 
	acessToken := //Token de acesso da API

	fmt.Println("Nr. WhatsApp:", whatsapp)
	fmt.Println("Cliente.....:", nome)

	payload := map[string]interface{}{
		"messaging_product": "whatsapp",
		"to":                whatsapp,
		"type":              "template",
		"template": map[string]interface{}{
			"name": "primeiro_contato_v2",
			"language": map[string]string{
				"code": "pt_BR",
			},
			"components": []map[string]interface{}{
				{
					"type": "body",
					"parameters": []map[string]string{
						{"type": "text", "text": nome},
						{"type": "text", "text": valorPagar},
						{"type": "text", "text": dataVencimento},
					},
				},
			},
		},
	}

	jsonData, err := json.Marshal(payload)
	if err != nil {
		log.Fatal(err)
	}

	url := fmt.Sprintf("https://graph.facebook.com/v22.0/%s/messages", phoneNumberID)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("Authorization", "Bearer "+acessToken)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)

	fmt.Println("StatusConde:", resp.StatusCode)
	fmt.Println("Response body:", string(body))

	//log
	arquivo, err := os.Create("C:\\temp\\envia_aviso_atraso.log")
	if err != nil {
		log.Fatal(err)
	}
	defer arquivo.Close()

	fmt.Fprintln(arquivo, "Status:", resp.StatusCode)
	fmt.Fprintln(arquivo, "Resposta:", string(body))
}

func main() {

	EnviarAviso()

}
