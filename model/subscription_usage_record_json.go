/*
 * goBilling
 *
 * goBilling is an open-source billing and payments platform
 *
 * API version: 1.0.0
 *
 */

package model

type SubscriptionUsageRecordJson struct {

	SubscriptionId string `json:"subscriptionId"`

	TrackingId string `json:"trackingId,omitempty"`

	UnitUsageRecords []UnitUsageRecordJson `json:"unitUsageRecords"`
}
