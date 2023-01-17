# ResourceCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name of the resource. | [optional] 
**ResourceID** | Pointer to **string** | The ID of the linked resource. | [optional] 

## Methods

### NewResourceCreate

`func NewResourceCreate() *ResourceCreate`

NewResourceCreate instantiates a new ResourceCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceCreateWithDefaults

`func NewResourceCreateWithDefaults() *ResourceCreate`

NewResourceCreateWithDefaults instantiates a new ResourceCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ResourceCreate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ResourceCreate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ResourceCreate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ResourceCreate) HasName() bool`

HasName returns a boolean if a field has been set.

### GetResourceID

`func (o *ResourceCreate) GetResourceID() string`

GetResourceID returns the ResourceID field if non-nil, zero value otherwise.

### GetResourceIDOk

`func (o *ResourceCreate) GetResourceIDOk() (*string, bool)`

GetResourceIDOk returns a tuple with the ResourceID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceID

`func (o *ResourceCreate) SetResourceID(v string)`

SetResourceID sets ResourceID field to given value.

### HasResourceID

`func (o *ResourceCreate) HasResourceID() bool`

HasResourceID returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
