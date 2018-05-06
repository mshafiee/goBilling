/*
 * goBilling
 *
 * goBilling is an open-source billing and payments platform
 *
 * API version: 1.0.0
 *
 */

package model

type BundleJson struct {

	AccountId string `json:"accountId"`

	BundleId string `json:"bundleId,omitempty"`

	ExternalKey string `json:"externalKey,omitempty"`

	Subscriptions []SubscriptionJson `json:"subscriptions,omitempty"`

	Timeline *BundleTimelineJson `json:"timeline,omitempty"`

	AuditLogs []AuditLogJson `json:"auditLogs,omitempty"`
}
