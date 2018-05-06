/*
 * goBilling
 *
 * goBilling is an open-source billing and payments platform
 *
 * API version: 1.0.0
 *
 */

package model

type EventSubscriptionJson struct {

	EventId string `json:"eventId,omitempty"`

	BillingPeriod string `json:"billingPeriod,omitempty"`

	EffectiveDate string `json:"effectiveDate,omitempty"`

	Plan string `json:"plan,omitempty"`

	Product string `json:"product,omitempty"`

	PriceList string `json:"priceList,omitempty"`

	EventType string `json:"eventType,omitempty"`

	IsBlockedBilling bool `json:"isBlockedBilling,omitempty"`

	IsBlockedEntitlement bool `json:"isBlockedEntitlement,omitempty"`

	ServiceName string `json:"serviceName,omitempty"`

	ServiceStateName string `json:"serviceStateName,omitempty"`

	Phase string `json:"phase,omitempty"`

	AuditLogs []AuditLogJson `json:"auditLogs,omitempty"`
}
