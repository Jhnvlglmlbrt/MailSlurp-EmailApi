package main

import (
	"context"
	"fmt"

	"github.com/antihax/optional"
	mailslurp "github.com/mailslurp/mailslurp-client-go"
)

// createClient создает клиента и контекст с API ключом
func createClient() (*mailslurp.APIClient, context.Context) {
	// Здесь нужно указать ваш API ключ
	apiKey := "Your API-key"

	// Создаем контекст с вашим API ключом
	ctx := context.WithValue(
		context.Background(),
		mailslurp.ContextAPIKey,
		mailslurp.APIKey{Key: apiKey},
	)

	// Создаем клиента MailSlurp
	config := mailslurp.NewConfiguration()
	return mailslurp.NewAPIClient(config), ctx
}

func main() {
	// Получаем клиента и контекст
	client, ctx := createClient()

	// Создаем несколько почтовых ящиков
	for i := 1; i <= 3; i++ {
		// Вызываем контроллер для создания почтового ящика с указанным именем
		opts := &mailslurp.CreateInboxOpts{
			Name: optional.NewString(fmt.Sprintf("hannibal%d", i)),
		}
		inbox, response, err := client.InboxControllerApi.CreateInbox(ctx, opts)

		if err != nil {
			fmt.Printf("Произошла ошибка при создании почтового ящика %d: %s\n", i, err)
			continue
		}

		if response.StatusCode != 201 {
			fmt.Printf("Не удалось создать почтовый ящик %d. Код статуса: %d\n", i, response.StatusCode)
			continue
		}

		fmt.Printf("Почтовый ящик %d с ID: %s и адресом: %s успешно создан.\n", i, inbox.Id, inbox.EmailAddress)
	}
}
