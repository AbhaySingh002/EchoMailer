package main

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"path/filepath"
	"sync"

	"github.com/AbhaySingh002/lamb-launch/config"
)

type Receiver struct{
	Name  string
	Email string
}

// AppConfig holds the application configuration
var AppConfig *config.Config

func init() {
	// Load configuration
	configPath := filepath.Join(".", "config.yaml")
	var err error
	AppConfig, err = config.LoadConfig(configPath)
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}
	
	// Log SMTP configuration
	fmt.Printf("SMTP Configuration loaded - Host: %s, Port: %d\n", 
		AppConfig.SMTPDef.Host, AppConfig.SMTPDef.Port)
}

func main() {
	fmt.Println("Welcome to the email dispatcher")
	recieverChannel := make(chan Receiver) // unbuffered channel
	go func(){
		err := loadRecipient("./email.csv",recieverChannel)
		if err != nil{
			fmt.Printf("Error loading recipients: %v\n", err)
	}
	}()

	// we need to lock the main until workers are finished

	var wg sync.WaitGroup

	nWorkers := 5
	for i:=1;i<=nWorkers;i++{
		wg.Add(1)
		go emailWorker(i,recieverChannel,&wg)
	}
	wg.Wait()
}


func executeTemplate ( r Receiver) (string, error) {
	t,err := template.ParseFiles("email.tmpl")
	if err != nil{
		return "",err
	}
	var tpl bytes.Buffer
	if err := t.Execute(&tpl,r); err != nil{
		return "",err
	}
	return tpl.String(),nil
}