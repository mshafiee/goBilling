/*
 * goBilling
 *
 * goBilling is an open-source billing and payments platform
 *
 * API version: 1.0.0
 *
 */

package model

type ComboPaymentTransactionJson struct {

	Account *AccountJson `json:"account,omitempty"`

	PaymentMethod *PaymentMethodJson `json:"paymentMethod,omitempty"`

	Transaction *PaymentTransactionJson `json:"transaction,omitempty"`

	PaymentMethodPluginProperties *IterablePluginPropertyJson `json:"paymentMethodPluginProperties,omitempty"`

	TransactionPluginProperties *IterablePluginPropertyJson `json:"transactionPluginProperties,omitempty"`

	AuditLogs []AuditLogJson `json:"auditLogs,omitempty"`

	TransactionType string `json:"transactionType,omitempty"`
}
