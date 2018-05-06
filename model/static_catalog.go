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

type StaticCatalog struct {

	EffectiveDate time.Time `json:"effectiveDate,omitempty"`

	AvailableBasePlanListings []Listing `json:"availableBasePlanListings,omitempty"`

	RecurringBillingMode string `json:"recurringBillingMode,omitempty"`

	CatalogName string `json:"catalogName,omitempty"`

	CurrentProducts []Product `json:"currentProducts,omitempty"`

	CurrentUnits []Unit `json:"currentUnits,omitempty"`

	CurrentSupportedCurrencies []string `json:"currentSupportedCurrencies,omitempty"`

	CurrentPlans []Plan `json:"currentPlans,omitempty"`
}
