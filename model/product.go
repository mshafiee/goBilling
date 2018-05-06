/*
 * goBilling
 *
 * goBilling is an open-source billing and payments platform
 *
 * API version: 1.0.0
 *
 */

package model

type Product struct {

	Category string `json:"category,omitempty"`

	Included []Product `json:"included,omitempty"`

	Available []Product `json:"available,omitempty"`

	Limits []Limit `json:"limits,omitempty"`

	CatalogName string `json:"catalogName,omitempty"`

	Name string `json:"name,omitempty"`
}
