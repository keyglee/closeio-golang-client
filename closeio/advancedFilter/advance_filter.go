package advancedfilter

import customfield "github.com/keyglee/closeio-golang-client/closeio/customField"

// Condition represents the condition object within a field_condition
type Condition struct {
	Mode  string `json:"mode"`
	Type  string `json:"type"`
	Value string `json:"value"`
}

// Field represents the field object within a field_condition
type Field struct {
	FieldName   string `json:"field_name"`
	CustomField customfield.Custom_Field
	ObjectType  string `json:"object_type"`
	Type        string `json:"type"`
}

// FieldCondition represents a field condition with its field and condition details
type FieldCondition struct {
	Negate    bool      `json:"negate"`
	Type      string    `json:"type"`
	Field     Field     `json:"field"`
	Condition Condition `json:"condition"`
}

// Query represents a single query within the queries array
type Query struct {
	Negate            bool            `json:"negate"`
	ObjectType        string          `json:"object_type,omitempty"`
	Type              string          `json:"type"`
	RelatedObjectType string          `json:"related_object_type,omitempty"`
	RelatedQuery      *Query          `json:"related_query,omitempty"`
	ThisObjectType    string          `json:"this_object_type,omitempty"`
	FieldCondition    *FieldCondition `json:"condition,omitempty"` // alias for field_condition for clarity
	Queries           []Query         `json:"queries,omitempty"`
}

// Root represents the entire JSON structure
type Root struct {
	IncludeCounts bool     `json:"include_counts"`
	Limit         *int     `json:"_limit,omitempty"`
	Query         Query    `json:"query"`
	ResultsLimit  *int     `json:"results_limit,omitempty"`
	Sort          []string `json:"sort"`
}

func (t *Query) SetFieldCondition(field_condition *FieldCondition) {
	t.FieldCondition = field_condition
}
