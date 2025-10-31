# ComputeAclListEntriesItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Prefix** | Pointer to **string** | An IP prefix defined in Classless Inter-Domain Routing (CIDR) format, i.e. a valid IP address (v4 or v6) followed by a forward slash (/) and a prefix length (0-32 or 0-128, depending on address family). | [optional] 
**Action** | Pointer to **string** | One of \&quot;ALLOW\&quot; or \&quot;BLOCK\&quot;. | [optional] 

## Methods

### NewComputeAclListEntriesItem

`func NewComputeAclListEntriesItem() *ComputeAclListEntriesItem`

NewComputeAclListEntriesItem instantiates a new ComputeAclListEntriesItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComputeAclListEntriesItemWithDefaults

`func NewComputeAclListEntriesItemWithDefaults() *ComputeAclListEntriesItem`

NewComputeAclListEntriesItemWithDefaults instantiates a new ComputeAclListEntriesItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrefix

`func (o *ComputeAclListEntriesItem) GetPrefix() string`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *ComputeAclListEntriesItem) GetPrefixOk() (*string, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *ComputeAclListEntriesItem) SetPrefix(v string)`

SetPrefix sets Prefix field to given value.

### HasPrefix

`func (o *ComputeAclListEntriesItem) HasPrefix() bool`

HasPrefix returns a boolean if a field has been set.

### GetAction

`func (o *ComputeAclListEntriesItem) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *ComputeAclListEntriesItem) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *ComputeAclListEntriesItem) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *ComputeAclListEntriesItem) HasAction() bool`

HasAction returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


