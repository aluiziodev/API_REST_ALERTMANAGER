package slack

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/slack-go/slack"
)

type SlackMessage struct {
	TextoAlerta string `json:"textoAlerta"`
}

func SendSlack(w http.ResponseWriter, r *http.Request) {
	token := os.Getenv("SLACK_TOKEN")
	if token == "" {
		log.Fatal("VARIAVEL SLACK_TOKEN NAO DEFINIDA!!")
	}
	channelID := os.Getenv("SLACK_CHANNEL_ID")
	if channelID == "" {
		log.Fatal("VARIAVEL SLACK_CHANNEL_ID NAO DEFINIDA!!")
	}
	slackMessage := SlackMessage{}
	err := json.NewDecoder(r.Body).Decode(&slackMessage)
	if err != nil {
		log.Fatalf("Erro ao decodificar os dados da messagem: %v", err)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	client := slack.New(token, slack.OptionDebug(true))
	attchament := slack.Attachment{
		Color:   "danger",
		Pretext: "Alerta de server down",
		Text:    slackMessage.TextoAlerta,
	}
	_, timeStamp, err := client.PostMessage(channelID, slack.MsgOptionAttachments(attchament))
	if err != nil {
		log.Fatalf("Erro ao enviar mensagem %v", err)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)

	}

	log.Printf("Mensagem enviada com sucesso %s as %s", channelID, timeStamp)

}
