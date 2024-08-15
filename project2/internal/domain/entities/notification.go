package entities

import "time"

type Notification struct {
	NotificationId string    `json:"notification_id"`
	UserId         string    `json:"user_id"`
	Message        string    `json:"message"`
	TimeToSend     time.Time `json:"time_to_send"`
}
