package transloadit

import (
	"encoding/json"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"reflect"
)

func AreEqualJSON(s1, s2 string) (bool, error) {
	var o1 interface{}
	var o2 interface{}

	var err error
	err = json.Unmarshal([]byte(s1), &o1)
	if err != nil {
		return false, fmt.Errorf("Error mashalling string 1 :: %s", err.Error())
	}
	err = json.Unmarshal([]byte(s2), &o2)
	if err != nil {
		return false, fmt.Errorf("Error mashalling string 2 :: %s", err.Error())
	}

	return reflect.DeepEqual(o1, o2), nil
}

type BatchSet map[string]interface{}

func saveToState(d *schema.ResourceData, m BatchSet) error {
	for key, value := range m {
		err := d.Set(key, value)
		if err != nil {
			return fmt.Errorf("Impossible to save %s = %s in state", key, value)
		}
	}
	return nil
}
