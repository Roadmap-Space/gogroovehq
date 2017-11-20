package gogroovehq

type webhooks struct {
	EndPointURL string
}

// Webhook represents a Groove webhook
// Create adds a new webhook in Groove
func (w *webhooks) Create(event, url string) (bool, error) {
	var data = new(struct {
		Event string `json:"event"`
		URL   string `json:"url"`
	})

	data.Event = event
	data.URL = url

	var result interface{}
	if err := apiClient.post(w.EndPointURL, data, &result); err != nil {
		return false, err
	}
	return true, nil
}
