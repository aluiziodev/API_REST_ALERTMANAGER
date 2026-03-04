package email

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/smtp"
	"os"
	"text/template"
)

/*func SendMail(to []string, subject string, server string, erro string, horario string,
	templatePath string) {
	from := "aluneto10@gmail.com"
	password := os.Getenv("GMAIL_PASSWORD")
	if password == "" {
		panic("GMAIL_PASSWORD environment variable is not set")
	}

	smtpHost := "smtp.gmail.com"
	smtpPort := "587"
	auth := smtp.PlainAuth("", from, password, smtpHost)
	t, _ := template.ParseFiles(templatePath)
	var body bytes.Buffer
	mimeHeaders := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	body.Write([]byte(fmt.Sprintf("Subject: %s\n%s\n\n", subject, mimeHeaders)))
	t.Execute(&body, struct {
		Server  string
		Error   string
		Horario string
	}{
		Server:  server,
		Error:   erro,
		Horario: horario,
	})

	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, body.Bytes())
	if err != nil {
		fmt.Printf("Erro ao enviar o email! %s", err)
		os.Exit(1)
	}
	fmt.Println("Email enviado com sucesso!!")

}*/

type ErrorMessage struct {
	Error string `json:"error"`
}

type Email struct {
	To      []string `json:"to"`
	Subject string   `json:"subject"`
	Server  string   `json:"server"`
	Horario string   `json:"horario"`
	Error   string   `json:"error"`
}

func SendMail(w http.ResponseWriter, r *http.Request) {
	var errorMessage ErrorMessage
	templatePath := os.Getenv("EMAIL_TEMPLATE_PATH")
	if templatePath == "" {
		log.Fatal("Template path nao esta definido!!")
	}
	server := os.Getenv("EMAIL_SERVER")
	if server == "" {
		log.Fatal("Email server nao esta definido!!")
	}
	porta := os.Getenv("EMAIL_PORT")
	if porta == "" {
		log.Fatal("Email port nao esta definido!!")
	}
	from := os.Getenv("EMAIL_ADRRESS")
	if from != "" {
		log.Fatal("Endereço de email nao esta definido!!")
	}
	password := os.Getenv("GMAIL_PASSWORD")
	if password == "" {
		log.Fatal("GMAIL_PASSWORD nao esta definido!!")
	}
	email := Email{}
	err := json.NewDecoder(r.Body).Decode(&email)
	if err != nil {
		log.Printf("Erro ao decodificar os dados do email: %s", err.Error())
		errorMessage.Error = "Error decoding email data"
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errorMessage)
		return
	}
	auth := smtp.PlainAuth("", from, password, server)
	t, err := template.ParseFiles(templatePath)
	if err != nil {
		log.Fatal("Erro ao abrir o template")
	}
	var body bytes.Buffer
	mimeHeaders := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	body.Write([]byte(fmt.Sprintf("Subject: %s\n%s\n\n", email.Subject, mimeHeaders)))
	t.Execute(&body, email)
	err = smtp.SendMail(server+":"+porta, auth, from, email.To, body.Bytes())
	if err != nil {
		log.Printf("Erro ao enviar o email %s", err.Error())
		errorMessage.Error = "Error ao enviar mail"
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errorMessage)
		return
	}
	log.Printf("Email enviado com sucesso!!")

}
