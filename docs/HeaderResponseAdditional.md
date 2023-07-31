# HeaderResponseAdditional

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IgnoreIfSet** | Pointer to **string** | Don&#39;t add the header if it is added already. Only applies to &#39;set&#39; action. Numerical value (\&quot;0\&quot; &#x3D; false, \&quot;1\&quot; &#x3D; true) | [optional] 
**Priority** | Pointer to **string** | Priority determines execution order. Lower numbers execute first. | [optional] [default to "100"]

## Methods

### NewHeaderResponseAdditional

`func NewHeaderResponseAdditional() *HeaderResponseAdditional`

NewHeaderResponseAdditional instantiates a new HeaderResponseAdditional object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHeaderResponseAdditionalWithDefaults

`func NewHeaderResponseAdditionalWithDefaults() *HeaderResponseAdditional`

NewHeaderResponseAdditionalWithDefaults instantiates a new HeaderResponseAdditional object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIgnoreIfSet

`func (o *HeaderResponseAdditional) GetIgnoreIfSet() string`

GetIgnoreIfSet returns the IgnoreIfSet field if non-nil, zero value otherwise.

### GetIgnoreIfSetOk

`func (o *HeaderResponseAdditional) GetIgnoreIfSetOk() (*string, bool)`

GetIgnoreIfSetOk returns a tuple with the IgnoreIfSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoreIfSet

`func (o *HeaderResponseAdditional) SetIgnoreIfSet(v string)`

SetIgnoreIfSet sets IgnoreIfSet field to given value.

### HasIgnoreIfSet

`func (o *HeaderResponseAdditional) HasIgnoreIfSet() bool`

HasIgnoreIfSet returns a boolean if a field has been set.

### GetPriority

`func (o *HeaderResponseAdditional) GetPriority() string`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *HeaderResponseAdditional) GetPriorityOk() (*string, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *HeaderResponseAdditional) SetPriority(v string)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *HeaderResponseAdditional) HasPriority() bool`

HasPriority returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
