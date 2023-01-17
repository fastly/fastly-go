# ServiceAuthorizationResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to [**TypeServiceAuthorization**](TypeServiceAuthorization.md) |  | [optional] [default to TYPESERVICEAUTHORIZATION_SERVICE_AUTHORIZATION]
**Attributes** | Pointer to [**Timestamps**](Timestamps.md) |  | [optional] 
**Relationships** | Pointer to [**ServiceAuthorizationDataRelationships**](ServiceAuthorizationDataRelationships.md) |  | [optional] 
**ID** | Pointer to **string** |  | [optional] [readonly] 

## Methods

### NewServiceAuthorizationResponseData

`func NewServiceAuthorizationResponseData() *ServiceAuthorizationResponseData`

NewServiceAuthorizationResponseData instantiates a new ServiceAuthorizationResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceAuthorizationResponseDataWithDefaults

`func NewServiceAuthorizationResponseDataWithDefaults() *ServiceAuthorizationResponseData`

NewServiceAuthorizationResponseDataWithDefaults instantiates a new ServiceAuthorizationResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ServiceAuthorizationResponseData) GetType() TypeServiceAuthorization`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ServiceAuthorizationResponseData) GetTypeOk() (*TypeServiceAuthorization, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ServiceAuthorizationResponseData) SetType(v TypeServiceAuthorization)`

SetType sets Type field to given value.

### HasType

`func (o *ServiceAuthorizationResponseData) HasType() bool`

HasType returns a boolean if a field has been set.

### GetAttributes

`func (o *ServiceAuthorizationResponseData) GetAttributes() Timestamps`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *ServiceAuthorizationResponseData) GetAttributesOk() (*Timestamps, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *ServiceAuthorizationResponseData) SetAttributes(v Timestamps)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *ServiceAuthorizationResponseData) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetRelationships

`func (o *ServiceAuthorizationResponseData) GetRelationships() ServiceAuthorizationDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *ServiceAuthorizationResponseData) GetRelationshipsOk() (*ServiceAuthorizationDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *ServiceAuthorizationResponseData) SetRelationships(v ServiceAuthorizationDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *ServiceAuthorizationResponseData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.

### GetID

`func (o *ServiceAuthorizationResponseData) GetID() string`

GetID returns the ID field if non-nil, zero value otherwise.

### GetIDOk

`func (o *ServiceAuthorizationResponseData) GetIDOk() (*string, bool)`

GetIDOk returns a tuple with the ID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetID

`func (o *ServiceAuthorizationResponseData) SetID(v string)`

SetID sets ID field to given value.

### HasID

`func (o *ServiceAuthorizationResponseData) HasID() bool`

HasID returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
