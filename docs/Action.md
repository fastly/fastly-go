# Action

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**ActionType**](ActionType.md) |  | 
**Value** | **string** | The value for the action. For `service` actions, this is the target Fastly service ID. | 

## Methods

### NewAction

`func NewAction(type_ ActionType, value string, ) *Action`

NewAction instantiates a new Action object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActionWithDefaults

`func NewActionWithDefaults() *Action`

NewActionWithDefaults instantiates a new Action object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *Action) GetType() ActionType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Action) GetTypeOk() (*ActionType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Action) SetType(v ActionType)`

SetType sets Type field to given value.


### GetValue

`func (o *Action) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *Action) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *Action) SetValue(v string)`

SetValue sets Value field to given value.



[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


