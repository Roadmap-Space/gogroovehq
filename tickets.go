package gogroovehq

import (
	"fmt"
	"time"
)

type tickets struct {
	EndpointURL string
}

// Ticket represents a GrooveHQ ticket
type Ticket struct {
	Ticket struct {
		AssignedGroup  interface{} `json:"assigned_group"`
		CreatedAt      time.Time   `json:"created_at"`
		Number         int         `json:"number"`
		Priority       interface{} `json:"priority"`
		ResolutionTime interface{} `json:"resolution_time"`
		Title          interface{} `json:"title"`
		UpdatedAt      time.Time   `json:"updated_at"`
		Href           string      `json:"href"`
		ClosedBy       string      `json:"closed_by"`
		Tags           []string    `json:"tags"`
		MessageCount   int         `json:"message_count"`
		Summary        string      `json:"summary"`
		Links          struct {
			Assignee struct {
				Href string `json:"href"`
			} `json:"assignee"`
			Customer struct {
				Href string `json:"href"`
			} `json:"customer"`
			State struct {
				Href string `json:"href"`
			} `json:"state"`
			Messages struct {
				Href string `json:"href"`
			} `json:"messages"`
		} `json:"links"`
	} `json:"ticket"`
}

// Get returns matching ticket by id
func (t *tickets) Get(ticketID string) (*Ticket, error) {
	path := fmt.Sprintf("%s/%s", t.EndpointURL, ticketID)

	var ticket Ticket
	if err := apiClient.get(path, &ticket); err != nil {
		return nil, err
	}
	return &ticket, nil
}
