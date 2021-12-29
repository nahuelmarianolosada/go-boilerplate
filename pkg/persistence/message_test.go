package persistence

import (
	"reflect"
	"testing"
	"time"

	"github.com/nahuelmarianolosada/go-boilerplate/pkg/models"
)

func TestCreateMessage_mustNotFail(t *testing.T) {
	now := time.Now()
	testMsg := models.Message{Sender: 0, Recipient: 0, Content: models.Content{Type: "text", Text: "test"}, LastUpdated: now}
	testMsgWanted := models.Message{ID: 1, Sender: 0, Recipient: 0, Content: models.Content{Type: "text", Text: "test"}, LastUpdated: now}
	type args struct {
		m models.Message
	}
	tests := []struct {
		name    string
		args    args
		want    *models.Message
		wantErr bool
	}{
		{"First test to check creation", args{testMsg}, &testMsgWanted, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CreateMessage(tt.args.m)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateMessage() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateMessage() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetAllMessages_mustReturnAnEmptySliceWithoutError(t *testing.T) {
	type args struct {
		recipient int
		start     int
		limit     int
	}
	tests := []struct {
		name    string
		args    args
		want    []models.Message
		wantErr bool
	}{
		{"First test to check nil error when DB is empty", args{}, []models.Message{}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetAllMessages(tt.args.recipient, tt.args.start, tt.args.limit)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAllMessages() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAllMessages() = %v, want %v", got, tt.want)
			}
		})
	}
}
func TestGetAllMessages_mustReturnResultsWithoutErrors(t *testing.T) {
	now := time.Now()
	testMsgExpected := models.Message{ID:2, Sender: 0, Recipient: 0, Content: models.Content{Type: "text", Text: "test"}, LastUpdated: now}

	type args struct {
		recipient int
		start     int
		limit     int
	}
	tests := []struct {
		name    string
		args    args
		want    []models.Message
		wantErr bool
	}{
		{"First test to check messages retribution.", args{0, 0, 10}, []models.Message{testMsgExpected}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetAllMessages(tt.args.recipient, tt.args.start, tt.args.limit)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAllMessages() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if len(got) != len(tt.want) {
				t.Errorf("GetAllMessages() = %v, want %v", got, tt.want)
			}
		})
	}
}
