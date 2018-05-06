/*
 * goBilling
 *
 * goBilling is an open-source billing and payments platform
 *
 * API version: 1.0.0
 *
 */

package model

type OverdueConditionJson struct {

	TimeSinceEarliestUnpaidInvoiceEqualsOrExceeds *DurationJson `json:"timeSinceEarliestUnpaidInvoiceEqualsOrExceeds,omitempty"`

	ControlTagInclusion string `json:"controlTagInclusion,omitempty"`

	ControlTagExclusion string `json:"controlTagExclusion,omitempty"`

	NumberOfUnpaidInvoicesEqualsOrExceeds int32 `json:"numberOfUnpaidInvoicesEqualsOrExceeds,omitempty"`

	ResponseForLastFailedPayment []string `json:"responseForLastFailedPayment,omitempty"`

	TotalUnpaidInvoiceBalanceEqualsOrExceeds float32 `json:"totalUnpaidInvoiceBalanceEqualsOrExceeds,omitempty"`
}
