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

type SessionJson struct {

	Id string `json:"id,omitempty"`

	StartDate time.Time `json:"startDate,omitempty"`

	LastAccessDate time.Time `json:"lastAccessDate,omitempty"`

	Timeout int64 `json:"timeout,omitempty"`

	Host string `json:"host,omitempty"`
}
