# HeaderResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | Pointer to **string** | Accepts a string value. | [optional] 
**CacheCondition** | Pointer to **NullableString** | Name of the cache condition controlling when this configuration applies. | [optional] 
**Dst** | Pointer to **string** | Header to set. | [optional] 
**Name** | Pointer to **string** | A handle to refer to this Header object. | [optional] 
**Regex** | Pointer to **NullableString** | Regular expression to use. Only applies to `regex` and `regex_repeat` actions. | [optional] 
**RequestCondition** | Pointer to **NullableString** | Condition which, if met, will select this configuration during a request. Optional. | [optional] 
**ResponseCondition** | Pointer to **NullableString** | Optional name of a response condition to apply. | [optional] 
**Src** | Pointer to **NullableString** | Variable to be used as a source for the header content. Does not apply to `delete` action. | [optional] 
**Substitution** | Pointer to **NullableString** | Value to substitute in place of regular expression. Only applies to `regex` and `regex_repeat` actions. | [optional] 
**Type** | Pointer to **string** | Accepts a string value. | [optional] 
**IgnoreIfSet** | Pointer to **string** | Don&#39;t add the header if it is added already. Only applies to &#39;set&#39; action. Numerical value (\&quot;0\&quot; &#x3D; false, \&quot;1\&quot; &#x3D; true) | [optional] 
**Priority** | Pointer to **string** | Priority determines execution order. Lower numbers execute first. | [optional] [default to "100"]
**ServiceID** | Pointer to **string** |  | [optional] [readonly] 
**Version** | Pointer to **string** |  | [optional] [readonly] 
**CreatedAt** | Pointer to **NullableTime** | Date and time in ISO 8601 format. | [optional] [readonly] 
**DeletedAt** | Pointer to **NullableTime** | Date and time in ISO 8601 format. | [optional] [readonly] 
**UpdatedAt** | Pointer to **NullableTime** | Date and time in ISO 8601 format. | [optional] [readonly] 

## Methods

### NewHeaderResponse

`func NewHeaderResponse() *HeaderResponse`

NewHeaderResponse instantiates a new HeaderResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHeaderResponseWithDefaults

`func NewHeaderResponseWithDefaults() *HeaderResponse`

NewHeaderResponseWithDefaults instantiates a new HeaderResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *HeaderResponse) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *HeaderResponse) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *HeaderResponse) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *HeaderResponse) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetCacheCondition

`func (o *HeaderResponse) GetCacheCondition() string`

GetCacheCondition returns the CacheCondition field if non-nil, zero value otherwise.

### GetCacheConditionOk

`func (o *HeaderResponse) GetCacheConditionOk() (*string, bool)`

GetCacheConditionOk returns a tuple with the CacheCondition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheCondition

`func (o *HeaderResponse) SetCacheCondition(v string)`

SetCacheCondition sets CacheCondition field to given value.

### HasCacheCondition

`func (o *HeaderResponse) HasCacheCondition() bool`

HasCacheCondition returns a boolean if a field has been set.

### SetCacheConditionNil

`func (o *HeaderResponse) SetCacheConditionNil(b bool)`

 SetCacheConditionNil sets the value for CacheCondition to be an explicit nil

### UnsetCacheCondition
`func (o *HeaderResponse) UnsetCacheCondition()`

UnsetCacheCondition ensures that no value is present for CacheCondition, not even an explicit nil
### GetDst

`func (o *HeaderResponse) GetDst() string`

GetDst returns the Dst field if non-nil, zero value otherwise.

### GetDstOk

`func (o *HeaderResponse) GetDstOk() (*string, bool)`

GetDstOk returns a tuple with the Dst field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDst

`func (o *HeaderResponse) SetDst(v string)`

SetDst sets Dst field to given value.

### HasDst

`func (o *HeaderResponse) HasDst() bool`

HasDst returns a boolean if a field has been set.

### GetName

