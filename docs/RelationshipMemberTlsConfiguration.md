# RelationshipMemberTlsConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to [**TypeTlsConfiguration**](TypeTlsConfiguration.md) |  | [optional] [default to TYPETLSCONFIGURATION_TLS_CONFIGURATION]
**Id** | Pointer to **string** |  | [optional] [readonly] 

## Methods

### NewRelationshipMemberTlsConfiguration

`func NewRelationshipMemberTlsConfiguration() *RelationshipMemberTlsConfiguration`

NewRelationshipMemberTlsConfiguration instantiates a new RelationshipMemberTlsConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRelationshipMemberTlsConfigurationWithDefaults

`func NewRelationshipMemberTlsConfigurationWithDefaults() *RelationshipMemberTlsConfiguration`

NewRelationshipMemberTlsConfigurationWithDefaults instantiates a new RelationshipMemberTlsConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *RelationshipMemberTlsConfiguration) GetType() TypeTlsConfiguration`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RelationshipMemberTlsConfiguration) GetTypeOk() (*TypeTlsConfiguration, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RelationshipMemberTlsConfiguration) SetType(v TypeTlsConfiguration)`

SetType sets Type field to given value.

### HasType

`func (o *RelationshipMemberTlsConfiguration) HasType() bool`

HasType returns a boolean if a field has been set.

### GetId

`func (o *RelationshipMemberTlsConfiguration) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RelationshipMemberTlsConfiguration) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RelationshipMemberTlsConfiguration) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *RelationshipMemberTlsConfiguration) HasId() bool`

HasId returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


