/*
 * goBilling
 *
 * goBilling is an open-source billing and payments platform
 *
 * API version: 1.0.0
 *
 */

package model

type PriceListJson struct {

	Name string `json:"name,omitempty"`

	Plans []string `json:"plans,omitempty"`
}
