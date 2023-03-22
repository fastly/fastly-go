# WafRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to [**TypeWafRule**](TypeWafRule.md) |  | [optional] [default to TYPEWAFRULE_WAF_RULE]
**ID** | Pointer to **string** |  | [optional] [readonly] 
**Attributes** | Pointer to [**WafRuleAttributes**](WafRuleAttributes.md) |  | [optional] 

## Methods

### NewWafRule

`func NewWafRule() *WafRule`

NewWafRule instantiates a new WafRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWafRuleWithDefaults

`func NewWafRuleWithDefaults() *WafRule`

NewWafRuleWithDefaults instantiates a new WafRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *WafRule) GetType() TypeWafRule`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *WafRule) GetTypeOk() (*TypeWafRule, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *WafRule) SetType(v TypeWafRule)`

SetType sets Type field to given value.

### HasType

`func (o *WafRule) HasType() bool`

HasType returns a boolean if a field has been set.

### GetID

`func (o *WafRule) GetID() string`

GetID returns the ID field if non-nil, zero value otherwise.

### GetIDOk

`func (o *WafRule) GetIDOk() (*string, bool)`

GetIDOk returns a tuple with the ID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetID

`func (o *WafRule) SetID(v string)`

SetID sets ID field to given value.

### HasID

`func (o *WafRule) HasID() bool`

HasID returns a boolean if a field has been set.

### GetAttributes

`func (o *WafRule) GetAttributes() WafRuleAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *WafRule) GetAttributesOk() (*WafRuleAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *WafRule) SetAttributes(v WafRuleAttributes)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *WafRule) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
