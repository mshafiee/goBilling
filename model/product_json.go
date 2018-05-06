/*
 * goBilling
 *
 * goBilling is an open-source billing and payments platform
 *
 * API version: 1.0.0
 *
 */

package model

type ProductJson struct {

	Type_ string `json:"type,omitempty"`

	Name string `json:"name,omitempty"`

	Plans []PlanJson `json:"plans,omitempty"`

	Included []string `json:"included,omitempty"`

	Available []string `json:"available,omitempty"`
}
