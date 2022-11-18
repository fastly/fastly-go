# WafExclusionResponseDataAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ID** | Pointer to **string** | Alphanumeric string identifying a WAF exclusion. | [optional] [readonly] 
**Attributes** | Pointer to [**WafExclusionResponseDataAttributes**](WafExclusionResponseDataAttributes.md) |  | [optional] 
**Relationships** | Pointer to [**WafExclusionResponseDataRelationships**](WafExclusionResponseDataRelationships.md) |  | [optional] 

## Methods

### NewWafExclusionResponseDataAllOf

`func NewWafExclusionResponseDataAllOf() *WafExclusionResponseDataAllOf`

NewWafExclusionResponseDataAllOf instantiates a new WafExclusionResponseDataAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWafExclusionResponseDataAllOfWithDefaults

`func NewWafExclusionResponseDataAllOfWithDefaults() *WafExclusionResponseDataAllOf`

NewWafExclusionResponseDataAllOfWithDefaults instantiates a new WafExclusionResponseDataAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetID

`func (o *WafExclusionResponseDataAllOf) GetID() string`

GetID returns the ID field if non-nil, zero value otherwise.

### GetIDOk

`func (o *WafExclusionResponseDataAllOf) GetIDOk() (*string, bool)`

GetIDOk returns a tuple with the ID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetID

`func (o *WafExclusionResponseDataAllOf) SetID(v string)`

SetID sets ID field to given value.

### HasID

`func (o *WafExclusionResponseDataAllOf) HasID() bool`

HasID returns a boolean if a field has been set.

### GetAttributes

`func (o *WafExclusionResponseDataAllOf) GetAttributes() WafExclusionResponseDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *WafExclusionResponseDataAllOf) GetAttributesOk() (*WafExclusionResponseDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *WafExclusionResponseDataAllOf) SetAttributes(v WafExclusionResponseDataAttributes)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *WafExclusionResponseDataAllOf) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetRelationships

`func (o *WafExclusionResponseDataAllOf) GetRelationships() WafExclusionResponseDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *WafExclusionResponseDataAllOf) GetRelationshipsOk() (*WafExclusionResponseDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *WafExclusionResponseDataAllOf) SetRelationships(v WafExclusionResponseDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *WafExclusionResponseDataAllOf) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
