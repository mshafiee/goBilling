/*
 * goBilling
 *
 * goBilling is an open-source billing and payments platform
 *
 * API version: 1.0.0
 *
 */

package model

type Fixed struct {

	Price *InternationalPrice `json:"price,omitempty"`

	Type_ string `json:"type,omitempty"`
}
