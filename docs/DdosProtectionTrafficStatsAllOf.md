# DdosProtectionTrafficStatsAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomerID** | Pointer to **string** | Alphanumeric string identifying the customer. | [optional] 
**ServiceID** | Pointer to **string** | Alphanumeric string identifying the service. | [optional] 
**Attributes** | Pointer to [**[]DdosProtectionAttributeStats**](DdosProtectionAttributeStats.md) |  | [optional] 

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

### GetCustomerID

`func (o *DdosProtectionTrafficStatsAllOf) GetCustomerID() string`

GetCustomerID returns the CustomerID field if non-nil, zero value otherwise.

### GetCustomerIDOk

`func (o *DdosProtectionTrafficStatsAllOf) GetCustomerIDOk() (*string, bool)`

GetCustomerIDOk returns a tuple with the CustomerID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerID

`func (o *DdosProtectionTrafficStatsAllOf) SetCustomerID(v string)`

SetCustomerID sets CustomerID field to given value.

### HasCustomerID

`func (o *DdosProtectionTrafficStatsAllOf) HasCustomerID() bool`

HasCustomerID returns a boolean if a field has been set.

### GetServiceID

`func (o *DdosProtectionTrafficStatsAllOf) GetServiceID() string`

GetServiceID returns the ServiceID field if non-nil, zero value otherwise.

### GetServiceIDOk

`func (o *DdosProtectionTrafficStatsAllOf) GetServiceIDOk() (*string, bool)`

GetServiceIDOk returns a tuple with the ServiceID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceID

`func (o *DdosProtectionTrafficStatsAllOf) SetServiceID(v string)`

SetServiceID sets ServiceID field to given value.

### HasServiceID

`func (o *DdosProtectionTrafficStatsAllOf) HasServiceID() bool`

HasServiceID returns a boolean if a field has been set.

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


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
