package main

import (
	"fmt"
	"gowhisper/internal/application/services"
	config "gowhisper/internal/infrastructure"
	"gowhisper/internal/infrastructure/repositories/mongodb"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func handleConnections(chatService *services.ChatService, messageService *services.MessageService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ws, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Printf("Error upgrading connection: %v", err)
			return
		}
		defer ws.Close()

		for {
			_, msg, err := ws.ReadMessage()
			if err != nil {
				log.Printf("Error reading message: %v", err)
				break
			}
			log.Printf("Received: %s", msg)

			err = ws.WriteMessage(websocket.TextMessage, msg)
			if err != nil {
				log.Printf("Error writing message: %v", err)
				break
			}
		}
	}
}

func main() {
	uri := os.Getenv("MONGO_URI")
	if uri == "" {
		log.Fatal("MONGO_URI not set")
	}

	db := config.ConnectMongoDB(uri)

	// userRepo := mongodb.NewMongoUserRepository(db)
	chatRepo := mongodb.NewMongoChatRepository(db)
	messageRepo := mongodb.NewMongoMessageRepository(db)

	// userService := services.NewUserService(userRepo)
	chatService := services.NewChatService(chatRepo)
	messageService := services.NewMessageService(messageRepo)

	http.HandleFunc("/ws", handleConnections(chatService, messageService))

	fmt.Println("Server started on :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
