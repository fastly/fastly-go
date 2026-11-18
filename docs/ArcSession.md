# ArcSession

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ID** | **string** | The unique identifier of the session. | [readonly] 
**VirtualKeyID** | **string** | The ID of the virtual key used for this session. | 
**VirtualKeyName** | Pointer to **string** | The display name of the virtual key used for this session. | [optional] 
**Model** | **string** | The AI model identifier used. | 
**Provider** | **string** | The name of the provider connection that ultimately served the session. | 
**ProviderConnectionID** | **string** | The ID of the provider connection that ultimately served the session. | 
**PrimaryProvider** | Pointer to **string** | The ID of the provider connection originally requested, before any failover occurred. | [optional] 
**Requests** | **int32** | The number of requests made in this session. | 
**InputTokens** | **int32** | The total input tokens consumed in this session. | 
**OutputTokens** | **int32** | The total output tokens generated in this session. | 
**Violations** | **int32** | The number of AI Firewall violations detected in this session. | 
**FirewallEnabled** | **bool** | Whether AI Firewall was enabled for this session&#39;s virtual key. | 
**FirewallOutcome** | Pointer to **string** | The outcome AI Firewall reached for this session. Present only when AI Firewall ran. | [optional] 
**FirewallStage** | Pointer to **string** | The AI Firewall stage that produced the outcome (for example, the request or response stage). Present only when AI Firewall ran. | [optional] 
**FirewallPolicy** | Pointer to **string** | The AI Firewall policy that was applied. Present only when AI Firewall ran. | [optional] 
**FirewallIncompleteReason** | Pointer to **string** | Why AI Firewall inspection was incomplete, if applicable. Present only when AI Firewall ran and inspection didn&#39;t complete. | [optional] 
**CreatedAt** | **time.Time** | The date and time the session was created, in RFC 3339 format. | [readonly] 
**UpdatedAt** | **time.Time** | The date and time the session was last updated, in RFC 3339 format. | [readonly] 
**Logs** | [**[]ArcSessionLogEntry**](ArcSessionLogEntry.md) | The request and response pairs recorded for this session. | 

## Methods

### NewArcSession

`func NewArcSession(id string, virtualKeyID string, model string, provider string, providerConnectionID string, requests int32, inputTokens int32, outputTokens int32, violations int32, firewallEnabled bool, createdAt time.Time, updatedAt time.Time, logs []ArcSessionLogEntry, ) *ArcSession`

NewArcSession instantiates a new ArcSession object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArcSessionWithDefaults

`func NewArcSessionWithDefaults() *ArcSession`

NewArcSessionWithDefaults instantiates a new ArcSession object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetID

`func (o *ArcSession) GetID() string`

GetID returns the ID field if non-nil, zero value otherwise.

### GetIDOk

`func (o *ArcSession) GetIDOk() (*string, bool)`

GetIDOk returns a tuple with the ID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetID

`func (o *ArcSession) SetID(v string)`

SetID sets ID field to given value.


### GetVirtualKeyID

`func (o *ArcSession) GetVirtualKeyID() string`

GetVirtualKeyID returns the VirtualKeyID field if non-nil, zero value otherwise.

### GetVirtualKeyIDOk

`func (o *ArcSession) GetVirtualKeyIDOk() (*string, bool)`

GetVirtualKeyIDOk returns a tuple with the VirtualKeyID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualKeyID

`func (o *ArcSession) SetVirtualKeyID(v string)`

SetVirtualKeyID sets VirtualKeyID field to given value.


### GetVirtualKeyName

`func (o *ArcSession) GetVirtualKeyName() string`

GetVirtualKeyName returns the VirtualKeyName field if non-nil, zero value otherwise.

### GetVirtualKeyNameOk

`func (o *ArcSession) GetVirtualKeyNameOk() (*string, bool)`

GetVirtualKeyNameOk returns a tuple with the VirtualKeyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualKeyName

`func (o *ArcSession) SetVirtualKeyName(v string)`

SetVirtualKeyName sets VirtualKeyName field to given value.

### HasVirtualKeyName

`func (o *ArcSession) HasVirtualKeyName() bool`

HasVirtualKeyName returns a boolean if a field has been set.

### GetModel

`func (o *ArcSession) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *ArcSession) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *ArcSession) SetModel(v string)`

SetModel sets Model field to given value.


### GetProvider

`func (o *ArcSession) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *ArcSession) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *ArcSession) SetProvider(v string)`

SetProvider sets Provider field to given value.


### GetProviderConnectionID

`func (o *ArcSession) GetProviderConnectionID() string`

GetProviderConnectionID returns the ProviderConnectionID field if non-nil, zero value otherwise.

### GetProviderConnectionIDOk

`func (o *ArcSession) GetProviderConnectionIDOk() (*string, bool)`

GetProviderConnectionIDOk returns a tuple with the ProviderConnectionID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderConnectionID

`func (o *ArcSession) SetProviderConnectionID(v string)`

SetProviderConnectionID sets ProviderConnectionID field to given value.


### GetPrimaryProvider

`func (o *ArcSession) GetPrimaryProvider() string`

GetPrimaryProvider returns the PrimaryProvider field if non-nil, zero value otherwise.

### GetPrimaryProviderOk

`func (o *ArcSession) GetPrimaryProviderOk() (*string, bool)`

GetPrimaryProviderOk returns a tuple with the PrimaryProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryProvider

`func (o *ArcSession) SetPrimaryProvider(v string)`

SetPrimaryProvider sets PrimaryProvider field to given value.

### HasPrimaryProvider

`func (o *ArcSession) HasPrimaryProvider() bool`

