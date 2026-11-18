# TLSSubscriptionData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to [**TypeTLSSubscription**](TypeTLSSubscription.md) |  | [optional] [default to TYPETLSSUBSCRIPTION_TLS_SUBSCRIPTION]
**Attributes** | Pointer to [**TLSSubscriptionDataAttributes**](TlsSubscriptionDataAttributes.md) |  | [optional] 
**Relationships** | Pointer to [**RelationshipsForTLSSubscription**](RelationshipsForTLSSubscription.md) |  | [optional] 

## Methods

### NewTLSSubscriptionData

`func NewTLSSubscriptionData() *TLSSubscriptionData`

NewTLSSubscriptionData instantiates a new TLSSubscriptionData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTLSSubscriptionDataWithDefaults

`func NewTLSSubscriptionDataWithDefaults() *TLSSubscriptionData`

NewTLSSubscriptionDataWithDefaults instantiates a new TLSSubscriptionData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *TLSSubscriptionData) GetType() TypeTLSSubscription`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TLSSubscriptionData) GetTypeOk() (*TypeTLSSubscription, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TLSSubscriptionData) SetType(v TypeTLSSubscription)`

SetType sets Type field to given value.

### HasType

`func (o *TLSSubscriptionData) HasType() bool`

HasType returns a boolean if a field has been set.

### GetAttributes

`func (o *TLSSubscriptionData) GetAttributes() TLSSubscriptionDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *TLSSubscriptionData) GetAttributesOk() (*TLSSubscriptionDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *TLSSubscriptionData) SetAttributes(v TLSSubscriptionDataAttributes)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *TLSSubscriptionData) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetRelationships

`func (o *TLSSubscriptionData) GetRelationships() RelationshipsForTLSSubscription`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *TLSSubscriptionData) GetRelationshipsOk() (*RelationshipsForTLSSubscription, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *TLSSubscriptionData) SetRelationships(v RelationshipsForTLSSubscription)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *TLSSubscriptionData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
