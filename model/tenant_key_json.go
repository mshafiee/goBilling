/*
 * goBilling
 *
 * goBilling is an open-source billing and payments platform
 *
 * API version: 1.0.0
 *
 */

package model

type TenantKeyJson struct {

	Key string `json:"key,omitempty"`

	Values []string `json:"values,omitempty"`
}
