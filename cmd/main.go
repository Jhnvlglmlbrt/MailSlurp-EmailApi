package main

import (
	"context"
	"fmt"

	"MailSlurpGo/internal/controllers"
	"MailSlurpGo/internal/services"

	mailslurp "github.com/mailslurp/mailslurp-client-go"
)

func main() {
	// Ваш API ключ
	apiKey := "b2cc7c8960b89031022e872e1876ddc8fc9a638ba81e0e30478e6b858e6698a6"

	// Initialize MailSlurp client
	client, ctx := createClient(apiKey)

	inboxController := &controllers.InboxController{Client: *client}
	inboxService := &services.InboxService{Controller: inboxController}

	// Create several inboxes
	for i := 1; i <= 2; i++ {
		name := fmt.Sprintf("hannibal%d", i)
		if err := inboxService.CreateInbox(ctx, name); err != nil {
			fmt.Printf("Error creating inbox %d: %s\n", i, err)
		}
	}
}

func createClient(apiKey string) (*mailslurp.APIClient, context.Context) {
	// Создаем контекст с вашим API ключом
	ctx := context.WithValue(context.Background(), mailslurp.ContextAPIKey, mailslurp.APIKey{Key: apiKey})

	// Создаем клиента MailSlurp
	config := mailslurp.NewConfiguration()

	return mailslurp.NewAPIClient(config), ctx
}
