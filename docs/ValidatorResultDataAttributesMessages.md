# ValidatorResultDataAttributesMessages

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Warning** | **bool** |  | 
**Message** | **string** |  | 
**Tokens** | [**[]map[string]TokensAdditionalProps**](map[string]TokensAdditionalProps.md) |  | 

## Methods

### NewValidatorResultDataAttributesMessages

`func NewValidatorResultDataAttributesMessages(type_ string, warning bool, message string, tokens []map[string]TokensAdditionalProps, ) *ValidatorResultDataAttributesMessages`

NewValidatorResultDataAttributesMessages instantiates a new ValidatorResultDataAttributesMessages object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewValidatorResultDataAttributesMessagesWithDefaults

`func NewValidatorResultDataAttributesMessagesWithDefaults() *ValidatorResultDataAttributesMessages`

NewValidatorResultDataAttributesMessagesWithDefaults instantiates a new ValidatorResultDataAttributesMessages object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ValidatorResultDataAttributesMessages) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ValidatorResultDataAttributesMessages) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ValidatorResultDataAttributesMessages) SetType(v string)`

SetType sets Type field to given value.


### GetWarning

`func (o *ValidatorResultDataAttributesMessages) GetWarning() bool`

GetWarning returns the Warning field if non-nil, zero value otherwise.

### GetWarningOk

`func (o *ValidatorResultDataAttributesMessages) GetWarningOk() (*bool, bool)`

GetWarningOk returns a tuple with the Warning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarning

`func (o *ValidatorResultDataAttributesMessages) SetWarning(v bool)`

SetWarning sets Warning field to given value.


### GetMessage

`func (o *ValidatorResultDataAttributesMessages) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ValidatorResultDataAttributesMessages) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ValidatorResultDataAttributesMessages) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetTokens

`func (o *ValidatorResultDataAttributesMessages) GetTokens() []map[string]TokensAdditionalProps`

GetTokens returns the Tokens field if non-nil, zero value otherwise.

### GetTokensOk

`func (o *ValidatorResultDataAttributesMessages) GetTokensOk() (*[]map[string]TokensAdditionalProps, bool)`

GetTokensOk returns a tuple with the Tokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokens

`func (o *ValidatorResultDataAttributesMessages) SetTokens(v []map[string]TokensAdditionalProps)`

SetTokens sets Tokens field to given value.



[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


