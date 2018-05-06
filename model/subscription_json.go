/*
 * goBilling
 *
 * goBilling is an open-source billing and payments platform
 *
 * API version: 1.0.0
 *
 */

package model

type SubscriptionJson struct {

	AccountId string `json:"accountId,omitempty"`

	BundleId string `json:"bundleId,omitempty"`

	SubscriptionId string `json:"subscriptionId,omitempty"`

	ExternalKey string `json:"externalKey,omitempty"`

	StartDate string `json:"startDate,omitempty"`

	ProductName string `json:"productName"`

	ProductCategory string `json:"productCategory"`

	BillingPeriod string `json:"billingPeriod"`

	PhaseType string `json:"phaseType,omitempty"`

	PriceList string `json:"priceList"`

	PlanName string `json:"planName"`

	State string `json:"state,omitempty"`

	SourceType string `json:"sourceType,omitempty"`

	CancelledDate string `json:"cancelledDate,omitempty"`

	ChargedThroughDate string `json:"chargedThroughDate,omitempty"`

	BillingStartDate string `json:"billingStartDate,omitempty"`

	BillingEndDate string `json:"billingEndDate,omitempty"`

	BillCycleDayLocal int32 `json:"billCycleDayLocal,omitempty"`

	Events []EventSubscriptionJson `json:"events,omitempty"`

	PriceOverrides []PhasePriceOverrideJson `json:"priceOverrides,omitempty"`

	AuditLogs []AuditLogJson `json:"auditLogs,omitempty"`
}
