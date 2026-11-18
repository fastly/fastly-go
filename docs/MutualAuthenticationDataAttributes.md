# MutualAuthenticationDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CertBundle** | Pointer to **string** | One or more certificates. Enter each individual certificate blob on a new line. Must be PEM-formatted. Required on create. You may optionally rotate the cert_bundle on update. | [optional] 
**Enforced** | Pointer to **bool** | Determines whether Mutual TLS will fail closed (enforced) or fail open. A true value will require a successful Mutual TLS handshake for the connection to continue and will fail closed if unsuccessful. A false value will fail open and allow the connection to proceed. Optional. Defaults to true. | [optional] 
**Name** | Pointer to **string** | A custom name for your mutual authentication. Optional. If name is not supplied we will auto-generate one. | [optional] 

## Methods

### NewMutualAuthenticationDataAttributes

`func NewMutualAuthenticationDataAttributes() *MutualAuthenticationDataAttributes`

NewMutualAuthenticationDataAttributes instantiates a new MutualAuthenticationDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMutualAuthenticationDataAttributesWithDefaults

`func NewMutualAuthenticationDataAttributesWithDefaults() *MutualAuthenticationDataAttributes`

NewMutualAuthenticationDataAttributesWithDefaults instantiates a new MutualAuthenticationDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCertBundle

`func (o *MutualAuthenticationDataAttributes) GetCertBundle() string`

GetCertBundle returns the CertBundle field if non-nil, zero value otherwise.

### GetCertBundleOk

`func (o *MutualAuthenticationDataAttributes) GetCertBundleOk() (*string, bool)`

GetCertBundleOk returns a tuple with the CertBundle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertBundle

`func (o *MutualAuthenticationDataAttributes) SetCertBundle(v string)`

SetCertBundle sets CertBundle field to given value.

### HasCertBundle

`func (o *MutualAuthenticationDataAttributes) HasCertBundle() bool`

HasCertBundle returns a boolean if a field has been set.

### GetEnforced

`func (o *MutualAuthenticationDataAttributes) GetEnforced() bool`

GetEnforced returns the Enforced field if non-nil, zero value otherwise.

### GetEnforcedOk

`func (o *MutualAuthenticationDataAttributes) GetEnforcedOk() (*bool, bool)`

GetEnforcedOk returns a tuple with the Enforced field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnforced

`func (o *MutualAuthenticationDataAttributes) SetEnforced(v bool)`

SetEnforced sets Enforced field to given value.

### HasEnforced

`func (o *MutualAuthenticationDataAttributes) HasEnforced() bool`

HasEnforced returns a boolean if a field has been set.

### GetName

`func (o *MutualAuthenticationDataAttributes) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MutualAuthenticationDataAttributes) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MutualAuthenticationDataAttributes) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MutualAuthenticationDataAttributes) HasName() bool`

HasName returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
