package secretstores

type SecretStore interface {
	Init(metadata Metadata) error
	GetSecret(req GetSecretRequest) (GetSecretResponse, error)
}
