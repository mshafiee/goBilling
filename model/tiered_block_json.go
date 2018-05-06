/*
 * goBilling
 *
 * goBilling is an open-source billing and payments platform
 *
 * API version: 1.0.0
 *
 */

package model

type TieredBlockJson struct {

	Unit string `json:"unit,omitempty"`

	Size string `json:"size,omitempty"`

	Max string `json:"max,omitempty"`

	Prices []PriceJson `json:"prices,omitempty"`
}
