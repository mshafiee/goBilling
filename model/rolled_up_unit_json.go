/*
 * goBilling
 *
 * goBilling is an open-source billing and payments platform
 *
 * API version: 1.0.0
 *
 */

package model

type RolledUpUnitJson struct {

	UnitType string `json:"unitType,omitempty"`

	Amount int64 `json:"amount,omitempty"`
}
