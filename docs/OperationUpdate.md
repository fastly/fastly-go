# OperationUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Method** | Pointer to **string** | The HTTP method for the operation. | [optional] 
**Domain** | Pointer to **string** | The domain for the operation. | [optional] 
**Path** | Pointer to **string** | The path for the operation, which may include path parameters. | [optional] 
**Description** | Pointer to **string** | A description of what the operation does. | [optional] 
**TagIds** | Pointer to **[]string** | An array of operation tag IDs associated with this operation. | [optional] 

## Methods

### NewOperationUpdate

`func NewOperationUpdate() *OperationUpdate`

NewOperationUpdate instantiates a new OperationUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOperationUpdateWithDefaults

`func NewOperationUpdateWithDefaults() *OperationUpdate`

NewOperationUpdateWithDefaults instantiates a new OperationUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMethod

`func (o *OperationUpdate) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *OperationUpdate) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *OperationUpdate) SetMethod(v string)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *OperationUpdate) HasMethod() bool`

HasMethod returns a boolean if a field has been set.

### GetDomain

`func (o *OperationUpdate) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *OperationUpdate) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *OperationUpdate) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *OperationUpdate) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetPath

`func (o *OperationUpdate) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *OperationUpdate) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *OperationUpdate) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *OperationUpdate) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetDescription

`func (o *OperationUpdate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *OperationUpdate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *OperationUpdate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *OperationUpdate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetTagIds

`func (o *OperationUpdate) GetTagIds() []string`

GetTagIds returns the TagIds field if non-nil, zero value otherwise.

### GetTagIdsOk

`func (o *OperationUpdate) GetTagIdsOk() (*[]string, bool)`

GetTagIdsOk returns a tuple with the TagIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagIds

`func (o *OperationUpdate) SetTagIds(v []string)`

SetTagIds sets TagIds field to given value.

### HasTagIds

`func (o *OperationUpdate) HasTagIds() bool`

HasTagIds returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


