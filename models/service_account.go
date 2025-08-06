package models

type ServiceAccount struct {
	ID                      uint `gorm:"primaryKey"`
	Type                    string
	ProjectID               string
	PrivateKeyID            string
	PrivateKey              string // encrypted
	ClientEmail             string
	ClientID                string
	AuthURI                 string
	TokenURI                string
	AuthProviderX509CertURL string
	ClientX509CertURL       string
	UniverseDomain          string
	EncryptedAt             string
}
