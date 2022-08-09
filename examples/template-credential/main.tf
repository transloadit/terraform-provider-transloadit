resource "transloadit_template_credential" "s3" {
  name    = "MY_S3_CREDENTIALS"
  type    = "s3"
  content = <<EOT
  {
	"bucket" : "mybucket",
	"bucket_region" : "us-east-1",
        "key" : "mykey",
	"secret" : "mysecret"
   }
EOT
}

resource "transloadit_template" "my-terraform-template" {
  name     = "my-terraform-template"
  template = <<EOT
{
  "steps": {
    ":original": {
      "robot": "/upload/handle"
    },
    "encoded": {
      "use": ":original",
      "robot": "/video/encode",
      "preset": "ipad-high",
      "ffmpeg_stack": "v3.3.3"
    },
    "thumbed": {
      "use": ":original",
      "robot": "/video/thumbs",
      "count": 4,
      "ffmpeg_stack": "v3.3.3"
    },
    "exported": {
      "use": [ ":original", "encoded", "thumbed"], 
      "credentials": "${transloadit_template_credential.s3.name}",
      "robot": "/s3/store"
    }
  }
}
EOT
}
