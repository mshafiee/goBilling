/*
 * goBilling
 *
 * goBilling is an open-source billing and payments platform
 *
 * API version: 1.0.0
 *
 */

package model

type AccountJson struct {

	AccountId string `json:"accountId,omitempty"`

	Name string `json:"name,omitempty"`

	FirstNameLength int32 `json:"firstNameLength,omitempty"`

	ExternalKey string `json:"externalKey,omitempty"`

	Email string `json:"email,omitempty"`

	BillCycleDayLocal int32 `json:"billCycleDayLocal,omitempty"`

	Currency string `json:"currency,omitempty"`

	ParentAccountId string `json:"parentAccountId,omitempty"`

	IsPaymentDelegatedToParent bool `json:"isPaymentDelegatedToParent,omitempty"`

	PaymentMethodId string `json:"paymentMethodId,omitempty"`

	TimeZone string `json:"timeZone,omitempty"`

	Address1 string `json:"address1,omitempty"`

	Address2 string `json:"address2,omitempty"`

	PostalCode string `json:"postalCode,omitempty"`

	Company string `json:"company,omitempty"`

	City string `json:"city,omitempty"`

	State string `json:"state,omitempty"`

	Country string `json:"country,omitempty"`

	Locale string `json:"locale,omitempty"`

	Phone string `json:"phone,omitempty"`

	Notes string `json:"notes,omitempty"`

	IsMigrated bool `json:"isMigrated,omitempty"`

	IsNotifiedForInvoices bool `json:"isNotifiedForInvoices,omitempty"`

	AccountBalance float32 `json:"accountBalance,omitempty"`

	AccountCBA float32 `json:"accountCBA,omitempty"`

	AuditLogs []AuditLogJson `json:"auditLogs,omitempty"`
}