`func (o *HeaderResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *HeaderResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *HeaderResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *HeaderResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRegex

`func (o *HeaderResponse) GetRegex() string`

GetRegex returns the Regex field if non-nil, zero value otherwise.

### GetRegexOk

`func (o *HeaderResponse) GetRegexOk() (*string, bool)`

GetRegexOk returns a tuple with the Regex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegex

`func (o *HeaderResponse) SetRegex(v string)`

SetRegex sets Regex field to given value.

### HasRegex

`func (o *HeaderResponse) HasRegex() bool`

HasRegex returns a boolean if a field has been set.

### SetRegexNil

`func (o *HeaderResponse) SetRegexNil(b bool)`

 SetRegexNil sets the value for Regex to be an explicit nil

### UnsetRegex
`func (o *HeaderResponse) UnsetRegex()`

UnsetRegex ensures that no value is present for Regex, not even an explicit nil
### GetRequestCondition

`func (o *HeaderResponse) GetRequestCondition() string`

GetRequestCondition returns the RequestCondition field if non-nil, zero value otherwise.

### GetRequestConditionOk

`func (o *HeaderResponse) GetRequestConditionOk() (*string, bool)`

GetRequestConditionOk returns a tuple with the RequestCondition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestCondition

`func (o *HeaderResponse) SetRequestCondition(v string)`

SetRequestCondition sets RequestCondition field to given value.

### HasRequestCondition

`func (o *HeaderResponse) HasRequestCondition() bool`

HasRequestCondition returns a boolean if a field has been set.

### SetRequestConditionNil

`func (o *HeaderResponse) SetRequestConditionNil(b bool)`

 SetRequestConditionNil sets the value for RequestCondition to be an explicit nil

### UnsetRequestCondition
`func (o *HeaderResponse) UnsetRequestCondition()`

UnsetRequestCondition ensures that no value is present for RequestCondition, not even an explicit nil
### GetResponseCondition

`func (o *HeaderResponse) GetResponseCondition() string`

GetResponseCondition returns the ResponseCondition field if non-nil, zero value otherwise.

### GetResponseConditionOk

`func (o *HeaderResponse) GetResponseConditionOk() (*string, bool)`

GetResponseConditionOk returns a tuple with the ResponseCondition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseCondition

`func (o *HeaderResponse) SetResponseCondition(v string)`

SetResponseCondition sets ResponseCondition field to given value.

### HasResponseCondition

`func (o *HeaderResponse) HasResponseCondition() bool`

HasResponseCondition returns a boolean if a field has been set.

### SetResponseConditionNil

`func (o *HeaderResponse) SetResponseConditionNil(b bool)`

 SetResponseConditionNil sets the value for ResponseCondition to be an explicit nil

### UnsetResponseCondition
`func (o *HeaderResponse) UnsetResponseCondition()`

UnsetResponseCondition ensures that no value is present for ResponseCondition, not even an explicit nil
### GetSrc

`func (o *HeaderResponse) GetSrc() string`

GetSrc returns the Src field if non-nil, zero value otherwise.

### GetSrcOk

`func (o *HeaderResponse) GetSrcOk() (*string, bool)`

GetSrcOk returns a tuple with the Src field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSrc

`func (o *HeaderResponse) SetSrc(v string)`

SetSrc sets Src field to given value.

### HasSrc

`func (o *HeaderResponse) HasSrc() bool`

HasSrc returns a boolean if a field has been set.

### SetSrcNil

`func (o *HeaderResponse) SetSrcNil(b bool)`

 SetSrcNil sets the value for Src to be an explicit nil

### UnsetSrc
`func (o *HeaderResponse) UnsetSrc()`

UnsetSrc ensures that no value is present for Src, not even an explicit nil
### GetSubstitution

`func (o *HeaderResponse) GetSubstitution() string`

GetSubstitution returns the Substitution field if non-nil, zero value otherwise.

### GetSubstitutionOk

`func (o *HeaderResponse) GetSubstitutionOk() (*string, bool)`

GetSubstitutionOk returns a tuple with the Substitution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubstitution

`func (o *HeaderResponse) SetSubstitution(v string)`

SetSubstitution sets Substitution field to given value.

### HasSubstitution

`func (o *HeaderResponse) HasSubstitution() bool`

HasSubstitution returns a boolean if a field has been set.

### SetSubstitutionNil

`func (o *HeaderResponse) SetSubstitutionNil(b bool)`

 SetSubstitutionNil sets the value for Substitution to be an explicit nil

### UnsetSubstitution
`func (o *HeaderResponse) UnsetSubstitution()`

UnsetSubstitution ensures that no value is present for Substitution, not even an explicit nil
### GetType

`func (o *HeaderResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *HeaderResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *HeaderResponse) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *HeaderResponse) HasType() bool`

HasType returns a boolean if a field has been set.

### GetIgnoreIfSet

`func (o *HeaderResponse) GetIgnoreIfSet() string`

GetIgnoreIfSet returns the IgnoreIfSet field if non-nil, zero value otherwise.

### GetIgnoreIfSetOk

`func (o *HeaderResponse) GetIgnoreIfSetOk() (*string, bool)`

GetIgnoreIfSetOk returns a tuple with the IgnoreIfSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoreIfSet

`func (o *HeaderResponse) SetIgnoreIfSet(v string)`

SetIgnoreIfSet sets IgnoreIfSet field to given value.

### HasIgnoreIfSet

`func (o *HeaderResponse) HasIgnoreIfSet() bool`

HasIgnoreIfSet returns a boolean if a field has been set.

### GetPriority

`func (o *HeaderResponse) GetPriority() string`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *HeaderResponse) GetPriorityOk() (*string, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *HeaderResponse) SetPriority(v string)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *HeaderResponse) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetServiceID

`func (o *HeaderResponse) GetServiceID() string`

GetServiceID returns the ServiceID field if non-nil, zero value otherwise.

### GetServiceIDOk

`func (o *HeaderResponse) GetServiceIDOk() (*string, bool)`

GetServiceIDOk returns a tuple with the ServiceID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceID

`func (o *HeaderResponse) SetServiceID(v string)`

SetServiceID sets ServiceID field to given value.

### HasServiceID

`func (o *HeaderResponse) HasServiceID() bool`

HasServiceID returns a boolean if a field has been set.

### GetVersion

`func (o *HeaderResponse) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *HeaderResponse) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *HeaderResponse) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *HeaderResponse) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetCreatedAt

`func (o *HeaderResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *HeaderResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *HeaderResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *HeaderResponse) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *HeaderResponse) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *HeaderResponse) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetDeletedAt

`func (o *HeaderResponse) GetDeletedAt() time.Time`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *HeaderResponse) GetDeletedAtOk() (*time.Time, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *HeaderResponse) SetDeletedAt(v time.Time)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *HeaderResponse) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### SetDeletedAtNil

`func (o *HeaderResponse) SetDeletedAtNil(b bool)`

 SetDeletedAtNil sets the value for DeletedAt to be an explicit nil

### UnsetDeletedAt
`func (o *HeaderResponse) UnsetDeletedAt()`

UnsetDeletedAt ensures that no value is present for DeletedAt, not even an explicit nil
### GetUpdatedAt

`func (o *HeaderResponse) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *HeaderResponse) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *HeaderResponse) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *HeaderResponse) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *HeaderResponse) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *HeaderResponse) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil

[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
