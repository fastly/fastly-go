# RelationshipsForInvitation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Customer** | Pointer to [**RelationshipCustomerCustomer**](RelationshipCustomerCustomer.md) |  | [optional] 
**ServiceInvitations** | Pointer to [**RelationshipServiceInvitationsServiceInvitations**](RelationshipServiceInvitationsServiceInvitations.md) |  | [optional] 

## Methods

### NewRelationshipsForInvitation

`func NewRelationshipsForInvitation() *RelationshipsForInvitation`

NewRelationshipsForInvitation instantiates a new RelationshipsForInvitation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRelationshipsForInvitationWithDefaults

`func NewRelationshipsForInvitationWithDefaults() *RelationshipsForInvitation`

NewRelationshipsForInvitationWithDefaults instantiates a new RelationshipsForInvitation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomer

`func (o *RelationshipsForInvitation) GetCustomer() RelationshipCustomerCustomer`

GetCustomer returns the Customer field if non-nil, zero value otherwise.

### GetCustomerOk

`func (o *RelationshipsForInvitation) GetCustomerOk() (*RelationshipCustomerCustomer, bool)`

GetCustomerOk returns a tuple with the Customer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomer

`func (o *RelationshipsForInvitation) SetCustomer(v RelationshipCustomerCustomer)`

SetCustomer sets Customer field to given value.

### HasCustomer

`func (o *RelationshipsForInvitation) HasCustomer() bool`

HasCustomer returns a boolean if a field has been set.

### GetServiceInvitations

`func (o *RelationshipsForInvitation) GetServiceInvitations() RelationshipServiceInvitationsServiceInvitations`

GetServiceInvitations returns the ServiceInvitations field if non-nil, zero value otherwise.

### GetServiceInvitationsOk

`func (o *RelationshipsForInvitation) GetServiceInvitationsOk() (*RelationshipServiceInvitationsServiceInvitations, bool)`

GetServiceInvitationsOk returns a tuple with the ServiceInvitations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceInvitations

`func (o *RelationshipsForInvitation) SetServiceInvitations(v RelationshipServiceInvitationsServiceInvitations)`

SetServiceInvitations sets ServiceInvitations field to given value.

### HasServiceInvitations

`func (o *RelationshipsForInvitation) HasServiceInvitations() bool`

HasServiceInvitations returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
