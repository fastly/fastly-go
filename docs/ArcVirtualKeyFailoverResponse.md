# ArcVirtualKeyFailoverResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Targets** | [**[]ArcFailoverTarget**](ArcFailoverTarget.md) | The provider connections to retry against, in order. | 
**TimeoutEnabled** | **bool** | Whether a request timeout triggers failover. | 
**RateLimit429Enabled** | **bool** | Whether a `429 Too Many Requests` response from the provider triggers failover. | 
**ServerError5xxEnabled** | **bool** | Whether a `5xx` server error response from the provider triggers failover. | 

## Methods

### NewArcVirtualKeyFailoverResponse

`func NewArcVirtualKeyFailoverResponse(targets []ArcFailoverTarget, timeoutEnabled bool, rateLimit429Enabled bool, serverError5xxEnabled bool, ) *ArcVirtualKeyFailoverResponse`

NewArcVirtualKeyFailoverResponse instantiates a new ArcVirtualKeyFailoverResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArcVirtualKeyFailoverResponseWithDefaults

`func NewArcVirtualKeyFailoverResponseWithDefaults() *ArcVirtualKeyFailoverResponse`

NewArcVirtualKeyFailoverResponseWithDefaults instantiates a new ArcVirtualKeyFailoverResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTargets

`func (o *ArcVirtualKeyFailoverResponse) GetTargets() []ArcFailoverTarget`

GetTargets returns the Targets field if non-nil, zero value otherwise.

### GetTargetsOk

`func (o *ArcVirtualKeyFailoverResponse) GetTargetsOk() (*[]ArcFailoverTarget, bool)`

GetTargetsOk returns a tuple with the Targets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargets

`func (o *ArcVirtualKeyFailoverResponse) SetTargets(v []ArcFailoverTarget)`

SetTargets sets Targets field to given value.


### GetTimeoutEnabled

`func (o *ArcVirtualKeyFailoverResponse) GetTimeoutEnabled() bool`

GetTimeoutEnabled returns the TimeoutEnabled field if non-nil, zero value otherwise.

### GetTimeoutEnabledOk

`func (o *ArcVirtualKeyFailoverResponse) GetTimeoutEnabledOk() (*bool, bool)`

GetTimeoutEnabledOk returns a tuple with the TimeoutEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeoutEnabled

`func (o *ArcVirtualKeyFailoverResponse) SetTimeoutEnabled(v bool)`

SetTimeoutEnabled sets TimeoutEnabled field to given value.


### GetRateLimit429Enabled

`func (o *ArcVirtualKeyFailoverResponse) GetRateLimit429Enabled() bool`

GetRateLimit429Enabled returns the RateLimit429Enabled field if non-nil, zero value otherwise.

### GetRateLimit429EnabledOk

`func (o *ArcVirtualKeyFailoverResponse) GetRateLimit429EnabledOk() (*bool, bool)`

GetRateLimit429EnabledOk returns a tuple with the RateLimit429Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateLimit429Enabled

`func (o *ArcVirtualKeyFailoverResponse) SetRateLimit429Enabled(v bool)`

SetRateLimit429Enabled sets RateLimit429Enabled field to given value.


### GetServerError5xxEnabled

`func (o *ArcVirtualKeyFailoverResponse) GetServerError5xxEnabled() bool`

GetServerError5xxEnabled returns the ServerError5xxEnabled field if non-nil, zero value otherwise.

### GetServerError5xxEnabledOk

`func (o *ArcVirtualKeyFailoverResponse) GetServerError5xxEnabledOk() (*bool, bool)`

GetServerError5xxEnabledOk returns a tuple with the ServerError5xxEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerError5xxEnabled

`func (o *ArcVirtualKeyFailoverResponse) SetServerError5xxEnabled(v bool)`

SetServerError5xxEnabled sets ServerError5xxEnabled field to given value.



[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


