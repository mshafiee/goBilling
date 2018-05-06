/*
 * goBilling
 *
 * goBilling is an open-source billing and payments platform
 *
 * API version: 1.0.0
 *
 */

package model

type CreditJson struct {

	CreditAmount float32 `json:"creditAmount"`

	Currency string `json:"currency,omitempty"`

	InvoiceId string `json:"invoiceId,omitempty"`

	InvoiceNumber string `json:"invoiceNumber,omitempty"`

	EffectiveDate string `json:"effectiveDate,omitempty"`

	AccountId string `json:"accountId"`

	Description string `json:"description,omitempty"`

	AuditLogs []AuditLogJson `json:"auditLogs,omitempty"`
}
