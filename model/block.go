/*
 * goBilling
 *
 * goBilling is an open-source billing and payments platform
 *
 * API version: 1.0.0
 *
 */

package model

type Block struct {

	Unit *Unit `json:"unit,omitempty"`

	Price *InternationalPrice `json:"price,omitempty"`

	MinTopUpCredit float64 `json:"minTopUpCredit,omitempty"`

	Type_ string `json:"type,omitempty"`

	Size float64 `json:"size,omitempty"`
}
