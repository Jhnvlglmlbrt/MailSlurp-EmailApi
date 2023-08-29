package controllers

import (
	"context"
	"fmt"
	"net/http"

	"github.com/antihax/optional"
	mailslurp "github.com/mailslurp/mailslurp-client-go"
)

type InboxController struct {
	Client mailslurp.APIClient
}

func (c *InboxController) CreateInbox(ctx context.Context, name string) error {
	opts := &mailslurp.CreateInboxOpts{
		Name: optional.NewString(name),
	}
	inbox, response, err := c.Client.InboxControllerApi.CreateInbox(ctx, opts)
	if err != nil {
		return err
	}
	if response.StatusCode != http.StatusCreated {
		return fmt.Errorf("failed to create inbox. Status code: %d", response.StatusCode)
	}
	fmt.Printf("Inbox with ID: %s and address: %s created successfully.\n", inbox.Id, inbox.EmailAddress)
	return nil
}
