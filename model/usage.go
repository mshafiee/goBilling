/*
 * goBilling
 *
 * goBilling is an open-source billing and payments platform
 *
 * API version: 1.0.0
 *
 */

package model

type Usage struct {

	BillingPeriod string `json:"billingPeriod,omitempty"`

	FixedPrice *InternationalPrice `json:"fixedPrice,omitempty"`

	RecurringPrice *InternationalPrice `json:"recurringPrice,omitempty"`

	Tiers []Tier `json:"tiers,omitempty"`

	Blocks []Block `json:"blocks,omitempty"`

	Limits []Limit `json:"limits,omitempty"`

	BillingMode string `json:"billingMode,omitempty"`

	UsageType string `json:"usageType,omitempty"`

	Name string `json:"name,omitempty"`
}
