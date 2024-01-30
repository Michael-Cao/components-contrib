package secretstores

const DefaultSecretRefKeyName = "_value"

type Metadata struct {
	Properties map[string]string `json:"properties,omitempty"`
}
