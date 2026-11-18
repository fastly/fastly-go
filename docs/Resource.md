# Resource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResourceID** | Pointer to **string** | The ID of the underlying linked resource. | [optional] 
**Name** | Pointer to **string** | The name of the resource link. Note this is separate from the resource store name and might not match the store name. | [optional] 

## Methods

### NewResource

`func NewResource() *Resource`

NewResource instantiates a new Resource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceWithDefaults

`func NewResourceWithDefaults() *Resource`

NewResourceWithDefaults instantiates a new Resource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResourceID

`func (o *Resource) GetResourceID() string`

GetResourceID returns the ResourceID field if non-nil, zero value otherwise.

### GetResourceIDOk

`func (o *Resource) GetResourceIDOk() (*string, bool)`

GetResourceIDOk returns a tuple with the ResourceID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceID

`func (o *Resource) SetResourceID(v string)`

SetResourceID sets ResourceID field to given value.

### HasResourceID

`func (o *Resource) HasResourceID() bool`

HasResourceID returns a boolean if a field has been set.

### GetName

`func (o *Resource) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Resource) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Resource) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Resource) HasName() bool`

HasName returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
