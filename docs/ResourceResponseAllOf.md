# ResourceResponseAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | An alphanumeric string identifying the resource link. | [optional] 
**Href** | Pointer to **string** | The path to the resource. | [optional] 
**ServiceId** | Pointer to **string** | Alphanumeric string identifying the service. | [optional] 
**Version** | Pointer to **int32** | Integer identifying a service version. | [optional] 
**ResourceType** | Pointer to [**TypeResource**](TypeResource.md) |  | [optional] 

## Methods

### NewResourceResponseAllOf

`func NewResourceResponseAllOf() *ResourceResponseAllOf`

NewResourceResponseAllOf instantiates a new ResourceResponseAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceResponseAllOfWithDefaults

`func NewResourceResponseAllOfWithDefaults() *ResourceResponseAllOf`

NewResourceResponseAllOfWithDefaults instantiates a new ResourceResponseAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ResourceResponseAllOf) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ResourceResponseAllOf) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ResourceResponseAllOf) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ResourceResponseAllOf) HasId() bool`

HasId returns a boolean if a field has been set.

### GetHref

`func (o *ResourceResponseAllOf) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *ResourceResponseAllOf) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *ResourceResponseAllOf) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *ResourceResponseAllOf) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetServiceId

`func (o *ResourceResponseAllOf) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *ResourceResponseAllOf) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *ResourceResponseAllOf) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.

### HasServiceId

`func (o *ResourceResponseAllOf) HasServiceId() bool`

HasServiceId returns a boolean if a field has been set.

### GetVersion

`func (o *ResourceResponseAllOf) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ResourceResponseAllOf) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ResourceResponseAllOf) SetVersion(v int32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *ResourceResponseAllOf) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetResourceType

`func (o *ResourceResponseAllOf) GetResourceType() TypeResource`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *ResourceResponseAllOf) GetResourceTypeOk() (*TypeResource, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *ResourceResponseAllOf) SetResourceType(v TypeResource)`

SetResourceType sets ResourceType field to given value.

### HasResourceType

`func (o *ResourceResponseAllOf) HasResourceType() bool`

HasResourceType returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


