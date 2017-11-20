package gogroovehq

import "fmt"

type customers struct {
	EndPointURL string
}

// Customer represents a GrooveHQ customer
type Customer struct {
	Customer struct {
		Href  string `json:"href"`
		Links struct {
			Tickets struct {
				Href string `json:"href"`
			} `json:"tickets"`
		} `json:"links"`
		Email            string      `json:"email"`
		Name             string      `json:"name"`
		About            interface{} `json:"about"`
		TwitterUsername  interface{} `json:"twitter_username"`
		Title            interface{} `json:"title"`
		CompanyName      interface{} `json:"company_name"`
		PhoneNumber      interface{} `json:"phone_number"`
		Location         interface{} `json:"location"`
		WebsiteURL       interface{} `json:"website_url"`
		LinkedinUsername interface{} `json:"linkedin_username"`
	} `json:"customer"`
}

// Get returns a customer by its id
func (c *customers) Get(customerID string) (*Customer, error) {
	path := fmt.Sprintf("%s/%s", c.EndPointURL, customerID)

	var customer Customer
	if err := apiClient.get(path, &customer); err != nil {
		return nil, err
	}
	return &customer, nil
}
