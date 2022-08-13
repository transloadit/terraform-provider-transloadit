package transloadit

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	transloadit "github.com/transloadit/go-sdk"
)

func resourceTemplate() *schema.Resource {
	return &schema.Resource{
		Create: resourceTemplateCreate,
		Read:   resourceTemplateRead,
		Delete: resourceTemplateDelete,
		Update: resourceTemplateUpdate,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    false,
				Description: "name",
			},
			"require_signature_auth": &schema.Schema{
				Type:        schema.TypeBool,
				Optional:    true,
				ForceNew:    false,
				Description: "Use true to deny requests that do not include a signature",
				Default:     false,
			},
			"template": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    false,
				Description: "template in json",
				DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					bool, _ := AreEqualJSON(old, new)
					return bool
				},
			},
		},
	}
}

func resourceTemplateCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*transloadit.Client)
	templateContent, err := jsonStringToTemplate(d.Get("template").(string))
	if err != nil {
		return err
	}
	payload := transloadit.Template{
		Name:                 d.Get("name").(string),
		Content:              *templateContent,
		RequireSignatureAuth: d.Get("require_signature_auth").(bool),
	}
	templateId, err := client.CreateTemplate(context.Background(), payload)
	if err != nil {
		return err
	}
	d.SetId(templateId)
	return resourceTemplateRead(d, meta)
}

func resourceTemplateDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*transloadit.Client)
	err := client.DeleteTemplate(context.Background(), d.Id())
	return err
}

func resourceTemplateRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*transloadit.Client)
	template, err := client.GetTemplate(context.Background(), d.Id())
	if err != nil {
		return err
	}
	d.Set("name", template.Name)
	result, err := json.Marshal(template.Content)
	if err != nil {
		return fmt.Errorf("Content field can't be marshalled to json +%v", template.Content)
	}
	d.Set("template", string(result))
	d.Set("require_signature_auth", template.RequireSignatureAuth)
	return nil
}

func resourceTemplateUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*transloadit.Client)
	templateContent, err := jsonStringToTemplate(d.Get("template").(string))
	if err != nil {
		return err
	}
	payload := transloadit.Template{
		Name:                 d.Get("name").(string),
		Content:              *templateContent,
		RequireSignatureAuth: d.Get("require_signature_auth").(bool),
	}
	err = client.UpdateTemplate(context.Background(), d.Id(), payload)
	return err
}

func jsonStringToTemplate(src string) (*transloadit.TemplateContent, error) {
	var content transloadit.TemplateContent
	err := json.Unmarshal([]byte(src), &content)
	if err != nil {
		return nil, err
	}

	return &content, nil
}
