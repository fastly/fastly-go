# ServiceAuthorizationData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to [**TypeServiceAuthorization**](TypeServiceAuthorization.md) |  | [optional] [default to TYPESERVICEAUTHORIZATION_SERVICE_AUTHORIZATION]
**Attributes** | Pointer to [**ServiceAuthorizationDataAttributes**](ServiceAuthorizationDataAttributes.md) |  | [optional] 
**Relationships** | Pointer to [**ServiceAuthorizationDataRelationships**](ServiceAuthorizationDataRelationships.md) |  | [optional] 

## Methods

### NewServiceAuthorizationData

`func NewServiceAuthorizationData() *ServiceAuthorizationData`

NewServiceAuthorizationData instantiates a new ServiceAuthorizationData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceAuthorizationDataWithDefaults

`func NewServiceAuthorizationDataWithDefaults() *ServiceAuthorizationData`

NewServiceAuthorizationDataWithDefaults instantiates a new ServiceAuthorizationData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ServiceAuthorizationData) GetType() TypeServiceAuthorization`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ServiceAuthorizationData) GetTypeOk() (*TypeServiceAuthorization, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ServiceAuthorizationData) SetType(v TypeServiceAuthorization)`

SetType sets Type field to given value.

### HasType

`func (o *ServiceAuthorizationData) HasType() bool`

HasType returns a boolean if a field has been set.

### GetAttributes

`func (o *ServiceAuthorizationData) GetAttributes() ServiceAuthorizationDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *ServiceAuthorizationData) GetAttributesOk() (*ServiceAuthorizationDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *ServiceAuthorizationData) SetAttributes(v ServiceAuthorizationDataAttributes)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *ServiceAuthorizationData) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetRelationships

`func (o *ServiceAuthorizationData) GetRelationships() ServiceAuthorizationDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *ServiceAuthorizationData) GetRelationshipsOk() (*ServiceAuthorizationDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *ServiceAuthorizationData) SetRelationships(v ServiceAuthorizationDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *ServiceAuthorizationData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
