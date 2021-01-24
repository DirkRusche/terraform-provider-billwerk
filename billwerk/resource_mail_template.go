package billwerk

import (
	"billwerk-go"
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceMailTemplate() *schema.Resource {
	return &schema.Resource{
		ReadContext:   readMailTemplate,
		CreateContext: createMailTemplate,
		UpdateContext: updateMailTemplate,
		DeleteContext: deleteMailTemplate,
		Importer: &schema.ResourceImporter{
			StateContext: importMailTemplate,
		},

		Schema: map[string]*schema.Schema{
			"internal_name": {
				Type: schema.TypeString,

				Required: true,
			},
			"external_id": {
				Type: schema.TypeString,

				Required: true,
			},
			"event_type": {
				Type: schema.TypeString,

				Required: true,
				ForceNew: true,
			},
			"language": {
				Type: schema.TypeSet,
				Required: true,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"language": {
							Type: schema.TypeString,
							Required: true,
						},
						"subject": {
							Type: schema.TypeString,
							Required: true,
						},
						"text": {
							Type:     schema.TypeString,
							Required: true,
						},
					},
				},
			},
		},
	}
}

func createMailTemplate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*billwerk.Client)

	newMailTemplate := billwerk.MailTemplateNew{
		EmailType: "CustomerEmailNotification",
		Attachments: []string{},
		EventType: d.Get("event_type").(string),
		InternalName: d.Get("internal_name").(string),
		ExternalId: d.Get("external_id").(string),
		Subject: map[string]string{},
		HtmlText: map[string]string{},
	}

	languageSet := d.Get("language")
	//languageSet := d.Get("language").(*schema.Set).List()

	for _, languageEntry := range languageSet.(*schema.Set).List() {
		entry := languageEntry.(map[string]interface{})

		language := entry["language"].(string)
		newMailTemplate.Subject[language] = entry["subject"].(string)
		newMailTemplate.HtmlText[language] = entry["text"].(string)
	}


	mailTemplate, err := c.CreateMailTemplate(newMailTemplate)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(mailTemplate.ID)

	return readMailTemplate(ctx, d, m)
}

func updateMailTemplate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*billwerk.Client)
	id := d.Id()

	update := billwerk.MailTemplateUpdate{
		ID: id,
		EmailType: "CustomerEmailNotification",
		Subject: map[string]string{},
		HtmlText: map[string]string{},
		Attachments:  []string{},
		EventType:    d.Get("event_type").(string),
		InternalName: d.Get("internal_name").(string),
		ExternalId:   d.Get("external_id").(string),
	}

	languageSet := d.Get("language")
	//languageSet := d.Get("language").(*schema.Set).List()

	for _, languageEntry := range languageSet.(*schema.Set).List() {
		entry := languageEntry.(map[string]interface{})

		language := entry["language"].(string)
		update.Subject[language] = entry["subject"].(string)
		update.HtmlText[language] = entry["text"].(string)
	}

	_, err := c.UpdateMailTemplate(update)

	if err != nil {
		return diag.FromErr(err)
	}

	return readMailTemplate(ctx, d, m)
}

func deleteMailTemplate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*billwerk.Client)
	id := d.Id()

	err := c.DeleteMailTemplate(id)

	if err != nil {
		return diag.FromErr(err)
	}
	return nil
}

func readMailTemplate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*billwerk.Client)
	id := d.Id()

	mailTemplate, err := c.GetMailTemplate(id)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	if err != nil {
		return diag.FromErr(err)
	}

	setData(d, mailTemplate)

	return diags
}

func setData(d *schema.ResourceData, mailTemplate *billwerk.MailTemplateResponse) {
	d.SetId(mailTemplate.ID)
	d.Set("internal_name", mailTemplate.InternalName)
	d.Set("external_id", mailTemplate.ExternalId)
	d.Set("event_type", mailTemplate.EventType)

	l2 := make([]map[string]interface{}, 0, 1)
	for key, value := range mailTemplate.Subject {
		entry := map[string]interface{}{}

		entry["language"] = key
		entry["subject"] = value
		entry["text"] = mailTemplate.HtmlText[key]

		l2 = append(l2, entry)
	}
	d.Set("language", l2)
}

func importMailTemplate(ctx context.Context, d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	c := m.(*billwerk.Client)
	id := d.Id()

	mailTemplate, err := c.GetMailTemplate(id)
	if err != nil {
		return nil, fmt.Errorf("Invalid id: %v", err)
	}

	setData(d, mailTemplate)

	return []*schema.ResourceData{d}, nil
}
