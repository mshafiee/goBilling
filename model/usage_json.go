/*
 * goBilling
 *
 * goBilling is an open-source billing and payments platform
 *
 * API version: 1.0.0
 *
 */

package model

type UsageJson struct {

	BillingPeriod string `json:"billingPeriod,omitempty"`

	Tiers []TierJson `json:"tiers,omitempty"`
}
