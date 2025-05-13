# DdosProtectionAttributeStats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to [**DdosProtectionTrafficAttribute**](DdosProtectionTrafficAttribute.md) |  | [optional] 
**Values** | Pointer to [**[]DdosProtectionAttributeValue**](DdosProtectionAttributeValue.md) | Values for traffic attribute. | [optional] 

## Methods

### NewDdosProtectionAttributeStats

`func NewDdosProtectionAttributeStats() *DdosProtectionAttributeStats`

NewDdosProtectionAttributeStats instantiates a new DdosProtectionAttributeStats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDdosProtectionAttributeStatsWithDefaults

`func NewDdosProtectionAttributeStatsWithDefaults() *DdosProtectionAttributeStats`

NewDdosProtectionAttributeStatsWithDefaults instantiates a new DdosProtectionAttributeStats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *DdosProtectionAttributeStats) GetName() DdosProtectionTrafficAttribute`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DdosProtectionAttributeStats) GetNameOk() (*DdosProtectionTrafficAttribute, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DdosProtectionAttributeStats) SetName(v DdosProtectionTrafficAttribute)`

SetName sets Name field to given value.

### HasName

`func (o *DdosProtectionAttributeStats) HasName() bool`

HasName returns a boolean if a field has been set.

### GetValues

`func (o *DdosProtectionAttributeStats) GetValues() []DdosProtectionAttributeValue`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *DdosProtectionAttributeStats) GetValuesOk() (*[]DdosProtectionAttributeValue, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *DdosProtectionAttributeStats) SetValues(v []DdosProtectionAttributeValue)`

SetValues sets Values field to given value.

### HasValues

`func (o *DdosProtectionAttributeStats) HasValues() bool`

HasValues returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
