# AclEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Negated** | Pointer to **int32** | Whether to negate the match. Useful primarily when creating individual exceptions to larger subnets. | [optional] [default to 0]
**Comment** | Pointer to **NullableString** | A freeform descriptive note. | [optional] 
**Ip** | Pointer to **string** | An IP address. | [optional] 
**Subnet** | Pointer to **NullableInt32** | Number of bits for the subnet mask applied to the IP address. For IPv4 addresses, a value of 32 represents the smallest subnet mask (1 address), 24 represents a class C subnet mask (256 addresses), 16 represents a class B subnet mask (65k addresses), and 8 is class A subnet mask (16m addresses). If not provided, no mask is applied. | [optional] 

## Methods

### NewAclEntry

`func NewAclEntry() *AclEntry`

NewAclEntry instantiates a new AclEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAclEntryWithDefaults

`func NewAclEntryWithDefaults() *AclEntry`

NewAclEntryWithDefaults instantiates a new AclEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNegated

`func (o *AclEntry) GetNegated() int32`

GetNegated returns the Negated field if non-nil, zero value otherwise.

### GetNegatedOk

`func (o *AclEntry) GetNegatedOk() (*int32, bool)`

GetNegatedOk returns a tuple with the Negated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNegated

`func (o *AclEntry) SetNegated(v int32)`

SetNegated sets Negated field to given value.

### HasNegated

`func (o *AclEntry) HasNegated() bool`

HasNegated returns a boolean if a field has been set.

### GetComment

`func (o *AclEntry) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *AclEntry) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *AclEntry) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *AclEntry) HasComment() bool`

HasComment returns a boolean if a field has been set.

### SetCommentNil

`func (o *AclEntry) SetCommentNil(b bool)`

 SetCommentNil sets the value for Comment to be an explicit nil

### UnsetComment
`func (o *AclEntry) UnsetComment()`

UnsetComment ensures that no value is present for Comment, not even an explicit nil
### GetIp

`func (o *AclEntry) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *AclEntry) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *AclEntry) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *AclEntry) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetSubnet

`func (o *AclEntry) GetSubnet() int32`

GetSubnet returns the Subnet field if non-nil, zero value otherwise.

### GetSubnetOk

`func (o *AclEntry) GetSubnetOk() (*int32, bool)`

GetSubnetOk returns a tuple with the Subnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnet

`func (o *AclEntry) SetSubnet(v int32)`

SetSubnet sets Subnet field to given value.

### HasSubnet

`func (o *AclEntry) HasSubnet() bool`

HasSubnet returns a boolean if a field has been set.

### SetSubnetNil

`func (o *AclEntry) SetSubnetNil(b bool)`

 SetSubnetNil sets the value for Subnet to be an explicit nil

### UnsetSubnet
`func (o *AclEntry) UnsetSubnet()`

UnsetSubnet ensures that no value is present for Subnet, not even an explicit nil

[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


