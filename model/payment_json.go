/*
 * goBilling
 *
 * goBilling is an open-source billing and payments platform
 *
 * API version: 1.0.0
 *
 */

package model

type PaymentJson struct {

	AccountId string `json:"accountId,omitempty"`

	PaymentId string `json:"paymentId,omitempty"`

	PaymentNumber string `json:"paymentNumber,omitempty"`

	PaymentExternalKey string `json:"paymentExternalKey,omitempty"`

	AuthAmount float32 `json:"authAmount,omitempty"`

	CapturedAmount float32 `json:"capturedAmount,omitempty"`

	PurchasedAmount float32 `json:"purchasedAmount,omitempty"`

	RefundedAmount float32 `json:"refundedAmount,omitempty"`

	CreditedAmount float32 `json:"creditedAmount,omitempty"`

	Currency string `json:"currency,omitempty"`

	PaymentMethodId string `json:"paymentMethodId,omitempty"`

	Transactions []PaymentTransactionJson `json:"transactions,omitempty"`

	PaymentAttempts []PaymentAttemptJson `json:"paymentAttempts,omitempty"`

	AuditLogs []AuditLogJson `json:"auditLogs,omitempty"`
}
