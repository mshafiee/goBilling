/*
 * goBilling
 *
 * goBilling is an open-source billing and payments platform
 *
 * API version: 1.0.0
 *
 */

package model

type OverdueStateConfigJson struct {

	Name string `json:"name,omitempty"`

	IsClearState bool `json:"isClearState,omitempty"`

	Condition *OverdueConditionJson `json:"condition,omitempty"`

	ExternalMessage string `json:"externalMessage,omitempty"`

	BlockChanges bool `json:"blockChanges,omitempty"`

	DisableEntitlement bool `json:"disableEntitlement,omitempty"`

	SubscriptionCancellationPolicy string `json:"subscriptionCancellationPolicy,omitempty"`

	AutoReevaluationIntervalDays int32 `json:"autoReevaluationIntervalDays,omitempty"`
}
