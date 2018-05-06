/*
 * goBilling
 *
 * goBilling is an open-source billing and payments platform
 *
 * API version: 1.0.0
 *
 */

package model

type TierJson struct {

	Limits []LimitJson `json:"limits,omitempty"`

	FixedPrice []PriceJson `json:"fixedPrice,omitempty"`

	RecurringPrice []PriceJson `json:"recurringPrice,omitempty"`

	Blocks []TieredBlockJson `json:"blocks,omitempty"`
}
