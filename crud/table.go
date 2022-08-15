package crud

import (
	"time"

	"gorm.io/gorm"
)

type AzureClient struct {
	gorm.Model
	SubscriptionName string
	SubscriptionID   string
	TenantID         string
	ClientID         string
	ClientSecret     string
	CreatedAt        time.Time
	UpdatedAt        time.Time
	DeletedAt        time.Time
}

func (t *AzureClient) TableName() string {
	return "azure_client_secret"
}
