# WafRuleRevision

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to [**TypeWafRuleRevision**](TypeWafRuleRevision.md) |  | [optional] [default to TYPEWAFRULEREVISION_WAF_RULE_REVISION]
**ID** | Pointer to **string** | Alphanumeric string identifying a WAF rule revision. | [optional] [readonly] 
**Attributes** | Pointer to [**WafRuleRevisionAttributes**](WafRuleRevisionAttributes.md) |  | [optional] 

## Methods

### NewWafRuleRevision

`func NewWafRuleRevision() *WafRuleRevision`

NewWafRuleRevision instantiates a new WafRuleRevision object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWafRuleRevisionWithDefaults

`func NewWafRuleRevisionWithDefaults() *WafRuleRevision`

NewWafRuleRevisionWithDefaults instantiates a new WafRuleRevision object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *WafRuleRevision) GetType() TypeWafRuleRevision`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *WafRuleRevision) GetTypeOk() (*TypeWafRuleRevision, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *WafRuleRevision) SetType(v TypeWafRuleRevision)`

SetType sets Type field to given value.

### HasType

`func (o *WafRuleRevision) HasType() bool`

HasType returns a boolean if a field has been set.

### GetID

`func (o *WafRuleRevision) GetID() string`

GetID returns the ID field if non-nil, zero value otherwise.

### GetIDOk

`func (o *WafRuleRevision) GetIDOk() (*string, bool)`

GetIDOk returns a tuple with the ID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetID

`func (o *WafRuleRevision) SetID(v string)`

SetID sets ID field to given value.

### HasID

`func (o *WafRuleRevision) HasID() bool`

HasID returns a boolean if a field has been set.

### GetAttributes

`func (o *WafRuleRevision) GetAttributes() WafRuleRevisionAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *WafRuleRevision) GetAttributesOk() (*WafRuleRevisionAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *WafRuleRevision) SetAttributes(v WafRuleRevisionAttributes)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *WafRuleRevision) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
