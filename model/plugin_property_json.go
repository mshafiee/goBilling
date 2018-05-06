/*
 * goBilling
 *
 * goBilling is an open-source billing and payments platform
 *
 * API version: 1.0.0
 *
 */

package model

type PluginPropertyJson struct {

	Key string `json:"key,omitempty"`

	Value string `json:"value,omitempty"`

	IsUpdatable bool `json:"isUpdatable,omitempty"`
}
