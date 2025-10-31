# TlsActivationResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to [**TypeTlsActivation**](TypeTlsActivation.md) |  | [optional] [default to TYPETLSACTIVATION_TLS_ACTIVATION]
**Relationships** | Pointer to [**RelationshipsForTlsActivation**](RelationshipsForTlsActivation.md) |  | [optional] 
**Id** | Pointer to **string** |  | [optional] [readonly] 
**Attributes** | Pointer to [**Timestamps**](Timestamps.md) |  | [optional] 

## Methods

### NewTlsActivationResponseData

`func NewTlsActivationResponseData() *TlsActivationResponseData`

NewTlsActivationResponseData instantiates a new TlsActivationResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTlsActivationResponseDataWithDefaults

`func NewTlsActivationResponseDataWithDefaults() *TlsActivationResponseData`

NewTlsActivationResponseDataWithDefaults instantiates a new TlsActivationResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *TlsActivationResponseData) GetType() TypeTlsActivation`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TlsActivationResponseData) GetTypeOk() (*TypeTlsActivation, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TlsActivationResponseData) SetType(v TypeTlsActivation)`

SetType sets Type field to given value.

### HasType

`func (o *TlsActivationResponseData) HasType() bool`

HasType returns a boolean if a field has been set.

### GetRelationships

`func (o *TlsActivationResponseData) GetRelationships() RelationshipsForTlsActivation`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *TlsActivationResponseData) GetRelationshipsOk() (*RelationshipsForTlsActivation, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *TlsActivationResponseData) SetRelationships(v RelationshipsForTlsActivation)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *TlsActivationResponseData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.

### GetId

`func (o *TlsActivationResponseData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TlsActivationResponseData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TlsActivationResponseData) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TlsActivationResponseData) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAttributes

`func (o *TlsActivationResponseData) GetAttributes() Timestamps`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *TlsActivationResponseData) GetAttributesOk() (*Timestamps, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *TlsActivationResponseData) SetAttributes(v Timestamps)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *TlsActivationResponseData) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


