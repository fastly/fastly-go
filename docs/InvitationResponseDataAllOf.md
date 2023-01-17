# InvitationResponseDataAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ID** | Pointer to **string** |  | [optional] [readonly] 
**Attributes** | Pointer to [**Timestamps**](Timestamps.md) |  | [optional] 
**Relationships** | Pointer to [**RelationshipsForInvitation**](RelationshipsForInvitation.md) |  | [optional] 

## Methods

### NewInvitationResponseDataAllOf

`func NewInvitationResponseDataAllOf() *InvitationResponseDataAllOf`

NewInvitationResponseDataAllOf instantiates a new InvitationResponseDataAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvitationResponseDataAllOfWithDefaults

`func NewInvitationResponseDataAllOfWithDefaults() *InvitationResponseDataAllOf`

NewInvitationResponseDataAllOfWithDefaults instantiates a new InvitationResponseDataAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetID

`func (o *InvitationResponseDataAllOf) GetID() string`

GetID returns the ID field if non-nil, zero value otherwise.

### GetIDOk

`func (o *InvitationResponseDataAllOf) GetIDOk() (*string, bool)`

GetIDOk returns a tuple with the ID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetID

`func (o *InvitationResponseDataAllOf) SetID(v string)`

SetID sets ID field to given value.

### HasID

`func (o *InvitationResponseDataAllOf) HasID() bool`

HasID returns a boolean if a field has been set.

### GetAttributes

`func (o *InvitationResponseDataAllOf) GetAttributes() Timestamps`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *InvitationResponseDataAllOf) GetAttributesOk() (*Timestamps, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *InvitationResponseDataAllOf) SetAttributes(v Timestamps)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *InvitationResponseDataAllOf) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetRelationships

`func (o *InvitationResponseDataAllOf) GetRelationships() RelationshipsForInvitation`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *InvitationResponseDataAllOf) GetRelationshipsOk() (*RelationshipsForInvitation, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *InvitationResponseDataAllOf) SetRelationships(v RelationshipsForInvitation)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *InvitationResponseDataAllOf) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
