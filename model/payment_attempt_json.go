/*
 * goBilling
 *
 * goBilling is an open-source billing and payments platform
 *
 * API version: 1.0.0
 *
 */

package model

import (
	"time"
)

// Payment attempt
type PaymentAttemptJson struct {

	AccountId string `json:"accountId,omitempty"`

	PaymentMethodId string `json:"paymentMethodId,omitempty"`

	PaymentExternalKey string `json:"paymentExternalKey,omitempty"`

	TransactionId string `json:"transactionId,omitempty"`

	TransactionExternalKey string `json:"transactionExternalKey,omitempty"`

	TransactionType string `json:"transactionType,omitempty"`

	EffectiveDate time.Time `json:"effectiveDate,omitempty"`

	StateName string `json:"stateName,omitempty"`

	// Transaction amount, required except for void operations
	Amount float32 `json:"amount,omitempty"`

	// Amount currency (account currency unless specified)
	Currency string `json:"currency,omitempty"`

	PluginName string `json:"pluginName,omitempty"`

	PluginProperties []PluginPropertyJson `json:"pluginProperties,omitempty"`

	AuditLogs []AuditLogJson `json:"auditLogs,omitempty"`
}
