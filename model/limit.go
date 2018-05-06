/*
 * goBilling
 *
 * goBilling is an open-source billing and payments platform
 *
 * API version: 1.0.0
 *
 */

package model

type Limit struct {

	Max float64 `json:"max,omitempty"`

	Unit *Unit `json:"unit,omitempty"`

	Min float64 `json:"min,omitempty"`
}
