/*
 * goBilling
 *
 * goBilling is an open-source billing and payments platform
 *
 * API version: 1.0.0
 *
 */

package model

type NodeCommandJson struct {

	SystemCommandType bool `json:"systemCommandType,omitempty"`

	NodeCommandType string `json:"nodeCommandType,omitempty"`

	NodeCommandProperties []NodeCommandPropertyJson `json:"nodeCommandProperties,omitempty"`
}
