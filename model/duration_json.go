/*
 * goBilling
 *
 * goBilling is an open-source billing and payments platform
 *
 * API version: 1.0.0
 *
 */

package model

type DurationJson struct {

	Unit string `json:"unit,omitempty"`

	Number int32 `json:"number,omitempty"`
}
