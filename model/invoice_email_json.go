/*
 * goBilling
 *
 * goBilling is an open-source billing and payments platform
 *
 * API version: 1.0.0
 *
 */

package model

type InvoiceEmailJson struct {

	AccountId string `json:"accountId,omitempty"`

	IsNotifiedForInvoices bool `json:"isNotifiedForInvoices,omitempty"`

	AuditLogs []AuditLogJson `json:"auditLogs,omitempty"`
}
