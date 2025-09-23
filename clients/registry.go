package clients

import (
	fieldConfig "field-service/clients/config"
	userConfig "field-service/clients/user"
	"field-service/config"
)

type RegistryClient struct{}

type IRegistryClient interface {
	GetUser() userConfig.IUserClient
}

func NewRegistryClient() IRegistryClient {
	return &RegistryClient{}
}

func (rc *RegistryClient) GetUser() userConfig.IUserClient {
	return userConfig.NewUserClient(
		fieldConfig.NewClientConfig(
			fieldConfig.WithBaseURL(config.Config.InternalService.User.Host),
			fieldConfig.WithSignatureKey(config.Config.SignatureKey),
		),
	)
}
