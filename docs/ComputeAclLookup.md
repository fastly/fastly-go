# ComputeACLLookup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Prefix** | Pointer to **string** | A valid IPv4 or IPv6 address and prefix in CIDR notation. | [optional] 
**Action** | Pointer to **string** | One of \&quot;ALLOW\&quot; or \&quot;BLOCK\&quot;. | [optional] 

## Methods

### NewComputeACLLookup

`func NewComputeACLLookup() *ComputeACLLookup`

NewComputeACLLookup instantiates a new ComputeACLLookup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComputeACLLookupWithDefaults

`func NewComputeACLLookupWithDefaults() *ComputeACLLookup`

NewComputeACLLookupWithDefaults instantiates a new ComputeACLLookup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrefix

`func (o *ComputeACLLookup) GetPrefix() string`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *ComputeACLLookup) GetPrefixOk() (*string, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *ComputeACLLookup) SetPrefix(v string)`

SetPrefix sets Prefix field to given value.

### HasPrefix

`func (o *ComputeACLLookup) HasPrefix() bool`

HasPrefix returns a boolean if a field has been set.

### GetAction

`func (o *ComputeACLLookup) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *ComputeACLLookup) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *ComputeACLLookup) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *ComputeACLLookup) HasAction() bool`

HasAction returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
