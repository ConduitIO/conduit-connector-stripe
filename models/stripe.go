// Copyright © 2022 Meroxa, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package models

import (
	"strings"

	"github.com/conduitio/conduit-connector-stripe/models/resources"
)

// A ResourceResponse represents a response resource data from Stripe.
type ResourceResponse struct {
	Data    []map[string]interface{} `json:"data"`
	HasMore bool                     `json:"has_more"`
}

// A EventResponse represents a response event data from Stripe.
type EventResponse struct {
	Data    []EventData `json:"data"`
	HasMore bool        `json:"has_more"`
}

// A EventData represents a data of event response.
type EventData struct {
	ID      string          `json:"id"`
	Created int64           `json:"created"`
	Data    EventDataObject `json:"data"`
	Type    string          `json:"type"`
}

// An EventDataObject represents an object of event's data.
// haris: is it a full object, or just the changed values?
type EventDataObject struct {
	Object map[string]interface{} `json:"object"`
}

// ResourcesMap represents a dictionary with valid resources,
// where the key is the object type and the value is the name
// of the API endpoints of that object type.
var ResourcesMap = map[string]string{
	resources.CreditNoteResource:                  resources.CreditNotesList,
	resources.BillingPortalConfigurationResource:  resources.BillingPortalConfigurationsList,
	resources.InvoiceResource:                     resources.InvoicesList,
	resources.InvoiceItemResource:                 resources.InvoiceItemsList,
	resources.PlanResource:                        resources.PlanList,
	resources.QuoteResource:                       resources.QuotesList,
	resources.SubscriptionResource:                resources.SubscriptionsList,
	resources.SubscriptionItemResource:            resources.SubscriptionItemsList,
	resources.SubscriptionScheduleResource:        resources.SubscriptionSchedulesList,
	resources.CheckoutSessionResource:             resources.CheckoutSessionsList,
	resources.AccountResource:                     resources.AccountsList,
	resources.ApplicationFeeResource:              resources.ApplicationFeesList,
	resources.CountrySpecResource:                 resources.CountrySpecsList,
	resources.TopUpResource:                       resources.TopUpsList,
	resources.TransferResource:                    resources.TransfersList,
	resources.BalanceTransactionResource:          resources.BalanceTransactionsList,
	resources.ChargeResource:                      resources.ChargesList,
	resources.CustomerResource:                    resources.CustomersList,
	resources.DisputeResource:                     resources.DisputesList,
	resources.EventResource:                       resources.EventsList,
	resources.FileResource:                        resources.FilesList,
	resources.FileLinkResource:                    resources.FileLinksList,
	resources.PaymentIntentResource:               resources.PaymentIntentsList,
	resources.SetupIntentResource:                 resources.SetupIntentsList,
	resources.SetupAttemptResource:                resources.SetupAttemptsList,
	resources.PayoutResource:                      resources.PayoutsList,
	resources.RefundResource:                      resources.RefundsList,
	resources.FinancialConnectionsAccountResource: resources.FinancialConnectionsAccountsResource,
	resources.RadarEarlyFraudWarningResource:      resources.RadarEarlyFraudWarningsList,
	resources.ReviewResource:                      resources.ReviewsList,
	resources.RadarValueListResource:              resources.RadarValueListsList,
	resources.RadarValueListItemResource:          resources.RadarValueListItemsList,
	resources.IdentityVerificationSessionResource: resources.IdentityVerificationSessionsList,
	resources.IdentityVerificationReportResource:  resources.IdentityVerificationReportsList,
	resources.IssuingAuthorizationResource:        resources.IssuingAuthorizationsList,
	resources.IssuingCardholderResource:           resources.IssuingCardholdersList,
	resources.IssuingCardResource:                 resources.IssuingCardsList,
	resources.IssuingDisputeResource:              resources.IssuingDisputesList,
	resources.FundingInstructionResource:          resources.FundingInstructionsList,
	resources.IssuingTransactionResource:          resources.IssuingTransactionsList,
	resources.OrderResource:                       resources.OrdersList,
	resources.PaymentLinkResource:                 resources.PaymentLinksList,
	resources.PaymentMethodResource:               resources.PaymentMethodsList,
	resources.ProductResource:                     resources.ProductsList,
	resources.PriceResource:                       resources.PricesList,
	resources.CouponResource:                      resources.CouponsList,
	resources.PromotionCodeResource:               resources.PromotionCodesList,
	resources.TaxCodeResource:                     resources.TaxCodesList,
	resources.TaxRateResource:                     resources.TaxRatesList,
	resources.ShippingRateResource:                resources.ShippingRatesList,
	resources.ReportingReportRunResource:          resources.ReportingReportRunsList,
	resources.ReportingReportTypeResource:         resources.ReportingReportTypesList,
	resources.ScheduledQueryRunResource:           resources.ScheduledQueryRunsList,
	resources.TerminalLocationResource:            resources.TerminalLocationsList,
	resources.TerminalReaderResource:              resources.TerminalReadersList,
	resources.TerminalConfigurationResource:       resources.TerminalConfigurationsList,
	resources.FinancialAccountResource:            resources.FinancialAccountsList,
	resources.TransactionResource:                 resources.TransactionsList,
	resources.TransactionEntryResource:            resources.TransactionEntriesList,
	resources.OutboundTransferResource:            resources.OutboundTransfersList,
	resources.OutboundPaymentResource:             resources.OutboundPaymentsList,
	resources.InboundTransferResource:             resources.InboundTransfersList,
	resources.ReceivedCreditResource:              resources.ReceivedCreditsList,
	resources.ReceivedDebitResource:               resources.ReceivedDebitsList,
	resources.CreditReversalResource:              resources.CreditReversalsList,
	resources.DebitReversalResource:               resources.DebitReversalsList,
	resources.WebhookEndpointResource:             resources.WebhookEndpointsList,
}

