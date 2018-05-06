/*
 * goBilling
 *
 * goBilling is an open-source billing and payments platform
 *
 * API version: 1.0.0
 *
 */

package model

type TagDefinitionJson struct {

	Id string `json:"id,omitempty"`

	IsControlTag bool `json:"isControlTag,omitempty"`

	Name string `json:"name"`

	Description string `json:"description"`

	ApplicableObjectTypes []string `json:"applicableObjectTypes,omitempty"`

	AuditLogs []AuditLogJson `json:"auditLogs,omitempty"`
}
