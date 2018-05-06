/*
 * goBilling
 *
 * goBilling is an open-source billing and payments platform
 *
 * API version: 1.0.0
 *
 */

package model

type PaymentMethodJson struct {

	PaymentMethodId string `json:"paymentMethodId,omitempty"`

	ExternalKey string `json:"externalKey,omitempty"`

	AccountId string `json:"accountId,omitempty"`

	IsDefault bool `json:"isDefault,omitempty"`

	PluginName string `json:"pluginName,omitempty"`

	PluginInfo *PaymentMethodPluginDetailJson `json:"pluginInfo,omitempty"`

	AuditLogs []AuditLogJson `json:"auditLogs,omitempty"`
}
