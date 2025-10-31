# RelationshipMemberTlsActivation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to [**TypeTlsActivation**](TypeTlsActivation.md) |  | [optional] [default to TYPETLSACTIVATION_TLS_ACTIVATION]
**Id** | Pointer to **string** |  | [optional] [readonly] 

## Methods

### NewRelationshipMemberTlsActivation

`func NewRelationshipMemberTlsActivation() *RelationshipMemberTlsActivation`

NewRelationshipMemberTlsActivation instantiates a new RelationshipMemberTlsActivation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRelationshipMemberTlsActivationWithDefaults

`func NewRelationshipMemberTlsActivationWithDefaults() *RelationshipMemberTlsActivation`

NewRelationshipMemberTlsActivationWithDefaults instantiates a new RelationshipMemberTlsActivation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *RelationshipMemberTlsActivation) GetType() TypeTlsActivation`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RelationshipMemberTlsActivation) GetTypeOk() (*TypeTlsActivation, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RelationshipMemberTlsActivation) SetType(v TypeTlsActivation)`

SetType sets Type field to given value.

### HasType

`func (o *RelationshipMemberTlsActivation) HasType() bool`

HasType returns a boolean if a field has been set.

### GetId

`func (o *RelationshipMemberTlsActivation) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RelationshipMemberTlsActivation) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RelationshipMemberTlsActivation) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *RelationshipMemberTlsActivation) HasId() bool`

HasId returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


