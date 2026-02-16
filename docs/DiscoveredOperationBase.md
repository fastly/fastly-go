# DiscoveredOperationBase

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Method** | Pointer to **string** | The HTTP method for the operation. | [optional] 
**Domain** | Pointer to **string** | The domain for the operation. | [optional] 
**Path** | Pointer to **string** | The path for the operation, which may include path parameters. | [optional] 

## Methods

### NewDiscoveredOperationBase

`func NewDiscoveredOperationBase() *DiscoveredOperationBase`

NewDiscoveredOperationBase instantiates a new DiscoveredOperationBase object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDiscoveredOperationBaseWithDefaults

`func NewDiscoveredOperationBaseWithDefaults() *DiscoveredOperationBase`

NewDiscoveredOperationBaseWithDefaults instantiates a new DiscoveredOperationBase object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMethod

`func (o *DiscoveredOperationBase) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *DiscoveredOperationBase) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *DiscoveredOperationBase) SetMethod(v string)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *DiscoveredOperationBase) HasMethod() bool`

HasMethod returns a boolean if a field has been set.

### GetDomain

`func (o *DiscoveredOperationBase) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *DiscoveredOperationBase) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *DiscoveredOperationBase) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *DiscoveredOperationBase) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetPath

`func (o *DiscoveredOperationBase) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *DiscoveredOperationBase) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *DiscoveredOperationBase) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *DiscoveredOperationBase) HasPath() bool`

HasPath returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


