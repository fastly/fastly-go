# SuccessfulResponseAsObjectAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Domain Identifier (UUID). | [optional] 
**Fqdn** | Pointer to **string** | The fully-qualified domain name for your domain. Can be created, but not updated. | [optional] 
**ServiceId** | Pointer to **NullableString** | The `service_id` associated with your domain or `null` if there is no association. | [optional] 
**Description** | Pointer to **string** | A freeform descriptive note. | [optional] 
**Activated** | Pointer to **bool** | Denotes if the domain has at least one TLS activation or not. | [optional] [readonly] 
**Verified** | Pointer to **bool** | Denotes that the customer has proven ownership over the domain by obtaining a Fastly-managed TLS certificate for it or by providing a their own TLS certificate from a publicly-trusted CA that includes the domain or matching wildcard.      | [optional] [readonly] 

## Methods

### NewSuccessfulResponseAsObjectAllOf

`func NewSuccessfulResponseAsObjectAllOf() *SuccessfulResponseAsObjectAllOf`

NewSuccessfulResponseAsObjectAllOf instantiates a new SuccessfulResponseAsObjectAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSuccessfulResponseAsObjectAllOfWithDefaults

`func NewSuccessfulResponseAsObjectAllOfWithDefaults() *SuccessfulResponseAsObjectAllOf`

NewSuccessfulResponseAsObjectAllOfWithDefaults instantiates a new SuccessfulResponseAsObjectAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SuccessfulResponseAsObjectAllOf) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SuccessfulResponseAsObjectAllOf) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SuccessfulResponseAsObjectAllOf) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SuccessfulResponseAsObjectAllOf) HasId() bool`

HasId returns a boolean if a field has been set.

### GetFqdn

`func (o *SuccessfulResponseAsObjectAllOf) GetFqdn() string`

GetFqdn returns the Fqdn field if non-nil, zero value otherwise.

### GetFqdnOk

`func (o *SuccessfulResponseAsObjectAllOf) GetFqdnOk() (*string, bool)`

GetFqdnOk returns a tuple with the Fqdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFqdn

`func (o *SuccessfulResponseAsObjectAllOf) SetFqdn(v string)`

SetFqdn sets Fqdn field to given value.

### HasFqdn

`func (o *SuccessfulResponseAsObjectAllOf) HasFqdn() bool`

HasFqdn returns a boolean if a field has been set.

### GetServiceId

`func (o *SuccessfulResponseAsObjectAllOf) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *SuccessfulResponseAsObjectAllOf) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *SuccessfulResponseAsObjectAllOf) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.

### HasServiceId

`func (o *SuccessfulResponseAsObjectAllOf) HasServiceId() bool`

HasServiceId returns a boolean if a field has been set.

### SetServiceIdNil

`func (o *SuccessfulResponseAsObjectAllOf) SetServiceIdNil(b bool)`

 SetServiceIdNil sets the value for ServiceId to be an explicit nil

### UnsetServiceId
`func (o *SuccessfulResponseAsObjectAllOf) UnsetServiceId()`

UnsetServiceId ensures that no value is present for ServiceId, not even an explicit nil
### GetDescription

`func (o *SuccessfulResponseAsObjectAllOf) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SuccessfulResponseAsObjectAllOf) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SuccessfulResponseAsObjectAllOf) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SuccessfulResponseAsObjectAllOf) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetActivated

`func (o *SuccessfulResponseAsObjectAllOf) GetActivated() bool`

GetActivated returns the Activated field if non-nil, zero value otherwise.

### GetActivatedOk

`func (o *SuccessfulResponseAsObjectAllOf) GetActivatedOk() (*bool, bool)`

GetActivatedOk returns a tuple with the Activated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivated

`func (o *SuccessfulResponseAsObjectAllOf) SetActivated(v bool)`

SetActivated sets Activated field to given value.

### HasActivated

`func (o *SuccessfulResponseAsObjectAllOf) HasActivated() bool`

HasActivated returns a boolean if a field has been set.

### GetVerified

`func (o *SuccessfulResponseAsObjectAllOf) GetVerified() bool`

GetVerified returns the Verified field if non-nil, zero value otherwise.

### GetVerifiedOk

`func (o *SuccessfulResponseAsObjectAllOf) GetVerifiedOk() (*bool, bool)`

GetVerifiedOk returns a tuple with the Verified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerified

`func (o *SuccessfulResponseAsObjectAllOf) SetVerified(v bool)`

SetVerified sets Verified field to given value.

### HasVerified

`func (o *SuccessfulResponseAsObjectAllOf) HasVerified() bool`

HasVerified returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


