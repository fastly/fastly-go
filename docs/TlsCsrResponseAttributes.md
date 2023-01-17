# TLSCsrResponseAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Content** | Pointer to **string** | The PEM encoded CSR. | [optional] 

## Methods

### NewTLSCsrResponseAttributes

`func NewTLSCsrResponseAttributes() *TLSCsrResponseAttributes`

NewTLSCsrResponseAttributes instantiates a new TLSCsrResponseAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTLSCsrResponseAttributesWithDefaults

`func NewTLSCsrResponseAttributesWithDefaults() *TLSCsrResponseAttributes`

NewTLSCsrResponseAttributesWithDefaults instantiates a new TLSCsrResponseAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContent

`func (o *TLSCsrResponseAttributes) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *TLSCsrResponseAttributes) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *TLSCsrResponseAttributes) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *TLSCsrResponseAttributes) HasContent() bool`

HasContent returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
