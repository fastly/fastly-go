# InitialVersion

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Activate** | Pointer to **bool** | Whether to activate the initial version on creation. | [optional] [default to false]
**Comment** | Pointer to **string** | A freeform comment for the initial version. | [optional] 
**Paths** | Pointer to [**[]InitialVersionPath**](InitialVersionPath.md) | The paths to create on the initial version. | [optional] 

## Methods

### NewInitialVersion

`func NewInitialVersion() *InitialVersion`

NewInitialVersion instantiates a new InitialVersion object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInitialVersionWithDefaults

`func NewInitialVersionWithDefaults() *InitialVersion`

NewInitialVersionWithDefaults instantiates a new InitialVersion object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActivate

`func (o *InitialVersion) GetActivate() bool`

GetActivate returns the Activate field if non-nil, zero value otherwise.

### GetActivateOk

`func (o *InitialVersion) GetActivateOk() (*bool, bool)`

GetActivateOk returns a tuple with the Activate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivate

`func (o *InitialVersion) SetActivate(v bool)`

SetActivate sets Activate field to given value.

### HasActivate

`func (o *InitialVersion) HasActivate() bool`

HasActivate returns a boolean if a field has been set.

### GetComment

`func (o *InitialVersion) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *InitialVersion) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *InitialVersion) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *InitialVersion) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetPaths

`func (o *InitialVersion) GetPaths() []InitialVersionPath`

GetPaths returns the Paths field if non-nil, zero value otherwise.

### GetPathsOk

`func (o *InitialVersion) GetPathsOk() (*[]InitialVersionPath, bool)`

GetPathsOk returns a tuple with the Paths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaths

`func (o *InitialVersion) SetPaths(v []InitialVersionPath)`

SetPaths sets Paths field to given value.

### HasPaths

`func (o *InitialVersion) HasPaths() bool`

HasPaths returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


