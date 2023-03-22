# RelationshipMemberTLSSubscription

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to [**TypeTLSSubscription**](TypeTLSSubscription.md) |  | [optional] [default to TYPETLSSUBSCRIPTION_TLS_SUBSCRIPTION]
**ID** | Pointer to **string** |  | [optional] [readonly] 

## Methods

### NewRelationshipMemberTLSSubscription

`func NewRelationshipMemberTLSSubscription() *RelationshipMemberTLSSubscription`

NewRelationshipMemberTLSSubscription instantiates a new RelationshipMemberTLSSubscription object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRelationshipMemberTLSSubscriptionWithDefaults

`func NewRelationshipMemberTLSSubscriptionWithDefaults() *RelationshipMemberTLSSubscription`

NewRelationshipMemberTLSSubscriptionWithDefaults instantiates a new RelationshipMemberTLSSubscription object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *RelationshipMemberTLSSubscription) GetType() TypeTLSSubscription`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RelationshipMemberTLSSubscription) GetTypeOk() (*TypeTLSSubscription, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RelationshipMemberTLSSubscription) SetType(v TypeTLSSubscription)`

SetType sets Type field to given value.

### HasType

`func (o *RelationshipMemberTLSSubscription) HasType() bool`

HasType returns a boolean if a field has been set.

### GetID

`func (o *RelationshipMemberTLSSubscription) GetID() string`

GetID returns the ID field if non-nil, zero value otherwise.

### GetIDOk

`func (o *RelationshipMemberTLSSubscription) GetIDOk() (*string, bool)`

GetIDOk returns a tuple with the ID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetID

`func (o *RelationshipMemberTLSSubscription) SetID(v string)`

SetID sets ID field to given value.

### HasID

`func (o *RelationshipMemberTLSSubscription) HasID() bool`

HasID returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
