package sms

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
)

type SmsMessage struct {
	Message string `json:"message"`
	Phone   string `json:"phone"`
}

func SendSMS(w http.ResponseWriter, r *http.Request) {
	endpoint := os.Getenv("SMS_ENDPOINT")
	if endpoint == "" {
		log.Fatalf("SMS_ENDPOINT nao foi definido!!")
	}

	data := url.Values{}
	api_key := os.Getenv("NEXMO_API_KEY")
	if api_key == "" {
		log.Fatalf("NEXMO_API_KEY nao foi definido!!")
	}
	api_secret := os.Getenv("NEXMO_API_SECRET")
	if api_secret == "" {
		log.Fatalf("NEXMO_API_SECRET nao foi definido!!")
	}
	smsMessage := SmsMessage{}
	err := json.NewDecoder(r.Body).Decode(&smsMessage)
	if err != nil {
		log.Fatalf("Erro ao decodificar os dados da messagem: %v", err)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	data.Set("api_key", api_key)
	data.Set("api_secret", api_secret)
	data.Set("to", smsMessage.Phone)
	data.Set("text", smsMessage.Message)
	data.Set("from", "Sistemas de Alertas XGH")

	client := &http.Client{}
	r, erro := http.NewRequest("POST", endpoint, strings.NewReader(data.Encode()))
	if erro != nil {
		log.Fatalf("Erro ao enviar sms: %v", erro)
	}
	r.Header.Add("Content-type", "application/x-www-form-urlenconded")
	r.Header.Add("Content-length", strconv.Itoa(len(data.Encode())))
	res, err := client.Do(r)
	if err != nil {
		log.Fatalf("Erro ao enviar sms: %v", err)
	}
	defer res.Body.Close()
	log.Println(res.Status)
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatalf("Erro ao enviar sms: %v", err)
	}
	log.Printf("%s", string(body))
}
