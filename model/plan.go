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

type Plan struct {

	RecurringBillingPeriod string `json:"recurringBillingPeriod,omitempty"`

	PriceListName string `json:"priceListName,omitempty"`

	Product *Product `json:"product,omitempty"`

	AllPhases []PlanPhase `json:"allPhases,omitempty"`

	FinalPhase *PlanPhase `json:"finalPhase,omitempty"`

	PlansAllowedInBundle int32 `json:"plansAllowedInBundle,omitempty"`

	EffectiveDateForExistingSubscriptions time.Time `json:"effectiveDateForExistingSubscriptions,omitempty"`

	InitialPhases []PlanPhase `json:"initialPhases,omitempty"`

	InitialPhaseIterator *IteratorPlanPhase `json:"initialPhaseIterator,omitempty"`

	Name string `json:"name,omitempty"`
}
