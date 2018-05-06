/*
 * goBilling
 *
 * goBilling is an open-source billing and payments platform
 *
 * API version: 1.0.0
 *
 */

package model

type PhasePriceOverrideJson struct {

	PlanName string `json:"planName,omitempty"`

	PhaseName string `json:"phaseName,omitempty"`

	PhaseType string `json:"phaseType,omitempty"`

	FixedPrice float32 `json:"fixedPrice,omitempty"`

	RecurringPrice float32 `json:"recurringPrice,omitempty"`
}
