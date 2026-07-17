# ScriptUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthorizationStatus** | Pointer to **string** | Script authorization status | [optional] 
**Justification** | Pointer to **string** | Reason for authorization decision | [optional] 
**AuthorizedHash** | Pointer to **string** | Hash of authorized script content | [optional] 

## Methods

### NewScriptUpdate

`func NewScriptUpdate() *ScriptUpdate`

NewScriptUpdate instantiates a new ScriptUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScriptUpdateWithDefaults

`func NewScriptUpdateWithDefaults() *ScriptUpdate`

NewScriptUpdateWithDefaults instantiates a new ScriptUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthorizationStatus

`func (o *ScriptUpdate) GetAuthorizationStatus() string`

GetAuthorizationStatus returns the AuthorizationStatus field if non-nil, zero value otherwise.

### GetAuthorizationStatusOk

`func (o *ScriptUpdate) GetAuthorizationStatusOk() (*string, bool)`

GetAuthorizationStatusOk returns a tuple with the AuthorizationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationStatus

`func (o *ScriptUpdate) SetAuthorizationStatus(v string)`

SetAuthorizationStatus sets AuthorizationStatus field to given value.

### HasAuthorizationStatus

`func (o *ScriptUpdate) HasAuthorizationStatus() bool`

HasAuthorizationStatus returns a boolean if a field has been set.

### GetJustification

`func (o *ScriptUpdate) GetJustification() string`

GetJustification returns the Justification field if non-nil, zero value otherwise.

### GetJustificationOk

`func (o *ScriptUpdate) GetJustificationOk() (*string, bool)`

GetJustificationOk returns a tuple with the Justification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJustification

`func (o *ScriptUpdate) SetJustification(v string)`

SetJustification sets Justification field to given value.

### HasJustification

`func (o *ScriptUpdate) HasJustification() bool`

HasJustification returns a boolean if a field has been set.

### GetAuthorizedHash

`func (o *ScriptUpdate) GetAuthorizedHash() string`

GetAuthorizedHash returns the AuthorizedHash field if non-nil, zero value otherwise.

### GetAuthorizedHashOk

`func (o *ScriptUpdate) GetAuthorizedHashOk() (*string, bool)`

GetAuthorizedHashOk returns a tuple with the AuthorizedHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizedHash

`func (o *ScriptUpdate) SetAuthorizedHash(v string)`

SetAuthorizedHash sets AuthorizedHash field to given value.

### HasAuthorizedHash

`func (o *ScriptUpdate) HasAuthorizedHash() bool`

HasAuthorizedHash returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


