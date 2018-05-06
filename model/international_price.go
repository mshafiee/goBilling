/*
 * goBilling
 *
 * goBilling is an open-source billing and payments platform
 *
 * API version: 1.0.0
 *
 */

package model

type InternationalPrice struct {

	Prices []Price `json:"prices,omitempty"`

	Zero bool `json:"zero,omitempty"`
}
