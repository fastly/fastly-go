# ServiceInvitationDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Permission** | Pointer to **string** | The permission the accepting user will have in relation to the service. | [optional] [default to "full"]

## Methods

### NewServiceInvitationDataAttributes

`func NewServiceInvitationDataAttributes() *ServiceInvitationDataAttributes`

NewServiceInvitationDataAttributes instantiates a new ServiceInvitationDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceInvitationDataAttributesWithDefaults

`func NewServiceInvitationDataAttributesWithDefaults() *ServiceInvitationDataAttributes`

NewServiceInvitationDataAttributesWithDefaults instantiates a new ServiceInvitationDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPermission

`func (o *ServiceInvitationDataAttributes) GetPermission() string`

GetPermission returns the Permission field if non-nil, zero value otherwise.

### GetPermissionOk

`func (o *ServiceInvitationDataAttributes) GetPermissionOk() (*string, bool)`

GetPermissionOk returns a tuple with the Permission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermission

`func (o *ServiceInvitationDataAttributes) SetPermission(v string)`

SetPermission sets Permission field to given value.

### HasPermission

`func (o *ServiceInvitationDataAttributes) HasPermission() bool`

HasPermission returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


