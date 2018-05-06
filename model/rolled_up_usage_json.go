/*
 * goBilling
 *
 * goBilling is an open-source billing and payments platform
 *
 * API version: 1.0.0
 *
 */

package model

type RolledUpUsageJson struct {

	SubscriptionId string `json:"subscriptionId,omitempty"`

	StartDate string `json:"startDate,omitempty"`

	EndDate string `json:"endDate,omitempty"`

	RolledUpUnits []RolledUpUnitJson `json:"rolledUpUnits,omitempty"`
}
