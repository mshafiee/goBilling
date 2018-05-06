/*
 * goBilling
 *
 * goBilling is an open-source billing and payments platform
 *
 * API version: 1.0.0
 *
 */

package model

type Recurring struct {

	BillingPeriod string `json:"billingPeriod,omitempty"`

	RecurringPrice *InternationalPrice `json:"recurringPrice,omitempty"`
}
