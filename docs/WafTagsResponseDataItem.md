# WafTagsResponseDataItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to [**TypeWafTag**](TypeWafTag.md) |  | [optional] [default to TYPEWAFTAG_WAF_TAG]
**ID** | Pointer to **string** | Alphanumeric string identifying a WAF tag. | [optional] [readonly] 
**Attributes** | Pointer to [**WafTagAttributes**](WafTagAttributes.md) |  | [optional] 
**Relationships** | Pointer to [**RelationshipWafRule**](RelationshipWafRule.md) |  | [optional] 

## Methods

### NewWafTagsResponseDataItem

`func NewWafTagsResponseDataItem() *WafTagsResponseDataItem`

NewWafTagsResponseDataItem instantiates a new WafTagsResponseDataItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWafTagsResponseDataItemWithDefaults

`func NewWafTagsResponseDataItemWithDefaults() *WafTagsResponseDataItem`

NewWafTagsResponseDataItemWithDefaults instantiates a new WafTagsResponseDataItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *WafTagsResponseDataItem) GetType() TypeWafTag`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *WafTagsResponseDataItem) GetTypeOk() (*TypeWafTag, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *WafTagsResponseDataItem) SetType(v TypeWafTag)`

SetType sets Type field to given value.

### HasType

`func (o *WafTagsResponseDataItem) HasType() bool`

HasType returns a boolean if a field has been set.

### GetID

`func (o *WafTagsResponseDataItem) GetID() string`

GetID returns the ID field if non-nil, zero value otherwise.

### GetIDOk

`func (o *WafTagsResponseDataItem) GetIDOk() (*string, bool)`

GetIDOk returns a tuple with the ID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetID

`func (o *WafTagsResponseDataItem) SetID(v string)`

SetID sets ID field to given value.

### HasID

`func (o *WafTagsResponseDataItem) HasID() bool`

HasID returns a boolean if a field has been set.

### GetAttributes

`func (o *WafTagsResponseDataItem) GetAttributes() WafTagAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *WafTagsResponseDataItem) GetAttributesOk() (*WafTagAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *WafTagsResponseDataItem) SetAttributes(v WafTagAttributes)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *WafTagsResponseDataItem) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetRelationships

`func (o *WafTagsResponseDataItem) GetRelationships() RelationshipWafRule`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *WafTagsResponseDataItem) GetRelationshipsOk() (*RelationshipWafRule, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *WafTagsResponseDataItem) SetRelationships(v RelationshipWafRule)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *WafTagsResponseDataItem) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
