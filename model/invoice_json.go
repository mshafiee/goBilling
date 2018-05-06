/*
 * goBilling
 *
 * goBilling is an open-source billing and payments platform
 *
 * API version: 1.0.0
 *
 */

package model

type InvoiceJson struct {

	Amount float32 `json:"amount,omitempty"`

	Currency string `json:"currency,omitempty"`

	Status string `json:"status,omitempty"`

	CreditAdj float32 `json:"creditAdj,omitempty"`

	RefundAdj float32 `json:"refundAdj,omitempty"`

	InvoiceId string `json:"invoiceId,omitempty"`

	InvoiceDate string `json:"invoiceDate,omitempty"`

	TargetDate string `json:"targetDate,omitempty"`

	InvoiceNumber string `json:"invoiceNumber,omitempty"`

	Balance float32 `json:"balance,omitempty"`

	AccountId string `json:"accountId,omitempty"`

	BundleKeys string `json:"bundleKeys,omitempty"`

	Credits []CreditJson `json:"credits,omitempty"`

	Items []InvoiceItemJson `json:"items,omitempty"`

	IsParentInvoice bool `json:"isParentInvoice,omitempty"`

	AuditLogs []AuditLogJson `json:"auditLogs,omitempty"`
}
