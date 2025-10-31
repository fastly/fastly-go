# DirectorBackend

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **NullableTime** | Date and time in ISO 8601 format. | [optional] [readonly] 
**DeletedAt** | Pointer to **NullableTime** | Date and time in ISO 8601 format. | [optional] [readonly] 
**UpdatedAt** | Pointer to **NullableTime** | Date and time in ISO 8601 format. | [optional] [readonly] 
**ServiceId** | Pointer to **string** |  | [optional] [readonly] 
**Version** | Pointer to **int32** |  | [optional] [readonly] 
**BackendName** | Pointer to **string** | The name of the backend. | [optional] 
**Director** | Pointer to **string** | Name for the Director. | [optional] 

## Methods

### NewDirectorBackend

`func NewDirectorBackend() *DirectorBackend`

NewDirectorBackend instantiates a new DirectorBackend object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDirectorBackendWithDefaults

`func NewDirectorBackendWithDefaults() *DirectorBackend`

NewDirectorBackendWithDefaults instantiates a new DirectorBackend object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *DirectorBackend) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *DirectorBackend) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *DirectorBackend) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *DirectorBackend) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *DirectorBackend) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *DirectorBackend) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetDeletedAt

`func (o *DirectorBackend) GetDeletedAt() time.Time`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *DirectorBackend) GetDeletedAtOk() (*time.Time, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *DirectorBackend) SetDeletedAt(v time.Time)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *DirectorBackend) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### SetDeletedAtNil

`func (o *DirectorBackend) SetDeletedAtNil(b bool)`

 SetDeletedAtNil sets the value for DeletedAt to be an explicit nil

### UnsetDeletedAt
`func (o *DirectorBackend) UnsetDeletedAt()`

UnsetDeletedAt ensures that no value is present for DeletedAt, not even an explicit nil
### GetUpdatedAt

`func (o *DirectorBackend) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *DirectorBackend) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *DirectorBackend) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *DirectorBackend) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *DirectorBackend) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *DirectorBackend) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
### GetServiceId

`func (o *DirectorBackend) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *DirectorBackend) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *DirectorBackend) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.

### HasServiceId

`func (o *DirectorBackend) HasServiceId() bool`

HasServiceId returns a boolean if a field has been set.

### GetVersion

`func (o *DirectorBackend) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *DirectorBackend) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *DirectorBackend) SetVersion(v int32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *DirectorBackend) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetBackendName

`func (o *DirectorBackend) GetBackendName() string`

GetBackendName returns the BackendName field if non-nil, zero value otherwise.

### GetBackendNameOk

`func (o *DirectorBackend) GetBackendNameOk() (*string, bool)`

GetBackendNameOk returns a tuple with the BackendName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackendName

`func (o *DirectorBackend) SetBackendName(v string)`

SetBackendName sets BackendName field to given value.

### HasBackendName

`func (o *DirectorBackend) HasBackendName() bool`

HasBackendName returns a boolean if a field has been set.

### GetDirector

`func (o *DirectorBackend) GetDirector() string`

GetDirector returns the Director field if non-nil, zero value otherwise.

### GetDirectorOk

`func (o *DirectorBackend) GetDirectorOk() (*string, bool)`

GetDirectorOk returns a tuple with the Director field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirector

`func (o *DirectorBackend) SetDirector(v string)`

SetDirector sets Director field to given value.

### HasDirector

`func (o *DirectorBackend) HasDirector() bool`

HasDirector returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


