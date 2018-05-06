/*
 * goBilling
 *
 * goBilling is an open-source billing and payments platform
 *
 * API version: 1.0.0
 *
 */

package model

type InvoiceDryRunJson struct {

	DryRunType string `json:"dryRunType,omitempty"`

	DryRunAction string `json:"dryRunAction,omitempty"`

	PhaseType string `json:"phaseType,omitempty"`

	ProductName string `json:"productName,omitempty"`

	ProductCategory string `json:"productCategory,omitempty"`

	BillingPeriod string `json:"billingPeriod,omitempty"`

	PriceListName string `json:"priceListName,omitempty"`

	SubscriptionId string `json:"subscriptionId,omitempty"`

	BundleId string `json:"bundleId,omitempty"`

	EffectiveDate string `json:"effectiveDate,omitempty"`

	BillingPolicy string `json:"billingPolicy,omitempty"`

	PriceOverrides []PhasePriceOverrideJson `json:"priceOverrides,omitempty"`
}
