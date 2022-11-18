# TLSPrivateKeyDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | A customizable name for your private key. Optional. | [optional] 
**Key** | Pointer to **string** | The contents of the private key. Must be a PEM-formatted key. Not returned in response body. Required. | [optional] 

## Methods

### NewTLSPrivateKeyDataAttributes

`func NewTLSPrivateKeyDataAttributes() *TLSPrivateKeyDataAttributes`

NewTLSPrivateKeyDataAttributes instantiates a new TLSPrivateKeyDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTLSPrivateKeyDataAttributesWithDefaults

`func NewTLSPrivateKeyDataAttributesWithDefaults() *TLSPrivateKeyDataAttributes`

NewTLSPrivateKeyDataAttributesWithDefaults instantiates a new TLSPrivateKeyDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *TLSPrivateKeyDataAttributes) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TLSPrivateKeyDataAttributes) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TLSPrivateKeyDataAttributes) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TLSPrivateKeyDataAttributes) HasName() bool`

HasName returns a boolean if a field has been set.

### GetKey

`func (o *TLSPrivateKeyDataAttributes) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *TLSPrivateKeyDataAttributes) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *TLSPrivateKeyDataAttributes) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *TLSPrivateKeyDataAttributes) HasKey() bool`

HasKey returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
