package billwerk

// Mail Template
type MailTemplateResponse struct {
	ID string `json:"Id"`
	Created string `json:"Created"`
	InternalName string `json:"InternalName"`
	ExternalId string `json:"ExternalId"`
	EventType string `json:"EventType"`
	Attachments []string `json:"Attachements"`
	UseIndividualEmailSettings bool `json:"UseIndividualEmailSettings"`
	EmailType string `json:"EmailType"`
	Subject map[string]string `json:"Subject"`
	HtmlText map[string]string `json:"HtmlText"`
}

type MailTemplateNew struct {
	EmailType string `json:"EmailType"`
	Subject map[string]string `json:"Subject"`
	HtmlText map[string]string `json:"HtmlText"`
	Attachments []string `json:"Attachements"`
	EventType string `json:"EventType"`
	InternalName string `json:"InternalName"`
	ExternalId string `json:"ExternalId"`
}

type MailTemplateUpdate struct {
	ID string `json:"Id"`
	EmailType string `json:"EmailType"`
	Subject map[string]string `json:"Subject"`
	HtmlText map[string]string `json:"HtmlText"`
	Attachments []string `json:"Attachements"`
	EventType string `json:"EventType"`
	InternalName string `json:"InternalName"`
	ExternalId string `json:"ExternalId"`
}
