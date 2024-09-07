package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"notification-service/models"
	"notification-service/service"
)

func SendNotification(res http.ResponseWriter, req *http.Request) {
	var notification models.Notification
	err := json.NewDecoder(req.Body).Decode(&notification)
	if err != nil {
		http.Error(res, "Invalid notification", http.StatusBadRequest)
		return
	}

	if notification.Email == "" || notification.Subject == "" || notification.Body == "" {
		http.Error(res, "Missing required fields", http.StatusBadRequest)
		return
	}

	err = service.SendMessage(notification.Email, notification.Subject, notification.Body)
	if err != nil {
		log.Printf("Failed to send email: %v", err)
		http.Error(res, "Failed to send email", http.StatusInternalServerError)
		return
	}

	log.Fatal(notification.Email, notification.Subject)

	res.WriteHeader(http.StatusOK)
	res.Write([]byte("Notification sent successfully"))
}