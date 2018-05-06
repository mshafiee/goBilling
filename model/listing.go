/*
 * goBilling
 *
 * goBilling is an open-source billing and payments platform
 *
 * API version: 1.0.0
 *
 */

package model

type Listing struct {

	PriceList *PriceList `json:"priceList,omitempty"`

	Plan *Plan `json:"plan,omitempty"`
}
