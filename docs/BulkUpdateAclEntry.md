# BulkUpdateAclEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Negated** | Pointer to **int32** | Whether to negate the match. Useful primarily when creating individual exceptions to larger subnets. | [optional] [default to 0]
**Comment** | Pointer to **NullableString** | A freeform descriptive note. | [optional] 
**Ip** | Pointer to **string** | An IP address. | [optional] 
**Subnet** | Pointer to **NullableInt32** | Number of bits for the subnet mask applied to the IP address. For IPv4 addresses, a value of 32 represents the smallest subnet mask (1 address), 24 represents a class C subnet mask (256 addresses), 16 represents a class B subnet mask (65k addresses), and 8 is class A subnet mask (16m addresses). If not provided, no mask is applied. | [optional] 
**Op** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] [readonly] 

## Methods

### NewBulkUpdateAclEntry

`func NewBulkUpdateAclEntry() *BulkUpdateAclEntry`

NewBulkUpdateAclEntry instantiates a new BulkUpdateAclEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkUpdateAclEntryWithDefaults

`func NewBulkUpdateAclEntryWithDefaults() *BulkUpdateAclEntry`

NewBulkUpdateAclEntryWithDefaults instantiates a new BulkUpdateAclEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNegated

`func (o *BulkUpdateAclEntry) GetNegated() int32`

GetNegated returns the Negated field if non-nil, zero value otherwise.

### GetNegatedOk

`func (o *BulkUpdateAclEntry) GetNegatedOk() (*int32, bool)`

GetNegatedOk returns a tuple with the Negated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNegated

`func (o *BulkUpdateAclEntry) SetNegated(v int32)`

SetNegated sets Negated field to given value.

### HasNegated

`func (o *BulkUpdateAclEntry) HasNegated() bool`

HasNegated returns a boolean if a field has been set.

### GetComment

`func (o *BulkUpdateAclEntry) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *BulkUpdateAclEntry) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *BulkUpdateAclEntry) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *BulkUpdateAclEntry) HasComment() bool`

HasComment returns a boolean if a field has been set.

### SetCommentNil

`func (o *BulkUpdateAclEntry) SetCommentNil(b bool)`

 SetCommentNil sets the value for Comment to be an explicit nil

### UnsetComment
`func (o *BulkUpdateAclEntry) UnsetComment()`

UnsetComment ensures that no value is present for Comment, not even an explicit nil
### GetIp

`func (o *BulkUpdateAclEntry) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *BulkUpdateAclEntry) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *BulkUpdateAclEntry) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *BulkUpdateAclEntry) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetSubnet

`func (o *BulkUpdateAclEntry) GetSubnet() int32`

GetSubnet returns the Subnet field if non-nil, zero value otherwise.

### GetSubnetOk

`func (o *BulkUpdateAclEntry) GetSubnetOk() (*int32, bool)`

GetSubnetOk returns a tuple with the Subnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnet

`func (o *BulkUpdateAclEntry) SetSubnet(v int32)`

SetSubnet sets Subnet field to given value.

### HasSubnet

`func (o *BulkUpdateAclEntry) HasSubnet() bool`

HasSubnet returns a boolean if a field has been set.

### SetSubnetNil

`func (o *BulkUpdateAclEntry) SetSubnetNil(b bool)`

 SetSubnetNil sets the value for Subnet to be an explicit nil

### UnsetSubnet
`func (o *BulkUpdateAclEntry) UnsetSubnet()`

UnsetSubnet ensures that no value is present for Subnet, not even an explicit nil
### GetOp

`func (o *BulkUpdateAclEntry) GetOp() string`

GetOp returns the Op field if non-nil, zero value otherwise.

### GetOpOk

`func (o *BulkUpdateAclEntry) GetOpOk() (*string, bool)`

GetOpOk returns a tuple with the Op field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOp

`func (o *BulkUpdateAclEntry) SetOp(v string)`

SetOp sets Op field to given value.

### HasOp

`func (o *BulkUpdateAclEntry) HasOp() bool`

HasOp returns a boolean if a field has been set.

### GetId

`func (o *BulkUpdateAclEntry) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BulkUpdateAclEntry) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BulkUpdateAclEntry) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *BulkUpdateAclEntry) HasId() bool`

HasId returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


