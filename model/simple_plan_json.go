/*
 * goBilling
 *
 * goBilling is an open-source billing and payments platform
 *
 * API version: 1.0.0
 *
 */

package model

type SimplePlanJson struct {

	PlanId string `json:"planId,omitempty"`

	ProductName string `json:"productName,omitempty"`

	ProductCategory string `json:"productCategory,omitempty"`

	Currency string `json:"currency,omitempty"`

	Amount float32 `json:"amount,omitempty"`

	BillingPeriod string `json:"billingPeriod,omitempty"`

	TrialLength int32 `json:"trialLength,omitempty"`

	TrialTimeUnit string `json:"trialTimeUnit,omitempty"`

	AvailableBaseProducts []string `json:"availableBaseProducts,omitempty"`
}
