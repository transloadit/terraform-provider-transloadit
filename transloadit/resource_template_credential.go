package transloadit

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	transloadit "github.com/transloadit/go-sdk"
)

func resourceTemplateCredential() *schema.Resource {
	return &schema.Resource{
		Create: resourceTemplateCredentialCreate,
		Read:   resourceTemplateCredentialRead,
		Delete: resourceTemplateCredentialDelete,
		Update: resourceTemplateCredentialUpdate,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    false,
				Description: "Template credential name",
			},
			"type": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Template credential type",
			},
			"content": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    false,
				Description: "Template credential content in JSON",
				DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					bool, _ := AreEqualJSON(old, new)
					return bool
				},
			},
			"created": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"modified": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceTemplateCredentialCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*transloadit.Client)
	templateCredentialContent, err := jsonStringToTemplateCredentialContent(d.Get("content").(string))
	if err != nil {
		return err
	}
	payload := transloadit.TemplateCredential{
		Name:    d.Get("name").(string),
		Type:    d.Get("type").(string),
		Content: *templateCredentialContent,
	}
	templateId, err := client.CreateTemplateCredential(context.Background(), payload)
	if err != nil {
		return err
	}
	d.SetId(templateId)
	return resourceTemplateCredentialRead(d, meta)
}

func resourceTemplateCredentialRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*transloadit.Client)
	templateCredential, err := client.GetTemplateCredential(context.Background(), d.Id())
	if err != nil {
		return err
	}
	toSave := BatchSet{
		"name":     templateCredential.Name,
		"type":     templateCredential.Type,
		"created":  templateCredential.Created,
		"modified": templateCredential.Modified,
	}
	result, err := json.Marshal(templateCredential.Content)
	if err != nil {
		return fmt.Errorf("Content field can't be marshalled to json +%v", templateCredential.Content)
	}
	toSave["content"] = string(result)
	return saveToState(d, toSave)
}

func resourceTemplateCredentialDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*transloadit.Client)
	err := client.DeleteTemplateCredential(context.Background(), d.Id())
	return err
}

func resourceTemplateCredentialUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*transloadit.Client)
	templateCredentialContent, err := jsonStringToTemplateCredentialContent(d.Get("content").(string))
	if err != nil {
		return err
	}
	payload := transloadit.TemplateCredential{
		Name:    d.Get("name").(string),
		Type:    d.Get("type").(string),
		Content: *templateCredentialContent,
	}
	err = client.UpdateTemplateCredential(context.Background(), d.Id(), payload)
	return err
}

func jsonStringToTemplateCredentialContent(src string) (*map[string]interface{}, error) {
	var content map[string]interface{}
	err := json.Unmarshal([]byte(src), &content)
	if err != nil {
		return nil, err
	}

	return &content, nil
}
