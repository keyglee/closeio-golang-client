package customfield

import (
	"strings"
)

type Custom_Field struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Value any    `json:"value"`
}

func (t Custom_Field) GetIDValue() (string, any) {
	return "custom." + t.Id, t.Value
}

func ExtractCustomFields(data map[string]interface{}) []Custom_Field {
	var customFields []Custom_Field

	for key, value := range data {
		// Skip non-custom fields
		if !strings.HasPrefix(key, "custom.") {
			continue
		}
		split_key := strings.Split(key, ".")
		key_name := split_key[len(split_key)-1]

		customFields = append(customFields, Custom_Field{Id: key_name, Value: value})
	}

	return customFields
}

func add_custom_fields_to_map(inInterface map[string]interface{}, custom_fields []Custom_Field) map[string]interface{} {

	for _, value := range custom_fields {
		key_id, key_value := value.GetIDValue()
		inInterface[key_id] = key_value
	}

	return inInterface
}
