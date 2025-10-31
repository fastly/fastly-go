# DdosProtectionTrafficStats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **NullableTime** | Date and time in ISO 8601 format. | [optional] [readonly] 
**UpdatedAt** | Pointer to **NullableTime** | Date and time in ISO 8601 format. | [optional] [readonly] 
**CustomerId** | Pointer to **string** | Alphanumeric string identifying the customer. | [optional] 
**ServiceId** | Pointer to **string** | Alphanumeric string identifying the service. | [optional] 
**Attributes** | Pointer to [**[]DdosProtectionAttributeStats**](DdosProtectionAttributeStats.md) |  | [optional] 
**TrafficPercentage** | Pointer to **float32** | Rule traffic percentage for the event. | [optional] 

## Methods

### NewDdosProtectionTrafficStats

`func NewDdosProtectionTrafficStats() *DdosProtectionTrafficStats`

NewDdosProtectionTrafficStats instantiates a new DdosProtectionTrafficStats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDdosProtectionTrafficStatsWithDefaults

`func NewDdosProtectionTrafficStatsWithDefaults() *DdosProtectionTrafficStats`

NewDdosProtectionTrafficStatsWithDefaults instantiates a new DdosProtectionTrafficStats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *DdosProtectionTrafficStats) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *DdosProtectionTrafficStats) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *DdosProtectionTrafficStats) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *DdosProtectionTrafficStats) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *DdosProtectionTrafficStats) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *DdosProtectionTrafficStats) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetUpdatedAt

`func (o *DdosProtectionTrafficStats) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *DdosProtectionTrafficStats) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *DdosProtectionTrafficStats) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *DdosProtectionTrafficStats) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *DdosProtectionTrafficStats) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *DdosProtectionTrafficStats) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
### GetCustomerId

`func (o *DdosProtectionTrafficStats) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *DdosProtectionTrafficStats) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *DdosProtectionTrafficStats) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.

### HasCustomerId

`func (o *DdosProtectionTrafficStats) HasCustomerId() bool`

HasCustomerId returns a boolean if a field has been set.

### GetServiceId

`func (o *DdosProtectionTrafficStats) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *DdosProtectionTrafficStats) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *DdosProtectionTrafficStats) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.

### HasServiceId

`func (o *DdosProtectionTrafficStats) HasServiceId() bool`

HasServiceId returns a boolean if a field has been set.

### GetAttributes

`func (o *DdosProtectionTrafficStats) GetAttributes() []DdosProtectionAttributeStats`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *DdosProtectionTrafficStats) GetAttributesOk() (*[]DdosProtectionAttributeStats, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *DdosProtectionTrafficStats) SetAttributes(v []DdosProtectionAttributeStats)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *DdosProtectionTrafficStats) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetTrafficPercentage

`func (o *DdosProtectionTrafficStats) GetTrafficPercentage() float32`

GetTrafficPercentage returns the TrafficPercentage field if non-nil, zero value otherwise.

### GetTrafficPercentageOk

`func (o *DdosProtectionTrafficStats) GetTrafficPercentageOk() (*float32, bool)`

GetTrafficPercentageOk returns a tuple with the TrafficPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrafficPercentage

`func (o *DdosProtectionTrafficStats) SetTrafficPercentage(v float32)`

SetTrafficPercentage sets TrafficPercentage field to given value.

### HasTrafficPercentage

`func (o *DdosProtectionTrafficStats) HasTrafficPercentage() bool`

HasTrafficPercentage returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


