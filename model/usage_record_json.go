/*
 * goBilling
 *
 * goBilling is an open-source billing and payments platform
 *
 * API version: 1.0.0
 *
 */

package model

type UsageRecordJson struct {

	RecordDate string `json:"recordDate,omitempty"`

	Amount int64 `json:"amount,omitempty"`
}
