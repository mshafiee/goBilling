/*
 * goBilling
 *
 * goBilling is an open-source billing and payments platform
 *
 * API version: 1.0.0
 *
 */

package model

type AdminPaymentJson struct {

	LastSuccessPaymentState string `json:"lastSuccessPaymentState,omitempty"`

	CurrentPaymentStateName string `json:"currentPaymentStateName,omitempty"`

	TransactionStatus string `json:"transactionStatus,omitempty"`
}
