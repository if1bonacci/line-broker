package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/line/line-bot-sdk-go/v8/linebot/messaging_api"
	"github.com/line/line-bot-sdk-go/v8/linebot/webhook"
)

const (
	secret = "11233"
	token  = "121212"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	bot, err := messaging_api.NewMessagingApiAPI(token)
	if err != nil {
		log.Fatal(err)
	}

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})

	r.Post("/api/v1/line/callback", func(w http.ResponseWriter, r *http.Request) {
		log.Println("/callback called...")

		cb, err := webhook.ParseRequest(secret, r)
		if err != nil {
			log.Printf("Cannot parse request: %+v\n", err)
			if errors.Is(err, webhook.ErrInvalidSignature) {
				w.WriteHeader(400)
			} else {
				w.WriteHeader(500)
			}
			return
		}

		log.Println("Handling events...")
		for _, event := range cb.Events {
			log.Printf("/callback called%+v...\n", event)

			switch e := event.(type) {
			case webhook.MessageEvent:
				switch message := e.Message.(type) {
				case webhook.TextMessageContent:
					if _, err = bot.ReplyMessage(
						&messaging_api.ReplyMessageRequest{
							ReplyToken: e.ReplyToken,
							Messages: []messaging_api.MessageInterface{
								messaging_api.TextMessage{
									Text: message.Text,
								},
							},
						},
					); err != nil {
						log.Print(err)
					} else {
						log.Println("Sent text reply.")
					}
				case webhook.StickerMessageContent:
					replyMessage := fmt.Sprintf(
						"Sticker message: sticker id is %s, stickerResourceType is %s", message.StickerId, message.StickerResourceType)
					if _, err = bot.ReplyMessage(
						&messaging_api.ReplyMessageRequest{
							ReplyToken: e.ReplyToken,
							Messages: []messaging_api.MessageInterface{
								messaging_api.TextMessage{
									Text: replyMessage,
								},
							},
						}); err != nil {
						log.Print(err)
					} else {
						log.Println("Sent sticker reply.")
					}
				case webhook.MemberJoinedEvent:
					log.Printf("Member joined: %s\n", e.Source.(webhook.UserSource).UserId)
				case webhook.MemberLeftEvent:
					log.Printf("Member joined: %s\n", e.Source.(webhook.UserSource).UserId)
				case webhook.FollowEvent:
					log.Printf("Follow event: %s\n", e.Source.(webhook.UserSource).UserId)
				case webhook.BeaconEvent:
					log.Printf("Beacon event: %s\n", e.Source.(webhook.UserSource).UserId)
				default:
					log.Printf("Unsupported message content: %T\n", e.Message)
				}
			default:
				log.Printf("Unsupported message: %T\n", event)
			}
		}

		w.WriteHeader(http.StatusOK)
	})

	http.ListenAndServe(":4040", r)
}
