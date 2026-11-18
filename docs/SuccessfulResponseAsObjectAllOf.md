# SuccessfulResponseAsObjectAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ID** | Pointer to **string** | Domain Identifier (UUID). | [optional] 
**Fqdn** | Pointer to **string** | The fully-qualified domain name for your domain. Can be created, but not updated. | [optional] 
**ServiceID** | Pointer to **NullableString** | The `service_id` associated with your domain or `null` if there is no association. | [optional] 
**RoutingConfigurationID** | Pointer to **NullableString** | The `routing_configuration_id` associated with your domain or `null` if there is no association. | [optional] 
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

### GetID

`func (o *SuccessfulResponseAsObjectAllOf) GetID() string`

GetID returns the ID field if non-nil, zero value otherwise.

### GetIDOk

`func (o *SuccessfulResponseAsObjectAllOf) GetIDOk() (*string, bool)`

GetIDOk returns a tuple with the ID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetID

`func (o *SuccessfulResponseAsObjectAllOf) SetID(v string)`

SetID sets ID field to given value.

### HasID

`func (o *SuccessfulResponseAsObjectAllOf) HasID() bool`

HasID returns a boolean if a field has been set.

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

### GetServiceID

`func (o *SuccessfulResponseAsObjectAllOf) GetServiceID() string`

GetServiceID returns the ServiceID field if non-nil, zero value otherwise.

### GetServiceIDOk

`func (o *SuccessfulResponseAsObjectAllOf) GetServiceIDOk() (*string, bool)`

GetServiceIDOk returns a tuple with the ServiceID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceID

`func (o *SuccessfulResponseAsObjectAllOf) SetServiceID(v string)`

SetServiceID sets ServiceID field to given value.

### HasServiceID

`func (o *SuccessfulResponseAsObjectAllOf) HasServiceID() bool`

HasServiceID returns a boolean if a field has been set.

### SetServiceIDNil

`func (o *SuccessfulResponseAsObjectAllOf) SetServiceIDNil(b bool)`

 SetServiceIDNil sets the value for ServiceID to be an explicit nil

### UnsetServiceID
`func (o *SuccessfulResponseAsObjectAllOf) UnsetServiceID()`

UnsetServiceID ensures that no value is present for ServiceID, not even an explicit nil
### GetRoutingConfigurationID

`func (o *SuccessfulResponseAsObjectAllOf) GetRoutingConfigurationID() string`

GetRoutingConfigurationID returns the RoutingConfigurationID field if non-nil, zero value otherwise.

### GetRoutingConfigurationIDOk

`func (o *SuccessfulResponseAsObjectAllOf) GetRoutingConfigurationIDOk() (*string, bool)`

GetRoutingConfigurationIDOk returns a tuple with the RoutingConfigurationID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutingConfigurationID

`func (o *SuccessfulResponseAsObjectAllOf) SetRoutingConfigurationID(v string)`

SetRoutingConfigurationID sets RoutingConfigurationID field to given value.

### HasRoutingConfigurationID

`func (o *SuccessfulResponseAsObjectAllOf) HasRoutingConfigurationID() bool`

HasRoutingConfigurationID returns a boolean if a field has been set.

### SetRoutingConfigurationIDNil

`func (o *SuccessfulResponseAsObjectAllOf) SetRoutingConfigurationIDNil(b bool)`

 SetRoutingConfigurationIDNil sets the value for RoutingConfigurationID to be an explicit nil

### UnsetRoutingConfigurationID
`func (o *SuccessfulResponseAsObjectAllOf) UnsetRoutingConfigurationID()`

UnsetRoutingConfigurationID ensures that no value is present for RoutingConfigurationID, not even an explicit nil
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
