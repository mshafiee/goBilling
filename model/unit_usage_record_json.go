/*
 * goBilling
 *
 * goBilling is an open-source billing and payments platform
 *
 * API version: 1.0.0
 *
 */

package model

type UnitUsageRecordJson struct {

	UnitType string `json:"unitType,omitempty"`

	UsageRecords []UsageRecordJson `json:"usageRecords,omitempty"`
}
