/*
 * goBilling
 *
 * goBilling is an open-source billing and payments platform
 *
 * API version: 1.0.0
 *
 */

package model

type UserRolesJson struct {

	Username string `json:"username"`

	Password string `json:"password"`

	Roles []string `json:"roles"`
}
