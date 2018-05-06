/*
 * goBilling
 *
 * goBilling is an open-source billing and payments platform
 *
 * API version: 1.0.0
 *
 */

package model

import (
	"time"
)

type ClockResource struct {

	CurrentUtcTime time.Time `json:"currentUtcTime,omitempty"`

	TimeZone string `json:"timeZone,omitempty"`

	LocalDate string `json:"localDate,omitempty"`
}
