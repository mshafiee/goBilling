/*
 * goBilling
 *
 * goBilling is an open-source billing and payments platform
 *
 * API version: 1.0.0
 *
 */

package model

type OverdueJson struct {

	InitialReevaluationInterval int32 `json:"initialReevaluationInterval,omitempty"`

	OverdueStates []OverdueStateConfigJson `json:"overdueStates,omitempty"`
}
