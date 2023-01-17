# InvitationResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to [**TypeInvitation**](TypeInvitation.md) |  | [optional] [default to TYPEINVITATION_INVITATION]
**Attributes** | Pointer to [**Timestamps**](Timestamps.md) |  | [optional] 
**Relationships** | Pointer to [**RelationshipsForInvitation**](RelationshipsForInvitation.md) |  | [optional] 
**ID** | Pointer to **string** |  | [optional] [readonly] 

## Methods

### NewInvitationResponseData

`func NewInvitationResponseData() *InvitationResponseData`

NewInvitationResponseData instantiates a new InvitationResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvitationResponseDataWithDefaults

`func NewInvitationResponseDataWithDefaults() *InvitationResponseData`

NewInvitationResponseDataWithDefaults instantiates a new InvitationResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *InvitationResponseData) GetType() TypeInvitation`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InvitationResponseData) GetTypeOk() (*TypeInvitation, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InvitationResponseData) SetType(v TypeInvitation)`

SetType sets Type field to given value.

### HasType

`func (o *InvitationResponseData) HasType() bool`

HasType returns a boolean if a field has been set.

### GetAttributes

`func (o *InvitationResponseData) GetAttributes() Timestamps`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *InvitationResponseData) GetAttributesOk() (*Timestamps, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *InvitationResponseData) SetAttributes(v Timestamps)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *InvitationResponseData) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetRelationships

`func (o *InvitationResponseData) GetRelationships() RelationshipsForInvitation`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *InvitationResponseData) GetRelationshipsOk() (*RelationshipsForInvitation, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *InvitationResponseData) SetRelationships(v RelationshipsForInvitation)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *InvitationResponseData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.

### GetID

`func (o *InvitationResponseData) GetID() string`

GetID returns the ID field if non-nil, zero value otherwise.

### GetIDOk

`func (o *InvitationResponseData) GetIDOk() (*string, bool)`

GetIDOk returns a tuple with the ID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetID

`func (o *InvitationResponseData) SetID(v string)`

SetID sets ID field to given value.

### HasID

`func (o *InvitationResponseData) HasID() bool`

HasID returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
