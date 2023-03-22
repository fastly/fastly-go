# RelationshipMemberWafActiveRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to [**TypeWafActiveRule**](TypeWafActiveRule.md) |  | [optional] [default to TYPEWAFACTIVERULE_WAF_ACTIVE_RULE]
**ID** | Pointer to **string** |  | [optional] [readonly] 

## Methods

### NewRelationshipMemberWafActiveRule

`func NewRelationshipMemberWafActiveRule() *RelationshipMemberWafActiveRule`

NewRelationshipMemberWafActiveRule instantiates a new RelationshipMemberWafActiveRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRelationshipMemberWafActiveRuleWithDefaults

`func NewRelationshipMemberWafActiveRuleWithDefaults() *RelationshipMemberWafActiveRule`

NewRelationshipMemberWafActiveRuleWithDefaults instantiates a new RelationshipMemberWafActiveRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *RelationshipMemberWafActiveRule) GetType() TypeWafActiveRule`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RelationshipMemberWafActiveRule) GetTypeOk() (*TypeWafActiveRule, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RelationshipMemberWafActiveRule) SetType(v TypeWafActiveRule)`

SetType sets Type field to given value.

### HasType

`func (o *RelationshipMemberWafActiveRule) HasType() bool`

HasType returns a boolean if a field has been set.

### GetID

`func (o *RelationshipMemberWafActiveRule) GetID() string`

GetID returns the ID field if non-nil, zero value otherwise.

### GetIDOk

`func (o *RelationshipMemberWafActiveRule) GetIDOk() (*string, bool)`

GetIDOk returns a tuple with the ID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetID

`func (o *RelationshipMemberWafActiveRule) SetID(v string)`

SetID sets ID field to given value.

### HasID

`func (o *RelationshipMemberWafActiveRule) HasID() bool`

HasID returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
