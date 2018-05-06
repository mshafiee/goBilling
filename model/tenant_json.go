/*
 * goBilling
 *
 * goBilling is an open-source billing and payments platform
 *
 * API version: 1.0.0
 *
 */

package model

type TenantJson struct {

	TenantId string `json:"tenantId,omitempty"`

	ExternalKey string `json:"externalKey,omitempty"`

	ApiKey string `json:"apiKey"`

	ApiSecret string `json:"apiSecret"`

	AuditLogs []AuditLogJson `json:"auditLogs,omitempty"`
}
