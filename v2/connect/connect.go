// Copyright 2019 James Cote
// All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Code generated by gen-esign; DO NOT EDIT.

// Package connect implements the DocuSign SDK
// category Connect.
//
// The Connect category enables your application to be called via HTTPS when an event of interest occurs.
//
// Use the Connect service to "end the polling madness." With Connect, there is no need for your application to poll DocuSign every 15 minutes to learn the latest about your envelopes.
//
// Instead, you register your interest in one or more types of envelope or recipient events. Then, when an interesting event occurs, the DocuSign platform will contact your application with the event's details and data. You can register interest in envelopes sent by particular users in your account, or for envelopes sent by any user.
//
// ## Incoming Connect Calls
// To use the Connect service, your application needs to provide an https url that can be called from the public internet. If your application runs on a server behind your organization's firewall, then you will need to create a "pinhole" in the firewall to allow the incoming Connect calls from DocuSign to reach your application. Other techniques for receiving the incoming calls including proxy servers and DMZ networking can also be used.
//
// ## Per-envelope Connect Configuration
// Instead of registering a general Connect configuration and listener, an individual envelope can have its own Connect configuration. See the `eventNotification` field for envelopes.
//
// ## Categories
// Use the Connect category for:
//
// * Programmatically creating Connect configurations. Connect configurations can be created manually by using the DocuSign web service, or programmatically via the API. Configurations created via the API can be seen and updated from the web service.
// * Retrieving and managing the event log for your Connect configurations.
// * Requesting that an event be re-published to the listener.
//
// Service Api documentation may be found at:
// https://developers.docusign.com/esign-rest-api/v2/reference/Connect
// Usage example:
//
//   import (
//       "github.com/ICGGroup/esign"
//       "github.com/ICGGroup/esign/v2/connect"
//       "github.com/ICGGroup/esign/v2/model"
//   )
//   ...
//   connectService := connect.New(esignCredential)
package connect // import "github.com/ICGGroup/esign/v2/connect"

import (
	"context"
	"fmt"
	"io"
	"net/url"
	"strings"
	"time"

	"github.com/ICGGroup/esign"
	"github.com/ICGGroup/esign/v2/model"
)

// Service implements DocuSign Connect Category API operations
type Service struct {
	credential esign.Credential
}

// New initializes a connect service using cred to authorize ops.
func New(cred esign.Credential) *Service {
	return &Service{credential: cred}
}

// ConfigurationsCreate creates a connect configuration for the specified account.
//
// https://developers.docusign.com/esign-rest-api/v2/reference/connect/connectconfigurations/create
//
// SDK Method Connect::createConfiguration
func (s *Service) ConfigurationsCreate(connectConfigurations *model.ConnectCustomConfiguration) *ConfigurationsCreateOp {
	return &ConfigurationsCreateOp{
		Credential: s.credential,
		Method:     "POST",
		Path:       "connect",
		Payload:    connectConfigurations,
		QueryOpts:  make(url.Values),
	}
}

// ConfigurationsCreateOp implements DocuSign API SDK Connect::createConfiguration
type ConfigurationsCreateOp esign.Op

// Do executes the op.  A nil context will return error.
func (op *ConfigurationsCreateOp) Do(ctx context.Context) (*model.ConnectCustomConfiguration, error) {
	var res *model.ConnectCustomConfiguration
	return res, ((*esign.Op)(op)).Do(ctx, &res)
}

// ConfigurationsDelete deletes the specified connect configuration.
//
// https://developers.docusign.com/esign-rest-api/v2/reference/connect/connectconfigurations/delete
//
// SDK Method Connect::deleteConfiguration
func (s *Service) ConfigurationsDelete(connectID string) *ConfigurationsDeleteOp {
	return &ConfigurationsDeleteOp{
		Credential: s.credential,
		Method:     "DELETE",
		Path:       strings.Join([]string{"connect", connectID}, "/"),
		QueryOpts:  make(url.Values),
	}
}

// ConfigurationsDeleteOp implements DocuSign API SDK Connect::deleteConfiguration
type ConfigurationsDeleteOp esign.Op

// Do executes the op.  A nil context will return error.
func (op *ConfigurationsDeleteOp) Do(ctx context.Context) error {
	return ((*esign.Op)(op)).Do(ctx, nil)
}

// ConfigurationsGet get information on a Connect Configuration
//
// https://developers.docusign.com/esign-rest-api/v2/reference/connect/connectconfigurations/get
//
// SDK Method Connect::getConfiguration
func (s *Service) ConfigurationsGet(connectID string) *ConfigurationsGetOp {
	return &ConfigurationsGetOp{
		Credential: s.credential,
		Method:     "GET",
		Path:       strings.Join([]string{"connect", connectID}, "/"),
		QueryOpts:  make(url.Values),
	}
}

