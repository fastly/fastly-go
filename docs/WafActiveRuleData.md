# WafActiveRuleData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to [**TypeWafActiveRule**](TypeWafActiveRule.md) |  | [optional] [default to TYPEWAFACTIVERULE_WAF_ACTIVE_RULE]
**Attributes** | Pointer to [**WafActiveRuleDataAttributes**](WafActiveRuleDataAttributes.md) |  | [optional] 
**Relationships** | Pointer to [**RelationshipsForWafActiveRule**](RelationshipsForWafActiveRule.md) |  | [optional] 

## Methods

### NewWafActiveRuleData

`func NewWafActiveRuleData() *WafActiveRuleData`

NewWafActiveRuleData instantiates a new WafActiveRuleData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWafActiveRuleDataWithDefaults

`func NewWafActiveRuleDataWithDefaults() *WafActiveRuleData`

NewWafActiveRuleDataWithDefaults instantiates a new WafActiveRuleData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *WafActiveRuleData) GetType() TypeWafActiveRule`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *WafActiveRuleData) GetTypeOk() (*TypeWafActiveRule, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *WafActiveRuleData) SetType(v TypeWafActiveRule)`

SetType sets Type field to given value.

### HasType

`func (o *WafActiveRuleData) HasType() bool`

HasType returns a boolean if a field has been set.

### GetAttributes

`func (o *WafActiveRuleData) GetAttributes() WafActiveRuleDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *WafActiveRuleData) GetAttributesOk() (*WafActiveRuleDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *WafActiveRuleData) SetAttributes(v WafActiveRuleDataAttributes)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *WafActiveRuleData) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetRelationships

`func (o *WafActiveRuleData) GetRelationships() RelationshipsForWafActiveRule`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *WafActiveRuleData) GetRelationshipsOk() (*RelationshipsForWafActiveRule, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *WafActiveRuleData) SetRelationships(v RelationshipsForWafActiveRule)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *WafActiveRuleData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
