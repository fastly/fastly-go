# RuleUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | Pointer to [**Action**](Action.md) |  | [optional] 
**Conditions** | Pointer to [**[]RoutingConfigCondition**](RoutingConfigCondition.md) |  | [optional] 
**Position** | Pointer to [**Position**](Position.md) |  | [optional] 

## Methods

### NewRuleUpdate

`func NewRuleUpdate() *RuleUpdate`

NewRuleUpdate instantiates a new RuleUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRuleUpdateWithDefaults

`func NewRuleUpdateWithDefaults() *RuleUpdate`

NewRuleUpdateWithDefaults instantiates a new RuleUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *RuleUpdate) GetAction() Action`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *RuleUpdate) GetActionOk() (*Action, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *RuleUpdate) SetAction(v Action)`

SetAction sets Action field to given value.

### HasAction

`func (o *RuleUpdate) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetConditions

`func (o *RuleUpdate) GetConditions() []RoutingConfigCondition`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *RuleUpdate) GetConditionsOk() (*[]RoutingConfigCondition, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *RuleUpdate) SetConditions(v []RoutingConfigCondition)`

SetConditions sets Conditions field to given value.

### HasConditions

`func (o *RuleUpdate) HasConditions() bool`

HasConditions returns a boolean if a field has been set.

### GetPosition

`func (o *RuleUpdate) GetPosition() Position`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *RuleUpdate) GetPositionOk() (*Position, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *RuleUpdate) SetPosition(v Position)`

SetPosition sets Position field to given value.

### HasPosition

`func (o *RuleUpdate) HasPosition() bool`

HasPosition returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


