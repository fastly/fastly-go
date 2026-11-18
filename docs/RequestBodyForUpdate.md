# RequestBodyForUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServiceID** | Pointer to **NullableString** | The `service_id` associated with your domain or `null` if there is no association. | [optional] 
**RoutingConfigurationID** | Pointer to **NullableString** | The `routing_configuration_id` associated with your domain or `null` if there is no association. | [optional] 
**Description** | Pointer to **string** | A freeform descriptive note. | [optional] 

## Methods

### NewRequestBodyForUpdate

`func NewRequestBodyForUpdate() *RequestBodyForUpdate`

NewRequestBodyForUpdate instantiates a new RequestBodyForUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestBodyForUpdateWithDefaults

`func NewRequestBodyForUpdateWithDefaults() *RequestBodyForUpdate`

NewRequestBodyForUpdateWithDefaults instantiates a new RequestBodyForUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServiceID

`func (o *RequestBodyForUpdate) GetServiceID() string`

GetServiceID returns the ServiceID field if non-nil, zero value otherwise.

### GetServiceIDOk

`func (o *RequestBodyForUpdate) GetServiceIDOk() (*string, bool)`

GetServiceIDOk returns a tuple with the ServiceID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceID

`func (o *RequestBodyForUpdate) SetServiceID(v string)`

SetServiceID sets ServiceID field to given value.

### HasServiceID

`func (o *RequestBodyForUpdate) HasServiceID() bool`

HasServiceID returns a boolean if a field has been set.

### SetServiceIDNil

`func (o *RequestBodyForUpdate) SetServiceIDNil(b bool)`

 SetServiceIDNil sets the value for ServiceID to be an explicit nil

### UnsetServiceID
`func (o *RequestBodyForUpdate) UnsetServiceID()`

UnsetServiceID ensures that no value is present for ServiceID, not even an explicit nil
### GetRoutingConfigurationID

`func (o *RequestBodyForUpdate) GetRoutingConfigurationID() string`

GetRoutingConfigurationID returns the RoutingConfigurationID field if non-nil, zero value otherwise.

### GetRoutingConfigurationIDOk

`func (o *RequestBodyForUpdate) GetRoutingConfigurationIDOk() (*string, bool)`

GetRoutingConfigurationIDOk returns a tuple with the RoutingConfigurationID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutingConfigurationID

`func (o *RequestBodyForUpdate) SetRoutingConfigurationID(v string)`

SetRoutingConfigurationID sets RoutingConfigurationID field to given value.

### HasRoutingConfigurationID

`func (o *RequestBodyForUpdate) HasRoutingConfigurationID() bool`

HasRoutingConfigurationID returns a boolean if a field has been set.

### SetRoutingConfigurationIDNil

`func (o *RequestBodyForUpdate) SetRoutingConfigurationIDNil(b bool)`

 SetRoutingConfigurationIDNil sets the value for RoutingConfigurationID to be an explicit nil

### UnsetRoutingConfigurationID
`func (o *RequestBodyForUpdate) UnsetRoutingConfigurationID()`

UnsetRoutingConfigurationID ensures that no value is present for RoutingConfigurationID, not even an explicit nil
### GetDescription

`func (o *RequestBodyForUpdate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RequestBodyForUpdate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RequestBodyForUpdate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *RequestBodyForUpdate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
