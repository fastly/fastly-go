# RelationshipMemberWafRuleRevision

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to [**TypeWafRuleRevision**](TypeWafRuleRevision.md) |  | [optional] [default to TYPEWAFRULEREVISION_WAF_RULE_REVISION]
**ID** | Pointer to **string** | Alphanumeric string identifying a WAF rule revision. | [optional] [readonly] 

## Methods

### NewRelationshipMemberWafRuleRevision

`func NewRelationshipMemberWafRuleRevision() *RelationshipMemberWafRuleRevision`

NewRelationshipMemberWafRuleRevision instantiates a new RelationshipMemberWafRuleRevision object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRelationshipMemberWafRuleRevisionWithDefaults

`func NewRelationshipMemberWafRuleRevisionWithDefaults() *RelationshipMemberWafRuleRevision`

NewRelationshipMemberWafRuleRevisionWithDefaults instantiates a new RelationshipMemberWafRuleRevision object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *RelationshipMemberWafRuleRevision) GetType() TypeWafRuleRevision`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RelationshipMemberWafRuleRevision) GetTypeOk() (*TypeWafRuleRevision, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RelationshipMemberWafRuleRevision) SetType(v TypeWafRuleRevision)`

SetType sets Type field to given value.

### HasType

`func (o *RelationshipMemberWafRuleRevision) HasType() bool`

HasType returns a boolean if a field has been set.

### GetID

`func (o *RelationshipMemberWafRuleRevision) GetID() string`

GetID returns the ID field if non-nil, zero value otherwise.

### GetIDOk

`func (o *RelationshipMemberWafRuleRevision) GetIDOk() (*string, bool)`

GetIDOk returns a tuple with the ID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetID

`func (o *RelationshipMemberWafRuleRevision) SetID(v string)`

SetID sets ID field to given value.

### HasID

`func (o *RelationshipMemberWafRuleRevision) HasID() bool`

HasID returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
