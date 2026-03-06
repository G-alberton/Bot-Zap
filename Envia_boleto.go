/*BY Gustavo Borghezan Alberton
  Version 1.0
  Esse programa é responsavel por enviar a mensagem junto com o Boleto
  ~ Necessario a utilização de mediaID, ele envia o arquivo para o servidor da Meta e utiliza o ID para informar no JSON
  */


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

func UploadArquivo(caminho string, phoneNumberID string, acessToken string) string{
	//O MediaID funciona para o envio de arquivos por API oficial 
	file, err := os.Open(caminho)
	if err != nil {
		log.Fatal(err)
	}
	def file.Close()

	body := &bytes.Buffer{}
	write := multipart.newWriter(body)

	part, err := write.CreateFormFile("File", "boleto.pdf")
	if err != nil{
		log.Fatal(err)
	}

	io.copy(part, file)

	write.writeField("messaging_product", "whatsapp")

	write.close()

	utl := fmt.Sprintf("https://graph.facebook.com/v22.0/%s/media", phoneNumberID)

	req, err := http.NewRequest("POST", url, body)
	if err != nil{
		log.Fatal(err)
	}

	req.Header.Set("Authorization", "Bearer "+acessToken)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	client := &http.Cliente{}
	resp, err := client.Do(req)
	if err != nil{
		log.Fatal(err)
	}

	defer resp.Body.Close()

	resBody, _ := io.ReadAll(resp.Body)

	fmt.Println("Upload responde:", string(resBody))

	var result map[string]interface{}
	json.Unmarshal(resBody, &result)

	mediaID := result["id"].(string)

	fmt.Println("MediaID:", mediaID)

	return mediaID
}

func EnviaBoleto() {
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
		"templete": map[string]interface{}{
			"name": "manda_boleto_v2",
			"language": map[string]string{
				"code": "pt_BR",
			},
			"components": []interface{}{
				map[string]interface{}{
					"type": "header",
					"parameters": []interface{}{
						map[string]interface{}{
							"type": "document",
							"document": map[string]string{
								"id": mediaID,
								"filename": "boleto.pdf",
							},
						},
					},
				},
				map[string]interface{}{
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
	if err !=nil {
		log.Fatal(err)
	}

	url := fmt.Sprintf("https://graph.facebook.com/v22.0/%s/messages", phoneNumberID)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		log.Fatal(err)
	}

	req.Header.set("Autorização", "Bearer "+acessToken)
	req.Header.set("content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil{
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	fmt.Fprintln(arquivo, "Status:", resp.StatusCode)
	fmt.Fprintln(arquivo, "Resposta", string(body))

	arquivo, err := os.Create("C\\temp\\envia_boleto.log")
	if err != nil{
		log.Fatal(err)
	}
	defer arquivo.Close()
}

func main() {
	if len(os.Args) != 6 {
		fmt.Println("Parametro insulficientes")
		fmt.Println("Uso:")
		fmt.Println("go run Enviar_boleto.go telefone nome valor vencimento mediaID")
		return
	}

	caminhoPDF := os.Args[5]

	phoneNumberID := //ID do numero que utilizara na API (Deve ser numero virgem)
	acessToken := //Token de acesso da API

	fmt.Println("Lendo arquivo:" caminhoPDF)

	mediaID := UploadArquivo(caminhoPDF, phoneNumberID, acessToken)

	EnviaBoleto()
}