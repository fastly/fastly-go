# TlsActivationData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to [**TypeTlsActivation**](TypeTlsActivation.md) |  | [optional] [default to TYPETLSACTIVATION_TLS_ACTIVATION]
**Relationships** | Pointer to [**RelationshipsForTlsActivation**](RelationshipsForTlsActivation.md) |  | [optional] 

## Methods

### NewTlsActivationData

`func NewTlsActivationData() *TlsActivationData`

NewTlsActivationData instantiates a new TlsActivationData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTlsActivationDataWithDefaults

`func NewTlsActivationDataWithDefaults() *TlsActivationData`

NewTlsActivationDataWithDefaults instantiates a new TlsActivationData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *TlsActivationData) GetType() TypeTlsActivation`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TlsActivationData) GetTypeOk() (*TypeTlsActivation, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TlsActivationData) SetType(v TypeTlsActivation)`

SetType sets Type field to given value.

### HasType

`func (o *TlsActivationData) HasType() bool`

HasType returns a boolean if a field has been set.

### GetRelationships

`func (o *TlsActivationData) GetRelationships() RelationshipsForTlsActivation`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *TlsActivationData) GetRelationshipsOk() (*RelationshipsForTlsActivation, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *TlsActivationData) SetRelationships(v RelationshipsForTlsActivation)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *TlsActivationData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


