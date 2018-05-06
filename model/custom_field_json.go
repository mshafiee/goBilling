/*
 * goBilling
 *
 * goBilling is an open-source billing and payments platform
 *
 * API version: 1.0.0
 *
 */

package model

type CustomFieldJson struct {

	CustomFieldId string `json:"customFieldId,omitempty"`

	ObjectId string `json:"objectId,omitempty"`

	ObjectType string `json:"objectType,omitempty"`

	Name string `json:"name"`

	Value string `json:"value"`

	AuditLogs []AuditLogJson `json:"auditLogs,omitempty"`
}
