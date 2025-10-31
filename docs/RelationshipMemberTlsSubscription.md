# RelationshipMemberTlsSubscription

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to [**TypeTlsSubscription**](TypeTlsSubscription.md) |  | [optional] [default to TYPETLSSUBSCRIPTION_TLS_SUBSCRIPTION]
**Id** | Pointer to **string** |  | [optional] [readonly] 

## Methods

### NewRelationshipMemberTlsSubscription

`func NewRelationshipMemberTlsSubscription() *RelationshipMemberTlsSubscription`

NewRelationshipMemberTlsSubscription instantiates a new RelationshipMemberTlsSubscription object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRelationshipMemberTlsSubscriptionWithDefaults

`func NewRelationshipMemberTlsSubscriptionWithDefaults() *RelationshipMemberTlsSubscription`

NewRelationshipMemberTlsSubscriptionWithDefaults instantiates a new RelationshipMemberTlsSubscription object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *RelationshipMemberTlsSubscription) GetType() TypeTlsSubscription`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RelationshipMemberTlsSubscription) GetTypeOk() (*TypeTlsSubscription, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RelationshipMemberTlsSubscription) SetType(v TypeTlsSubscription)`

SetType sets Type field to given value.

### HasType

`func (o *RelationshipMemberTlsSubscription) HasType() bool`

HasType returns a boolean if a field has been set.

### GetId

`func (o *RelationshipMemberTlsSubscription) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RelationshipMemberTlsSubscription) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RelationshipMemberTlsSubscription) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *RelationshipMemberTlsSubscription) HasId() bool`

HasId returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


