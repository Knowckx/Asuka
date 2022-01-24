package crud

import (
	"time"

	"gorm.io/gorm"
)

type AzureClientSecret struct {
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

func (t *AzureClientSecret) TableName() string {
	return "azure_client_secret"
}
