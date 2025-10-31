# InvitationCreateData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to [**TypeInvitation**](TypeInvitation.md) |  | [optional] [default to TYPEINVITATION_INVITATION]
**Attributes** | Pointer to [**InvitationDataAttributes**](InvitationDataAttributes.md) |  | [optional] 
**Relationships** | Pointer to [**RelationshipServiceInvitationsCreate**](RelationshipServiceInvitationsCreate.md) |  | [optional] 

## Methods

### NewInvitationCreateData

`func NewInvitationCreateData() *InvitationCreateData`

NewInvitationCreateData instantiates a new InvitationCreateData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvitationCreateDataWithDefaults

`func NewInvitationCreateDataWithDefaults() *InvitationCreateData`

NewInvitationCreateDataWithDefaults instantiates a new InvitationCreateData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *InvitationCreateData) GetType() TypeInvitation`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InvitationCreateData) GetTypeOk() (*TypeInvitation, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InvitationCreateData) SetType(v TypeInvitation)`

SetType sets Type field to given value.

### HasType

`func (o *InvitationCreateData) HasType() bool`

HasType returns a boolean if a field has been set.

### GetAttributes

`func (o *InvitationCreateData) GetAttributes() InvitationDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *InvitationCreateData) GetAttributesOk() (*InvitationDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *InvitationCreateData) SetAttributes(v InvitationDataAttributes)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *InvitationCreateData) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetRelationships

`func (o *InvitationCreateData) GetRelationships() RelationshipServiceInvitationsCreate`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *InvitationCreateData) GetRelationshipsOk() (*RelationshipServiceInvitationsCreate, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *InvitationCreateData) SetRelationships(v RelationshipServiceInvitationsCreate)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *InvitationCreateData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


