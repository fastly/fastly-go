# ComputeACLUpdateEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Op** | Pointer to **string** | One of \&quot;create\&quot;, \&quot;update\&quot;, or \&quot;delete\&quot; indicating the operation to perform on the update. | [optional] 
**Prefix** | Pointer to **string** | An IP prefix defined in Classless Inter-Domain Routing (CIDR) format, i.e. a valid IP address (v4 or v6) followed by a forward slash (/) and a prefix length (0-32 or 0-128, depending on address family). | [optional] 
**Action** | Pointer to **string** | The action taken on the IP address, one of \&quot;BLOCK\&quot; or \&quot;ALLOW\&quot;. If using the \&quot;delete\&quot; operation, no action should be specified. | [optional] 

## Methods

### NewComputeACLUpdateEntry

`func NewComputeACLUpdateEntry() *ComputeACLUpdateEntry`

NewComputeACLUpdateEntry instantiates a new ComputeACLUpdateEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComputeACLUpdateEntryWithDefaults

`func NewComputeACLUpdateEntryWithDefaults() *ComputeACLUpdateEntry`

NewComputeACLUpdateEntryWithDefaults instantiates a new ComputeACLUpdateEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOp

`func (o *ComputeACLUpdateEntry) GetOp() string`

GetOp returns the Op field if non-nil, zero value otherwise.

### GetOpOk

`func (o *ComputeACLUpdateEntry) GetOpOk() (*string, bool)`

GetOpOk returns a tuple with the Op field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOp

`func (o *ComputeACLUpdateEntry) SetOp(v string)`

SetOp sets Op field to given value.

### HasOp

`func (o *ComputeACLUpdateEntry) HasOp() bool`

HasOp returns a boolean if a field has been set.

### GetPrefix

`func (o *ComputeACLUpdateEntry) GetPrefix() string`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *ComputeACLUpdateEntry) GetPrefixOk() (*string, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *ComputeACLUpdateEntry) SetPrefix(v string)`

SetPrefix sets Prefix field to given value.

### HasPrefix

`func (o *ComputeACLUpdateEntry) HasPrefix() bool`

HasPrefix returns a boolean if a field has been set.

### GetAction

`func (o *ComputeACLUpdateEntry) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *ComputeACLUpdateEntry) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *ComputeACLUpdateEntry) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *ComputeACLUpdateEntry) HasAction() bool`

HasAction returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
