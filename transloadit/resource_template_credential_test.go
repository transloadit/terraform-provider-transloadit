package transloadit

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccTemplateCredential_basic(t *testing.T) {

	randInt := rand.Int() % 1000
	template1_name := fmt.Sprintf("templatebasic%d", randInt)
	template2_name := fmt.Sprintf("templatebasic%d", randInt+1)
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: nil,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(testAccTemplateCredential_basic_1, template1_name, template2_name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckExistingResource("transloadit_template_credential.s3"),
					testAccCheckExistingResource("transloadit_template_credential.backblaze"),
					resource.TestCheckResourceAttr("transloadit_template_credential.s3", "name", template1_name),
					resource.TestCheckResourceAttr("transloadit_template_credential.s3", "type", "s3"),
					resource.TestCheckResourceAttr("transloadit_template_credential.s3", "content", "{\"bucket\":\"mybucket\",\"bucket_region\":\"us-east-1\",\"key\":\"mykey\",\"secret\":\"mysecret\"}"),
					resource.TestCheckResourceAttr("transloadit_template_credential.backblaze", "name", template2_name),
					resource.TestCheckResourceAttr("transloadit_template_credential.backblaze", "type", "backblaze"),
					resource.TestCheckResourceAttr("transloadit_template_credential.backblaze", "content", "{\"app_key\":\"mykey\",\"app_key_id\":\"mykeyid\",\"bucket\":\"mybucket\"}"),
				),
			},
		},
	})
}

const testAccTemplateCredential_basic_1 = `
resource "transloadit_template_credential" "s3" {
  name = "%s"
  type = "s3"
  content = <<EOT
  {
	"bucket" : "mybucket",
	"bucket_region" : "us-east-1",
    "key" : "mykey",
	"secret" : "mysecret"
   }
EOT
}

resource "transloadit_template_credential" "backblaze" {
  name = "%s"
  type = "backblaze"
  content = <<EOT
  {
	"bucket" : "mybucket",
	"app_key_id" : "mykeyid",
    "app_key" : "mykey"
   }
EOT
}
`
