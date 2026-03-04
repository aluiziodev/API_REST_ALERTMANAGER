# API_REST_ALERTMANAGER
Repositório referente a uma API REST de envio de alertas/mensagens para sms, email, telegram e slack.


---


## Como Iniciar a API?


1. Definir todas as variáveis de ambiente em um .env
```env
export EMAIL_TEMPLATE_PATH=
export EMAIL_SERVER=
export EMAIL_PORT=
export EMAIL_ADRRESS=
export GMAIL_PASSWORD=
export SLACK_TOKEN=
export SLACK_CHANNEL_ID=
export SMS_ENDPOINT=
export NEXMO_API_KEY=
export NEXMO_API_SECRET=
export TELEGRAM_BOT_API=
```
2. Gerar um binário para a aplicação
```bash
go build
```
3. Rodar o programa
```bash
./alertmanager
```


---


Após a primeira inicialização, dentro do diretório será criado um .log onde é possível visualizar os status da API


O programa utilizado para realização das requisições http foi o Postman no endereço http://localhost:8080
