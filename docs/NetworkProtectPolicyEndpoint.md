# NetworkProtectPolicyEndpoint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EndpointId** | **string** | A unique identifier for a Network Protect resource. | 
**Preference** | **int32** | Delivery preference within a policy. Lower values are preferred first. Two endpoints in the same policy cannot share a value. | 
**CreatedAt** | Pointer to **time.Time** | An ISO 8601 timestamp in UTC. | [optional] [readonly] 
**UpdatedAt** | Pointer to **time.Time** | An ISO 8601 timestamp in UTC. | [optional] [readonly] 

## Methods

### NewNetworkProtectPolicyEndpoint

`func NewNetworkProtectPolicyEndpoint(endpointId string, preference int32, ) *NetworkProtectPolicyEndpoint`

NewNetworkProtectPolicyEndpoint instantiates a new NetworkProtectPolicyEndpoint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkProtectPolicyEndpointWithDefaults

`func NewNetworkProtectPolicyEndpointWithDefaults() *NetworkProtectPolicyEndpoint`

NewNetworkProtectPolicyEndpointWithDefaults instantiates a new NetworkProtectPolicyEndpoint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEndpointId

`func (o *NetworkProtectPolicyEndpoint) GetEndpointId() string`

GetEndpointId returns the EndpointId field if non-nil, zero value otherwise.

### GetEndpointIdOk

`func (o *NetworkProtectPolicyEndpoint) GetEndpointIdOk() (*string, bool)`

GetEndpointIdOk returns a tuple with the EndpointId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpointId

`func (o *NetworkProtectPolicyEndpoint) SetEndpointId(v string)`

SetEndpointId sets EndpointId field to given value.


### GetPreference

`func (o *NetworkProtectPolicyEndpoint) GetPreference() int32`

GetPreference returns the Preference field if non-nil, zero value otherwise.

### GetPreferenceOk

`func (o *NetworkProtectPolicyEndpoint) GetPreferenceOk() (*int32, bool)`

GetPreferenceOk returns a tuple with the Preference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreference

`func (o *NetworkProtectPolicyEndpoint) SetPreference(v int32)`

SetPreference sets Preference field to given value.


### GetCreatedAt

`func (o *NetworkProtectPolicyEndpoint) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *NetworkProtectPolicyEndpoint) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *NetworkProtectPolicyEndpoint) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *NetworkProtectPolicyEndpoint) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *NetworkProtectPolicyEndpoint) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *NetworkProtectPolicyEndpoint) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *NetworkProtectPolicyEndpoint) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *NetworkProtectPolicyEndpoint) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


