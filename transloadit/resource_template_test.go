package transloadit

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"testing"
)

func TestAccTemplate_basic(t *testing.T) {

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: nil,
		Steps: []resource.TestStep{
			{
				Config: testAccTemplate_basic_1,
				Check:  nil,
			},
		},
	})
}

const testAccTemplate_basic_1 = `
resource "transloadit_template" "test" {
	name = "templatebasic1"
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
