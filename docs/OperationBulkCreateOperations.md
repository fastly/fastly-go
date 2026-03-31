# OperationBulkCreateOperations

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Method** | **string** | The HTTP method for the operation. | 
**Domain** | **string** | The domain for the operation. | 
**Path** | **string** | The path for the operation. | 
**Description** | Pointer to **string** | A description of what the operation does. | [optional] 
**TagIds** | Pointer to **[]string** | An array of tag IDs to associate with this operation. | [optional] 
**Status** | Pointer to **string** | The status to assign to the operation. Defaults to SAVED if omitted. | [optional] [default to "SAVED"]

## Methods

### NewOperationBulkCreateOperations

`func NewOperationBulkCreateOperations(method string, domain string, path string, ) *OperationBulkCreateOperations`

NewOperationBulkCreateOperations instantiates a new OperationBulkCreateOperations object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOperationBulkCreateOperationsWithDefaults

`func NewOperationBulkCreateOperationsWithDefaults() *OperationBulkCreateOperations`

NewOperationBulkCreateOperationsWithDefaults instantiates a new OperationBulkCreateOperations object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMethod

`func (o *OperationBulkCreateOperations) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *OperationBulkCreateOperations) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *OperationBulkCreateOperations) SetMethod(v string)`

SetMethod sets Method field to given value.


### GetDomain

`func (o *OperationBulkCreateOperations) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *OperationBulkCreateOperations) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *OperationBulkCreateOperations) SetDomain(v string)`

SetDomain sets Domain field to given value.


### GetPath

`func (o *OperationBulkCreateOperations) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *OperationBulkCreateOperations) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *OperationBulkCreateOperations) SetPath(v string)`

SetPath sets Path field to given value.


### GetDescription

`func (o *OperationBulkCreateOperations) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *OperationBulkCreateOperations) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *OperationBulkCreateOperations) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *OperationBulkCreateOperations) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetTagIds

`func (o *OperationBulkCreateOperations) GetTagIds() []string`

GetTagIds returns the TagIds field if non-nil, zero value otherwise.

### GetTagIdsOk

`func (o *OperationBulkCreateOperations) GetTagIdsOk() (*[]string, bool)`

GetTagIdsOk returns a tuple with the TagIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagIds

`func (o *OperationBulkCreateOperations) SetTagIds(v []string)`

SetTagIds sets TagIds field to given value.

### HasTagIds

`func (o *OperationBulkCreateOperations) HasTagIds() bool`

HasTagIds returns a boolean if a field has been set.

### GetStatus

`func (o *OperationBulkCreateOperations) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *OperationBulkCreateOperations) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *OperationBulkCreateOperations) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *OperationBulkCreateOperations) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


