# BulkUpdateACLEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Negated** | Pointer to **int32** | Whether to negate the match. Useful primarily when creating individual exceptions to larger subnets. | [optional] [default to 0]
**Comment** | Pointer to **NullableString** | A freeform descriptive note. | [optional] 
**IP** | Pointer to **string** | An IP address. | [optional] 
**Subnet** | Pointer to **NullableInt32** | Number of bits for the subnet mask applied to the IP address. For IPv4 addresses, a value of 32 represents the smallest subnet mask (1 address), 24 represents a class C subnet mask (256 addresses), 16 represents a class B subnet mask (65k addresses), and 8 is class A subnet mask (16m addresses). If not provided, no mask is applied. | [optional] 
**Op** | Pointer to **string** |  | [optional] 
**ID** | Pointer to **string** |  | [optional] [readonly] 

## Methods

### NewBulkUpdateACLEntry

`func NewBulkUpdateACLEntry() *BulkUpdateACLEntry`

NewBulkUpdateACLEntry instantiates a new BulkUpdateACLEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkUpdateACLEntryWithDefaults

`func NewBulkUpdateACLEntryWithDefaults() *BulkUpdateACLEntry`

NewBulkUpdateACLEntryWithDefaults instantiates a new BulkUpdateACLEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNegated

`func (o *BulkUpdateACLEntry) GetNegated() int32`

GetNegated returns the Negated field if non-nil, zero value otherwise.

### GetNegatedOk

`func (o *BulkUpdateACLEntry) GetNegatedOk() (*int32, bool)`

GetNegatedOk returns a tuple with the Negated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNegated

`func (o *BulkUpdateACLEntry) SetNegated(v int32)`

SetNegated sets Negated field to given value.

### HasNegated

`func (o *BulkUpdateACLEntry) HasNegated() bool`

HasNegated returns a boolean if a field has been set.

### GetComment

`func (o *BulkUpdateACLEntry) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *BulkUpdateACLEntry) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *BulkUpdateACLEntry) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *BulkUpdateACLEntry) HasComment() bool`

HasComment returns a boolean if a field has been set.

### SetCommentNil

`func (o *BulkUpdateACLEntry) SetCommentNil(b bool)`

 SetCommentNil sets the value for Comment to be an explicit nil

### UnsetComment
`func (o *BulkUpdateACLEntry) UnsetComment()`

UnsetComment ensures that no value is present for Comment, not even an explicit nil
### GetIP

`func (o *BulkUpdateACLEntry) GetIP() string`

GetIP returns the IP field if non-nil, zero value otherwise.

### GetIPOk

`func (o *BulkUpdateACLEntry) GetIPOk() (*string, bool)`

GetIPOk returns a tuple with the IP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIP

`func (o *BulkUpdateACLEntry) SetIP(v string)`

SetIP sets IP field to given value.

### HasIP

`func (o *BulkUpdateACLEntry) HasIP() bool`

HasIP returns a boolean if a field has been set.

### GetSubnet

`func (o *BulkUpdateACLEntry) GetSubnet() int32`

GetSubnet returns the Subnet field if non-nil, zero value otherwise.

### GetSubnetOk

`func (o *BulkUpdateACLEntry) GetSubnetOk() (*int32, bool)`

GetSubnetOk returns a tuple with the Subnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnet

`func (o *BulkUpdateACLEntry) SetSubnet(v int32)`

SetSubnet sets Subnet field to given value.

### HasSubnet

`func (o *BulkUpdateACLEntry) HasSubnet() bool`

HasSubnet returns a boolean if a field has been set.

### SetSubnetNil

`func (o *BulkUpdateACLEntry) SetSubnetNil(b bool)`

 SetSubnetNil sets the value for Subnet to be an explicit nil

### UnsetSubnet
`func (o *BulkUpdateACLEntry) UnsetSubnet()`

UnsetSubnet ensures that no value is present for Subnet, not even an explicit nil
### GetOp

`func (o *BulkUpdateACLEntry) GetOp() string`

GetOp returns the Op field if non-nil, zero value otherwise.

### GetOpOk

`func (o *BulkUpdateACLEntry) GetOpOk() (*string, bool)`

GetOpOk returns a tuple with the Op field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOp

`func (o *BulkUpdateACLEntry) SetOp(v string)`

SetOp sets Op field to given value.

### HasOp

`func (o *BulkUpdateACLEntry) HasOp() bool`

HasOp returns a boolean if a field has been set.

### GetID

`func (o *BulkUpdateACLEntry) GetID() string`

GetID returns the ID field if non-nil, zero value otherwise.

### GetIDOk

`func (o *BulkUpdateACLEntry) GetIDOk() (*string, bool)`

GetIDOk returns a tuple with the ID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetID

`func (o *BulkUpdateACLEntry) SetID(v string)`

SetID sets ID field to given value.

### HasID

`func (o *BulkUpdateACLEntry) HasID() bool`

HasID returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
