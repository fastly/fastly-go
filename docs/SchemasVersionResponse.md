# SchemasVersionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | Pointer to **bool** | Whether this is the active version or not. | [optional] [default to false]
**Comment** | Pointer to **NullableString** | A freeform descriptive note. | [optional] 
**Deployed** | Pointer to **bool** | Unused at this time. | [optional] 
**Locked** | Pointer to **bool** | Whether this version is locked or not. Objects can not be added or edited on locked versions. | [optional] [default to false]
**Number** | Pointer to **int32** | The number of this version. | [optional] [readonly] 
**Staging** | Pointer to **bool** | Unused at this time. | [optional] [default to false]
**Testing** | Pointer to **bool** | Unused at this time. | [optional] [default to false]
**CreatedAt** | Pointer to **NullableTime** | Date and time in ISO 8601 format. | [optional] [readonly] 
**DeletedAt** | Pointer to **NullableTime** | Date and time in ISO 8601 format. | [optional] [readonly] 
**UpdatedAt** | Pointer to **NullableTime** | Date and time in ISO 8601 format. | [optional] [readonly] 
**ServiceID** | Pointer to **string** |  | [optional] [readonly] 
**Environments** | Pointer to [**[]Environment**](Environment.md) | A list of environments where the service has been deployed. | [optional] 

## Methods

### NewSchemasVersionResponse

`func NewSchemasVersionResponse() *SchemasVersionResponse`

NewSchemasVersionResponse instantiates a new SchemasVersionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSchemasVersionResponseWithDefaults

`func NewSchemasVersionResponseWithDefaults() *SchemasVersionResponse`

NewSchemasVersionResponseWithDefaults instantiates a new SchemasVersionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActive

`func (o *SchemasVersionResponse) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *SchemasVersionResponse) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *SchemasVersionResponse) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *SchemasVersionResponse) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetComment

`func (o *SchemasVersionResponse) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *SchemasVersionResponse) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *SchemasVersionResponse) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *SchemasVersionResponse) HasComment() bool`

HasComment returns a boolean if a field has been set.

### SetCommentNil

`func (o *SchemasVersionResponse) SetCommentNil(b bool)`

 SetCommentNil sets the value for Comment to be an explicit nil

### UnsetComment
`func (o *SchemasVersionResponse) UnsetComment()`

UnsetComment ensures that no value is present for Comment, not even an explicit nil
### GetDeployed

`func (o *SchemasVersionResponse) GetDeployed() bool`

GetDeployed returns the Deployed field if non-nil, zero value otherwise.

### GetDeployedOk

`func (o *SchemasVersionResponse) GetDeployedOk() (*bool, bool)`

GetDeployedOk returns a tuple with the Deployed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeployed

`func (o *SchemasVersionResponse) SetDeployed(v bool)`

SetDeployed sets Deployed field to given value.

### HasDeployed

`func (o *SchemasVersionResponse) HasDeployed() bool`

HasDeployed returns a boolean if a field has been set.

### GetLocked

`func (o *SchemasVersionResponse) GetLocked() bool`

GetLocked returns the Locked field if non-nil, zero value otherwise.

### GetLockedOk

`func (o *SchemasVersionResponse) GetLockedOk() (*bool, bool)`

GetLockedOk returns a tuple with the Locked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocked

`func (o *SchemasVersionResponse) SetLocked(v bool)`

SetLocked sets Locked field to given value.

### HasLocked

`func (o *SchemasVersionResponse) HasLocked() bool`

HasLocked returns a boolean if a field has been set.

### GetNumber

`func (o *SchemasVersionResponse) GetNumber() int32`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *SchemasVersionResponse) GetNumberOk() (*int32, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *SchemasVersionResponse) SetNumber(v int32)`

SetNumber sets Number field to given value.

### HasNumber

`func (o *SchemasVersionResponse) HasNumber() bool`

HasNumber returns a boolean if a field has been set.

### GetStaging

`func (o *SchemasVersionResponse) GetStaging() bool`

GetStaging returns the Staging field if non-nil, zero value otherwise.

### GetStagingOk

`func (o *SchemasVersionResponse) GetStagingOk() (*bool, bool)`

GetStagingOk returns a tuple with the Staging field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStaging

`func (o *SchemasVersionResponse) SetStaging(v bool)`

SetStaging sets Staging field to given value.

### HasStaging

`func (o *SchemasVersionResponse) HasStaging() bool`

HasStaging returns a boolean if a field has been set.

### GetTesting

`func (o *SchemasVersionResponse) GetTesting() bool`

GetTesting returns the Testing field if non-nil, zero value otherwise.

### GetTestingOk

`func (o *SchemasVersionResponse) GetTestingOk() (*bool, bool)`

GetTestingOk returns a tuple with the Testing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTesting

`func (o *SchemasVersionResponse) SetTesting(v bool)`

SetTesting sets Testing field to given value.

### HasTesting

`func (o *SchemasVersionResponse) HasTesting() bool`

HasTesting returns a boolean if a field has been set.

### GetCreatedAt

`func (o *SchemasVersionResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *SchemasVersionResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *SchemasVersionResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *SchemasVersionResponse) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *SchemasVersionResponse) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *SchemasVersionResponse) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetDeletedAt

`func (o *SchemasVersionResponse) GetDeletedAt() time.Time`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *SchemasVersionResponse) GetDeletedAtOk() (*time.Time, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *SchemasVersionResponse) SetDeletedAt(v time.Time)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *SchemasVersionResponse) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### SetDeletedAtNil

`func (o *SchemasVersionResponse) SetDeletedAtNil(b bool)`

 SetDeletedAtNil sets the value for DeletedAt to be an explicit nil

### UnsetDeletedAt
`func (o *SchemasVersionResponse) UnsetDeletedAt()`

UnsetDeletedAt ensures that no value is present for DeletedAt, not even an explicit nil
### GetUpdatedAt

`func (o *SchemasVersionResponse) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *SchemasVersionResponse) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *SchemasVersionResponse) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *SchemasVersionResponse) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *SchemasVersionResponse) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *SchemasVersionResponse) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
### GetServiceID

`func (o *SchemasVersionResponse) GetServiceID() string`

GetServiceID returns the ServiceID field if non-nil, zero value otherwise.

### GetServiceIDOk

`func (o *SchemasVersionResponse) GetServiceIDOk() (*string, bool)`

GetServiceIDOk returns a tuple with the ServiceID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceID

`func (o *SchemasVersionResponse) SetServiceID(v string)`

SetServiceID sets ServiceID field to given value.

### HasServiceID

`func (o *SchemasVersionResponse) HasServiceID() bool`

HasServiceID returns a boolean if a field has been set.

### GetEnvironments

`func (o *SchemasVersionResponse) GetEnvironments() []Environment`

GetEnvironments returns the Environments field if non-nil, zero value otherwise.

### GetEnvironmentsOk

`func (o *SchemasVersionResponse) GetEnvironmentsOk() (*[]Environment, bool)`

GetEnvironmentsOk returns a tuple with the Environments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironments

`func (o *SchemasVersionResponse) SetEnvironments(v []Environment)`

SetEnvironments sets Environments field to given value.

### HasEnvironments

`func (o *SchemasVersionResponse) HasEnvironments() bool`

HasEnvironments returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
