package main

import (
	"fmt"
	"log"
	"net/smtp"
	"strconv"
	"sync"
	"time"
)

func emailWorker(id int, ch chan Receiver,wg *sync.WaitGroup) {
	defer wg.Done()
	for receiv:= range ch{
		// Use SMTP settings from configuration
		smtpHost := AppConfig.SMTPDef.Host
		smtpPort := AppConfig.SMTPDef.Port
		serverAddress := smtpHost + ":" + strconv.Itoa(smtpPort)


		// formatedMsg := fmt.Sprintf("To: %s\r\nSubject: Test Email\r\n\n%s\r\n",receiv.Email,"Just testing our email")
		// msg := []byte(formatedMsg)
		// fmt.Printf("Worker %d: Sending email to %s <%s> using SMTP server %s\n", 
		// 	id, receiv.Name, receiv.Email, serverAddress)


		msg ,err := executeTemplate(receiv)
		if err != nil{
			fmt.Printf("Worker %d has found error parsing template %s",id,receiv.Email)
			// Adding to the dead Queue
			continue
		}
		fmt.Printf("Worker %d: Sending email to %s <%s> using SMTP server %s\n", id, receiv.Name, receiv.Email, serverAddress)

		time.Sleep(time.Millisecond * 50)
		if err := smtp.SendMail(serverAddress,nil,"abhaykumarss9110@gmail.com", []string{receiv.Email},[]byte(msg)); err != nil{
			log.Fatal(err)
		}
		fmt.Printf("Sent the mail to %s:",receiv.Email)
	}
}