HasPrimaryProvider returns a boolean if a field has been set.

### GetRequests

`func (o *ArcSession) GetRequests() int32`

GetRequests returns the Requests field if non-nil, zero value otherwise.

### GetRequestsOk

`func (o *ArcSession) GetRequestsOk() (*int32, bool)`

GetRequestsOk returns a tuple with the Requests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequests

`func (o *ArcSession) SetRequests(v int32)`

SetRequests sets Requests field to given value.


### GetInputTokens

`func (o *ArcSession) GetInputTokens() int32`

GetInputTokens returns the InputTokens field if non-nil, zero value otherwise.

### GetInputTokensOk

`func (o *ArcSession) GetInputTokensOk() (*int32, bool)`

GetInputTokensOk returns a tuple with the InputTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputTokens

`func (o *ArcSession) SetInputTokens(v int32)`

SetInputTokens sets InputTokens field to given value.


### GetOutputTokens

`func (o *ArcSession) GetOutputTokens() int32`

GetOutputTokens returns the OutputTokens field if non-nil, zero value otherwise.

### GetOutputTokensOk

`func (o *ArcSession) GetOutputTokensOk() (*int32, bool)`

GetOutputTokensOk returns a tuple with the OutputTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputTokens

`func (o *ArcSession) SetOutputTokens(v int32)`

SetOutputTokens sets OutputTokens field to given value.


### GetViolations

`func (o *ArcSession) GetViolations() int32`

GetViolations returns the Violations field if non-nil, zero value otherwise.

### GetViolationsOk

`func (o *ArcSession) GetViolationsOk() (*int32, bool)`

GetViolationsOk returns a tuple with the Violations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViolations

`func (o *ArcSession) SetViolations(v int32)`

SetViolations sets Violations field to given value.


### GetFirewallEnabled

`func (o *ArcSession) GetFirewallEnabled() bool`

GetFirewallEnabled returns the FirewallEnabled field if non-nil, zero value otherwise.

### GetFirewallEnabledOk

`func (o *ArcSession) GetFirewallEnabledOk() (*bool, bool)`

GetFirewallEnabledOk returns a tuple with the FirewallEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirewallEnabled

`func (o *ArcSession) SetFirewallEnabled(v bool)`

SetFirewallEnabled sets FirewallEnabled field to given value.


### GetFirewallOutcome

`func (o *ArcSession) GetFirewallOutcome() string`

GetFirewallOutcome returns the FirewallOutcome field if non-nil, zero value otherwise.

### GetFirewallOutcomeOk

`func (o *ArcSession) GetFirewallOutcomeOk() (*string, bool)`

GetFirewallOutcomeOk returns a tuple with the FirewallOutcome field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirewallOutcome

`func (o *ArcSession) SetFirewallOutcome(v string)`

SetFirewallOutcome sets FirewallOutcome field to given value.

### HasFirewallOutcome

`func (o *ArcSession) HasFirewallOutcome() bool`

HasFirewallOutcome returns a boolean if a field has been set.

### GetFirewallStage

`func (o *ArcSession) GetFirewallStage() string`

GetFirewallStage returns the FirewallStage field if non-nil, zero value otherwise.

### GetFirewallStageOk

`func (o *ArcSession) GetFirewallStageOk() (*string, bool)`

GetFirewallStageOk returns a tuple with the FirewallStage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirewallStage

`func (o *ArcSession) SetFirewallStage(v string)`

SetFirewallStage sets FirewallStage field to given value.

### HasFirewallStage

`func (o *ArcSession) HasFirewallStage() bool`

HasFirewallStage returns a boolean if a field has been set.

### GetFirewallPolicy

`func (o *ArcSession) GetFirewallPolicy() string`

GetFirewallPolicy returns the FirewallPolicy field if non-nil, zero value otherwise.

### GetFirewallPolicyOk

`func (o *ArcSession) GetFirewallPolicyOk() (*string, bool)`

GetFirewallPolicyOk returns a tuple with the FirewallPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirewallPolicy

`func (o *ArcSession) SetFirewallPolicy(v string)`

SetFirewallPolicy sets FirewallPolicy field to given value.

### HasFirewallPolicy

`func (o *ArcSession) HasFirewallPolicy() bool`

HasFirewallPolicy returns a boolean if a field has been set.

### GetFirewallIncompleteReason

`func (o *ArcSession) GetFirewallIncompleteReason() string`

GetFirewallIncompleteReason returns the FirewallIncompleteReason field if non-nil, zero value otherwise.

### GetFirewallIncompleteReasonOk

`func (o *ArcSession) GetFirewallIncompleteReasonOk() (*string, bool)`

GetFirewallIncompleteReasonOk returns a tuple with the FirewallIncompleteReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirewallIncompleteReason

`func (o *ArcSession) SetFirewallIncompleteReason(v string)`

SetFirewallIncompleteReason sets FirewallIncompleteReason field to given value.

### HasFirewallIncompleteReason

`func (o *ArcSession) HasFirewallIncompleteReason() bool`

HasFirewallIncompleteReason returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ArcSession) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ArcSession) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ArcSession) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *ArcSession) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ArcSession) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ArcSession) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetLogs

`func (o *ArcSession) GetLogs() []ArcSessionLogEntry`

GetLogs returns the Logs field if non-nil, zero value otherwise.

### GetLogsOk

`func (o *ArcSession) GetLogsOk() (*[]ArcSessionLogEntry, bool)`

GetLogsOk returns a tuple with the Logs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogs

`func (o *ArcSession) SetLogs(v []ArcSessionLogEntry)`

SetLogs sets Logs field to given value.



[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
