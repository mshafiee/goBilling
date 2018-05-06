/*
 * goBilling
 *
 * goBilling is an open-source billing and payments platform
 *
 * API version: 1.0.0
 *
 */

package model

type PriceList struct {

	Plans []Plan `json:"plans,omitempty"`

	Name string `json:"name,omitempty"`
}
