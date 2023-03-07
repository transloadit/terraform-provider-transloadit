package transloadit

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccTemplate_basic(t *testing.T) {

	randInt := rand.Int() % 1000
	template1_name := fmt.Sprintf("templatebasic%d", randInt)
	template2_name := fmt.Sprintf("templatebasic%d", randInt+1)
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: nil,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(testAccTemplate_basic_1, template1_name, template2_name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckExistingResource("transloadit_template.test"),
					testAccCheckExistingResource("transloadit_template.require_auth"),
					resource.TestCheckResourceAttr("transloadit_template.test", "name", template1_name),
					resource.TestCheckResourceAttr("transloadit_template.test", "require_signature_auth", "false"),
					resource.TestCheckResourceAttr("transloadit_template.require_auth", "name", template2_name),
					resource.TestCheckResourceAttr("transloadit_template.require_auth", "require_signature_auth", "true"),
				),
			},
		},
	})
}

const testAccTemplate_basic_1 = `
resource "transloadit_template" "test" {
  name = "%s"
  template = <<EOT
    {
  "steps": {
    ":original": {
      "robot": "/upload/handle"
    },
    "encoded": {
      "preset": "ipad-high",
      "robot": "/video/encode",
      "use": ":original",
      "ffmpeg_stack": "v3.3.3"
    },
    "exported": {
      "credentials": "YOUR_S3_CREDENTIALS",
      "robot": "/s3/store",
      "use": [
        "encoded",
        "thumbed"
      ]
    },
    "thumbed": {
      "count": 4,
      "robot": "/video/thumbs",
      "use": ":original",
      "ffmpeg_stack": "v3.3.3"
    }
  },
  "notify_url": "https://example.com/"
}
EOT
}

resource "transloadit_template" "require_auth" {
  name = "%s"
  require_signature_auth = true
  template = <<EOT
    {
  "steps": {
    ":original": {
      "robot": "/upload/handle"
    }
  }
}
EOT
}
`
