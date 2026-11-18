# RequestBodyForCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Fqdn** | **string** | The fully-qualified domain name for your domain. Can be created, but not updated. | 
**ServiceID** | Pointer to **NullableString** | The `service_id` associated with your domain or `null` if there is no association. | [optional] 
**RoutingConfigurationID** | Pointer to **NullableString** | The `routing_configuration_id` associated with your domain or `null` if there is no association. | [optional] 
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


### GetServiceID

`func (o *RequestBodyForCreate) GetServiceID() string`

GetServiceID returns the ServiceID field if non-nil, zero value otherwise.

### GetServiceIDOk

`func (o *RequestBodyForCreate) GetServiceIDOk() (*string, bool)`

GetServiceIDOk returns a tuple with the ServiceID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceID

`func (o *RequestBodyForCreate) SetServiceID(v string)`

SetServiceID sets ServiceID field to given value.

### HasServiceID

`func (o *RequestBodyForCreate) HasServiceID() bool`

HasServiceID returns a boolean if a field has been set.

### SetServiceIDNil

`func (o *RequestBodyForCreate) SetServiceIDNil(b bool)`

 SetServiceIDNil sets the value for ServiceID to be an explicit nil

### UnsetServiceID
`func (o *RequestBodyForCreate) UnsetServiceID()`

UnsetServiceID ensures that no value is present for ServiceID, not even an explicit nil
### GetRoutingConfigurationID

`func (o *RequestBodyForCreate) GetRoutingConfigurationID() string`

GetRoutingConfigurationID returns the RoutingConfigurationID field if non-nil, zero value otherwise.

### GetRoutingConfigurationIDOk

`func (o *RequestBodyForCreate) GetRoutingConfigurationIDOk() (*string, bool)`

GetRoutingConfigurationIDOk returns a tuple with the RoutingConfigurationID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutingConfigurationID

`func (o *RequestBodyForCreate) SetRoutingConfigurationID(v string)`

SetRoutingConfigurationID sets RoutingConfigurationID field to given value.

### HasRoutingConfigurationID

`func (o *RequestBodyForCreate) HasRoutingConfigurationID() bool`

HasRoutingConfigurationID returns a boolean if a field has been set.

### SetRoutingConfigurationIDNil

`func (o *RequestBodyForCreate) SetRoutingConfigurationIDNil(b bool)`

 SetRoutingConfigurationIDNil sets the value for RoutingConfigurationID to be an explicit nil

### UnsetRoutingConfigurationID
`func (o *RequestBodyForCreate) UnsetRoutingConfigurationID()`

UnsetRoutingConfigurationID ensures that no value is present for RoutingConfigurationID, not even an explicit nil
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
