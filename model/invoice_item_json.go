/*
 * goBilling
 *
 * goBilling is an open-source billing and payments platform
 *
 * API version: 1.0.0
 *
 */

package model

type InvoiceItemJson struct {

	InvoiceItemId string `json:"invoiceItemId"`

	InvoiceId string `json:"invoiceId,omitempty"`

	LinkedInvoiceItemId string `json:"linkedInvoiceItemId,omitempty"`

	AccountId string `json:"accountId"`

	ChildAccountId string `json:"childAccountId,omitempty"`

	BundleId string `json:"bundleId,omitempty"`

	SubscriptionId string `json:"subscriptionId,omitempty"`

	PlanName string `json:"planName,omitempty"`

	PhaseName string `json:"phaseName,omitempty"`

	UsageName string `json:"usageName,omitempty"`

	ItemType string `json:"itemType,omitempty"`

	Description string `json:"description,omitempty"`

	StartDate string `json:"startDate,omitempty"`

	EndDate string `json:"endDate,omitempty"`

	Amount float32 `json:"amount,omitempty"`

	Currency string `json:"currency,omitempty"`

	ChildItems []InvoiceItemJson `json:"childItems,omitempty"`

	AuditLogs []AuditLogJson `json:"auditLogs,omitempty"`
}
