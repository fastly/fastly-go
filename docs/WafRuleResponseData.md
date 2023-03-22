# WafRuleResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to [**TypeWafRule**](TypeWafRule.md) |  | [optional] [default to TYPEWAFRULE_WAF_RULE]
**ID** | Pointer to **string** |  | [optional] [readonly] 
**Attributes** | Pointer to [**WafRuleAttributes**](WafRuleAttributes.md) |  | [optional] 
**Relationships** | Pointer to [**RelationshipsForWafRule**](RelationshipsForWafRule.md) |  | [optional] 

## Methods

### NewWafRuleResponseData

`func NewWafRuleResponseData() *WafRuleResponseData`

NewWafRuleResponseData instantiates a new WafRuleResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWafRuleResponseDataWithDefaults

`func NewWafRuleResponseDataWithDefaults() *WafRuleResponseData`

NewWafRuleResponseDataWithDefaults instantiates a new WafRuleResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *WafRuleResponseData) GetType() TypeWafRule`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *WafRuleResponseData) GetTypeOk() (*TypeWafRule, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *WafRuleResponseData) SetType(v TypeWafRule)`

SetType sets Type field to given value.

### HasType

`func (o *WafRuleResponseData) HasType() bool`

HasType returns a boolean if a field has been set.

### GetID

`func (o *WafRuleResponseData) GetID() string`

GetID returns the ID field if non-nil, zero value otherwise.

### GetIDOk

`func (o *WafRuleResponseData) GetIDOk() (*string, bool)`

GetIDOk returns a tuple with the ID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetID

`func (o *WafRuleResponseData) SetID(v string)`

SetID sets ID field to given value.

### HasID

`func (o *WafRuleResponseData) HasID() bool`

HasID returns a boolean if a field has been set.

### GetAttributes

`func (o *WafRuleResponseData) GetAttributes() WafRuleAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *WafRuleResponseData) GetAttributesOk() (*WafRuleAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *WafRuleResponseData) SetAttributes(v WafRuleAttributes)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *WafRuleResponseData) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetRelationships

`func (o *WafRuleResponseData) GetRelationships() RelationshipsForWafRule`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *WafRuleResponseData) GetRelationshipsOk() (*RelationshipsForWafRule, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *WafRuleResponseData) SetRelationships(v RelationshipsForWafRule)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *WafRuleResponseData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
