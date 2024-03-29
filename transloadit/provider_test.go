package transloadit

import (
	"fmt"
	"math/rand"
	"os"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

var testAccProviders map[string]*schema.Provider

func init() {
	testAccProviders = map[string]*schema.Provider{
		"transloadit": Provider(),
	}
	rand.Seed(time.Now().UTC().UnixNano())
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

func testAccCheckExistingResource(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Can't find resource %s", n)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("Primary Id is not set")
		}
		return nil
	}
}
