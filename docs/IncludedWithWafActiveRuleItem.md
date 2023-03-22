# IncludedWithWafActiveRuleItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**SchemasWafFirewallVersionData**](SchemasWafFirewallVersionData.md) |  | [optional] 
**Type** | Pointer to [**TypeWafRuleRevision**](TypeWafRuleRevision.md) |  | [optional] [default to TYPEWAFRULEREVISION_WAF_RULE_REVISION]
**ID** | Pointer to **string** | Alphanumeric string identifying a WAF rule revision. | [optional] [readonly] 
**Attributes** | Pointer to [**WafRuleRevisionAttributes**](WafRuleRevisionAttributes.md) |  | [optional] 

## Methods

### NewIncludedWithWafActiveRuleItem

`func NewIncludedWithWafActiveRuleItem() *IncludedWithWafActiveRuleItem`

NewIncludedWithWafActiveRuleItem instantiates a new IncludedWithWafActiveRuleItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIncludedWithWafActiveRuleItemWithDefaults

`func NewIncludedWithWafActiveRuleItemWithDefaults() *IncludedWithWafActiveRuleItem`

NewIncludedWithWafActiveRuleItemWithDefaults instantiates a new IncludedWithWafActiveRuleItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *IncludedWithWafActiveRuleItem) GetData() SchemasWafFirewallVersionData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *IncludedWithWafActiveRuleItem) GetDataOk() (*SchemasWafFirewallVersionData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *IncludedWithWafActiveRuleItem) SetData(v SchemasWafFirewallVersionData)`

SetData sets Data field to given value.

### HasData

`func (o *IncludedWithWafActiveRuleItem) HasData() bool`

HasData returns a boolean if a field has been set.

### GetType

`func (o *IncludedWithWafActiveRuleItem) GetType() TypeWafRuleRevision`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *IncludedWithWafActiveRuleItem) GetTypeOk() (*TypeWafRuleRevision, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *IncludedWithWafActiveRuleItem) SetType(v TypeWafRuleRevision)`

SetType sets Type field to given value.

### HasType

`func (o *IncludedWithWafActiveRuleItem) HasType() bool`

HasType returns a boolean if a field has been set.

### GetID

`func (o *IncludedWithWafActiveRuleItem) GetID() string`

GetID returns the ID field if non-nil, zero value otherwise.

### GetIDOk

`func (o *IncludedWithWafActiveRuleItem) GetIDOk() (*string, bool)`

GetIDOk returns a tuple with the ID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetID

`func (o *IncludedWithWafActiveRuleItem) SetID(v string)`

SetID sets ID field to given value.

### HasID

`func (o *IncludedWithWafActiveRuleItem) HasID() bool`

HasID returns a boolean if a field has been set.

### GetAttributes

`func (o *IncludedWithWafActiveRuleItem) GetAttributes() WafRuleRevisionAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *IncludedWithWafActiveRuleItem) GetAttributesOk() (*WafRuleRevisionAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *IncludedWithWafActiveRuleItem) SetAttributes(v WafRuleRevisionAttributes)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *IncludedWithWafActiveRuleItem) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
