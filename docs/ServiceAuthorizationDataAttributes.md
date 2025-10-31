# ServiceAuthorizationDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Permission** | Pointer to [**Permission**](Permission.md) |  | [optional] [default to PERMISSION_FULL]

## Methods

### NewServiceAuthorizationDataAttributes

`func NewServiceAuthorizationDataAttributes() *ServiceAuthorizationDataAttributes`

NewServiceAuthorizationDataAttributes instantiates a new ServiceAuthorizationDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceAuthorizationDataAttributesWithDefaults

`func NewServiceAuthorizationDataAttributesWithDefaults() *ServiceAuthorizationDataAttributes`

NewServiceAuthorizationDataAttributesWithDefaults instantiates a new ServiceAuthorizationDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPermission

`func (o *ServiceAuthorizationDataAttributes) GetPermission() Permission`

GetPermission returns the Permission field if non-nil, zero value otherwise.

### GetPermissionOk

`func (o *ServiceAuthorizationDataAttributes) GetPermissionOk() (*Permission, bool)`

GetPermissionOk returns a tuple with the Permission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermission

`func (o *ServiceAuthorizationDataAttributes) SetPermission(v Permission)`

SetPermission sets Permission field to given value.

### HasPermission

`func (o *ServiceAuthorizationDataAttributes) HasPermission() bool`

HasPermission returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


