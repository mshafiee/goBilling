/*
 * goBilling
 *
 * goBilling is an open-source billing and payments platform
 *
 * API version: 1.0.0
 *
 */

package model

type AccountEmailJson struct {

	AccountId string `json:"accountId,omitempty"`

	Email string `json:"email"`

	AuditLogs []AuditLogJson `json:"auditLogs,omitempty"`
}
