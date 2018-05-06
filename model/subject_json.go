/*
 * goBilling
 *
 * goBilling is an open-source billing and payments platform
 *
 * API version: 1.0.0
 *
 */

package model

type SubjectJson struct {

	Principal string `json:"principal,omitempty"`

	IsAuthenticated bool `json:"isAuthenticated,omitempty"`

	IsRemembered bool `json:"isRemembered,omitempty"`

	Session *SessionJson `json:"session,omitempty"`
}
