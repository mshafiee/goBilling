/*
 * goBilling
 *
 * goBilling is an open-source billing and payments platform
 *
 * API version: 1.0.0
 *
 */

package model

type PlanJson struct {

	Name string `json:"name,omitempty"`

	BillingPeriod string `json:"billingPeriod,omitempty"`

	Phases []PhaseJson `json:"phases,omitempty"`
}
