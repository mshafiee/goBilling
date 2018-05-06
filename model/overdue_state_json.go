/*
 * goBilling
 *
 * goBilling is an open-source billing and payments platform
 *
 * API version: 1.0.0
 *
 */

package model

type OverdueStateJson struct {

	Name string `json:"name,omitempty"`

	ExternalMessage string `json:"externalMessage,omitempty"`

	DaysBetweenPaymentRetries []int32 `json:"daysBetweenPaymentRetries,omitempty"`

	DisableEntitlementAndChangesBlocked bool `json:"disableEntitlementAndChangesBlocked,omitempty"`

	BlockChanges bool `json:"blockChanges,omitempty"`

	ClearState bool `json:"clearState,omitempty"`

	ReevaluationIntervalDays int32 `json:"reevaluationIntervalDays,omitempty"`
}
