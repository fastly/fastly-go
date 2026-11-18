# NetworkProtectRuleSetResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultAction** | [**NetworkProtectRuleSetDefaultAction**](NetworkProtectRuleSetDefaultAction.md) |  | 
**Rules** | [**[]NetworkProtectRuleResponse**](NetworkProtectRuleResponse.md) | Rules in evaluation order with read-only timestamps. | 
**ID** | **string** | A read-only unique identifier for a Network Protect resource. | [readonly] 
**CreatedAt** | **time.Time** | An ISO 8601 timestamp in UTC. | [readonly] 
**UpdatedAt** | **time.Time** | An ISO 8601 timestamp in UTC. | [readonly] 

## Methods

### NewNetworkProtectRuleSetResponse

`func NewNetworkProtectRuleSetResponse(defaultAction NetworkProtectRuleSetDefaultAction, rules []NetworkProtectRuleResponse, id string, createdAt time.Time, updatedAt time.Time, ) *NetworkProtectRuleSetResponse`

NewNetworkProtectRuleSetResponse instantiates a new NetworkProtectRuleSetResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkProtectRuleSetResponseWithDefaults

`func NewNetworkProtectRuleSetResponseWithDefaults() *NetworkProtectRuleSetResponse`

NewNetworkProtectRuleSetResponseWithDefaults instantiates a new NetworkProtectRuleSetResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefaultAction

`func (o *NetworkProtectRuleSetResponse) GetDefaultAction() NetworkProtectRuleSetDefaultAction`

GetDefaultAction returns the DefaultAction field if non-nil, zero value otherwise.

### GetDefaultActionOk

`func (o *NetworkProtectRuleSetResponse) GetDefaultActionOk() (*NetworkProtectRuleSetDefaultAction, bool)`

GetDefaultActionOk returns a tuple with the DefaultAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultAction

`func (o *NetworkProtectRuleSetResponse) SetDefaultAction(v NetworkProtectRuleSetDefaultAction)`

SetDefaultAction sets DefaultAction field to given value.


### GetRules

`func (o *NetworkProtectRuleSetResponse) GetRules() []NetworkProtectRuleResponse`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *NetworkProtectRuleSetResponse) GetRulesOk() (*[]NetworkProtectRuleResponse, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *NetworkProtectRuleSetResponse) SetRules(v []NetworkProtectRuleResponse)`

SetRules sets Rules field to given value.


### GetID

`func (o *NetworkProtectRuleSetResponse) GetID() string`

GetID returns the ID field if non-nil, zero value otherwise.

### GetIDOk

`func (o *NetworkProtectRuleSetResponse) GetIDOk() (*string, bool)`

GetIDOk returns a tuple with the ID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetID

`func (o *NetworkProtectRuleSetResponse) SetID(v string)`

SetID sets ID field to given value.


### GetCreatedAt

`func (o *NetworkProtectRuleSetResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *NetworkProtectRuleSetResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *NetworkProtectRuleSetResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *NetworkProtectRuleSetResponse) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *NetworkProtectRuleSetResponse) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *NetworkProtectRuleSetResponse) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
