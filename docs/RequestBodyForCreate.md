# RequestBodyForCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Fqdn** | **string** | The fully-qualified domain name for your domain. Can be created, but not updated. | 
**ServiceId** | Pointer to **NullableString** | The `service_id` associated with your domain or `null` if there is no association. | [optional] 
**Description** | Pointer to **string** | A freeform descriptive note. | [optional] 

## Methods

### NewRequestBodyForCreate

`func NewRequestBodyForCreate(fqdn string, ) *RequestBodyForCreate`

NewRequestBodyForCreate instantiates a new RequestBodyForCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestBodyForCreateWithDefaults

`func NewRequestBodyForCreateWithDefaults() *RequestBodyForCreate`

NewRequestBodyForCreateWithDefaults instantiates a new RequestBodyForCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFqdn

`func (o *RequestBodyForCreate) GetFqdn() string`

GetFqdn returns the Fqdn field if non-nil, zero value otherwise.

### GetFqdnOk

`func (o *RequestBodyForCreate) GetFqdnOk() (*string, bool)`

GetFqdnOk returns a tuple with the Fqdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFqdn

`func (o *RequestBodyForCreate) SetFqdn(v string)`

SetFqdn sets Fqdn field to given value.


### GetServiceId

`func (o *RequestBodyForCreate) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *RequestBodyForCreate) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *RequestBodyForCreate) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.

### HasServiceId

`func (o *RequestBodyForCreate) HasServiceId() bool`

HasServiceId returns a boolean if a field has been set.

### SetServiceIdNil

`func (o *RequestBodyForCreate) SetServiceIdNil(b bool)`

 SetServiceIdNil sets the value for ServiceId to be an explicit nil

### UnsetServiceId
`func (o *RequestBodyForCreate) UnsetServiceId()`

UnsetServiceId ensures that no value is present for ServiceId, not even an explicit nil
### GetDescription

`func (o *RequestBodyForCreate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RequestBodyForCreate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RequestBodyForCreate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *RequestBodyForCreate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


