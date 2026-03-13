package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"net/textproto"
	"os"
)

func UploadImage(caminho string, phoneNumberID string, acessToken string) string {

	file, err := os.Open(caminho)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	header := make(textproto.MIMEHeader)
	header.Set("Content-Disposition", `form-data; name="file"; filename="qrcode.jpeg"`)
	header.Set("Content-Type", "image/jpeg")

	part, err := writer.CreatePart(header)
	if err != nil {
		log.Fatal(err)
	}

	_, err = io.Copy(part, file)
	if err != nil {
		log.Fatal(err)
	}

	err = writer.WriteField("messaging_product", "whatsapp")
	if err != nil {
		log.Fatal(err)
	}

	writer.Close()

	url := fmt.Sprintf("https://graph.facebook.com/v22.0/%s/media", phoneNumberID)

	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("Authorization", "Bearer "+acessToken)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	resBody, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Upload responde:", string(resBody))

	if resp.StatusCode != 200 {
		log.Fatal("Erro no upload:", string(resBody))
	}

	var result map[string]interface{}
	err = json.Unmarshal(resBody, &result)
	if err != nil {
		log.Fatal(err)
	}

	mediaID, ok := result["id"].(string)
	if !ok {
		log.Fatal("MediaID não encontrado:", string(resBody))
	}

	return mediaID
}

func EnviarAviso() {

	if len(os.Args) < 5 {
		log.Fatal("Uso: programa whatsapp nome caminhoImagem pixCopiaCola")
	}

	whatsapp := os.Args[1]
	nome := os.Args[2]
	caminhoImagem := os.Args[3]
	pixCopiaCola := os.Args[4]

	phoneNumberID := //ID do numero que utilizara na API (Deve ser numero virgem)
	acessToken := //Token de acesso da API

	mediaID := UploadImage(caminhoImagem, phoneNumberID, acessToken)

	fmt.Println("Nr. Whatsapp:", whatsapp)
	fmt.Println("Cliente.....:", nome)

	payload := map[string]interface{}{
		"messaging_product": "whatsapp",
		"to":                whatsapp,
		"type":              "template",
		"template": map[string]interface{}{
			"name": "envia_pix_v2",
			"language": map[string]string{
				"code": "pt_BR",
			},
			"components": []interface{}{
				map[string]interface{}{
					"type": "header",
					"parameters": []interface{}{
						map[string]interface{}{
							"type": "image",
							"image": map[string]string{
								"id": mediaID,
							},
						},
					},
				},
				map[string]interface{}{
					"type": "body",
					"parameters": []map[string]string{
						{"type": "text", "text": nome},
						{"type": "text", "text": pixCopiaCola},
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

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	arquivo, err := os.Create("C:\\temp\\envia_pix.log")
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
