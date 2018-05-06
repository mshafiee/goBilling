/*
 * goBilling
 *
 * goBilling is an open-source billing and payments platform
 *
 * API version: 1.0.0
 *
 */

package model

type AccountTimelineJson struct {

	Account *AccountJson `json:"account,omitempty"`

	Bundles []BundleJson `json:"bundles,omitempty"`

	Invoices []InvoiceJson `json:"invoices,omitempty"`

	Payments []InvoicePaymentJson `json:"payments,omitempty"`
}
