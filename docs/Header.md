# Header

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

## Methods

### NewHeader

`func NewHeader() *Header`

NewHeader instantiates a new Header object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHeaderWithDefaults

`func NewHeaderWithDefaults() *Header`

NewHeaderWithDefaults instantiates a new Header object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *Header) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *Header) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *Header) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *Header) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetCacheCondition

`func (o *Header) GetCacheCondition() string`

GetCacheCondition returns the CacheCondition field if non-nil, zero value otherwise.

### GetCacheConditionOk

`func (o *Header) GetCacheConditionOk() (*string, bool)`

GetCacheConditionOk returns a tuple with the CacheCondition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheCondition

`func (o *Header) SetCacheCondition(v string)`

SetCacheCondition sets CacheCondition field to given value.

### HasCacheCondition

`func (o *Header) HasCacheCondition() bool`

HasCacheCondition returns a boolean if a field has been set.

### SetCacheConditionNil

`func (o *Header) SetCacheConditionNil(b bool)`

 SetCacheConditionNil sets the value for CacheCondition to be an explicit nil

### UnsetCacheCondition
`func (o *Header) UnsetCacheCondition()`

UnsetCacheCondition ensures that no value is present for CacheCondition, not even an explicit nil
### GetDst

`func (o *Header) GetDst() string`

GetDst returns the Dst field if non-nil, zero value otherwise.

### GetDstOk

`func (o *Header) GetDstOk() (*string, bool)`

GetDstOk returns a tuple with the Dst field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDst

`func (o *Header) SetDst(v string)`

SetDst sets Dst field to given value.

### HasDst

`func (o *Header) HasDst() bool`

HasDst returns a boolean if a field has been set.

### GetName

`func (o *Header) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Header) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Header) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Header) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRegex

`func (o *Header) GetRegex() string`

GetRegex returns the Regex field if non-nil, zero value otherwise.

### GetRegexOk

`func (o *Header) GetRegexOk() (*string, bool)`

GetRegexOk returns a tuple with the Regex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegex

`func (o *Header) SetRegex(v string)`

SetRegex sets Regex field to given value.

### HasRegex

`func (o *Header) HasRegex() bool`

HasRegex returns a boolean if a field has been set.

### SetRegexNil

`func (o *Header) SetRegexNil(b bool)`

 SetRegexNil sets the value for Regex to be an explicit nil

### UnsetRegex
`func (o *Header) UnsetRegex()`

UnsetRegex ensures that no value is present for Regex, not even an explicit nil
### GetRequestCondition

`func (o *Header) GetRequestCondition() string`

GetRequestCondition returns the RequestCondition field if non-nil, zero value otherwise.

### GetRequestConditionOk

`func (o *Header) GetRequestConditionOk() (*string, bool)`

GetRequestConditionOk returns a tuple with the RequestCondition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestCondition

`func (o *Header) SetRequestCondition(v string)`

SetRequestCondition sets RequestCondition field to given value.

### HasRequestCondition

`func (o *Header) HasRequestCondition() bool`

HasRequestCondition returns a boolean if a field has been set.

### SetRequestConditionNil

`func (o *Header) SetRequestConditionNil(b bool)`

 SetRequestConditionNil sets the value for RequestCondition to be an explicit nil

### UnsetRequestCondition
`func (o *Header) UnsetRequestCondition()`

UnsetRequestCondition ensures that no value is present for RequestCondition, not even an explicit nil
### GetResponseCondition

`func (o *Header) GetResponseCondition() string`

GetResponseCondition returns the ResponseCondition field if non-nil, zero value otherwise.

### GetResponseConditionOk

`func (o *Header) GetResponseConditionOk() (*string, bool)`

GetResponseConditionOk returns a tuple with the ResponseCondition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseCondition

`func (o *Header) SetResponseCondition(v string)`

SetResponseCondition sets ResponseCondition field to given value.

### HasResponseCondition

`func (o *Header) HasResponseCondition() bool`

HasResponseCondition returns a boolean if a field has been set.

### SetResponseConditionNil

`func (o *Header) SetResponseConditionNil(b bool)`

 SetResponseConditionNil sets the value for ResponseCondition to be an explicit nil

### UnsetResponseCondition
`func (o *Header) UnsetResponseCondition()`

UnsetResponseCondition ensures that no value is present for ResponseCondition, not even an explicit nil
### GetSrc

`func (o *Header) GetSrc() string`

GetSrc returns the Src field if non-nil, zero value otherwise.

### GetSrcOk

`func (o *Header) GetSrcOk() (*string, bool)`

GetSrcOk returns a tuple with the Src field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSrc

`func (o *Header) SetSrc(v string)`

SetSrc sets Src field to given value.

### HasSrc

`func (o *Header) HasSrc() bool`

HasSrc returns a boolean if a field has been set.

### SetSrcNil

`func (o *Header) SetSrcNil(b bool)`

 SetSrcNil sets the value for Src to be an explicit nil

### UnsetSrc
`func (o *Header) UnsetSrc()`

UnsetSrc ensures that no value is present for Src, not even an explicit nil
### GetSubstitution

`func (o *Header) GetSubstitution() string`

GetSubstitution returns the Substitution field if non-nil, zero value otherwise.

### GetSubstitutionOk

`func (o *Header) GetSubstitutionOk() (*string, bool)`

GetSubstitutionOk returns a tuple with the Substitution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubstitution

`func (o *Header) SetSubstitution(v string)`

SetSubstitution sets Substitution field to given value.

### HasSubstitution

`func (o *Header) HasSubstitution() bool`

HasSubstitution returns a boolean if a field has been set.

### SetSubstitutionNil

`func (o *Header) SetSubstitutionNil(b bool)`

 SetSubstitutionNil sets the value for Substitution to be an explicit nil

### UnsetSubstitution
`func (o *Header) UnsetSubstitution()`

UnsetSubstitution ensures that no value is present for Substitution, not even an explicit nil
### GetType

`func (o *Header) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Header) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Header) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Header) HasType() bool`

HasType returns a boolean if a field has been set.

### GetIgnoreIfSet

`func (o *Header) GetIgnoreIfSet() string`

GetIgnoreIfSet returns the IgnoreIfSet field if non-nil, zero value otherwise.

### GetIgnoreIfSetOk

`func (o *Header) GetIgnoreIfSetOk() (*string, bool)`

GetIgnoreIfSetOk returns a tuple with the IgnoreIfSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoreIfSet

`func (o *Header) SetIgnoreIfSet(v string)`

SetIgnoreIfSet sets IgnoreIfSet field to given value.

### HasIgnoreIfSet

`func (o *Header) HasIgnoreIfSet() bool`

HasIgnoreIfSet returns a boolean if a field has been set.

### GetPriority

`func (o *Header) GetPriority() string`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *Header) GetPriorityOk() (*string, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *Header) SetPriority(v string)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *Header) HasPriority() bool`

HasPriority returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


