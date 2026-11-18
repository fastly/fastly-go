# ValidatorResultDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Msg** | Pointer to **NullableString** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Errors** | Pointer to **[]string** |  | [optional] 
**Warnings** | Pointer to **[]string** |  | [optional] 
**Messages** | Pointer to [**[]ValidatorResultDataAttributesMessages**](ValidatorResultDataAttributesMessages.md) |  | [optional] 

## Methods

### NewValidatorResultDataAttributes

`func NewValidatorResultDataAttributes() *ValidatorResultDataAttributes`

NewValidatorResultDataAttributes instantiates a new ValidatorResultDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewValidatorResultDataAttributesWithDefaults

`func NewValidatorResultDataAttributesWithDefaults() *ValidatorResultDataAttributes`

NewValidatorResultDataAttributesWithDefaults instantiates a new ValidatorResultDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMsg

`func (o *ValidatorResultDataAttributes) GetMsg() string`

GetMsg returns the Msg field if non-nil, zero value otherwise.

### GetMsgOk

`func (o *ValidatorResultDataAttributes) GetMsgOk() (*string, bool)`

GetMsgOk returns a tuple with the Msg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsg

`func (o *ValidatorResultDataAttributes) SetMsg(v string)`

SetMsg sets Msg field to given value.

### HasMsg

`func (o *ValidatorResultDataAttributes) HasMsg() bool`

HasMsg returns a boolean if a field has been set.

### SetMsgNil

`func (o *ValidatorResultDataAttributes) SetMsgNil(b bool)`

 SetMsgNil sets the value for Msg to be an explicit nil

### UnsetMsg
`func (o *ValidatorResultDataAttributes) UnsetMsg()`

UnsetMsg ensures that no value is present for Msg, not even an explicit nil
### GetStatus

`func (o *ValidatorResultDataAttributes) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ValidatorResultDataAttributes) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ValidatorResultDataAttributes) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ValidatorResultDataAttributes) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetErrors

`func (o *ValidatorResultDataAttributes) GetErrors() []string`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *ValidatorResultDataAttributes) GetErrorsOk() (*[]string, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *ValidatorResultDataAttributes) SetErrors(v []string)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *ValidatorResultDataAttributes) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetWarnings

`func (o *ValidatorResultDataAttributes) GetWarnings() []string`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *ValidatorResultDataAttributes) GetWarningsOk() (*[]string, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *ValidatorResultDataAttributes) SetWarnings(v []string)`

SetWarnings sets Warnings field to given value.

### HasWarnings

`func (o *ValidatorResultDataAttributes) HasWarnings() bool`

HasWarnings returns a boolean if a field has been set.

### GetMessages

`func (o *ValidatorResultDataAttributes) GetMessages() []ValidatorResultDataAttributesMessages`

GetMessages returns the Messages field if non-nil, zero value otherwise.

### GetMessagesOk

`func (o *ValidatorResultDataAttributes) GetMessagesOk() (*[]ValidatorResultDataAttributesMessages, bool)`

GetMessagesOk returns a tuple with the Messages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessages

`func (o *ValidatorResultDataAttributes) SetMessages(v []ValidatorResultDataAttributesMessages)`

SetMessages sets Messages field to given value.

### HasMessages

`func (o *ValidatorResultDataAttributes) HasMessages() bool`

HasMessages returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
