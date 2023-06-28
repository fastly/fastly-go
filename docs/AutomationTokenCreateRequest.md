# AutomationTokenCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attributes** | Pointer to [**AutomationTokenCreateRequestAttributes**](AutomationTokenCreateRequestAttributes.md) |  | [optional] 

## Methods

### NewAutomationTokenCreateRequest

`func NewAutomationTokenCreateRequest() *AutomationTokenCreateRequest`

NewAutomationTokenCreateRequest instantiates a new AutomationTokenCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutomationTokenCreateRequestWithDefaults

`func NewAutomationTokenCreateRequestWithDefaults() *AutomationTokenCreateRequest`

NewAutomationTokenCreateRequestWithDefaults instantiates a new AutomationTokenCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttributes

`func (o *AutomationTokenCreateRequest) GetAttributes() AutomationTokenCreateRequestAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *AutomationTokenCreateRequest) GetAttributesOk() (*AutomationTokenCreateRequestAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *AutomationTokenCreateRequest) SetAttributes(v AutomationTokenCreateRequestAttributes)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *AutomationTokenCreateRequest) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
