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

type AuditLogJson struct {

	ChangeType string `json:"changeType,omitempty"`

	ChangeDate time.Time `json:"changeDate,omitempty"`

	ChangedBy string `json:"changedBy,omitempty"`

	ReasonCode string `json:"reasonCode,omitempty"`

	Comments string `json:"comments,omitempty"`

	UserToken string `json:"userToken,omitempty"`
}
