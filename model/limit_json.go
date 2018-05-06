/*
 * goBilling
 *
 * goBilling is an open-source billing and payments platform
 *
 * API version: 1.0.0
 *
 */

package model

type LimitJson struct {

	Unit string `json:"unit,omitempty"`

	Max string `json:"max,omitempty"`

	Min string `json:"min,omitempty"`
}
