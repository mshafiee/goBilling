/*
 * goBilling
 *
 * goBilling is an open-source billing and payments platform
 *
 * API version: 1.0.0
 *
 */

package model

type HostedPaymentPageFieldsJson struct {

	AuditLogs []AuditLogJson `json:"auditLogs,omitempty"`

	CustomFields []PluginPropertyJson `json:"customFields,omitempty"`
}
