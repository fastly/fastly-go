# SchemasVersion

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

## Methods

### NewSchemasVersion

`func NewSchemasVersion() *SchemasVersion`

NewSchemasVersion instantiates a new SchemasVersion object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSchemasVersionWithDefaults

`func NewSchemasVersionWithDefaults() *SchemasVersion`

NewSchemasVersionWithDefaults instantiates a new SchemasVersion object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActive

`func (o *SchemasVersion) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *SchemasVersion) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *SchemasVersion) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *SchemasVersion) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetComment

`func (o *SchemasVersion) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *SchemasVersion) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *SchemasVersion) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *SchemasVersion) HasComment() bool`

HasComment returns a boolean if a field has been set.

### SetCommentNil

`func (o *SchemasVersion) SetCommentNil(b bool)`

 SetCommentNil sets the value for Comment to be an explicit nil

### UnsetComment
`func (o *SchemasVersion) UnsetComment()`

UnsetComment ensures that no value is present for Comment, not even an explicit nil
### GetDeployed

`func (o *SchemasVersion) GetDeployed() bool`

GetDeployed returns the Deployed field if non-nil, zero value otherwise.

### GetDeployedOk

`func (o *SchemasVersion) GetDeployedOk() (*bool, bool)`

GetDeployedOk returns a tuple with the Deployed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeployed

`func (o *SchemasVersion) SetDeployed(v bool)`

SetDeployed sets Deployed field to given value.

### HasDeployed

`func (o *SchemasVersion) HasDeployed() bool`

HasDeployed returns a boolean if a field has been set.

### GetLocked

`func (o *SchemasVersion) GetLocked() bool`

GetLocked returns the Locked field if non-nil, zero value otherwise.

### GetLockedOk

`func (o *SchemasVersion) GetLockedOk() (*bool, bool)`

GetLockedOk returns a tuple with the Locked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocked

`func (o *SchemasVersion) SetLocked(v bool)`

SetLocked sets Locked field to given value.

### HasLocked

`func (o *SchemasVersion) HasLocked() bool`

HasLocked returns a boolean if a field has been set.

### GetNumber

`func (o *SchemasVersion) GetNumber() int32`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *SchemasVersion) GetNumberOk() (*int32, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *SchemasVersion) SetNumber(v int32)`

SetNumber sets Number field to given value.

### HasNumber

`func (o *SchemasVersion) HasNumber() bool`

HasNumber returns a boolean if a field has been set.

### GetStaging

`func (o *SchemasVersion) GetStaging() bool`

GetStaging returns the Staging field if non-nil, zero value otherwise.

### GetStagingOk

`func (o *SchemasVersion) GetStagingOk() (*bool, bool)`

GetStagingOk returns a tuple with the Staging field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStaging

`func (o *SchemasVersion) SetStaging(v bool)`

SetStaging sets Staging field to given value.

### HasStaging

`func (o *SchemasVersion) HasStaging() bool`

HasStaging returns a boolean if a field has been set.

### GetTesting

`func (o *SchemasVersion) GetTesting() bool`

GetTesting returns the Testing field if non-nil, zero value otherwise.

### GetTestingOk

`func (o *SchemasVersion) GetTestingOk() (*bool, bool)`

GetTestingOk returns a tuple with the Testing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTesting

`func (o *SchemasVersion) SetTesting(v bool)`

SetTesting sets Testing field to given value.

### HasTesting

`func (o *SchemasVersion) HasTesting() bool`

HasTesting returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


