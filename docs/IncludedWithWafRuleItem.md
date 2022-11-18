# IncludedWithWafRuleItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to [**TypeWafRuleRevision**](TypeWafRuleRevision.md) |  | [optional] [default to TYPEWAFRULEREVISION_WAF_RULE_REVISION]
**ID** | Pointer to **string** | Alphanumeric string identifying a WAF rule revision. | [optional] [readonly] 
**Attributes** | Pointer to [**WafRuleRevisionAttributes**](WafRuleRevisionAttributes.md) |  | [optional] 

## Methods

### NewIncludedWithWafRuleItem

`func NewIncludedWithWafRuleItem() *IncludedWithWafRuleItem`

NewIncludedWithWafRuleItem instantiates a new IncludedWithWafRuleItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIncludedWithWafRuleItemWithDefaults

`func NewIncludedWithWafRuleItemWithDefaults() *IncludedWithWafRuleItem`

NewIncludedWithWafRuleItemWithDefaults instantiates a new IncludedWithWafRuleItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *IncludedWithWafRuleItem) GetType() TypeWafRuleRevision`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *IncludedWithWafRuleItem) GetTypeOk() (*TypeWafRuleRevision, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *IncludedWithWafRuleItem) SetType(v TypeWafRuleRevision)`

SetType sets Type field to given value.

### HasType

`func (o *IncludedWithWafRuleItem) HasType() bool`

HasType returns a boolean if a field has been set.

### GetID

`func (o *IncludedWithWafRuleItem) GetID() string`

GetID returns the ID field if non-nil, zero value otherwise.

### GetIDOk

`func (o *IncludedWithWafRuleItem) GetIDOk() (*string, bool)`

GetIDOk returns a tuple with the ID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetID

`func (o *IncludedWithWafRuleItem) SetID(v string)`

SetID sets ID field to given value.

### HasID

`func (o *IncludedWithWafRuleItem) HasID() bool`

HasID returns a boolean if a field has been set.

### GetAttributes

`func (o *IncludedWithWafRuleItem) GetAttributes() WafRuleRevisionAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *IncludedWithWafRuleItem) GetAttributesOk() (*WafRuleRevisionAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *IncludedWithWafRuleItem) SetAttributes(v WafRuleRevisionAttributes)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *IncludedWithWafRuleItem) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
