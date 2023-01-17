# InvitationData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to [**TypeInvitation**](TypeInvitation.md) |  | [optional] [default to TYPEINVITATION_INVITATION]
**Attributes** | Pointer to [**InvitationDataAttributes**](InvitationDataAttributes.md) |  | [optional] 
**Relationships** | Pointer to [**RelationshipServiceInvitationsCreate**](RelationshipServiceInvitationsCreate.md) |  | [optional] 

## Methods

### NewInvitationData

`func NewInvitationData() *InvitationData`

NewInvitationData instantiates a new InvitationData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvitationDataWithDefaults

`func NewInvitationDataWithDefaults() *InvitationData`

NewInvitationDataWithDefaults instantiates a new InvitationData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *InvitationData) GetType() TypeInvitation`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InvitationData) GetTypeOk() (*TypeInvitation, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InvitationData) SetType(v TypeInvitation)`

SetType sets Type field to given value.

### HasType

`func (o *InvitationData) HasType() bool`

HasType returns a boolean if a field has been set.

### GetAttributes

`func (o *InvitationData) GetAttributes() InvitationDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *InvitationData) GetAttributesOk() (*InvitationDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *InvitationData) SetAttributes(v InvitationDataAttributes)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *InvitationData) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetRelationships

`func (o *InvitationData) GetRelationships() RelationshipServiceInvitationsCreate`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *InvitationData) GetRelationshipsOk() (*RelationshipServiceInvitationsCreate, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *InvitationData) SetRelationships(v RelationshipServiceInvitationsCreate)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *InvitationData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
