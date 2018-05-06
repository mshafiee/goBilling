/*
 * goBilling
 *
 * goBilling is an open-source billing and payments platform
 *
 * API version: 1.0.0
 *
 */

package model

type PlanDetailJson struct {

	Product string `json:"product,omitempty"`

	Plan string `json:"plan,omitempty"`

	PriceList string `json:"priceList,omitempty"`

	FinalPhaseBillingPeriod string `json:"finalPhaseBillingPeriod,omitempty"`

	FinalPhaseRecurringPrice []PriceJson `json:"finalPhaseRecurringPrice,omitempty"`
}
