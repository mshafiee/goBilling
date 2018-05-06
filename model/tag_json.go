/*
 * goBilling
 *
 * goBilling is an open-source billing and payments platform
 *
 * API version: 1.0.0
 *
 */

package model

type TagJson struct {

	TagId string `json:"tagId,omitempty"`

	ObjectType string `json:"objectType,omitempty"`

	ObjectId string `json:"objectId,omitempty"`

	TagDefinitionId string `json:"tagDefinitionId,omitempty"`

	TagDefinitionName string `json:"tagDefinitionName,omitempty"`

	AuditLogs []AuditLogJson `json:"auditLogs,omitempty"`
}
