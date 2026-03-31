# OperationCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Method** | **string** | The HTTP method for the operation. | 
**Domain** | **string** | The domain for the operation. | 
**Path** | **string** | The path for the operation, which may include path parameters. | 
**Description** | Pointer to **string** | A description of what the operation does. | [optional] 
**TagIds** | Pointer to **[]string** | An array of operation tag IDs associated with this operation. | [optional] 
**Status** | Pointer to **string** | The status to assign to the operation. Defaults to SAVED if omitted. | [optional] [default to "SAVED"]

## Methods

### NewOperationCreate

`func NewOperationCreate(method string, domain string, path string, ) *OperationCreate`

NewOperationCreate instantiates a new OperationCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOperationCreateWithDefaults

`func NewOperationCreateWithDefaults() *OperationCreate`

NewOperationCreateWithDefaults instantiates a new OperationCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMethod

`func (o *OperationCreate) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *OperationCreate) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *OperationCreate) SetMethod(v string)`

SetMethod sets Method field to given value.


### GetDomain

`func (o *OperationCreate) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *OperationCreate) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *OperationCreate) SetDomain(v string)`

SetDomain sets Domain field to given value.


### GetPath

`func (o *OperationCreate) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *OperationCreate) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *OperationCreate) SetPath(v string)`

SetPath sets Path field to given value.


### GetDescription

`func (o *OperationCreate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *OperationCreate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *OperationCreate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *OperationCreate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetTagIds

`func (o *OperationCreate) GetTagIds() []string`

GetTagIds returns the TagIds field if non-nil, zero value otherwise.

### GetTagIdsOk

`func (o *OperationCreate) GetTagIdsOk() (*[]string, bool)`

GetTagIdsOk returns a tuple with the TagIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagIds

`func (o *OperationCreate) SetTagIds(v []string)`

SetTagIds sets TagIds field to given value.

### HasTagIds

`func (o *OperationCreate) HasTagIds() bool`

HasTagIds returns a boolean if a field has been set.

### GetStatus

`func (o *OperationCreate) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *OperationCreate) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *OperationCreate) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *OperationCreate) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


