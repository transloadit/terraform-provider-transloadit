package transloadit

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	transloadit "gopkg.in/transloadit/go-sdk.v1"
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
		Name:    d.Get("name").(string),
		Content: *templateContent,
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
	return nil
}

func resourceTemplateUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*transloadit.Client)
	templateContent, err := jsonStringToTemplate(d.Get("template").(string))
	if err != nil {
		return err
	}
	payload := transloadit.Template{
		Name:    d.Get("name").(string),
		Content: *templateContent,
	}
	err = client.UpdateTemplate(context.Background(), d.Id(), payload)
	return err
}

func jsonStringToTemplate(src string) (*transloadit.TemplateContent, error) {
	var v interface{}
	err := json.Unmarshal([]byte(src), &v)
	if err != nil {
		return nil, err
	}
	content, ok := v.(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("Content must be an JSON Object (Map)")
	}
	if value, ok := content["steps"]; ok {
		steps, ok := value.(map[string]interface{})
		if !ok {
			return nil, fmt.Errorf("Steps value must be an Json Object")
		}
		return &transloadit.TemplateContent{Steps: steps}, nil
	} else {
		return nil, fmt.Errorf("Template must have a top-level key 'steps'")
	}
}
