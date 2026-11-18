# ValidatorResultData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ID** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Attributes** | Pointer to [**ValidatorResultDataAttributes**](ValidatorResultDataAttributes.md) |  | [optional] 

## Methods

### NewValidatorResultData

`func NewValidatorResultData() *ValidatorResultData`

NewValidatorResultData instantiates a new ValidatorResultData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewValidatorResultDataWithDefaults

`func NewValidatorResultDataWithDefaults() *ValidatorResultData`

NewValidatorResultDataWithDefaults instantiates a new ValidatorResultData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetID

`func (o *ValidatorResultData) GetID() string`

GetID returns the ID field if non-nil, zero value otherwise.

### GetIDOk

`func (o *ValidatorResultData) GetIDOk() (*string, bool)`

GetIDOk returns a tuple with the ID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetID

`func (o *ValidatorResultData) SetID(v string)`

SetID sets ID field to given value.

### HasID

`func (o *ValidatorResultData) HasID() bool`

HasID returns a boolean if a field has been set.

### GetType

`func (o *ValidatorResultData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ValidatorResultData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ValidatorResultData) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ValidatorResultData) HasType() bool`

HasType returns a boolean if a field has been set.

### GetAttributes

`func (o *ValidatorResultData) GetAttributes() ValidatorResultDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *ValidatorResultData) GetAttributesOk() (*ValidatorResultDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *ValidatorResultData) SetAttributes(v ValidatorResultDataAttributes)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *ValidatorResultData) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
