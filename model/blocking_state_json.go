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

type BlockingStateJson struct {

	BlockedId string `json:"blockedId,omitempty"`

	StateName string `json:"stateName,omitempty"`

	Service string `json:"service,omitempty"`

	BlockChange bool `json:"blockChange,omitempty"`

	BlockEntitlement bool `json:"blockEntitlement,omitempty"`

	BlockBilling bool `json:"blockBilling,omitempty"`

	EffectiveDate time.Time `json:"effectiveDate,omitempty"`

	Type_ string `json:"type,omitempty"`

	AuditLogs []AuditLogJson `json:"auditLogs,omitempty"`
}
