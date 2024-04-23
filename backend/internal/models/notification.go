package models

import (
	"time"

	"github.com/mattermost/mattermost-server/v6/model"
)

type NotificationTime struct {
	Id    string
	Type  string `yaml:"type"`
	Title string
	Time  time.Duration `yaml:"time"`
	// Date time.Time
	// Add  time.Time
	Date time.Duration `yaml:"date"`
	Add  time.Duration `yaml:"add"`
	Sub  time.Duration `yaml:"sub"`
	// Sub  NotificationSubTime
}
type NotificationSubTime struct {
	// чет похоже это не правильно
	TimeId string
	// Value  time.Time
	Value time.Duration
}

type Notification struct {
	MostId  string       `json:"userId"`
	Type    string       `json:"type"`
	Status  string       `json:"status"`
	Message string       `json:"message"`
	SI      []SelectedSI `json:"si"`
}

type SelectedSI struct {
	Id            string `json:"id"`
	Name          string `json:"name"`
	FactoryNumber string `json:"factoryNumber"`
	Person        string `json:"person"`
}

type Confirmation struct {
	UserID      string              `json:"user_id"`
	UserName    string              `json:"user_name"`
	ChannelID   string              `json:"channel_id"`
	ChannelName string              `json:"channel_name"`
	TeamID      string              `json:"team_id"`
	TeamDomain  string              `json:"team_domain"`
	PostID      string              `json:"post_id"`
	TriggerID   string              `json:"trigger_id"`
	Type        string              `json:"type"`
	DataSource  string              `json:"data_source"`
	Context     ConfirmationContext `json:"context" binding:"required"`
}
type ConfirmationContext struct {
	InstrumentIds []string `json:"instrumentIds"`
	Status        string   `json:"status"`
	Type          string   `json:"type"`
}
type FormField struct {
	Id       string `json:"name"`
	Title    string `json:"display_name"`
	Name     string `json:"placeholder"`
	Type     string `json:"type"`
	Default  string `json:"default"`
	Optional bool   `json:"optional"`

	// display_name	String	Display name of the field shown to the user in the dialog. Maximum 24 characters.
	// name	String	Name of the field element used by the integration. Maximum 300 characters. You should use unique name fields in the same dialog.
	// type	String	Set this value to bool for a checkbox element.
	// default	String	(Optional) Set a default value for this form element. true or false.
}

type CreatePostDTO struct {
	UserId      string                   `json:"userId" binding:"required"`
	Message     string                   `json:"message" binding:"required"`
	Props       []*Props                 `json:"props"`
	Actions     []*model.PostAction      `json:"actions"`
	Attachments []*model.SlackAttachment `json:"attachments"`
}
type UpdatePostDTO struct {
	PostId      string                   `json:"postId" binding:"required"`
	Message     string                   `json:"message" binding:"required"`
	Props       []*Props                 `json:"props"`
	Actions     []*model.PostAction      `json:"actions"`
	Attachments []*model.SlackAttachment `json:"attachments"`
}

type UpdatePostData struct {
	PostID      string
	Status      string
	Missing     []SelectedSI
	Instruments []SelectedSI
}

type Props struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type DialogResponse struct {
	Type       string          `json:"type"`
	CallbackID string          `json:"callback_id"`
	State      string          `json:"state"`
	UserID     string          `json:"user_id"`
	ChannelID  string          `json:"channel_id"`
	TeamID     string          `json:"team_id"`
	Submission map[string]bool `json:"submission"`
	Cancelled  bool            `json:"cancelled"`
}
