package billwerk

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
)

func (c *Client) GetMailTemplate(templateId string) (*MailTemplateResponse, error) {
	req, err := http.NewRequest(
		"GET",
		fmt.Sprintf("%s/api/v1/emailNotificationsTemplates?id=%s", c.HostURL, templateId),
		nil,
	)
	if err != nil {
		return nil, err
	}

	body, _, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	template := MailTemplateResponse{}
	err = json.Unmarshal(body, &template)
	if err != nil {
		return nil, err
	}

	return &template, nil
}

func (c *Client) CreateMailTemplate(mailTemplate MailTemplateNew) (*MailTemplateResponse, error) {
	rb, err := json.Marshal(mailTemplate)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(
		"POST",
		fmt.Sprintf("%s/api/v1/emailNotificationsTemplates", c.HostURL),
		strings.NewReader(string(rb)),
	)
	if err != nil {
		return nil, err
	}

	body, _, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	mailTemplateResponse := MailTemplateResponse{}
	err = json.Unmarshal(body, &mailTemplateResponse)
	if err != nil {
		return nil, err
	}

	return &mailTemplateResponse, nil
}

func (c *Client) UpdateMailTemplate(mailTemplate MailTemplateUpdate) (*MailTemplateResponse, error) {
	rb, err := json.Marshal(mailTemplate)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(
		"PUT",
		fmt.Sprintf("%s/api/v1/emailNotificationsTemplates?id=%s", c.HostURL, mailTemplate.ID),
		strings.NewReader(string(rb)),
	)
	if err != nil {
		return nil, err
	}

	body, status, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}
	if status != http.StatusOK {
		return nil, errors.New(string(body))
	}

	mailTemplateResponse := MailTemplateResponse{}
	err = json.Unmarshal(body, &mailTemplateResponse)
	if err != nil {
		return nil, err
	}

	return &mailTemplateResponse, nil
}

func (c *Client) DeleteMailTemplate(templateId string) error {
	req, err := http.NewRequest(
		"DELETE",
		fmt.Sprintf("%s/api/v1/emailNotificationsTemplates?id=%s", c.HostURL, templateId),
		nil,
	)

	if err != nil {
		return err
	}

	body, status, err := c.doRequest(req)
	if err != nil {
		return err
	}

	if status != 204 {
		return errors.New(string(body))
	}

	return nil
}
