# WafRuleRevisionResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to [**TypeWafRuleRevision**](TypeWafRuleRevision.md) |  | [optional] [default to TYPEWAFRULEREVISION_WAF_RULE_REVISION]
**ID** | Pointer to **string** | Alphanumeric string identifying a WAF rule revision. | [optional] [readonly] 
**Attributes** | Pointer to [**WafRuleRevisionAttributes**](WafRuleRevisionAttributes.md) |  | [optional] 
**Relationships** | Pointer to [**RelationshipWafRule**](RelationshipWafRule.md) |  | [optional] 

## Methods

### NewWafRuleRevisionResponseData

`func NewWafRuleRevisionResponseData() *WafRuleRevisionResponseData`

NewWafRuleRevisionResponseData instantiates a new WafRuleRevisionResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWafRuleRevisionResponseDataWithDefaults

`func NewWafRuleRevisionResponseDataWithDefaults() *WafRuleRevisionResponseData`

NewWafRuleRevisionResponseDataWithDefaults instantiates a new WafRuleRevisionResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *WafRuleRevisionResponseData) GetType() TypeWafRuleRevision`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *WafRuleRevisionResponseData) GetTypeOk() (*TypeWafRuleRevision, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *WafRuleRevisionResponseData) SetType(v TypeWafRuleRevision)`

SetType sets Type field to given value.

### HasType

`func (o *WafRuleRevisionResponseData) HasType() bool`

HasType returns a boolean if a field has been set.

### GetID

`func (o *WafRuleRevisionResponseData) GetID() string`

GetID returns the ID field if non-nil, zero value otherwise.

### GetIDOk

`func (o *WafRuleRevisionResponseData) GetIDOk() (*string, bool)`

GetIDOk returns a tuple with the ID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetID

`func (o *WafRuleRevisionResponseData) SetID(v string)`

SetID sets ID field to given value.

### HasID

`func (o *WafRuleRevisionResponseData) HasID() bool`

HasID returns a boolean if a field has been set.

### GetAttributes

`func (o *WafRuleRevisionResponseData) GetAttributes() WafRuleRevisionAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *WafRuleRevisionResponseData) GetAttributesOk() (*WafRuleRevisionAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *WafRuleRevisionResponseData) SetAttributes(v WafRuleRevisionAttributes)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *WafRuleRevisionResponseData) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetRelationships

`func (o *WafRuleRevisionResponseData) GetRelationships() RelationshipWafRule`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *WafRuleRevisionResponseData) GetRelationshipsOk() (*RelationshipWafRule, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *WafRuleRevisionResponseData) SetRelationships(v RelationshipWafRule)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *WafRuleRevisionResponseData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