// ConfigurationsGetOp implements DocuSign API SDK Connect::getConfiguration
type ConfigurationsGetOp esign.Op

// Do executes the op.  A nil context will return error.
func (op *ConfigurationsGetOp) Do(ctx context.Context) (*model.ConnectConfigResults, error) {
	var res *model.ConnectConfigResults
	return res, ((*esign.Op)(op)).Do(ctx, &res)
}

// ConfigurationsList get Connect Configuration Information
//
// https://developers.docusign.com/esign-rest-api/v2/reference/connect/connectconfigurations/list
//
// SDK Method Connect::listConfigurations
func (s *Service) ConfigurationsList() *ConfigurationsListOp {
	return &ConfigurationsListOp{
		Credential: s.credential,
		Method:     "GET",
		Path:       "connect",
		QueryOpts:  make(url.Values),
	}
}

// ConfigurationsListOp implements DocuSign API SDK Connect::listConfigurations
type ConfigurationsListOp esign.Op

// Do executes the op.  A nil context will return error.
func (op *ConfigurationsListOp) Do(ctx context.Context) (*model.ConnectConfigResults, error) {
	var res *model.ConnectConfigResults
	return res, ((*esign.Op)(op)).Do(ctx, &res)
}

// ConfigurationsListUsers returns users from the configured Connect service.
//
// https://developers.docusign.com/esign-rest-api/v2/reference/connect/connectconfigurations/listusers
//
// SDK Method Connect::connectUsers
func (s *Service) ConfigurationsListUsers(connectID string) *ConfigurationsListUsersOp {
	return &ConfigurationsListUsersOp{
		Credential: s.credential,
		Method:     "GET",
		Path:       strings.Join([]string{"connect", connectID, "users"}, "/"),
		QueryOpts:  make(url.Values),
	}
}

// ConfigurationsListUsersOp implements DocuSign API SDK Connect::connectUsers
type ConfigurationsListUsersOp esign.Op

// Do executes the op.  A nil context will return error.
func (op *ConfigurationsListUsersOp) Do(ctx context.Context) (*model.IntegratedUserInfoList, error) {
	var res *model.IntegratedUserInfoList
	return res, ((*esign.Op)(op)).Do(ctx, &res)
}

// Count optional. Number of items to return.
func (op *ConfigurationsListUsersOp) Count(val int) *ConfigurationsListUsersOp {
	if op != nil {
		op.QueryOpts.Set("count", fmt.Sprintf("%d", val))
	}
	return op
}

// EmailSubstring filters the returned user records by the email address or a sub-string of email address.
func (op *ConfigurationsListUsersOp) EmailSubstring(val string) *ConfigurationsListUsersOp {
	if op != nil {
		op.QueryOpts.Set("email_substring", val)
	}
	return op
}

// ListIncludedUsers set the call query parameter list_included_users
func (op *ConfigurationsListUsersOp) ListIncludedUsers() *ConfigurationsListUsersOp {
	if op != nil {
		op.QueryOpts.Set("list_included_users", "true")
	}
	return op
}

// StartPosition is the position within the total result set from which to start returning values. The value **thumbnail** may be used to return the page image.
func (op *ConfigurationsListUsersOp) StartPosition(val int) *ConfigurationsListUsersOp {
	if op != nil {
		op.QueryOpts.Set("start_position", fmt.Sprintf("%d", val))
	}
	return op
}

// Status filters the results by user status.
// You can specify a comma-separated
// list of the following statuses:
//
// * ActivationRequired
// * ActivationSent
// * Active
// * Closed
// * Disabled
func (op *ConfigurationsListUsersOp) Status(val ...string) *ConfigurationsListUsersOp {
	if op != nil {
		op.QueryOpts.Set("status", strings.Join(val, ","))
	}
	return op
}

// UserNameSubstring filters the user records returned by the user name or a sub-string of user name.
func (op *ConfigurationsListUsersOp) UserNameSubstring(val string) *ConfigurationsListUsersOp {
	if op != nil {
		op.QueryOpts.Set("user_name_substring", val)
	}
	return op
}

// ConfigurationsUpdate updates a specified Connect configuration.
//
// https://developers.docusign.com/esign-rest-api/v2/reference/connect/connectconfigurations/update
//
// SDK Method Connect::updateConfiguration
func (s *Service) ConfigurationsUpdate(connectConfigurations *model.ConnectCustomConfiguration) *ConfigurationsUpdateOp {
	return &ConfigurationsUpdateOp{
		Credential: s.credential,
		Method:     "PUT",
		Path:       "connect",
		Payload:    connectConfigurations,
		QueryOpts:  make(url.Values),
	}
}

