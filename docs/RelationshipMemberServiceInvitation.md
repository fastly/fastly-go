# RelationshipMemberServiceInvitation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to [**TypeServiceInvitation**](TypeServiceInvitation.md) |  | [optional] [default to TYPESERVICEINVITATION_SERVICE_INVITATION]
**Id** | Pointer to **string** | Alphanumeric string identifying a service invitation. | [optional] [readonly] 

## Methods

### NewRelationshipMemberServiceInvitation

`func NewRelationshipMemberServiceInvitation() *RelationshipMemberServiceInvitation`

NewRelationshipMemberServiceInvitation instantiates a new RelationshipMemberServiceInvitation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRelationshipMemberServiceInvitationWithDefaults

`func NewRelationshipMemberServiceInvitationWithDefaults() *RelationshipMemberServiceInvitation`

NewRelationshipMemberServiceInvitationWithDefaults instantiates a new RelationshipMemberServiceInvitation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *RelationshipMemberServiceInvitation) GetType() TypeServiceInvitation`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RelationshipMemberServiceInvitation) GetTypeOk() (*TypeServiceInvitation, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RelationshipMemberServiceInvitation) SetType(v TypeServiceInvitation)`

SetType sets Type field to given value.

### HasType

`func (o *RelationshipMemberServiceInvitation) HasType() bool`

HasType returns a boolean if a field has been set.

### GetId

`func (o *RelationshipMemberServiceInvitation) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RelationshipMemberServiceInvitation) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RelationshipMemberServiceInvitation) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *RelationshipMemberServiceInvitation) HasId() bool`

HasId returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


