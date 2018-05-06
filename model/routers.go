/*
 * goBilling
 *
 * goBilling is an open-source billing and payments platform
 *
 * API version: 1.0.0
 *
 */

package model

import (
	"fmt"
	"net/http"
	"strings"
	"github.com/gorilla/mux"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	return router
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to goBilling!")
}

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},

	Route{
		"AddAccountBlockingState",
		strings.ToUpper("Put"),
		"/api/v1/accounts/{accountId}/block",
		AddAccountBlockingState,
	},

	Route{
		"AddEmail",
		strings.ToUpper("Post"),
		"/api/v1/accounts/{accountId}/emails",
		AddEmail,
	},

	Route{
		"CloseAccount",
		strings.ToUpper("Delete"),
		"/api/v1/accounts/{accountId}",
		CloseAccount,
	},

	Route{
		"CreateAccount",
		strings.ToUpper("Post"),
		"/api/v1/accounts",
		CreateAccount,
	},

	Route{
		"CreateCustomFields",
		strings.ToUpper("Post"),
		"/api/v1/accounts/{accountId}/customFields",
		CreateCustomFields,
	},

	Route{
		"CreatePaymentMethod",
		strings.ToUpper("Post"),
		"/api/v1/accounts/{accountId}/paymentMethods",
		CreatePaymentMethod,
	},

	Route{
		"CreateTags",
		strings.ToUpper("Post"),
		"/api/v1/accounts/{accountId}/tags",
		CreateTags,
	},

	Route{
		"DeleteCustomFields",
		strings.ToUpper("Delete"),
		"/api/v1/accounts/{accountId}/customFields",
		DeleteCustomFields,
	},

	Route{
		"DeleteTags",
		strings.ToUpper("Delete"),
		"/api/v1/accounts/{accountId}/tags",
		DeleteTags,
	},

	Route{
		"GetAccount",
		strings.ToUpper("Get"),
		"/api/v1/accounts/{accountId}",
		GetAccount,
	},

	Route{
		"GetAccountBundles",
		strings.ToUpper("Get"),
		"/api/v1/accounts/{accountId}/bundles",
		GetAccountBundles,
	},

	Route{
		"GetAccountByKey",
		strings.ToUpper("Get"),
		"/api/v1/accounts",
		GetAccountByKey,
	},

	Route{
		"GetAccountTimeline",
		strings.ToUpper("Get"),
		"/api/v1/accounts/{accountId}/timeline",
		GetAccountTimeline,
	},

	Route{
		"GetAccounts",
		strings.ToUpper("Get"),
		"/api/v1/accounts/pagination",
		GetAccounts,
	},

	Route{
		"GetAllCustomFields",
		strings.ToUpper("Get"),
		"/api/v1/accounts/{accountId}/allCustomFields",
		GetAllCustomFields,
	},

	Route{
		"GetAllTags",
		strings.ToUpper("Get"),
		"/api/v1/accounts/{accountId}/allTags",
		GetAllTags,
	},

	Route{
		"GetBlockingStates",
		strings.ToUpper("Get"),
		"/api/v1/accounts/{accountId}/block",
		GetBlockingStates,
	},

	Route{
		"GetChildrenAccounts",
		strings.ToUpper("Get"),
		"/api/v1/accounts/{accountId}/children",
		GetChildrenAccounts,
	},

	Route{
		"GetCustomFields",
		strings.ToUpper("Get"),
		"/api/v1/accounts/{accountId}/customFields",
		GetCustomFields,
	},

	Route{
		"GetEmailNotificationsForAccount",
		strings.ToUpper("Get"),
		"/api/v1/accounts/{accountId}/emailNotifications",
		GetEmailNotificationsForAccount,
	},

	Route{
		"GetEmails",
		strings.ToUpper("Get"),
		"/api/v1/accounts/{accountId}/emails",
		GetEmails,
	},

	Route{
		"GetInvoicePayments",
		strings.ToUpper("Get"),
		"/api/v1/accounts/{accountId}/invoicePayments",
		GetInvoicePayments,
	},

	Route{
		"GetInvoices",
		strings.ToUpper("Get"),
		"/api/v1/accounts/{accountId}/invoices",
		GetInvoices,
	},

	Route{
		"GetOverdueAccount",
		strings.ToUpper("Get"),
		"/api/v1/accounts/{accountId}/overdue",
		GetOverdueAccount,
	},

	Route{
		"GetPaymentMethods",
		strings.ToUpper("Get"),
		"/api/v1/accounts/{accountId}/paymentMethods",
		GetPaymentMethods,
	},

	Route{
		"GetPayments",
		strings.ToUpper("Get"),
		"/api/v1/accounts/{accountId}/payments",
		GetPayments,
	},

	Route{
		"GetTags",
		strings.ToUpper("Get"),
		"/api/v1/accounts/{accountId}/tags",
		GetTags,
	},

	Route{
		"PayAllInvoices",
		strings.ToUpper("Post"),
		"/api/v1/accounts/{accountId}/invoicePayments",
		PayAllInvoices,
	},

	Route{
		"ProcessPayment",
		strings.ToUpper("Post"),
		"/api/v1/accounts/{accountId}/payments",
		ProcessPayment,
	},

	Route{
		"ProcessPaymentByExternalKey",
		strings.ToUpper("Post"),
		"/api/v1/accounts/payments",
		ProcessPaymentByExternalKey,
	},

	Route{
		"RebalanceExistingCBAOnAccount",
		strings.ToUpper("Post"),
		"/api/v1/accounts/{accountId}/cbaRebalancing",
		RebalanceExistingCBAOnAccount,
	},

	Route{
		"RefreshPaymentMethods",
		strings.ToUpper("Post"),
		"/api/v1/accounts/{accountId}/paymentMethods/refresh",
		RefreshPaymentMethods,
	},

	Route{
		"RemoveEmail",
		strings.ToUpper("Delete"),
		"/api/v1/accounts/{accountId}/emails/{email}",
		RemoveEmail,
	},

	Route{
		"SearchAccounts",
		strings.ToUpper("Get"),
		"/api/v1/accounts/search/{searchKey}",
		SearchAccounts,
	},

	Route{
		"SetDefaultPaymentMethod",
		strings.ToUpper("Put"),
		"/api/v1/accounts/{accountId}/paymentMethods/{paymentMethodId}/setDefault",
		SetDefaultPaymentMethod,
	},

	Route{
		"SetEmailNotificationsForAccount",
		strings.ToUpper("Put"),
		"/api/v1/accounts/{accountId}/emailNotifications",
		SetEmailNotificationsForAccount,
	},

	Route{
		"TransferChildCreditToParent",
		strings.ToUpper("Post"),
		"/api/v1/accounts/{childAccountId}/transferCredit",
		TransferChildCreditToParent,
	},

	Route{
		"UpdateAccount",
		strings.ToUpper("Put"),
		"/api/v1/accounts/{accountId}",
		UpdateAccount,
	},

	Route{
		"GetQueueEntries",
		strings.ToUpper("Get"),
		"/api/v1/admin/queues",
		GetQueueEntries,
	},

	Route{
		"InvalidatesCache",
		strings.ToUpper("Delete"),
		"/api/v1/admin/cache",
		InvalidatesCache,
	},

	Route{
		"InvalidatesCacheByAccount",
		strings.ToUpper("Delete"),
		"/api/v1/admin/cache/accounts/{accountId}",
		InvalidatesCacheByAccount,
	},

	Route{
		"InvalidatesCacheByTenant",
		strings.ToUpper("Delete"),
		"/api/v1/admin/cache/tenants",
		InvalidatesCacheByTenant,
	},

	Route{
		"PutInRotation",
		strings.ToUpper("Post"),
		"/api/v1/admin/healthcheck",
		PutInRotation,
	},

	Route{
		"PutOutOfRotation",
		strings.ToUpper("Delete"),
		"/api/v1/admin/healthcheck",
		PutOutOfRotation,
	},

	Route{
		"TriggerInvoiceGenerationForParkedAccounts",
		strings.ToUpper("Post"),
		"/api/v1/admin/invoices",
		TriggerInvoiceGenerationForParkedAccounts,
	},

	Route{
		"UpdatePaymentTransactionState",
		strings.ToUpper("Put"),
		"/api/v1/admin/payments/{paymentId}/transactions/{paymentTransactionId}",
		UpdatePaymentTransactionState,
	},

	Route{
		"AddBundleBlockingState",
		strings.ToUpper("Put"),
		"/api/v1/bundles/{bundleId}/block",
		AddBundleBlockingState,
	},

	Route{
		"CreateCustomFields",
		strings.ToUpper("Post"),
		"/api/v1/bundles/{bundleId}/customFields",
		CreateCustomFields,
	},

	Route{
		"CreateTags",
		strings.ToUpper("Post"),
		"/api/v1/bundles/{bundleId}/tags",
		CreateTags,
	},

	Route{
		"DeleteCustomFields",
		strings.ToUpper("Delete"),
		"/api/v1/bundles/{bundleId}/customFields",
		DeleteCustomFields,
	},

	Route{
		"DeleteTags",
		strings.ToUpper("Delete"),
		"/api/v1/bundles/{bundleId}/tags",
		DeleteTags,
	},

	Route{
		"GetBundle",
		strings.ToUpper("Get"),
		"/api/v1/bundles/{bundleId}",
		GetBundle,
	},

	Route{
		"GetBundleByKey",
		strings.ToUpper("Get"),
		"/api/v1/bundles",
		GetBundleByKey,
	},

	Route{
		"GetBundles",
		strings.ToUpper("Get"),
		"/api/v1/bundles/pagination",
		GetBundles,
	},

	Route{
		"GetCustomFields",
		strings.ToUpper("Get"),
		"/api/v1/bundles/{bundleId}/customFields",
		GetCustomFields,
	},

	Route{
		"GetTags",
		strings.ToUpper("Get"),
		"/api/v1/bundles/{bundleId}/tags",
		GetTags,
	},

	Route{
		"PauseBundle",
		strings.ToUpper("Put"),
		"/api/v1/bundles/{bundleId}/pause",
		PauseBundle,
	},

	Route{
		"ResumeBundle",
		strings.ToUpper("Put"),
		"/api/v1/bundles/{bundleId}/resume",
		ResumeBundle,
	},

	Route{
		"SearchBundles",
		strings.ToUpper("Get"),
		"/api/v1/bundles/search/{searchKey}",
		SearchBundles,
	},

	Route{
		"TransferBundle",
		strings.ToUpper("Put"),
		"/api/v1/bundles/{bundleId}",
		TransferBundle,
	},

	Route{
		"AddSimplePlan",
		strings.ToUpper("Post"),
		"/api/v1/catalog/simplePlan",
		AddSimplePlan,
	},

	Route{
		"GetAvailableAddons",
		strings.ToUpper("Get"),
		"/api/v1/catalog/availableAddons",
		GetAvailableAddons,
	},

	Route{
		"GetAvailableBasePlans",
		strings.ToUpper("Get"),
		"/api/v1/catalog/availableBasePlans",
		GetAvailableBasePlans,
	},

	Route{
		"GetCatalogJson",
		strings.ToUpper("Get"),
		"/api/v1/catalog",
		GetCatalogJson,
	},

	Route{
		"GetPhaseForSubscriptionAndDate",
		strings.ToUpper("Get"),
		"/api/v1/catalog/phase",
		GetPhaseForSubscriptionAndDate,
	},

	Route{
		"GetPlanForSubscriptionAndDate",
		strings.ToUpper("Get"),
		"/api/v1/catalog/plan",
		GetPlanForSubscriptionAndDate,
	},

	Route{
		"GetPriceListForSubscriptionAndDate",
		strings.ToUpper("Get"),
		"/api/v1/catalog/priceList",
		GetPriceListForSubscriptionAndDate,
	},

	Route{
		"GetProductForSubscriptionAndDate",
		strings.ToUpper("Get"),
		"/api/v1/catalog/product",
		GetProductForSubscriptionAndDate,
	},

	Route{
		"UploadCatalogXml",
		strings.ToUpper("Post"),
		"/api/v1/catalog",
		UploadCatalogXml,
	},

	Route{
		"CreateCredit",
		strings.ToUpper("Post"),
		"/api/v1/credits",
		CreateCredit,
	},

	Route{
		"GetCredit",
		strings.ToUpper("Get"),
		"/api/v1/credits/{creditId}",
		GetCredit,
	},

	Route{
		"GetCustomFields",
		strings.ToUpper("Get"),
		"/api/v1/customFields/pagination",
		GetCustomFields,
	},

	Route{
		"SearchCustomFields",
		strings.ToUpper("Get"),
		"/api/v1/customFields/search/{searchKey}",
		SearchCustomFields,
	},

	Route{
		"ExportDataForAccount",
		strings.ToUpper("Get"),
		"/api/v1/export/{accountId}",
		ExportDataForAccount,
	},

	Route{
		"CreateChargeback",
		strings.ToUpper("Post"),
		"/api/v1/invoicePayments/{paymentId}/chargebacks",
		CreateChargeback,
	},

	Route{
		"CreateChargebackReversal",
		strings.ToUpper("Post"),
		"/api/v1/invoicePayments/{paymentId}/chargebackReversals",
		CreateChargebackReversal,
	},

	Route{
		"CreateCustomFields",
		strings.ToUpper("Post"),
		"/api/v1/invoicePayments/{paymentId}/customFields",
		CreateCustomFields,
	},

	Route{
		"CreateRefundWithAdjustments",
		strings.ToUpper("Post"),
		"/api/v1/invoicePayments/{paymentId}/refunds",
		CreateRefundWithAdjustments,
	},

	Route{
		"CreateTags",
		strings.ToUpper("Post"),
		"/api/v1/invoicePayments/{paymentId}/tags",
		CreateTags,
	},

	Route{
		"DeleteCustomFields",
		strings.ToUpper("Delete"),
		"/api/v1/invoicePayments/{paymentId}/customFields",
		DeleteCustomFields,
	},

	Route{
		"DeleteTags",
		strings.ToUpper("Delete"),
		"/api/v1/invoicePayments/{paymentId}/tags",
		DeleteTags,
	},

	Route{
		"GetCustomFields",
		strings.ToUpper("Get"),
		"/api/v1/invoicePayments/{paymentId}/customFields",
		GetCustomFields,
	},

	Route{
		"GetInvoicePayment",
		strings.ToUpper("Get"),
		"/api/v1/invoicePayments/{paymentId}",
		GetInvoicePayment,
	},

	Route{
		"GetTags",
		strings.ToUpper("Get"),
		"/api/v1/invoicePayments/{paymentId}/tags",
		GetTags,
	},

	Route{
		"AdjustInvoiceItem",
		strings.ToUpper("Post"),
		"/api/v1/invoices/{invoiceId}",
		AdjustInvoiceItem,
	},

	Route{
		"CommitInvoice",
		strings.ToUpper("Put"),
		"/api/v1/invoices/{invoiceId}/commitInvoice",
		CommitInvoice,
	},

	Route{
		"CreateCustomFields",
		strings.ToUpper("Post"),
		"/api/v1/invoices/{invoiceId}/customFields",
		CreateCustomFields,
	},

	Route{
		"CreateExternalCharges",
		strings.ToUpper("Post"),
		"/api/v1/invoices/charges/{accountId}",
		CreateExternalCharges,
	},

	Route{
		"CreateFutureInvoice",
		strings.ToUpper("Post"),
		"/api/v1/invoices",
		CreateFutureInvoice,
	},

	Route{
		"CreateInstantPayment",
		strings.ToUpper("Post"),
		"/api/v1/invoices/{invoiceId}/payments",
		CreateInstantPayment,
	},

	Route{
		"CreateMigrationInvoice",
		strings.ToUpper("Post"),
		"/api/v1/invoices/migration/{accountId}",
		CreateMigrationInvoice,
	},

	Route{
		"CreateTags",
		strings.ToUpper("Post"),
		"/api/v1/invoices/{invoiceId}/tags",
		CreateTags,
	},

	Route{
		"DeleteCBA",
		strings.ToUpper("Delete"),
		"/api/v1/invoices/{invoiceId}/{invoiceItemId}/cba",
		DeleteCBA,
	},

	Route{
		"DeleteCustomFields",
		strings.ToUpper("Delete"),
		"/api/v1/invoices/{invoiceId}/customFields",
		DeleteCustomFields,
	},

	Route{
		"DeleteTags",
		strings.ToUpper("Delete"),
		"/api/v1/invoices/{invoiceId}/tags",
		DeleteTags,
	},

	Route{
		"GenerateDryRunInvoice",
		strings.ToUpper("Post"),
		"/api/v1/invoices/dryRun",
		GenerateDryRunInvoice,
	},

	Route{
		"GetCustomFields",
		strings.ToUpper("Get"),
		"/api/v1/invoices/{invoiceId}/customFields",
		GetCustomFields,
	},

	Route{
		"GetInvoice",
		strings.ToUpper("Get"),
		"/api/v1/invoices/{invoiceId}",
		GetInvoice,
	},

	Route{
		"GetInvoiceAsHTML",
		strings.ToUpper("Get"),
		"/api/v1/invoices/{invoiceId}/html",
		GetInvoiceAsHTML,
	},

	Route{
		"GetInvoiceByNumber",
		strings.ToUpper("Get"),
		"/api/v1/invoices/{invoiceNumber}",
		GetInvoiceByNumber,
	},

	Route{
		"GetInvoices",
		strings.ToUpper("Get"),
		"/api/v1/invoices/pagination",
		GetInvoices,
	},

	Route{
		"GetPayments",
		strings.ToUpper("Get"),
		"/api/v1/invoices/{invoiceId}/payments",
		GetPayments,
	},

	Route{
		"GetTags",
		strings.ToUpper("Get"),
		"/api/v1/invoices/{invoiceId}/tags",
		GetTags,
	},

	Route{
		"SearchInvoices",
		strings.ToUpper("Get"),
		"/api/v1/invoices/search/{searchKey}",
		SearchInvoices,
	},

	Route{
		"TriggerEmailNotificationForInvoice",
		strings.ToUpper("Post"),
		"/api/v1/invoices/{invoiceId}/emailNotifications",
		TriggerEmailNotificationForInvoice,
	},

	Route{
		"UploadCatalogTranslation",
		strings.ToUpper("Post"),
		"/api/v1/invoices/catalogTranslation/{locale}",
		UploadCatalogTranslation,
	},

	Route{
		"UploadInvoiceMPTemplate",
		strings.ToUpper("Post"),
		"/api/v1/invoices/manualPayTemplate",
		UploadInvoiceMPTemplate,
	},

	Route{
		"UploadInvoiceTemplate",
		strings.ToUpper("Post"),
		"/api/v1/invoices/template",
		UploadInvoiceTemplate,
	},

	Route{
		"UploadInvoiceTranslation",
		strings.ToUpper("Post"),
		"/api/v1/invoices/translation/{locale}",
		UploadInvoiceTranslation,
	},

	Route{
		"GetNodesInfo",
		strings.ToUpper("Get"),
		"/api/v1/nodesInfo",
		GetNodesInfo,
	},

	Route{
		"TriggerNodeCommand",
		strings.ToUpper("Post"),
		"/api/v1/nodesInfo",
		TriggerNodeCommand,
	},

	Route{
		"GetOverdueConfigJson",
		strings.ToUpper("Get"),
		"/api/v1/overdue",
		GetOverdueConfigJson,
	},

	Route{
		"UploadOverdueConfigJson",
		strings.ToUpper("Post"),
		"/api/v1/overdue",
		UploadOverdueConfigJson,
	},

	Route{
		"BuildComboFormDescriptor",
		strings.ToUpper("Post"),
		"/api/v1/paymentGateways/hosted/form",
		BuildComboFormDescriptor,
	},

	Route{
		"BuildFormDescriptor",
		strings.ToUpper("Post"),
		"/api/v1/paymentGateways/hosted/form/{accountId}",
		BuildFormDescriptor,
	},

	Route{
		"ProcessNotification",
		strings.ToUpper("Post"),
		"/api/v1/paymentGateways/notification/{pluginName}",
		ProcessNotification,
	},

	Route{
		"CreateCustomFields",
		strings.ToUpper("Post"),
		"/api/v1/paymentMethods/{paymentMethodId}/customFields",
		CreateCustomFields,
	},

	Route{
		"DeleteCustomFields",
		strings.ToUpper("Delete"),
		"/api/v1/paymentMethods/{paymentMethodId}/customFields",
		DeleteCustomFields,
	},

	Route{
		"DeletePaymentMethod",
		strings.ToUpper("Delete"),
		"/api/v1/paymentMethods/{paymentMethodId}",
		DeletePaymentMethod,
	},

	Route{
		"GetCustomFields",
		strings.ToUpper("Get"),
		"/api/v1/paymentMethods/{paymentMethodId}/customFields",
		GetCustomFields,
	},

	Route{
		"GetPaymentMethod",
		strings.ToUpper("Get"),
		"/api/v1/paymentMethods/{paymentMethodId}",
		GetPaymentMethod,
	},

	Route{
		"GetPaymentMethodByKey",
		strings.ToUpper("Get"),
		"/api/v1/paymentMethods",
		GetPaymentMethodByKey,
	},

	Route{
		"GetPaymentMethods",
		strings.ToUpper("Get"),
		"/api/v1/paymentMethods/pagination",
		GetPaymentMethods,
	},

	Route{
		"SearchPaymentMethods",
		strings.ToUpper("Get"),
		"/api/v1/paymentMethods/search/{searchKey}",
		SearchPaymentMethods,
	},

	Route{
		"CreateCustomFields",
		strings.ToUpper("Post"),
		"/api/v1/paymentTransactions/{transactionId}/customFields",
		CreateCustomFields,
	},

	Route{
		"CreateTags",
		strings.ToUpper("Post"),
		"/api/v1/paymentTransactions/{transactionId}/tags",
		CreateTags,
	},

	Route{
		"DeleteCustomFields",
		strings.ToUpper("Delete"),
		"/api/v1/paymentTransactions/{transactionId}/customFields",
		DeleteCustomFields,
	},

	Route{
		"DeleteTags",
		strings.ToUpper("Delete"),
		"/api/v1/paymentTransactions/{transactionId}/tags",
		DeleteTags,
	},

	Route{
		"GetCustomFields",
		strings.ToUpper("Get"),
		"/api/v1/paymentTransactions/{transactionId}/customFields",
		GetCustomFields,
	},

	Route{
		"GetPaymentByTransactionId",
		strings.ToUpper("Get"),
		"/api/v1/paymentTransactions/{transactionId}",
		GetPaymentByTransactionId,
	},

	Route{
		"GetTags",
		strings.ToUpper("Get"),
		"/api/v1/paymentTransactions/{transactionId}/tags",
		GetTags,
	},

	Route{
		"NotifyStateChanged",
		strings.ToUpper("Post"),
		"/api/v1/paymentTransactions/{transactionId}",
		NotifyStateChanged,
	},

	Route{
		"CancelScheduledPaymentTransactionByExternalKey",
		strings.ToUpper("Delete"),
		"/api/v1/payments/cancelScheduledPaymentTransaction",
		CancelScheduledPaymentTransactionByExternalKey,
	},

	Route{
		"CancelScheduledPaymentTransactionById",
		strings.ToUpper("Delete"),
		"/api/v1/payments/{paymentTransactionId}/cancelScheduledPaymentTransaction",
		CancelScheduledPaymentTransactionById,
	},

	Route{
		"CaptureAuthorization",
		strings.ToUpper("Post"),
		"/api/v1/payments/{paymentId}",
		CaptureAuthorization,
	},

	Route{
		"CaptureAuthorizationByExternalKey",
		strings.ToUpper("Post"),
		"/api/v1/payments",
		CaptureAuthorizationByExternalKey,
	},

	Route{
		"ChargebackPayment",
		strings.ToUpper("Post"),
		"/api/v1/payments/{paymentId}/chargebacks",
		ChargebackPayment,
	},

	Route{
		"ChargebackPaymentByExternalKey",
		strings.ToUpper("Post"),
		"/api/v1/payments/chargebacks",
		ChargebackPaymentByExternalKey,
	},

	Route{
		"ChargebackReversalPayment",
		strings.ToUpper("Post"),
		"/api/v1/payments/{paymentId}/chargebackReversals",
		ChargebackReversalPayment,
	},

	Route{
		"ChargebackReversalPaymentByExternalKey",
		strings.ToUpper("Post"),
		"/api/v1/payments/chargebackReversals",
		ChargebackReversalPaymentByExternalKey,
	},

	Route{
		"CompleteTransaction",
		strings.ToUpper("Put"),
		"/api/v1/payments/{paymentId}",
		CompleteTransaction,
	},

	Route{
		"CompleteTransactionByExternalKey",
		strings.ToUpper("Put"),
		"/api/v1/payments",
		CompleteTransactionByExternalKey,
	},

	Route{
		"CreateComboPayment",
		strings.ToUpper("Post"),
		"/api/v1/payments/combo",
		CreateComboPayment,
	},

	Route{
		"CreateCustomFields",
		strings.ToUpper("Post"),
		"/api/v1/payments/{paymentId}/customFields",
		CreateCustomFields,
	},

	Route{
		"CreateTags",
		strings.ToUpper("Post"),
		"/api/v1/payments/{paymentId}/tags",
		CreateTags,
	},

	Route{
		"DeleteCustomFields",
		strings.ToUpper("Delete"),
		"/api/v1/payments/{paymentId}/customFields",
		DeleteCustomFields,
	},

	Route{
		"DeleteTags",
		strings.ToUpper("Delete"),
		"/api/v1/payments/{paymentId}/tags",
		DeleteTags,
	},

	Route{
		"GetCustomFields",
		strings.ToUpper("Get"),
		"/api/v1/payments/{paymentId}/customFields",
		GetCustomFields,
	},

	Route{
		"GetPayment",
		strings.ToUpper("Get"),
		"/api/v1/payments/{paymentId}",
		GetPayment,
	},

	Route{
		"GetPaymentByExternalKey",
		strings.ToUpper("Get"),
		"/api/v1/payments",
		GetPaymentByExternalKey,
	},

	Route{
		"GetPayments",
		strings.ToUpper("Get"),
		"/api/v1/payments/pagination",
		GetPayments,
	},

	Route{
		"GetTags",
		strings.ToUpper("Get"),
		"/api/v1/payments/{paymentId}/tags",
		GetTags,
	},

	Route{
		"RefundPayment",
		strings.ToUpper("Post"),
		"/api/v1/payments/{paymentId}/refunds",
		RefundPayment,
	},

	Route{
		"RefundPaymentByExternalKey",
		strings.ToUpper("Post"),
		"/api/v1/payments/refunds",
		RefundPaymentByExternalKey,
	},

	Route{
		"SearchPayments",
		strings.ToUpper("Get"),
		"/api/v1/payments/search/{searchKey}",
		SearchPayments,
	},

	Route{
		"VoidPayment",
		strings.ToUpper("Delete"),
		"/api/v1/payments/{paymentId}",
		VoidPayment,
	},

	Route{
		"VoidPaymentByExternalKey",
		strings.ToUpper("Delete"),
		"/api/v1/payments",
		VoidPaymentByExternalKey,
	},

	Route{
		"GetPluginsInfo",
		strings.ToUpper("Get"),
		"/api/v1/pluginsInfo",
		GetPluginsInfo,
	},

	Route{
		"AddRoleDefinition",
		strings.ToUpper("Post"),
		"/api/v1/security/roles",
		AddRoleDefinition,
	},

	Route{
		"AddUserRoles",
		strings.ToUpper("Post"),
		"/api/v1/security/users",
		AddUserRoles,
	},

	Route{
		"GetCurrentUserPermissions",
		strings.ToUpper("Get"),
		"/api/v1/security/permissions",
		GetCurrentUserPermissions,
	},

	Route{
		"GetCurrentUserSubject",
		strings.ToUpper("Get"),
		"/api/v1/security/subject",
		GetCurrentUserSubject,
	},

	Route{
		"GetUserRoles",
		strings.ToUpper("Get"),
		"/api/v1/security/users/{username}/roles",
		GetUserRoles,
	},

	Route{
		"InvalidateUser",
		strings.ToUpper("Delete"),
		"/api/v1/security/users/{username}",
		InvalidateUser,
	},

	Route{
		"UpdateUserPassword",
		strings.ToUpper("Put"),
		"/api/v1/security/users/{username}/password",
		UpdateUserPassword,
	},

	Route{
		"UpdateUserRoles",
		strings.ToUpper("Put"),
		"/api/v1/security/users/{username}/roles",
		UpdateUserRoles,
	},

	Route{
		"AddSubscriptionBlockingState",
		strings.ToUpper("Put"),
		"/api/v1/subscriptions/{subscriptionId}/block",
		AddSubscriptionBlockingState,
	},

	Route{
		"CancelEntitlementPlan",
		strings.ToUpper("Delete"),
		"/api/v1/subscriptions/{subscriptionId}",
		CancelEntitlementPlan,
	},

	Route{
		"ChangeEntitlementPlan",
		strings.ToUpper("Put"),
		"/api/v1/subscriptions/{subscriptionId}",
		ChangeEntitlementPlan,
	},

	Route{
		"CreateCustomFields",
		strings.ToUpper("Post"),
		"/api/v1/subscriptions/{subscriptionId}/customFields",
		CreateCustomFields,
	},

	Route{
		"CreateEntitlement",
		strings.ToUpper("Post"),
		"/api/v1/subscriptions",
		CreateEntitlement,
	},

	Route{
		"CreateEntitlementWithAddOns",
		strings.ToUpper("Post"),
		"/api/v1/subscriptions/createEntitlementWithAddOns",
		CreateEntitlementWithAddOns,
	},

	Route{
		"CreateEntitlementsWithAddOns",
		strings.ToUpper("Post"),
		"/api/v1/subscriptions/createEntitlementsWithAddOns",
		CreateEntitlementsWithAddOns,
	},

	Route{
		"CreateTags",
		strings.ToUpper("Post"),
		"/api/v1/subscriptions/{subscriptionId}/tags",
		CreateTags,
	},

	Route{
		"DeleteCustomFields",
		strings.ToUpper("Delete"),
		"/api/v1/subscriptions/{subscriptionId}/customFields",
		DeleteCustomFields,
	},

	Route{
		"DeleteTags",
		strings.ToUpper("Delete"),
		"/api/v1/subscriptions/{subscriptionId}/tags",
		DeleteTags,
	},

	Route{
		"GetCustomFields",
		strings.ToUpper("Get"),
		"/api/v1/subscriptions/{subscriptionId}/customFields",
		GetCustomFields,
	},

	Route{
		"GetEntitlement",
		strings.ToUpper("Get"),
		"/api/v1/subscriptions/{subscriptionId}",
		GetEntitlement,
	},

	Route{
		"GetTags",
		strings.ToUpper("Get"),
		"/api/v1/subscriptions/{subscriptionId}/tags",
		GetTags,
	},

	Route{
		"UncancelEntitlementPlan",
		strings.ToUpper("Put"),
		"/api/v1/subscriptions/{subscriptionId}/uncancel",
		UncancelEntitlementPlan,
	},

	Route{
		"UpdateSubscriptionBCD",
		strings.ToUpper("Put"),
		"/api/v1/subscriptions/{subscriptionId}/bcd",
		UpdateSubscriptionBCD,
	},

	Route{
		"CreateTagDefinition",
		strings.ToUpper("Post"),
		"/api/v1/tagDefinitions",
		CreateTagDefinition,
	},

	Route{
		"DeleteTagDefinition",
		strings.ToUpper("Delete"),
		"/api/v1/tagDefinitions/{tagDefinitionId}",
		DeleteTagDefinition,
	},

	Route{
		"GetTagDefinition",
		strings.ToUpper("Get"),
		"/api/v1/tagDefinitions/{tagDefinitionId}",
		GetTagDefinition,
	},

	Route{
		"GetTagDefinitions",
		strings.ToUpper("Get"),
		"/api/v1/tagDefinitions",
		GetTagDefinitions,
	},

	Route{
		"GetTags",
		strings.ToUpper("Get"),
		"/api/v1/tags/pagination",
		GetTags,
	},

	Route{
		"SearchTags",
		strings.ToUpper("Get"),
		"/api/v1/tags/search/{searchKey}",
		SearchTags,
	},

	Route{
		"CreateTenant",
		strings.ToUpper("Post"),
		"/api/v1/tenants",
		CreateTenant,
	},

	Route{
		"DeletePerTenantConfiguration",
		strings.ToUpper("Delete"),
		"/api/v1/tenants/uploadPerTenantConfig",
		DeletePerTenantConfiguration,
	},

	Route{
		"DeletePluginConfiguration",
		strings.ToUpper("Delete"),
		"/api/v1/tenants/uploadPluginConfig/{pluginName}",
		DeletePluginConfiguration,
	},

	Route{
		"DeletePluginPaymentStateMachineConfig",
		strings.ToUpper("Delete"),
		"/api/v1/tenants/uploadPluginPaymentStateMachineConfig/{pluginName}",
		DeletePluginPaymentStateMachineConfig,
	},

	Route{
		"DeletePushNotificationCallbacks",
		strings.ToUpper("Delete"),
		"/api/v1/tenants/registerNotificationCallback",
		DeletePushNotificationCallbacks,
	},

	Route{
		"DeleteUserKeyValue",
		strings.ToUpper("Delete"),
		"/api/v1/tenants/userKeyValue/{keyName}",
		DeleteUserKeyValue,
	},

	Route{
		"GetAllPluginConfiguration",
		strings.ToUpper("Get"),
		"/api/v1/tenants/uploadPerTenantConfig/{keyPrefix}/search",
		GetAllPluginConfiguration,
	},

	Route{
		"GetPerTenantConfiguration",
		strings.ToUpper("Get"),
		"/api/v1/tenants/uploadPerTenantConfig",
		GetPerTenantConfiguration,
	},

	Route{
		"GetPluginConfiguration",
		strings.ToUpper("Get"),
		"/api/v1/tenants/uploadPluginConfig/{pluginName}",
		GetPluginConfiguration,
	},

	Route{
		"GetPluginPaymentStateMachineConfig",
		strings.ToUpper("Get"),
		"/api/v1/tenants/uploadPluginPaymentStateMachineConfig/{pluginName}",
		GetPluginPaymentStateMachineConfig,
	},

	Route{
		"GetPushNotificationCallbacks",
		strings.ToUpper("Get"),
		"/api/v1/tenants/registerNotificationCallback",
		GetPushNotificationCallbacks,
	},

	Route{
		"GetTenant",
		strings.ToUpper("Get"),
		"/api/v1/tenants/{tenantId}",
		GetTenant,
	},

	Route{
		"GetTenantByApiKey",
		strings.ToUpper("Get"),
		"/api/v1/tenants",
		GetTenantByApiKey,
	},

	Route{
		"GetUserKeyValue",
		strings.ToUpper("Get"),
		"/api/v1/tenants/userKeyValue/{keyName}",
		GetUserKeyValue,
	},

	Route{
		"InsertUserKeyValue",
		strings.ToUpper("Post"),
		"/api/v1/tenants/userKeyValue/{keyName}",
		InsertUserKeyValue,
	},

	Route{
		"RegisterPushNotificationCallback",
		strings.ToUpper("Post"),
		"/api/v1/tenants/registerNotificationCallback",
		RegisterPushNotificationCallback,
	},

	Route{
		"UploadPerTenantConfiguration",
		strings.ToUpper("Post"),
		"/api/v1/tenants/uploadPerTenantConfig",
		UploadPerTenantConfiguration,
	},

	Route{
		"UploadPluginConfiguration",
		strings.ToUpper("Post"),
		"/api/v1/tenants/uploadPluginConfig/{pluginName}",
		UploadPluginConfiguration,
	},

	Route{
		"UploadPluginPaymentStateMachineConfig",
		strings.ToUpper("Post"),
		"/api/v1/tenants/uploadPluginPaymentStateMachineConfig/{pluginName}",
		UploadPluginPaymentStateMachineConfig,
	},

	Route{
		"GetAllUsage",
		strings.ToUpper("Get"),
		"/api/v1/usages/{subscriptionId}",
		GetAllUsage,
	},

	Route{
		"GetUsage",
		strings.ToUpper("Get"),
		"/api/v1/usages/{subscriptionId}/{unitType}",
		GetUsage,
	},

	Route{
		"RecordUsage",
		strings.ToUpper("Post"),
		"/api/v1/usages",
		RecordUsage,
	},
}
