/*
 * goBilling
 *
 * goBilling is an open-source billing and payments platform
 *
 * API version: 1.0.0
 *
 */

package model

type PluginInfoJson struct {

	BundleSymbolicName string `json:"bundleSymbolicName,omitempty"`

	PluginKey string `json:"pluginKey,omitempty"`

	PluginName string `json:"pluginName,omitempty"`

	Version string `json:"version,omitempty"`

	State string `json:"state,omitempty"`

	IsSelectedForStart bool `json:"isSelectedForStart,omitempty"`

	Services []PluginServiceInfoJson `json:"services,omitempty"`
}
