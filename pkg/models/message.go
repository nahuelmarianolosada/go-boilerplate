package models

import "time"

type Message struct {
	// TODO: Implement Message model

	ID    int64 `json:"id"`
	Sender    int64 `json:"sender" validate:"required"`
	Recipient int64 `json:"recipient" validate:"required"`
	Content Content `json:"content"`	
	LastUpdated time.Time `json:"last_updated"`	
}

type Content  struct {
	Type string `json:"type" validate:"required"`//,oneof='text video image'
	Text string `json:"text"`
} 
