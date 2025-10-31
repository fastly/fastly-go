# DdosProtectionTrafficStatsAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomerId** | Pointer to **string** | Alphanumeric string identifying the customer. | [optional] 
**ServiceId** | Pointer to **string** | Alphanumeric string identifying the service. | [optional] 
**Attributes** | Pointer to [**[]DdosProtectionAttributeStats**](DdosProtectionAttributeStats.md) |  | [optional] 
**TrafficPercentage** | Pointer to **float32** | Rule traffic percentage for the event. | [optional] 

## Methods

### NewDdosProtectionTrafficStatsAllOf

`func NewDdosProtectionTrafficStatsAllOf() *DdosProtectionTrafficStatsAllOf`

NewDdosProtectionTrafficStatsAllOf instantiates a new DdosProtectionTrafficStatsAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDdosProtectionTrafficStatsAllOfWithDefaults

`func NewDdosProtectionTrafficStatsAllOfWithDefaults() *DdosProtectionTrafficStatsAllOf`

NewDdosProtectionTrafficStatsAllOfWithDefaults instantiates a new DdosProtectionTrafficStatsAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomerId

`func (o *DdosProtectionTrafficStatsAllOf) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *DdosProtectionTrafficStatsAllOf) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *DdosProtectionTrafficStatsAllOf) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.

### HasCustomerId

`func (o *DdosProtectionTrafficStatsAllOf) HasCustomerId() bool`

HasCustomerId returns a boolean if a field has been set.

### GetServiceId

`func (o *DdosProtectionTrafficStatsAllOf) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *DdosProtectionTrafficStatsAllOf) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *DdosProtectionTrafficStatsAllOf) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.

### HasServiceId

`func (o *DdosProtectionTrafficStatsAllOf) HasServiceId() bool`

HasServiceId returns a boolean if a field has been set.

### GetAttributes

`func (o *DdosProtectionTrafficStatsAllOf) GetAttributes() []DdosProtectionAttributeStats`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *DdosProtectionTrafficStatsAllOf) GetAttributesOk() (*[]DdosProtectionAttributeStats, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *DdosProtectionTrafficStatsAllOf) SetAttributes(v []DdosProtectionAttributeStats)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *DdosProtectionTrafficStatsAllOf) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetTrafficPercentage

`func (o *DdosProtectionTrafficStatsAllOf) GetTrafficPercentage() float32`

GetTrafficPercentage returns the TrafficPercentage field if non-nil, zero value otherwise.

### GetTrafficPercentageOk

`func (o *DdosProtectionTrafficStatsAllOf) GetTrafficPercentageOk() (*float32, bool)`

GetTrafficPercentageOk returns a tuple with the TrafficPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrafficPercentage

`func (o *DdosProtectionTrafficStatsAllOf) SetTrafficPercentage(v float32)`

SetTrafficPercentage sets TrafficPercentage field to given value.

### HasTrafficPercentage

`func (o *DdosProtectionTrafficStatsAllOf) HasTrafficPercentage() bool`

HasTrafficPercentage returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


