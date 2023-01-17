# TLSActivationData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to [**TypeTLSActivation**](TypeTLSActivation.md) |  | [optional] [default to TYPETLSACTIVATION_TLS_ACTIVATION]
**Relationships** | Pointer to [**RelationshipsForTLSActivation**](RelationshipsForTLSActivation.md) |  | [optional] 

## Methods

### NewTLSActivationData

`func NewTLSActivationData() *TLSActivationData`

NewTLSActivationData instantiates a new TLSActivationData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTLSActivationDataWithDefaults

`func NewTLSActivationDataWithDefaults() *TLSActivationData`

NewTLSActivationDataWithDefaults instantiates a new TLSActivationData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *TLSActivationData) GetType() TypeTLSActivation`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TLSActivationData) GetTypeOk() (*TypeTLSActivation, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TLSActivationData) SetType(v TypeTLSActivation)`

SetType sets Type field to given value.

### HasType

`func (o *TLSActivationData) HasType() bool`

HasType returns a boolean if a field has been set.

### GetRelationships

`func (o *TLSActivationData) GetRelationships() RelationshipsForTLSActivation`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *TLSActivationData) GetRelationshipsOk() (*RelationshipsForTLSActivation, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *TLSActivationData) SetRelationships(v RelationshipsForTLSActivation)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *TLSActivationData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
