# RequestSettingsAdditional

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | Pointer to **NullableString** | Allows you to terminate request handling and immediately perform an action. | [optional] 
**DefaultHost** | Pointer to **NullableString** | Sets the host header. | [optional] 
**HashKeys** | Pointer to **NullableString** | Comma separated list of varnish request object fields that should be in the hash key. | [optional] 
**Name** | Pointer to **string** | Name for the request settings. | [optional] 
**RequestCondition** | Pointer to **NullableString** | Condition which, if met, will select this configuration during a request. Optional. | [optional] 
**Xff** | Pointer to **NullableString** | Short for X-Forwarded-For. | [optional] 

## Methods

### NewRequestSettingsAdditional

`func NewRequestSettingsAdditional() *RequestSettingsAdditional`

NewRequestSettingsAdditional instantiates a new RequestSettingsAdditional object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestSettingsAdditionalWithDefaults

`func NewRequestSettingsAdditionalWithDefaults() *RequestSettingsAdditional`

NewRequestSettingsAdditionalWithDefaults instantiates a new RequestSettingsAdditional object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *RequestSettingsAdditional) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *RequestSettingsAdditional) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *RequestSettingsAdditional) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *RequestSettingsAdditional) HasAction() bool`

HasAction returns a boolean if a field has been set.

### SetActionNil

`func (o *RequestSettingsAdditional) SetActionNil(b bool)`

 SetActionNil sets the value for Action to be an explicit nil

### UnsetAction
`func (o *RequestSettingsAdditional) UnsetAction()`

UnsetAction ensures that no value is present for Action, not even an explicit nil
### GetDefaultHost

`func (o *RequestSettingsAdditional) GetDefaultHost() string`

GetDefaultHost returns the DefaultHost field if non-nil, zero value otherwise.

### GetDefaultHostOk

`func (o *RequestSettingsAdditional) GetDefaultHostOk() (*string, bool)`

GetDefaultHostOk returns a tuple with the DefaultHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultHost

`func (o *RequestSettingsAdditional) SetDefaultHost(v string)`

SetDefaultHost sets DefaultHost field to given value.

### HasDefaultHost

`func (o *RequestSettingsAdditional) HasDefaultHost() bool`

HasDefaultHost returns a boolean if a field has been set.

### SetDefaultHostNil

`func (o *RequestSettingsAdditional) SetDefaultHostNil(b bool)`

 SetDefaultHostNil sets the value for DefaultHost to be an explicit nil

### UnsetDefaultHost
`func (o *RequestSettingsAdditional) UnsetDefaultHost()`

UnsetDefaultHost ensures that no value is present for DefaultHost, not even an explicit nil
### GetHashKeys

`func (o *RequestSettingsAdditional) GetHashKeys() string`

GetHashKeys returns the HashKeys field if non-nil, zero value otherwise.

### GetHashKeysOk

`func (o *RequestSettingsAdditional) GetHashKeysOk() (*string, bool)`

GetHashKeysOk returns a tuple with the HashKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHashKeys

`func (o *RequestSettingsAdditional) SetHashKeys(v string)`

SetHashKeys sets HashKeys field to given value.

### HasHashKeys

`func (o *RequestSettingsAdditional) HasHashKeys() bool`

HasHashKeys returns a boolean if a field has been set.

### SetHashKeysNil

`func (o *RequestSettingsAdditional) SetHashKeysNil(b bool)`

 SetHashKeysNil sets the value for HashKeys to be an explicit nil

### UnsetHashKeys
`func (o *RequestSettingsAdditional) UnsetHashKeys()`

UnsetHashKeys ensures that no value is present for HashKeys, not even an explicit nil
### GetName

`func (o *RequestSettingsAdditional) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RequestSettingsAdditional) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RequestSettingsAdditional) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *RequestSettingsAdditional) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRequestCondition

`func (o *RequestSettingsAdditional) GetRequestCondition() string`

GetRequestCondition returns the RequestCondition field if non-nil, zero value otherwise.

### GetRequestConditionOk

`func (o *RequestSettingsAdditional) GetRequestConditionOk() (*string, bool)`

GetRequestConditionOk returns a tuple with the RequestCondition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestCondition

`func (o *RequestSettingsAdditional) SetRequestCondition(v string)`

SetRequestCondition sets RequestCondition field to given value.

### HasRequestCondition

`func (o *RequestSettingsAdditional) HasRequestCondition() bool`

HasRequestCondition returns a boolean if a field has been set.

### SetRequestConditionNil

`func (o *RequestSettingsAdditional) SetRequestConditionNil(b bool)`

 SetRequestConditionNil sets the value for RequestCondition to be an explicit nil

### UnsetRequestCondition
`func (o *RequestSettingsAdditional) UnsetRequestCondition()`

UnsetRequestCondition ensures that no value is present for RequestCondition, not even an explicit nil
### GetXff

`func (o *RequestSettingsAdditional) GetXff() string`

GetXff returns the Xff field if non-nil, zero value otherwise.

### GetXffOk

`func (o *RequestSettingsAdditional) GetXffOk() (*string, bool)`

GetXffOk returns a tuple with the Xff field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXff

`func (o *RequestSettingsAdditional) SetXff(v string)`

SetXff sets Xff field to given value.

### HasXff

`func (o *RequestSettingsAdditional) HasXff() bool`

HasXff returns a boolean if a field has been set.

### SetXffNil

`func (o *RequestSettingsAdditional) SetXffNil(b bool)`

 SetXffNil sets the value for Xff to be an explicit nil

### UnsetXff
`func (o *RequestSettingsAdditional) UnsetXff()`

UnsetXff ensures that no value is present for Xff, not even an explicit nil

[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
