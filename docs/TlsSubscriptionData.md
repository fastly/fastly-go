# TlsSubscriptionData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to [**TypeTlsSubscription**](TypeTlsSubscription.md) |  | [optional] [default to TYPETLSSUBSCRIPTION_TLS_SUBSCRIPTION]
**Attributes** | Pointer to [**TlsSubscriptionDataAttributes**](TlsSubscriptionDataAttributes.md) |  | [optional] 
**Relationships** | Pointer to [**RelationshipsForTlsSubscription**](RelationshipsForTlsSubscription.md) |  | [optional] 

## Methods

### NewTlsSubscriptionData

`func NewTlsSubscriptionData() *TlsSubscriptionData`

NewTlsSubscriptionData instantiates a new TlsSubscriptionData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTlsSubscriptionDataWithDefaults

`func NewTlsSubscriptionDataWithDefaults() *TlsSubscriptionData`

NewTlsSubscriptionDataWithDefaults instantiates a new TlsSubscriptionData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *TlsSubscriptionData) GetType() TypeTlsSubscription`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TlsSubscriptionData) GetTypeOk() (*TypeTlsSubscription, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TlsSubscriptionData) SetType(v TypeTlsSubscription)`

SetType sets Type field to given value.

### HasType

`func (o *TlsSubscriptionData) HasType() bool`

HasType returns a boolean if a field has been set.

### GetAttributes

`func (o *TlsSubscriptionData) GetAttributes() TlsSubscriptionDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *TlsSubscriptionData) GetAttributesOk() (*TlsSubscriptionDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *TlsSubscriptionData) SetAttributes(v TlsSubscriptionDataAttributes)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *TlsSubscriptionData) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetRelationships

`func (o *TlsSubscriptionData) GetRelationships() RelationshipsForTlsSubscription`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *TlsSubscriptionData) GetRelationshipsOk() (*RelationshipsForTlsSubscription, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *TlsSubscriptionData) SetRelationships(v RelationshipsForTlsSubscription)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *TlsSubscriptionData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


