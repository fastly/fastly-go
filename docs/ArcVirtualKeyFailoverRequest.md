# ArcVirtualKeyFailoverRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Targets** | [**[]ArcFailoverTarget**](ArcFailoverTarget.md) | The provider connections to retry against, in order. | 
**TimeoutEnabled** | Pointer to **bool** | Whether a request timeout triggers failover. | [optional] 
**RateLimit429Enabled** | Pointer to **bool** | Whether a `429 Too Many Requests` response from the provider triggers failover. | [optional] 
**ServerError5xxEnabled** | Pointer to **bool** | Whether a `5xx` server error response from the provider triggers failover. | [optional] 

## Methods

### NewArcVirtualKeyFailoverRequest

`func NewArcVirtualKeyFailoverRequest(targets []ArcFailoverTarget, ) *ArcVirtualKeyFailoverRequest`

NewArcVirtualKeyFailoverRequest instantiates a new ArcVirtualKeyFailoverRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArcVirtualKeyFailoverRequestWithDefaults

`func NewArcVirtualKeyFailoverRequestWithDefaults() *ArcVirtualKeyFailoverRequest`

NewArcVirtualKeyFailoverRequestWithDefaults instantiates a new ArcVirtualKeyFailoverRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTargets

`func (o *ArcVirtualKeyFailoverRequest) GetTargets() []ArcFailoverTarget`

GetTargets returns the Targets field if non-nil, zero value otherwise.

### GetTargetsOk

`func (o *ArcVirtualKeyFailoverRequest) GetTargetsOk() (*[]ArcFailoverTarget, bool)`

GetTargetsOk returns a tuple with the Targets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargets

`func (o *ArcVirtualKeyFailoverRequest) SetTargets(v []ArcFailoverTarget)`

SetTargets sets Targets field to given value.


### GetTimeoutEnabled

`func (o *ArcVirtualKeyFailoverRequest) GetTimeoutEnabled() bool`

GetTimeoutEnabled returns the TimeoutEnabled field if non-nil, zero value otherwise.

### GetTimeoutEnabledOk

`func (o *ArcVirtualKeyFailoverRequest) GetTimeoutEnabledOk() (*bool, bool)`

GetTimeoutEnabledOk returns a tuple with the TimeoutEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeoutEnabled

`func (o *ArcVirtualKeyFailoverRequest) SetTimeoutEnabled(v bool)`

SetTimeoutEnabled sets TimeoutEnabled field to given value.

### HasTimeoutEnabled

`func (o *ArcVirtualKeyFailoverRequest) HasTimeoutEnabled() bool`

HasTimeoutEnabled returns a boolean if a field has been set.

### GetRateLimit429Enabled

`func (o *ArcVirtualKeyFailoverRequest) GetRateLimit429Enabled() bool`

GetRateLimit429Enabled returns the RateLimit429Enabled field if non-nil, zero value otherwise.

### GetRateLimit429EnabledOk

`func (o *ArcVirtualKeyFailoverRequest) GetRateLimit429EnabledOk() (*bool, bool)`

GetRateLimit429EnabledOk returns a tuple with the RateLimit429Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateLimit429Enabled

`func (o *ArcVirtualKeyFailoverRequest) SetRateLimit429Enabled(v bool)`

SetRateLimit429Enabled sets RateLimit429Enabled field to given value.

### HasRateLimit429Enabled

`func (o *ArcVirtualKeyFailoverRequest) HasRateLimit429Enabled() bool`

HasRateLimit429Enabled returns a boolean if a field has been set.

### GetServerError5xxEnabled

`func (o *ArcVirtualKeyFailoverRequest) GetServerError5xxEnabled() bool`

GetServerError5xxEnabled returns the ServerError5xxEnabled field if non-nil, zero value otherwise.

### GetServerError5xxEnabledOk

`func (o *ArcVirtualKeyFailoverRequest) GetServerError5xxEnabledOk() (*bool, bool)`

GetServerError5xxEnabledOk returns a tuple with the ServerError5xxEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerError5xxEnabled

`func (o *ArcVirtualKeyFailoverRequest) SetServerError5xxEnabled(v bool)`

SetServerError5xxEnabled sets ServerError5xxEnabled field to given value.

### HasServerError5xxEnabled

`func (o *ArcVirtualKeyFailoverRequest) HasServerError5xxEnabled() bool`

HasServerError5xxEnabled returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
