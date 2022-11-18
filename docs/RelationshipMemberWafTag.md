# RelationshipMemberWafTag

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to [**TypeWafTag**](TypeWafTag.md) |  | [optional] [default to TYPEWAFTAG_WAF_TAG]
**ID** | Pointer to **string** | Alphanumeric string identifying a WAF tag. | [optional] [readonly] 

## Methods

### NewRelationshipMemberWafTag

`func NewRelationshipMemberWafTag() *RelationshipMemberWafTag`

NewRelationshipMemberWafTag instantiates a new RelationshipMemberWafTag object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRelationshipMemberWafTagWithDefaults

`func NewRelationshipMemberWafTagWithDefaults() *RelationshipMemberWafTag`

NewRelationshipMemberWafTagWithDefaults instantiates a new RelationshipMemberWafTag object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *RelationshipMemberWafTag) GetType() TypeWafTag`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RelationshipMemberWafTag) GetTypeOk() (*TypeWafTag, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RelationshipMemberWafTag) SetType(v TypeWafTag)`

SetType sets Type field to given value.

### HasType

`func (o *RelationshipMemberWafTag) HasType() bool`

HasType returns a boolean if a field has been set.

### GetID

`func (o *RelationshipMemberWafTag) GetID() string`

GetID returns the ID field if non-nil, zero value otherwise.

### GetIDOk

`func (o *RelationshipMemberWafTag) GetIDOk() (*string, bool)`

GetIDOk returns a tuple with the ID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetID

`func (o *RelationshipMemberWafTag) SetID(v string)`

SetID sets ID field to given value.

### HasID

`func (o *RelationshipMemberWafTag) HasID() bool`

HasID returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
