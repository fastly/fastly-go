# InvitationDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | Pointer to **string** | The email address of the invitee. | [optional] 
**LimitServices** | Pointer to **bool** | Indicates the user has limited access to the customer&#39;s services. | [optional] 
**Role** | Pointer to [**RoleUser**](RoleUser.md) |  | [optional] 
**StatusCode** | Pointer to **int32** | Indicates whether or not the invitation is active. | [optional] 

## Methods

### NewInvitationDataAttributes

`func NewInvitationDataAttributes() *InvitationDataAttributes`

NewInvitationDataAttributes instantiates a new InvitationDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvitationDataAttributesWithDefaults

`func NewInvitationDataAttributesWithDefaults() *InvitationDataAttributes`

NewInvitationDataAttributesWithDefaults instantiates a new InvitationDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *InvitationDataAttributes) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *InvitationDataAttributes) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *InvitationDataAttributes) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *InvitationDataAttributes) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetLimitServices

`func (o *InvitationDataAttributes) GetLimitServices() bool`

GetLimitServices returns the LimitServices field if non-nil, zero value otherwise.

### GetLimitServicesOk

`func (o *InvitationDataAttributes) GetLimitServicesOk() (*bool, bool)`

GetLimitServicesOk returns a tuple with the LimitServices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimitServices

`func (o *InvitationDataAttributes) SetLimitServices(v bool)`

SetLimitServices sets LimitServices field to given value.

### HasLimitServices

`func (o *InvitationDataAttributes) HasLimitServices() bool`

HasLimitServices returns a boolean if a field has been set.

### GetRole

`func (o *InvitationDataAttributes) GetRole() RoleUser`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *InvitationDataAttributes) GetRoleOk() (*RoleUser, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *InvitationDataAttributes) SetRole(v RoleUser)`

SetRole sets Role field to given value.

### HasRole

`func (o *InvitationDataAttributes) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetStatusCode

`func (o *InvitationDataAttributes) GetStatusCode() int32`

GetStatusCode returns the StatusCode field if non-nil, zero value otherwise.

### GetStatusCodeOk

`func (o *InvitationDataAttributes) GetStatusCodeOk() (*int32, bool)`

GetStatusCodeOk returns a tuple with the StatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCode

`func (o *InvitationDataAttributes) SetStatusCode(v int32)`

SetStatusCode sets StatusCode field to given value.

### HasStatusCode

`func (o *InvitationDataAttributes) HasStatusCode() bool`

HasStatusCode returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
