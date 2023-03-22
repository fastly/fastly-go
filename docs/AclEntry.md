# ACLEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Negated** | Pointer to **int32** | Whether to negate the match. Useful primarily when creating individual exceptions to larger subnets. | [optional] [default to 0]
**Comment** | Pointer to **NullableString** | A freeform descriptive note. | [optional] 
**IP** | Pointer to **string** | An IP address. | [optional] 
**Subnet** | Pointer to **NullableInt32** | Number of bits for the subnet mask applied to the IP address. For IPv4 addresses, a value of 32 represents the smallest subnet mask (1 address), 24 represents a class C subnet mask (256 addresses), 16 represents a class B subnet mask (65k addresses), and 8 is class A subnet mask (16m addresses). If not provided, no mask is applied. | [optional] 

## Methods

### NewACLEntry

`func NewACLEntry() *ACLEntry`

NewACLEntry instantiates a new ACLEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewACLEntryWithDefaults

`func NewACLEntryWithDefaults() *ACLEntry`

NewACLEntryWithDefaults instantiates a new ACLEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNegated

`func (o *ACLEntry) GetNegated() int32`

GetNegated returns the Negated field if non-nil, zero value otherwise.

### GetNegatedOk

`func (o *ACLEntry) GetNegatedOk() (*int32, bool)`

GetNegatedOk returns a tuple with the Negated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNegated

`func (o *ACLEntry) SetNegated(v int32)`

SetNegated sets Negated field to given value.

### HasNegated

`func (o *ACLEntry) HasNegated() bool`

HasNegated returns a boolean if a field has been set.

### GetComment

`func (o *ACLEntry) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *ACLEntry) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *ACLEntry) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *ACLEntry) HasComment() bool`

HasComment returns a boolean if a field has been set.

### SetCommentNil

`func (o *ACLEntry) SetCommentNil(b bool)`

 SetCommentNil sets the value for Comment to be an explicit nil

### UnsetComment
`func (o *ACLEntry) UnsetComment()`

UnsetComment ensures that no value is present for Comment, not even an explicit nil
### GetIP

`func (o *ACLEntry) GetIP() string`

GetIP returns the IP field if non-nil, zero value otherwise.

### GetIPOk

`func (o *ACLEntry) GetIPOk() (*string, bool)`

GetIPOk returns a tuple with the IP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIP

`func (o *ACLEntry) SetIP(v string)`

SetIP sets IP field to given value.

### HasIP

`func (o *ACLEntry) HasIP() bool`

HasIP returns a boolean if a field has been set.

### GetSubnet

`func (o *ACLEntry) GetSubnet() int32`

GetSubnet returns the Subnet field if non-nil, zero value otherwise.

### GetSubnetOk

`func (o *ACLEntry) GetSubnetOk() (*int32, bool)`

GetSubnetOk returns a tuple with the Subnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnet

`func (o *ACLEntry) SetSubnet(v int32)`

SetSubnet sets Subnet field to given value.

### HasSubnet

`func (o *ACLEntry) HasSubnet() bool`

HasSubnet returns a boolean if a field has been set.

### SetSubnetNil

`func (o *ACLEntry) SetSubnetNil(b bool)`

 SetSubnetNil sets the value for Subnet to be an explicit nil

### UnsetSubnet
`func (o *ACLEntry) UnsetSubnet()`

UnsetSubnet ensures that no value is present for Subnet, not even an explicit nil

[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
