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

// Payment transaction
type PaymentTransactionJson struct {

	TransactionId string `json:"transactionId,omitempty"`

	TransactionExternalKey string `json:"transactionExternalKey,omitempty"`

	// Associated payment id, required when notifying state transitions
	PaymentId string `json:"paymentId,omitempty"`

	PaymentExternalKey string `json:"paymentExternalKey,omitempty"`

	TransactionType string `json:"transactionType,omitempty"`

	// Transaction amount, required except for void operations
	Amount float32 `json:"amount,omitempty"`

	// Amount currency (account currency unless specified)
	Currency string `json:"currency,omitempty"`

	EffectiveDate time.Time `json:"effectiveDate,omitempty"`

	ProcessedAmount float32 `json:"processedAmount,omitempty"`

	ProcessedCurrency string `json:"processedCurrency,omitempty"`

	// Transaction status, required for state change notifications
	Status string `json:"status,omitempty"`

	GatewayErrorCode string `json:"gatewayErrorCode,omitempty"`

	GatewayErrorMsg string `json:"gatewayErrorMsg,omitempty"`

	FirstPaymentReferenceId string `json:"firstPaymentReferenceId,omitempty"`

	SecondPaymentReferenceId string `json:"secondPaymentReferenceId,omitempty"`

	Properties []PluginPropertyJson `json:"properties,omitempty"`

	AuditLogs []AuditLogJson `json:"auditLogs,omitempty"`
}
