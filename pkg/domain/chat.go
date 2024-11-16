package domain

import "time"

const (
	ChatMessageStatusSent     = "sent"
	ChatMessageStatusReceived = "received"
)

type Chat struct {
	ID        int64 `json:"id"`
	DoctorID  int64 `json:"doctor_id"`
	PatientID int64 `json:"patient_id"`
}

type ChatMessage struct {
	ID        int64     `json:"id"`
	ChatID    int64     `json:"chat_id"`
	Message   string    `json:"message"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
