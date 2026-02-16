# OperationGet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Method** | **string** | The HTTP method for the operation. | 
**Domain** | **string** | The domain for the operation. | 
**Path** | **string** | The path for the operation, which may include path parameters. | 
**Description** | Pointer to **string** | A description of what the operation does. | [optional] 
**TagIds** | Pointer to **[]string** | An array of operation tag IDs associated with this operation. | [optional] 
**Id** | **string** | The unique identifier of the operation. | [readonly] 
**CreatedAt** | Pointer to **time.Time** | The timestamp when the operation was created. | [optional] [readonly] 
**UpdatedAt** | **time.Time** | The timestamp when the operation was last updated. | [readonly] 
**LastSeenAt** | Pointer to **time.Time** | The timestamp when the operation was last seen in traffic. | [optional] [readonly] 

## Methods

### NewOperationGet

`func NewOperationGet(method string, domain string, path string, id string, updatedAt time.Time, ) *OperationGet`

NewOperationGet instantiates a new OperationGet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOperationGetWithDefaults

`func NewOperationGetWithDefaults() *OperationGet`

NewOperationGetWithDefaults instantiates a new OperationGet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMethod

`func (o *OperationGet) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *OperationGet) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *OperationGet) SetMethod(v string)`

SetMethod sets Method field to given value.


### GetDomain

`func (o *OperationGet) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *OperationGet) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *OperationGet) SetDomain(v string)`

SetDomain sets Domain field to given value.


### GetPath

`func (o *OperationGet) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *OperationGet) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *OperationGet) SetPath(v string)`

SetPath sets Path field to given value.


### GetDescription

`func (o *OperationGet) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *OperationGet) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *OperationGet) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *OperationGet) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetTagIds

`func (o *OperationGet) GetTagIds() []string`

GetTagIds returns the TagIds field if non-nil, zero value otherwise.

### GetTagIdsOk

`func (o *OperationGet) GetTagIdsOk() (*[]string, bool)`

GetTagIdsOk returns a tuple with the TagIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagIds

`func (o *OperationGet) SetTagIds(v []string)`

SetTagIds sets TagIds field to given value.

### HasTagIds

`func (o *OperationGet) HasTagIds() bool`

HasTagIds returns a boolean if a field has been set.

### GetId

`func (o *OperationGet) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OperationGet) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OperationGet) SetId(v string)`

SetId sets Id field to given value.


### GetCreatedAt

`func (o *OperationGet) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *OperationGet) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *OperationGet) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *OperationGet) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *OperationGet) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *OperationGet) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *OperationGet) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetLastSeenAt

`func (o *OperationGet) GetLastSeenAt() time.Time`

GetLastSeenAt returns the LastSeenAt field if non-nil, zero value otherwise.

### GetLastSeenAtOk

`func (o *OperationGet) GetLastSeenAtOk() (*time.Time, bool)`

GetLastSeenAtOk returns a tuple with the LastSeenAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSeenAt

`func (o *OperationGet) SetLastSeenAt(v time.Time)`

SetLastSeenAt sets LastSeenAt field to given value.

### HasLastSeenAt

`func (o *OperationGet) HasLastSeenAt() bool`

HasLastSeenAt returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


