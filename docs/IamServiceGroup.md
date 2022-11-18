# IamServiceGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **NullableTime** | Date and time in ISO 8601 format. | [optional] [readonly] 
**UpdatedAt** | Pointer to **NullableTime** | Date and time in ISO 8601 format. | [optional] [readonly] 
**ID** | Pointer to **string** | Alphanumeric string identifying the service group. | [optional] 
**Object** | Pointer to **string** | The type of the object. | [optional] 
**Name** | Pointer to **string** | Name of the service group. | [optional] 
**Description** | Pointer to **string** | Description of the service group. | [optional] 
**ServicesCount** | Pointer to **int32** | Number of services in the service group. | [optional] 

## Methods

### NewIamServiceGroup

`func NewIamServiceGroup() *IamServiceGroup`

NewIamServiceGroup instantiates a new IamServiceGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIamServiceGroupWithDefaults

`func NewIamServiceGroupWithDefaults() *IamServiceGroup`

NewIamServiceGroupWithDefaults instantiates a new IamServiceGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *IamServiceGroup) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *IamServiceGroup) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *IamServiceGroup) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *IamServiceGroup) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *IamServiceGroup) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *IamServiceGroup) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetUpdatedAt

`func (o *IamServiceGroup) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *IamServiceGroup) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *IamServiceGroup) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *IamServiceGroup) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *IamServiceGroup) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *IamServiceGroup) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
### GetID

`func (o *IamServiceGroup) GetID() string`

GetID returns the ID field if non-nil, zero value otherwise.

### GetIDOk

`func (o *IamServiceGroup) GetIDOk() (*string, bool)`

GetIDOk returns a tuple with the ID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetID

`func (o *IamServiceGroup) SetID(v string)`

SetID sets ID field to given value.

### HasID

`func (o *IamServiceGroup) HasID() bool`

HasID returns a boolean if a field has been set.

### GetObject

`func (o *IamServiceGroup) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *IamServiceGroup) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *IamServiceGroup) SetObject(v string)`

SetObject sets Object field to given value.

### HasObject

`func (o *IamServiceGroup) HasObject() bool`

HasObject returns a boolean if a field has been set.

### GetName

`func (o *IamServiceGroup) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IamServiceGroup) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IamServiceGroup) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *IamServiceGroup) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *IamServiceGroup) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *IamServiceGroup) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *IamServiceGroup) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *IamServiceGroup) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetServicesCount

`func (o *IamServiceGroup) GetServicesCount() int32`

GetServicesCount returns the ServicesCount field if non-nil, zero value otherwise.

### GetServicesCountOk

`func (o *IamServiceGroup) GetServicesCountOk() (*int32, bool)`

GetServicesCountOk returns a tuple with the ServicesCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicesCount

`func (o *IamServiceGroup) SetServicesCount(v int32)`

SetServicesCount sets ServicesCount field to given value.

### HasServicesCount

`func (o *IamServiceGroup) HasServicesCount() bool`

HasServicesCount returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
