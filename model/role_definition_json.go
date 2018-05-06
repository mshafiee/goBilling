/*
 * goBilling
 *
 * goBilling is an open-source billing and payments platform
 *
 * API version: 1.0.0
 *
 */

package model

type RoleDefinitionJson struct {

	Role string `json:"role"`

	Permissions []string `json:"permissions"`
}
