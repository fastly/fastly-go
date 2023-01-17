# RelationshipMemberTLSActivation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to [**TypeTLSActivation**](TypeTLSActivation.md) |  | [optional] [default to TYPETLSACTIVATION_TLS_ACTIVATION]
**ID** | Pointer to **string** |  | [optional] [readonly] 

## Methods

### NewRelationshipMemberTLSActivation

`func NewRelationshipMemberTLSActivation() *RelationshipMemberTLSActivation`

NewRelationshipMemberTLSActivation instantiates a new RelationshipMemberTLSActivation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRelationshipMemberTLSActivationWithDefaults

`func NewRelationshipMemberTLSActivationWithDefaults() *RelationshipMemberTLSActivation`

NewRelationshipMemberTLSActivationWithDefaults instantiates a new RelationshipMemberTLSActivation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *RelationshipMemberTLSActivation) GetType() TypeTLSActivation`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RelationshipMemberTLSActivation) GetTypeOk() (*TypeTLSActivation, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RelationshipMemberTLSActivation) SetType(v TypeTLSActivation)`

SetType sets Type field to given value.

### HasType

`func (o *RelationshipMemberTLSActivation) HasType() bool`

HasType returns a boolean if a field has been set.

### GetID

`func (o *RelationshipMemberTLSActivation) GetID() string`

GetID returns the ID field if non-nil, zero value otherwise.

### GetIDOk

`func (o *RelationshipMemberTLSActivation) GetIDOk() (*string, bool)`

GetIDOk returns a tuple with the ID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetID

`func (o *RelationshipMemberTLSActivation) SetID(v string)`

SetID sets ID field to given value.

### HasID

`func (o *RelationshipMemberTLSActivation) HasID() bool`

HasID returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
