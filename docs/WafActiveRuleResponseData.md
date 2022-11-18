# WafActiveRuleResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to [**TypeWafActiveRule**](TypeWafActiveRule.md) |  | [optional] [default to TYPEWAFACTIVERULE_WAF_ACTIVE_RULE]
**Attributes** | Pointer to [**WafActiveRuleResponseDataAttributes**](WafActiveRuleResponseDataAttributes.md) |  | [optional] 
**Relationships** | Pointer to [**WafActiveRuleResponseDataRelationships**](WafActiveRuleResponseDataRelationships.md) |  | [optional] 
**ID** | Pointer to **string** |  | [optional] [readonly] 

## Methods

### NewWafActiveRuleResponseData

`func NewWafActiveRuleResponseData() *WafActiveRuleResponseData`

NewWafActiveRuleResponseData instantiates a new WafActiveRuleResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWafActiveRuleResponseDataWithDefaults

`func NewWafActiveRuleResponseDataWithDefaults() *WafActiveRuleResponseData`

NewWafActiveRuleResponseDataWithDefaults instantiates a new WafActiveRuleResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *WafActiveRuleResponseData) GetType() TypeWafActiveRule`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *WafActiveRuleResponseData) GetTypeOk() (*TypeWafActiveRule, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *WafActiveRuleResponseData) SetType(v TypeWafActiveRule)`

SetType sets Type field to given value.

### HasType

`func (o *WafActiveRuleResponseData) HasType() bool`

HasType returns a boolean if a field has been set.

### GetAttributes

`func (o *WafActiveRuleResponseData) GetAttributes() WafActiveRuleResponseDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *WafActiveRuleResponseData) GetAttributesOk() (*WafActiveRuleResponseDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *WafActiveRuleResponseData) SetAttributes(v WafActiveRuleResponseDataAttributes)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *WafActiveRuleResponseData) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetRelationships

`func (o *WafActiveRuleResponseData) GetRelationships() WafActiveRuleResponseDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *WafActiveRuleResponseData) GetRelationshipsOk() (*WafActiveRuleResponseDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *WafActiveRuleResponseData) SetRelationships(v WafActiveRuleResponseDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *WafActiveRuleResponseData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.

### GetID

`func (o *WafActiveRuleResponseData) GetID() string`

GetID returns the ID field if non-nil, zero value otherwise.

### GetIDOk

`func (o *WafActiveRuleResponseData) GetIDOk() (*string, bool)`

GetIDOk returns a tuple with the ID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetID

`func (o *WafActiveRuleResponseData) SetID(v string)`

SetID sets ID field to given value.

### HasID

`func (o *WafActiveRuleResponseData) HasID() bool`

HasID returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
