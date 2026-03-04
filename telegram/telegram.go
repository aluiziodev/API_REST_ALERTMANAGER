package telegram

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type Message struct {
	Text    string `json:"text"`
	GroupID int64  `json:"groupid"`
}

type ErrorMessage struct {
	Error string `json:"error"`
}

func SendTelegram(w http.ResponseWriter, r *http.Request) {
	token := os.Getenv("TELEGRAM_BOT_API")
	if token == "" {
		log.Fatal("token nao definido!!")
	}
	var errormessage ErrorMessage
	message := Message{}

	err := json.NewDecoder(r.Body).Decode(&message)
	if err != nil {
		log.Printf("Erro ao decodificar a mensagem %s", err.Error())
		errormessage.Error = fmt.Sprintf("Erro ao decodificar a mensagem %s", err.Error())
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errormessage)
		return
	}
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Printf("Erro ao criar o bot do Telegram %s", err.Error())
		errormessage.Error = fmt.Sprintf("Erro ao criar o bot do Telegram %s", err.Error())
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errormessage)
		return
	}
	alertText := tgbotapi.NewMessage(message.GroupID, message.Text)
	retorno, err := bot.Send(alertText)
	if err != nil {
		log.Printf("Erro ao enviar a mensagem para o Telegram %s", err.Error())
		errormessage.Error = fmt.Sprintf("Erro ao enviar a mensagem para o Telegram %s", err.Error())
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errormessage)
		return
	}
	log.Printf("Mensagem enviada com sucesso: %d", retorno.MessageID)

}
