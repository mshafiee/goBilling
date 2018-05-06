/*
 * goBilling
 *
 * goBilling is an open-source billing and payments platform
 *
 * API version: 1.0.0
 *
 */

package model

type PhaseJson struct {

	Type_ string `json:"type,omitempty"`

	Prices []PriceJson `json:"prices,omitempty"`

	FixedPrices []PriceJson `json:"fixedPrices,omitempty"`

	Duration *DurationJson `json:"duration,omitempty"`

	Usages []UsageJson `json:"usages,omitempty"`
}
