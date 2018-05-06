/*
 * goBilling
 *
 * goBilling is an open-source billing and payments platform
 *
 * API version: 1.0.0
 *
 */

package model

type PaymentMethodPluginDetailJson struct {

	ExternalPaymentMethodId string `json:"externalPaymentMethodId,omitempty"`

	IsDefaultPaymentMethod bool `json:"isDefaultPaymentMethod,omitempty"`

	Properties []PluginPropertyJson `json:"properties,omitempty"`
}
