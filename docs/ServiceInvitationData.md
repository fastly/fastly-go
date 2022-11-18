# ServiceInvitationData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to [**TypeServiceInvitation**](TypeServiceInvitation.md) |  | [optional] [default to TYPESERVICEINVITATION_SERVICE_INVITATION]
**Attributes** | Pointer to [**ServiceInvitationDataAttributes**](ServiceInvitationDataAttributes.md) |  | [optional] 
**Relationships** | Pointer to [**ServiceInvitationDataRelationships**](ServiceInvitationDataRelationships.md) |  | [optional] 

## Methods

### NewServiceInvitationData

`func NewServiceInvitationData() *ServiceInvitationData`

NewServiceInvitationData instantiates a new ServiceInvitationData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceInvitationDataWithDefaults

`func NewServiceInvitationDataWithDefaults() *ServiceInvitationData`

NewServiceInvitationDataWithDefaults instantiates a new ServiceInvitationData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ServiceInvitationData) GetType() TypeServiceInvitation`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ServiceInvitationData) GetTypeOk() (*TypeServiceInvitation, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ServiceInvitationData) SetType(v TypeServiceInvitation)`

SetType sets Type field to given value.

### HasType

`func (o *ServiceInvitationData) HasType() bool`

HasType returns a boolean if a field has been set.

### GetAttributes

`func (o *ServiceInvitationData) GetAttributes() ServiceInvitationDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *ServiceInvitationData) GetAttributesOk() (*ServiceInvitationDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *ServiceInvitationData) SetAttributes(v ServiceInvitationDataAttributes)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *ServiceInvitationData) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetRelationships

`func (o *ServiceInvitationData) GetRelationships() ServiceInvitationDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *ServiceInvitationData) GetRelationshipsOk() (*ServiceInvitationDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *ServiceInvitationData) SetRelationships(v ServiceInvitationDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *ServiceInvitationData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