// EventsMap represents a dictionary with all events in each resource,
// where the key is the resource and the value is a slice of events.
var EventsMap = map[string][]string{
	resources.CreditNoteResource:                  resources.CreditNoteEvents,
	resources.InvoiceResource:                     resources.InvoiceEvents,
	resources.InvoiceItemResource:                 resources.InvoiceItemEvents,
	resources.PlanResource:                        resources.PlanEvents,
	resources.QuoteResource:                       resources.QuoteEvents,
	resources.SubscriptionResource:                resources.SubscriptionEvents,
	resources.SubscriptionScheduleResource:        resources.SubscriptionScheduleEvents,
	resources.CheckoutSessionResource:             resources.CheckoutSessionEvents,
	resources.AccountResource:                     resources.AccountEvents,
	resources.ApplicationFeeResource:              resources.ApplicationFeeEvents,
	resources.TopUpResource:                       resources.TopUpEvents,
	resources.TransferResource:                    resources.TransferEvents,
	resources.ChargeResource:                      resources.ChargeEvents,
	resources.CustomerResource:                    resources.CustomerEvents,
	resources.DisputeResource:                     resources.DisputeEvents,
	resources.FileResource:                        resources.FileEvents,
	resources.PaymentIntentResource:               resources.PaymentIntentEvents,
	resources.SetupIntentResource:                 resources.SetupIntentEvents,
	resources.PayoutResource:                      resources.PayoutEvents,
	resources.RefundResource:                      resources.RefundEvents,
	resources.RadarEarlyFraudWarningResource:      resources.RadarEarlyFraudWarningEvents,
	resources.ReviewResource:                      resources.ReviewEvents,
	resources.IdentityVerificationSessionResource: resources.IdentityVerificationSessionEvents,
	resources.IssuingAuthorizationResource:        resources.IssuingAuthorizationEvents,
	resources.IssuingCardholderResource:           resources.IssuingCardholderEvents,
	resources.IssuingCardResource:                 resources.IssuingCardEvents,
	resources.IssuingDisputeResource:              resources.IssuingDisputeEvents,
	resources.IssuingTransactionResource:          resources.IssuingTransactionEvents,
	resources.OrderResource:                       resources.OrderEvents,
	resources.PaymentLinkResource:                 resources.PaymentLinkEvents,
	resources.PaymentMethodResource:               resources.PaymentMethodEvents,
	resources.ProductResource:                     resources.ProductEvents,
	resources.PriceResource:                       resources.PriceEvents,
	resources.CouponResource:                      resources.CouponEvents,
	resources.PromotionCodeResource:               resources.PromotionCodeEvents,
	resources.TaxRateResource:                     resources.TaxRateEvents,
	resources.ReportingReportRunResource:          resources.ReportingReportRunEvents,
	resources.ReportingReportTypeResource:         resources.ReportingReportTypeEvents,
	resources.ScheduledQueryRunResource:           resources.ScheduledQueryRunEvents,
	resources.TerminalReaderResource:              resources.TerminalReaderEvents,
}

// EventsAction represents a dictionary with actions of events,
// where the key is an event and the value is an action.
var EventsAction = (func() map[string]string {
	eventsAction := make(map[string]string)

	for _, events := range EventsMap {
		for _, event := range events {
			switch e := event; {
			case strings.Contains(e, eventKeyCreated):
				eventsAction[e] = InsertAction
			case strings.Contains(e, eventKeyDeleted):
				eventsAction[e] = DeleteAction
			default:
				eventsAction[e] = UpdateAction
			}
		}
	}

	return eventsAction
})()
