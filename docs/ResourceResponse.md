# ResourceResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResourceID** | Pointer to **string** | The ID of the underlying linked resource. | [optional] 
**Name** | Pointer to **string** | The name of the resource link. | [optional] 
**CreatedAt** | Pointer to **NullableTime** | Date and time in ISO 8601 format. | [optional] [readonly] 
**DeletedAt** | Pointer to **NullableTime** | Date and time in ISO 8601 format. | [optional] [readonly] 
**UpdatedAt** | Pointer to **NullableTime** | Date and time in ISO 8601 format. | [optional] [readonly] 
**ID** | Pointer to **string** | An alphanumeric string identifying the resource link. | [optional] 
**Href** | Pointer to **string** | The path to the resource. | [optional] 
**ServiceID** | Pointer to **string** | Alphanumeric string identifying the service. | [optional] 
**Version** | Pointer to **int32** | Integer identifying a service version. | [optional] 
**ResourceType** | Pointer to [**TypeResource**](TypeResource.md) |  | [optional] 

## Methods

### NewResourceResponse

`func NewResourceResponse() *ResourceResponse`

NewResourceResponse instantiates a new ResourceResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceResponseWithDefaults

`func NewResourceResponseWithDefaults() *ResourceResponse`

NewResourceResponseWithDefaults instantiates a new ResourceResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResourceID

`func (o *ResourceResponse) GetResourceID() string`

GetResourceID returns the ResourceID field if non-nil, zero value otherwise.

### GetResourceIDOk

`func (o *ResourceResponse) GetResourceIDOk() (*string, bool)`

GetResourceIDOk returns a tuple with the ResourceID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceID

`func (o *ResourceResponse) SetResourceID(v string)`

SetResourceID sets ResourceID field to given value.

### HasResourceID

`func (o *ResourceResponse) HasResourceID() bool`

HasResourceID returns a boolean if a field has been set.

### GetName

`func (o *ResourceResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ResourceResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ResourceResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ResourceResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ResourceResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ResourceResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ResourceResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ResourceResponse) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *ResourceResponse) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *ResourceResponse) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetDeletedAt

`func (o *ResourceResponse) GetDeletedAt() time.Time`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *ResourceResponse) GetDeletedAtOk() (*time.Time, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *ResourceResponse) SetDeletedAt(v time.Time)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *ResourceResponse) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### SetDeletedAtNil

`func (o *ResourceResponse) SetDeletedAtNil(b bool)`

 SetDeletedAtNil sets the value for DeletedAt to be an explicit nil

### UnsetDeletedAt
`func (o *ResourceResponse) UnsetDeletedAt()`

UnsetDeletedAt ensures that no value is present for DeletedAt, not even an explicit nil
### GetUpdatedAt

`func (o *ResourceResponse) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ResourceResponse) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ResourceResponse) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ResourceResponse) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *ResourceResponse) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *ResourceResponse) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
### GetID

`func (o *ResourceResponse) GetID() string`

GetID returns the ID field if non-nil, zero value otherwise.

### GetIDOk

`func (o *ResourceResponse) GetIDOk() (*string, bool)`

GetIDOk returns a tuple with the ID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetID

`func (o *ResourceResponse) SetID(v string)`

SetID sets ID field to given value.

### HasID

`func (o *ResourceResponse) HasID() bool`

HasID returns a boolean if a field has been set.

### GetHref

`func (o *ResourceResponse) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *ResourceResponse) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *ResourceResponse) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *ResourceResponse) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetServiceID

`func (o *ResourceResponse) GetServiceID() string`

GetServiceID returns the ServiceID field if non-nil, zero value otherwise.

### GetServiceIDOk

`func (o *ResourceResponse) GetServiceIDOk() (*string, bool)`

GetServiceIDOk returns a tuple with the ServiceID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceID

`func (o *ResourceResponse) SetServiceID(v string)`

SetServiceID sets ServiceID field to given value.

### HasServiceID

`func (o *ResourceResponse) HasServiceID() bool`

HasServiceID returns a boolean if a field has been set.

### GetVersion

`func (o *ResourceResponse) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ResourceResponse) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ResourceResponse) SetVersion(v int32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *ResourceResponse) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetResourceType

`func (o *ResourceResponse) GetResourceType() TypeResource`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *ResourceResponse) GetResourceTypeOk() (*TypeResource, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *ResourceResponse) SetResourceType(v TypeResource)`

SetResourceType sets ResourceType field to given value.

### HasResourceType

`func (o *ResourceResponse) HasResourceType() bool`

HasResourceType returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
