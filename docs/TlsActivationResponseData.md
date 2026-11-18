# TLSActivationResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to [**TypeTLSActivation**](TypeTLSActivation.md) |  | [optional] [default to TYPETLSACTIVATION_TLS_ACTIVATION]
**Relationships** | Pointer to [**RelationshipsForTLSActivation**](RelationshipsForTLSActivation.md) |  | [optional] 
**ID** | Pointer to **string** |  | [optional] [readonly] 
**Attributes** | Pointer to [**Timestamps**](Timestamps.md) |  | [optional] 

## Methods

### NewTLSActivationResponseData

`func NewTLSActivationResponseData() *TLSActivationResponseData`

NewTLSActivationResponseData instantiates a new TLSActivationResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTLSActivationResponseDataWithDefaults

`func NewTLSActivationResponseDataWithDefaults() *TLSActivationResponseData`

NewTLSActivationResponseDataWithDefaults instantiates a new TLSActivationResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *TLSActivationResponseData) GetType() TypeTLSActivation`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TLSActivationResponseData) GetTypeOk() (*TypeTLSActivation, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TLSActivationResponseData) SetType(v TypeTLSActivation)`

SetType sets Type field to given value.

### HasType

`func (o *TLSActivationResponseData) HasType() bool`

HasType returns a boolean if a field has been set.

### GetRelationships

`func (o *TLSActivationResponseData) GetRelationships() RelationshipsForTLSActivation`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *TLSActivationResponseData) GetRelationshipsOk() (*RelationshipsForTLSActivation, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *TLSActivationResponseData) SetRelationships(v RelationshipsForTLSActivation)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *TLSActivationResponseData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.

### GetID

`func (o *TLSActivationResponseData) GetID() string`

GetID returns the ID field if non-nil, zero value otherwise.

### GetIDOk

`func (o *TLSActivationResponseData) GetIDOk() (*string, bool)`

GetIDOk returns a tuple with the ID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetID

`func (o *TLSActivationResponseData) SetID(v string)`

SetID sets ID field to given value.

### HasID

`func (o *TLSActivationResponseData) HasID() bool`

HasID returns a boolean if a field has been set.

### GetAttributes

`func (o *TLSActivationResponseData) GetAttributes() Timestamps`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *TLSActivationResponseData) GetAttributesOk() (*Timestamps, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *TLSActivationResponseData) SetAttributes(v Timestamps)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *TLSActivationResponseData) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
