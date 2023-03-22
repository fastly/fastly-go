# RelationshipMemberWafRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to [**TypeWafRule**](TypeWafRule.md) |  | [optional] [default to TYPEWAFRULE_WAF_RULE]
**ID** | Pointer to **string** |  | [optional] [readonly] 

## Methods

### NewRelationshipMemberWafRule

`func NewRelationshipMemberWafRule() *RelationshipMemberWafRule`

NewRelationshipMemberWafRule instantiates a new RelationshipMemberWafRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRelationshipMemberWafRuleWithDefaults

`func NewRelationshipMemberWafRuleWithDefaults() *RelationshipMemberWafRule`

NewRelationshipMemberWafRuleWithDefaults instantiates a new RelationshipMemberWafRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *RelationshipMemberWafRule) GetType() TypeWafRule`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RelationshipMemberWafRule) GetTypeOk() (*TypeWafRule, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RelationshipMemberWafRule) SetType(v TypeWafRule)`

SetType sets Type field to given value.

### HasType

`func (o *RelationshipMemberWafRule) HasType() bool`

HasType returns a boolean if a field has been set.

### GetID

`func (o *RelationshipMemberWafRule) GetID() string`

GetID returns the ID field if non-nil, zero value otherwise.

### GetIDOk

`func (o *RelationshipMemberWafRule) GetIDOk() (*string, bool)`

GetIDOk returns a tuple with the ID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetID

`func (o *RelationshipMemberWafRule) SetID(v string)`

SetID sets ID field to given value.

### HasID

`func (o *RelationshipMemberWafRule) HasID() bool`

HasID returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
