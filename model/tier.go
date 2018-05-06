/*
 * goBilling
 *
 * goBilling is an open-source billing and payments platform
 *
 * API version: 1.0.0
 *
 */

package model

type Tier struct {

	FixedPrice *InternationalPrice `json:"fixedPrice,omitempty"`

	RecurringPrice *InternationalPrice `json:"recurringPrice,omitempty"`

	Limits []Limit `json:"limits,omitempty"`

	TieredBlocks []TieredBlock `json:"tieredBlocks,omitempty"`
}
