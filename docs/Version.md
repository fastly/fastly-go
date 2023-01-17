# Version

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

### NewVersion

`func NewVersion() *Version`

NewVersion instantiates a new Version object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVersionWithDefaults

`func NewVersionWithDefaults() *Version`

NewVersionWithDefaults instantiates a new Version object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActive

`func (o *Version) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *Version) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *Version) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *Version) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetComment

`func (o *Version) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *Version) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *Version) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *Version) HasComment() bool`

HasComment returns a boolean if a field has been set.

### SetCommentNil

`func (o *Version) SetCommentNil(b bool)`

 SetCommentNil sets the value for Comment to be an explicit nil

### UnsetComment
`func (o *Version) UnsetComment()`

UnsetComment ensures that no value is present for Comment, not even an explicit nil
### GetDeployed

`func (o *Version) GetDeployed() bool`

GetDeployed returns the Deployed field if non-nil, zero value otherwise.

### GetDeployedOk

`func (o *Version) GetDeployedOk() (*bool, bool)`

GetDeployedOk returns a tuple with the Deployed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeployed

`func (o *Version) SetDeployed(v bool)`

SetDeployed sets Deployed field to given value.

### HasDeployed

`func (o *Version) HasDeployed() bool`

HasDeployed returns a boolean if a field has been set.

### GetLocked

`func (o *Version) GetLocked() bool`

GetLocked returns the Locked field if non-nil, zero value otherwise.

### GetLockedOk

`func (o *Version) GetLockedOk() (*bool, bool)`

GetLockedOk returns a tuple with the Locked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocked

`func (o *Version) SetLocked(v bool)`

SetLocked sets Locked field to given value.

### HasLocked

`func (o *Version) HasLocked() bool`

HasLocked returns a boolean if a field has been set.

### GetNumber

`func (o *Version) GetNumber() int32`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *Version) GetNumberOk() (*int32, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *Version) SetNumber(v int32)`

SetNumber sets Number field to given value.

### HasNumber

`func (o *Version) HasNumber() bool`

HasNumber returns a boolean if a field has been set.

### GetStaging

`func (o *Version) GetStaging() bool`

GetStaging returns the Staging field if non-nil, zero value otherwise.

### GetStagingOk

`func (o *Version) GetStagingOk() (*bool, bool)`

GetStagingOk returns a tuple with the Staging field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStaging

`func (o *Version) SetStaging(v bool)`

SetStaging sets Staging field to given value.

### HasStaging

`func (o *Version) HasStaging() bool`

HasStaging returns a boolean if a field has been set.

### GetTesting

`func (o *Version) GetTesting() bool`

GetTesting returns the Testing field if non-nil, zero value otherwise.

### GetTestingOk

`func (o *Version) GetTestingOk() (*bool, bool)`

GetTestingOk returns a tuple with the Testing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTesting

`func (o *Version) SetTesting(v bool)`

SetTesting sets Testing field to given value.

### HasTesting

`func (o *Version) HasTesting() bool`

HasTesting returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
