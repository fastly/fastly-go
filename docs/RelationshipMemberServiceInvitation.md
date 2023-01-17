# RelationshipMemberServiceInvitation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to [**TypeServiceInvitation**](TypeServiceInvitation.md) |  | [optional] [default to TYPESERVICEINVITATION_SERVICE_INVITATION]
**ID** | Pointer to **string** | Alphanumeric string identifying a service invitation. | [optional] [readonly] 

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

### GetID

`func (o *RelationshipMemberServiceInvitation) GetID() string`

GetID returns the ID field if non-nil, zero value otherwise.

### GetIDOk

`func (o *RelationshipMemberServiceInvitation) GetIDOk() (*string, bool)`

GetIDOk returns a tuple with the ID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetID

`func (o *RelationshipMemberServiceInvitation) SetID(v string)`

SetID sets ID field to given value.

### HasID

`func (o *RelationshipMemberServiceInvitation) HasID() bool`

HasID returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
