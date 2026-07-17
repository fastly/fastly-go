# RuleCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | [**Action**](Action.md) |  | 
**Conditions** | [**[]RoutingConfigCondition**](RoutingConfigCondition.md) | The conditions a request must satisfy for this rule to match. An empty array indicates the default rule for the path. | 
**Position** | Pointer to [**Position**](Position.md) |  | [optional] 

## Methods

### NewRuleCreate

`func NewRuleCreate(action Action, conditions []RoutingConfigCondition, ) *RuleCreate`

NewRuleCreate instantiates a new RuleCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRuleCreateWithDefaults

`func NewRuleCreateWithDefaults() *RuleCreate`

NewRuleCreateWithDefaults instantiates a new RuleCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *RuleCreate) GetAction() Action`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *RuleCreate) GetActionOk() (*Action, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *RuleCreate) SetAction(v Action)`

SetAction sets Action field to given value.


### GetConditions

`func (o *RuleCreate) GetConditions() []RoutingConfigCondition`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *RuleCreate) GetConditionsOk() (*[]RoutingConfigCondition, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *RuleCreate) SetConditions(v []RoutingConfigCondition)`

SetConditions sets Conditions field to given value.


### GetPosition

`func (o *RuleCreate) GetPosition() Position`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *RuleCreate) GetPositionOk() (*Position, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *RuleCreate) SetPosition(v Position)`

SetPosition sets Position field to given value.

### HasPosition

`func (o *RuleCreate) HasPosition() bool`

HasPosition returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


