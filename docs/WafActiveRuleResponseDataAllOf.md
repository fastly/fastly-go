# WafActiveRuleResponseDataAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ID** | Pointer to **string** |  | [optional] [readonly] 
**Attributes** | Pointer to [**WafActiveRuleResponseDataAttributes**](WafActiveRuleResponseDataAttributes.md) |  | [optional] 
**Relationships** | Pointer to [**WafActiveRuleResponseDataRelationships**](WafActiveRuleResponseDataRelationships.md) |  | [optional] 

## Methods

### NewWafActiveRuleResponseDataAllOf

`func NewWafActiveRuleResponseDataAllOf() *WafActiveRuleResponseDataAllOf`

NewWafActiveRuleResponseDataAllOf instantiates a new WafActiveRuleResponseDataAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWafActiveRuleResponseDataAllOfWithDefaults

`func NewWafActiveRuleResponseDataAllOfWithDefaults() *WafActiveRuleResponseDataAllOf`

NewWafActiveRuleResponseDataAllOfWithDefaults instantiates a new WafActiveRuleResponseDataAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetID

`func (o *WafActiveRuleResponseDataAllOf) GetID() string`

GetID returns the ID field if non-nil, zero value otherwise.

### GetIDOk

`func (o *WafActiveRuleResponseDataAllOf) GetIDOk() (*string, bool)`

GetIDOk returns a tuple with the ID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetID

`func (o *WafActiveRuleResponseDataAllOf) SetID(v string)`

SetID sets ID field to given value.

### HasID

`func (o *WafActiveRuleResponseDataAllOf) HasID() bool`

HasID returns a boolean if a field has been set.

### GetAttributes

`func (o *WafActiveRuleResponseDataAllOf) GetAttributes() WafActiveRuleResponseDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *WafActiveRuleResponseDataAllOf) GetAttributesOk() (*WafActiveRuleResponseDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *WafActiveRuleResponseDataAllOf) SetAttributes(v WafActiveRuleResponseDataAttributes)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *WafActiveRuleResponseDataAllOf) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetRelationships

`func (o *WafActiveRuleResponseDataAllOf) GetRelationships() WafActiveRuleResponseDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *WafActiveRuleResponseDataAllOf) GetRelationshipsOk() (*WafActiveRuleResponseDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *WafActiveRuleResponseDataAllOf) SetRelationships(v WafActiveRuleResponseDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *WafActiveRuleResponseDataAllOf) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
