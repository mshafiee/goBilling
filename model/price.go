/*
 * goBilling
 *
 * goBilling is an open-source billing and payments platform
 *
 * API version: 1.0.0
 *
 */

package model

type Price struct {

	Currency string `json:"currency,omitempty"`

	Value float32 `json:"value,omitempty"`
}
