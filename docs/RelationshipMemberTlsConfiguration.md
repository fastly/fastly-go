# RelationshipMemberTLSConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to [**TypeTLSConfiguration**](TypeTLSConfiguration.md) |  | [optional] [default to TYPETLSCONFIGURATION_TLS_CONFIGURATION]
**ID** | Pointer to **string** |  | [optional] [readonly] 

## Methods

### NewRelationshipMemberTLSConfiguration

`func NewRelationshipMemberTLSConfiguration() *RelationshipMemberTLSConfiguration`

NewRelationshipMemberTLSConfiguration instantiates a new RelationshipMemberTLSConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRelationshipMemberTLSConfigurationWithDefaults

`func NewRelationshipMemberTLSConfigurationWithDefaults() *RelationshipMemberTLSConfiguration`

NewRelationshipMemberTLSConfigurationWithDefaults instantiates a new RelationshipMemberTLSConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *RelationshipMemberTLSConfiguration) GetType() TypeTLSConfiguration`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RelationshipMemberTLSConfiguration) GetTypeOk() (*TypeTLSConfiguration, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RelationshipMemberTLSConfiguration) SetType(v TypeTLSConfiguration)`

SetType sets Type field to given value.

### HasType

`func (o *RelationshipMemberTLSConfiguration) HasType() bool`

HasType returns a boolean if a field has been set.

### GetID

`func (o *RelationshipMemberTLSConfiguration) GetID() string`

GetID returns the ID field if non-nil, zero value otherwise.

### GetIDOk

`func (o *RelationshipMemberTLSConfiguration) GetIDOk() (*string, bool)`

GetIDOk returns a tuple with the ID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetID

`func (o *RelationshipMemberTLSConfiguration) SetID(v string)`

SetID sets ID field to given value.

### HasID

`func (o *RelationshipMemberTLSConfiguration) HasID() bool`

HasID returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
