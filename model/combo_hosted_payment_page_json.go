/*
 * goBilling
 *
 * goBilling is an open-source billing and payments platform
 *
 * API version: 1.0.0
 *
 */

package model

type ComboHostedPaymentPageJson struct {

	Account *AccountJson `json:"account,omitempty"`

	PaymentMethod *PaymentMethodJson `json:"paymentMethod,omitempty"`

	PaymentMethodPluginProperties *IterablePluginPropertyJson `json:"paymentMethodPluginProperties,omitempty"`

	AuditLogs []AuditLogJson `json:"auditLogs,omitempty"`

	HostedPaymentPageFieldsJson *HostedPaymentPageFieldsJson `json:"hostedPaymentPageFieldsJson,omitempty"`
}
