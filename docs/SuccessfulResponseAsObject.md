# SuccessfulResponseAsObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **NullableTime** | Date and time in ISO 8601 format. | [optional] [readonly] 
**UpdatedAt** | Pointer to **NullableTime** | Date and time in ISO 8601 format. | [optional] [readonly] 
**Id** | Pointer to **string** | Domain Identifier (UUID). | [optional] 
**Fqdn** | Pointer to **string** | The fully-qualified domain name for your domain. Can be created, but not updated. | [optional] 
**ServiceId** | Pointer to **NullableString** | The `service_id` associated with your domain or `null` if there is no association. | [optional] 
**Description** | Pointer to **string** | A freeform descriptive note. | [optional] 
**Activated** | Pointer to **bool** | Denotes if the domain has at least one TLS activation or not. | [optional] [readonly] 
**Verified** | Pointer to **bool** | Denotes that the customer has proven ownership over the domain by obtaining a Fastly-managed TLS certificate for it or by providing a their own TLS certificate from a publicly-trusted CA that includes the domain or matching wildcard.      | [optional] [readonly] 

## Methods

### NewSuccessfulResponseAsObject

`func NewSuccessfulResponseAsObject() *SuccessfulResponseAsObject`

NewSuccessfulResponseAsObject instantiates a new SuccessfulResponseAsObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSuccessfulResponseAsObjectWithDefaults

`func NewSuccessfulResponseAsObjectWithDefaults() *SuccessfulResponseAsObject`

NewSuccessfulResponseAsObjectWithDefaults instantiates a new SuccessfulResponseAsObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *SuccessfulResponseAsObject) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *SuccessfulResponseAsObject) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *SuccessfulResponseAsObject) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *SuccessfulResponseAsObject) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *SuccessfulResponseAsObject) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *SuccessfulResponseAsObject) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetUpdatedAt

`func (o *SuccessfulResponseAsObject) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *SuccessfulResponseAsObject) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *SuccessfulResponseAsObject) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *SuccessfulResponseAsObject) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *SuccessfulResponseAsObject) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *SuccessfulResponseAsObject) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
### GetId

`func (o *SuccessfulResponseAsObject) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SuccessfulResponseAsObject) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SuccessfulResponseAsObject) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SuccessfulResponseAsObject) HasId() bool`

HasId returns a boolean if a field has been set.

### GetFqdn

`func (o *SuccessfulResponseAsObject) GetFqdn() string`

GetFqdn returns the Fqdn field if non-nil, zero value otherwise.

### GetFqdnOk

`func (o *SuccessfulResponseAsObject) GetFqdnOk() (*string, bool)`

GetFqdnOk returns a tuple with the Fqdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFqdn

`func (o *SuccessfulResponseAsObject) SetFqdn(v string)`

SetFqdn sets Fqdn field to given value.

### HasFqdn

`func (o *SuccessfulResponseAsObject) HasFqdn() bool`

HasFqdn returns a boolean if a field has been set.

### GetServiceId

`func (o *SuccessfulResponseAsObject) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *SuccessfulResponseAsObject) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *SuccessfulResponseAsObject) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.

### HasServiceId

`func (o *SuccessfulResponseAsObject) HasServiceId() bool`

HasServiceId returns a boolean if a field has been set.

### SetServiceIdNil

`func (o *SuccessfulResponseAsObject) SetServiceIdNil(b bool)`

 SetServiceIdNil sets the value for ServiceId to be an explicit nil

### UnsetServiceId
`func (o *SuccessfulResponseAsObject) UnsetServiceId()`

UnsetServiceId ensures that no value is present for ServiceId, not even an explicit nil
### GetDescription

`func (o *SuccessfulResponseAsObject) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SuccessfulResponseAsObject) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SuccessfulResponseAsObject) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SuccessfulResponseAsObject) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetActivated

`func (o *SuccessfulResponseAsObject) GetActivated() bool`

GetActivated returns the Activated field if non-nil, zero value otherwise.

### GetActivatedOk

`func (o *SuccessfulResponseAsObject) GetActivatedOk() (*bool, bool)`

GetActivatedOk returns a tuple with the Activated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivated

`func (o *SuccessfulResponseAsObject) SetActivated(v bool)`

SetActivated sets Activated field to given value.

### HasActivated

`func (o *SuccessfulResponseAsObject) HasActivated() bool`

HasActivated returns a boolean if a field has been set.

### GetVerified

`func (o *SuccessfulResponseAsObject) GetVerified() bool`

GetVerified returns the Verified field if non-nil, zero value otherwise.

### GetVerifiedOk

`func (o *SuccessfulResponseAsObject) GetVerifiedOk() (*bool, bool)`

GetVerifiedOk returns a tuple with the Verified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerified

`func (o *SuccessfulResponseAsObject) SetVerified(v bool)`

SetVerified sets Verified field to given value.

### HasVerified

`func (o *SuccessfulResponseAsObject) HasVerified() bool`

HasVerified returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


