/*
 * goBilling
 *
 * goBilling is an open-source billing and payments platform
 *
 * API version: 1.0.0
 *
 */

package model

type HostedPaymentPageFormDescriptorJson struct {

	KbAccountId string `json:"kbAccountId,omitempty"`

	FormMethod string `json:"formMethod,omitempty"`

	FormUrl string `json:"formUrl,omitempty"`

	FormFields map[string]interface{} `json:"formFields,omitempty"`

	Properties map[string]interface{} `json:"properties,omitempty"`

	AuditLogs []AuditLogJson `json:"auditLogs,omitempty"`
}
