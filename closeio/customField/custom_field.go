package customfield

import (
	"strings"
)

type CustomField struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Value any    `json:"value"`
}

func (t CustomField) GetIDValue() (string, any) {
	return "custom." + t.ID, t.Value
}

func ExtractCustomFields(data map[string]interface{}) []CustomField {
	var customFields []CustomField

	for key, value := range data {
		// Skip non-custom fields
		if !strings.HasPrefix(key, "custom.") {
			continue
		}
		split_key := strings.Split(key, ".")
		key_name := split_key[len(split_key)-1]

		customFields = append(customFields, CustomField{ID: key_name, Value: value})
	}

	return customFields
}

func add_custom_fields_to_map(inInterface map[string]interface{}, custom_fields []CustomField) map[string]interface{} {

	for _, value := range custom_fields {
		key_id, key_value := value.GetIDValue()
		inInterface[key_id] = key_value
	}

	return inInterface
}