// ConfigurationsUpdateOp implements DocuSign API SDK Connect::updateConfiguration
type ConfigurationsUpdateOp esign.Op

// Do executes the op.  A nil context will return error.
func (op *ConfigurationsUpdateOp) Do(ctx context.Context) (*model.ConnectCustomConfiguration, error) {
	var res *model.ConnectCustomConfiguration
	return res, ((*esign.Op)(op)).Do(ctx, &res)
}

// EventsDelete deletes a specified Connect log entry.
//
// https://developers.docusign.com/esign-rest-api/v2/reference/connect/connectevents/delete
//
// SDK Method Connect::deleteEventLog
func (s *Service) EventsDelete(logID string) *EventsDeleteOp {
	return &EventsDeleteOp{
		Credential: s.credential,
		Method:     "DELETE",
		Path:       strings.Join([]string{"connect", "logs", logID}, "/"),
		QueryOpts:  make(url.Values),
	}
}

// EventsDeleteOp implements DocuSign API SDK Connect::deleteEventLog
type EventsDeleteOp esign.Op

// Do executes the op.  A nil context will return error.
func (op *EventsDeleteOp) Do(ctx context.Context) error {
	return ((*esign.Op)(op)).Do(ctx, nil)
}

// EventsDeleteFailure deletes a Connect failure log entry.
//
// https://developers.docusign.com/esign-rest-api/v2/reference/connect/connectevents/deletefailure
//
// SDK Method Connect::deleteEventFailureLog
func (s *Service) EventsDeleteFailure(failureID string) *EventsDeleteFailureOp {
	return &EventsDeleteFailureOp{
		Credential: s.credential,
		Method:     "DELETE",
		Path:       strings.Join([]string{"connect", "failures", failureID}, "/"),
		QueryOpts:  make(url.Values),
	}
}

// EventsDeleteFailureOp implements DocuSign API SDK Connect::deleteEventFailureLog
type EventsDeleteFailureOp esign.Op

// Do executes the op.  A nil context will return error.
func (op *EventsDeleteFailureOp) Do(ctx context.Context) error {
	return ((*esign.Op)(op)).Do(ctx, nil)
}

// EventsDeleteList gets a list of Connect log entries.
//
// https://developers.docusign.com/esign-rest-api/v2/reference/connect/connectevents/deletelist
//
// SDK Method Connect::deleteEventLogs
func (s *Service) EventsDeleteList() *EventsDeleteListOp {
	return &EventsDeleteListOp{
		Credential: s.credential,
		Method:     "DELETE",
		Path:       "connect/logs",
		QueryOpts:  make(url.Values),
	}
}

// EventsDeleteListOp implements DocuSign API SDK Connect::deleteEventLogs
type EventsDeleteListOp esign.Op

// Do executes the op.  A nil context will return error.
func (op *EventsDeleteListOp) Do(ctx context.Context) error {
	return ((*esign.Op)(op)).Do(ctx, nil)
}

// EventsGet get the specified Connect log entry.
//
// https://developers.docusign.com/esign-rest-api/v2/reference/connect/connectevents/get
//
// SDK Method Connect::getEventLog
func (s *Service) EventsGet(logID string) *EventsGetOp {
	return &EventsGetOp{
		Credential: s.credential,
		Method:     "GET",
		Path:       strings.Join([]string{"connect", "logs", logID}, "/"),
		QueryOpts:  make(url.Values),
	}
}

// EventsGetOp implements DocuSign API SDK Connect::getEventLog
type EventsGetOp esign.Op

// Do executes the op.  A nil context will return error.
func (op *EventsGetOp) Do(ctx context.Context) (*model.ConnectLog, error) {
	var res *model.ConnectLog
	return res, ((*esign.Op)(op)).Do(ctx, &res)
}

// AdditionalInfo when true, the connectDebugLog information is included in the response.
func (op *EventsGetOp) AdditionalInfo() *EventsGetOp {
	if op != nil {
		op.QueryOpts.Set("additional_info", "true")
	}
	return op
}

// EventsList gets the Connect log.
//
// https://developers.docusign.com/esign-rest-api/v2/reference/connect/connectevents/list
//
// SDK Method Connect::listEventLogs
func (s *Service) EventsList() *EventsListOp {
	return &EventsListOp{
		Credential: s.credential,
		Method:     "GET",
		Path:       "connect/logs",
		QueryOpts:  make(url.Values),
	}
}

// EventsListOp implements DocuSign API SDK Connect::listEventLogs
type EventsListOp esign.Op

