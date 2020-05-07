package transloadit

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"math/rand"
	"testing"
)

func TestAccTemplate_basic(t *testing.T) {

	resource_name := fmt.Sprintf("templatebasic%d", rand.Int()%1000)
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: nil,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(testAccTemplate_basic_1, resource_name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckExistingResource("transloadit_template.test"),
					resource.TestCheckResourceAttr("transloadit_template.test", "name", resource_name),
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
  }
}
EOT
}
`
