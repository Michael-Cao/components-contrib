package secretstores

type GetSecretRequest struct {
	Name     string            `json:"name"`
	Metadata map[string]string `json:"metadata"`
}
