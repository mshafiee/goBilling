/*
 * goBilling
 *
 * goBilling is an open-source billing and payments platform
 *
 * API version: 1.0.0
 *
 */

package model

type BundleTimelineJson struct {

	AccountId string `json:"accountId,omitempty"`

	BundleId string `json:"bundleId,omitempty"`

	ExternalKey string `json:"externalKey,omitempty"`

	Events []EventSubscriptionJson `json:"events,omitempty"`

	AuditLogs []AuditLogJson `json:"auditLogs,omitempty"`
}
