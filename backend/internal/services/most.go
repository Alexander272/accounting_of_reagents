package services

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/Alexander272/accounting_of_reagents/backend/internal/models"
	"github.com/goccy/go-json"
)

type MostService struct {
	url       string
	channelId string
}

func NewMostService(url string, channelId string) *MostService {
	return &MostService{
		url:       url,
		channelId: channelId,
	}
}

type Most interface {
	Send(ctx context.Context, notification *models.Notification) error
}

func (s *MostService) Send(ctx context.Context, notification *models.Notification) error {
	var post models.CreatePostDTO
	apiPath := "/api/posts"

	list := []string{"```"}
	for _, n := range notification.Data {
		list = append(list, fmt.Sprintf("%s %s %s", n.Name, n.Document, n.Purity))
	}
	list = append(list, "```")

	post.Message = strings.Join(list, "\n")
	if notification.Message != "" {
		post.Message = fmt.Sprintf("#### %s\n%s", notification.Message, post.Message)
	}
	if notification.UserId == "" {
		post.ChannelId = s.channelId
	} else {
		post.UserId = notification.UserId
	}
	// post.ChannelId = notification.ChannelId
	post.Props = []*models.Props{
		{Key: "service", Value: "reagents"},
	}

	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(post); err != nil {
		return fmt.Errorf("failed to encode notification data. error: %w", err)
	}

	_, err := http.Post(s.url+apiPath, "application/json", &buf)
	if err != nil {
		return fmt.Errorf("failed to send data to bot. error: %w", err)
	}

	// if !strings.HasPrefix(resp.Status, "2") {
	// 	body, err := io.ReadAll(resp.Body)
	// 	if err != nil {
	// 		return fmt.Errorf("client: could not read response body: %w", err)
	// 	}

	// 	return fmt.Errorf("request returned an error: %s", string(body))
	// }

	return nil
}
