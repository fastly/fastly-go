# SigningKey

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SigningKey** | Pointer to **string** | A Base64-encoded Ed25519 public key that can be used to verify signatures of client keys. | [optional] 

## Methods

### NewSigningKey

`func NewSigningKey() *SigningKey`

NewSigningKey instantiates a new SigningKey object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSigningKeyWithDefaults

`func NewSigningKeyWithDefaults() *SigningKey`

NewSigningKeyWithDefaults instantiates a new SigningKey object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSigningKey

`func (o *SigningKey) GetSigningKey() string`

GetSigningKey returns the SigningKey field if non-nil, zero value otherwise.

### GetSigningKeyOk

`func (o *SigningKey) GetSigningKeyOk() (*string, bool)`

GetSigningKeyOk returns a tuple with the SigningKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSigningKey

`func (o *SigningKey) SetSigningKey(v string)`

SetSigningKey sets SigningKey field to given value.

### HasSigningKey

`func (o *SigningKey) HasSigningKey() bool`

HasSigningKey returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


