package transloadit

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"os"
	"testing"
)

var testAccProviders map[string]terraform.ResourceProvider

func init() {
	testAccProviders = map[string]terraform.ResourceProvider{
		"transloadit": Provider().(*schema.Provider),
	}
}

func testAccPreCheck(t *testing.T) {
	testVariable("TRANSLOADIT_AUTH_KEY", t)
	testVariable("TRANSLOADIT_AUTH_SECRET", t)
}

func testVariable(envVar string, t *testing.T) {
	v := os.Getenv(envVar)
	if v == "" {
		t.Fatalf("%s environment variable must be set", envVar)
	}
}
