# IncludedWithWafExclusionItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to [**TypeWafRuleRevision**](TypeWafRuleRevision.md) |  | [optional] [default to TYPEWAFRULEREVISION_WAF_RULE_REVISION]
**ID** | Pointer to **string** | Alphanumeric string identifying a WAF rule revision. | [optional] [readonly] 
**Attributes** | Pointer to [**WafRuleRevisionAttributes**](WafRuleRevisionAttributes.md) |  | [optional] 

## Methods

### NewIncludedWithWafExclusionItem

`func NewIncludedWithWafExclusionItem() *IncludedWithWafExclusionItem`

NewIncludedWithWafExclusionItem instantiates a new IncludedWithWafExclusionItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIncludedWithWafExclusionItemWithDefaults

`func NewIncludedWithWafExclusionItemWithDefaults() *IncludedWithWafExclusionItem`

NewIncludedWithWafExclusionItemWithDefaults instantiates a new IncludedWithWafExclusionItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *IncludedWithWafExclusionItem) GetType() TypeWafRuleRevision`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *IncludedWithWafExclusionItem) GetTypeOk() (*TypeWafRuleRevision, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *IncludedWithWafExclusionItem) SetType(v TypeWafRuleRevision)`

SetType sets Type field to given value.

### HasType

`func (o *IncludedWithWafExclusionItem) HasType() bool`

HasType returns a boolean if a field has been set.

### GetID

`func (o *IncludedWithWafExclusionItem) GetID() string`

GetID returns the ID field if non-nil, zero value otherwise.

### GetIDOk

`func (o *IncludedWithWafExclusionItem) GetIDOk() (*string, bool)`

GetIDOk returns a tuple with the ID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetID

`func (o *IncludedWithWafExclusionItem) SetID(v string)`

SetID sets ID field to given value.

### HasID

`func (o *IncludedWithWafExclusionItem) HasID() bool`

HasID returns a boolean if a field has been set.

### GetAttributes

`func (o *IncludedWithWafExclusionItem) GetAttributes() WafRuleRevisionAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *IncludedWithWafExclusionItem) GetAttributesOk() (*WafRuleRevisionAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *IncludedWithWafExclusionItem) SetAttributes(v WafRuleRevisionAttributes)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *IncludedWithWafExclusionItem) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