// Do executes the op.  A nil context will return error.
func (op *EventsListOp) Do(ctx context.Context) (*model.ConnectLogs, error) {
	var res *model.ConnectLogs
	return res, ((*esign.Op)(op)).Do(ctx, &res)
}

// FromDate start of the search date range. Only returns templates created on or after this date/time. If no value is specified, there is no limit on the earliest date created.
func (op *EventsListOp) FromDate(val time.Time) *EventsListOp {
	if op != nil {
		op.QueryOpts.Set("from_date", val.Format(time.RFC3339))
	}
	return op
}

// ToDate end of the search date range. Only returns templates created up to this date/time. If no value is provided, this defaults to the current date.
func (op *EventsListOp) ToDate(val time.Time) *EventsListOp {
	if op != nil {
		op.QueryOpts.Set("to_date", val.Format(time.RFC3339))
	}
	return op
}

// EventsListFailures gets the Connect failure log information.
//
// https://developers.docusign.com/esign-rest-api/v2/reference/connect/connectevents/listfailures
//
// SDK Method Connect::listEventFailureLogs
func (s *Service) EventsListFailures() *EventsListFailuresOp {
	return &EventsListFailuresOp{
		Credential: s.credential,
		Method:     "GET",
		Path:       "connect/failures",
		QueryOpts:  make(url.Values),
	}
}

// EventsListFailuresOp implements DocuSign API SDK Connect::listEventFailureLogs
type EventsListFailuresOp esign.Op

// Do executes the op.  A nil context will return error.
func (op *EventsListFailuresOp) Do(ctx context.Context) (*model.ConnectLogs, error) {
	var res *model.ConnectLogs
	return res, ((*esign.Op)(op)).Do(ctx, &res)
}

// FromDate start of the search date range. Only returns templates created on or after this date/time. If no value is specified, there is no limit on the earliest date created.
func (op *EventsListFailuresOp) FromDate(val time.Time) *EventsListFailuresOp {
	if op != nil {
		op.QueryOpts.Set("from_date", val.Format(time.RFC3339))
	}
	return op
}

// ToDate end of the search date range. Only returns templates created up to this date/time. If no value is provided, this defaults to the current date.
func (op *EventsListFailuresOp) ToDate(val time.Time) *EventsListFailuresOp {
	if op != nil {
		op.QueryOpts.Set("to_date", val.Format(time.RFC3339))
	}
	return op
}

// EventsRetryForEnvelope republishes Connect information for the specified envelope.
// If media is an io.ReadCloser, Do() will close media.
//
// https://developers.docusign.com/esign-rest-api/v2/reference/connect/connectevents/retryforenvelope
//
// SDK Method Connect::retryEventForEnvelope
func (s *Service) EventsRetryForEnvelope(envelopeID string, media io.Reader, mimeType string) *EventsRetryForEnvelopeOp {
	return &EventsRetryForEnvelopeOp{
		Credential: s.credential,
		Method:     "PUT",
		Path:       strings.Join([]string{"connect", "envelopes", envelopeID, "retry_queue"}, "/"),
		Payload:    &esign.UploadFile{Reader: media, ContentType: mimeType},
		QueryOpts:  make(url.Values),
	}
}

// EventsRetryForEnvelopeOp implements DocuSign API SDK Connect::retryEventForEnvelope
type EventsRetryForEnvelopeOp esign.Op

// Do executes the op.  A nil context will return error.
func (op *EventsRetryForEnvelopeOp) Do(ctx context.Context) (*model.ConnectFailureResults, error) {
	var res *model.ConnectFailureResults
	return res, ((*esign.Op)(op)).Do(ctx, &res)
}

// EventsRetryForEnvelopes republishes Connect information for multiple envelopes.
//
// https://developers.docusign.com/esign-rest-api/v2/reference/connect/connectevents/retryforenvelopes
//
// SDK Method Connect::retryEventForEnvelopes
func (s *Service) EventsRetryForEnvelopes(connectFailureFilter *model.ConnectFailureFilter) *EventsRetryForEnvelopesOp {
	return &EventsRetryForEnvelopesOp{
		Credential: s.credential,
		Method:     "PUT",
		Path:       "connect/envelopes/retry_queue",
		Payload:    connectFailureFilter,
		QueryOpts:  make(url.Values),
	}
}

// EventsRetryForEnvelopesOp implements DocuSign API SDK Connect::retryEventForEnvelopes
type EventsRetryForEnvelopesOp esign.Op

// Do executes the op.  A nil context will return error.
func (op *EventsRetryForEnvelopesOp) Do(ctx context.Context) (*model.ConnectFailureResults, error) {
	var res *model.ConnectFailureResults
	return res, ((*esign.Op)(op)).Do(ctx, &res)
}
