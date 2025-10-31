# InvitationResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to [**TypeInvitation**](TypeInvitation.md) |  | [optional] [default to TYPEINVITATION_INVITATION]
**Attributes** | Pointer to [**Timestamps**](Timestamps.md) |  | [optional] 
**Id** | Pointer to **string** |  | [optional] [readonly] 
**Relationships** | Pointer to [**RelationshipsForInvitation**](RelationshipsForInvitation.md) |  | [optional] 

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

### GetId

`func (o *InvitationResponseData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InvitationResponseData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InvitationResponseData) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *InvitationResponseData) HasId() bool`

HasId returns a boolean if a field has been set.

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


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


