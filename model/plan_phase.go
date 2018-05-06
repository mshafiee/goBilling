/*
 * goBilling
 *
 * goBilling is an open-source billing and payments platform
 *
 * API version: 1.0.0
 *
 */

package model

type PlanPhase struct {

	Duration *Duration `json:"duration,omitempty"`

	Usages []Usage `json:"usages,omitempty"`

	PhaseType string `json:"phaseType,omitempty"`

	Recurring *Recurring `json:"recurring,omitempty"`

	Fixed *Fixed `json:"fixed,omitempty"`

	Name string `json:"name,omitempty"`
}
