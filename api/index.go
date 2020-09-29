package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/kataras/go-mailer"
)

type Receiver struct {
	Name     string
	LastName string `json:"last_name"`
	Email    string
	Subject  string
	Message  string
}

func Handler(w http.ResponseWriter, r *http.Request) {

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "{}")
		return
	}

	config := mailer.Config{
		Host:       os.Getenv("MAILER_HOST"),
		Username:   os.Getenv("MAILER_USERNAME"),
		Password:   os.Getenv("MAILER_PASSWORD"),
		Port:       587,
		UseCommand: false,
		FromAddr:   "ifrancisco.iglesias@gmail.com",
	}
	sender := mailer.New(config)
	var receiver Receiver
	err := json.NewDecoder(r.Body).Decode(&receiver)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	subject := "Thanks for your email"

	resp, err := http.Get("https://go-mailer.vercel.app/email.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	parsedTemplate := template.Must(template.New("email-template").Parse(string(bodyBytes)))
	var templateBuffer bytes.Buffer
	parsedTemplate.Execute(&templateBuffer, receiver)

	to := []string{"ifrancisco.iglesias@gmail.com", receiver.Email}

	err = sender.Send(subject, templateBuffer.String(), to...)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = sender.Send(receiver.Subject, receiver.Message, "ifrancisco.iglesias@gmail.com")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "{}")
}
