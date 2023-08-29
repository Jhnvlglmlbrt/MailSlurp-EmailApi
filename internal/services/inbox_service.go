package services

import (
	"context"

	"MailSlurpGo/internal/controllers"
)

type InboxService struct {
	Controller *controllers.InboxController
}

func (s *InboxService) CreateInbox(ctx context.Context, name string) error {
	return s.Controller.CreateInbox(ctx, name)
}
