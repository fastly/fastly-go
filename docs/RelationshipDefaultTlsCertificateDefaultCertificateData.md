# RelationshipDefaultTLSCertificateDefaultCertificateData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to [**TypeTLSCertificate**](TypeTLSCertificate.md) |  | [optional] [default to TYPETLSCERTIFICATE_TLS_CERTIFICATE]
**ID** | Pointer to **string** | Alphanumeric string identifying the default TLS certificate. | [optional] 

## Methods

### NewRelationshipDefaultTLSCertificateDefaultCertificateData

`func NewRelationshipDefaultTLSCertificateDefaultCertificateData() *RelationshipDefaultTLSCertificateDefaultCertificateData`

NewRelationshipDefaultTLSCertificateDefaultCertificateData instantiates a new RelationshipDefaultTLSCertificateDefaultCertificateData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRelationshipDefaultTLSCertificateDefaultCertificateDataWithDefaults

`func NewRelationshipDefaultTLSCertificateDefaultCertificateDataWithDefaults() *RelationshipDefaultTLSCertificateDefaultCertificateData`

NewRelationshipDefaultTLSCertificateDefaultCertificateDataWithDefaults instantiates a new RelationshipDefaultTLSCertificateDefaultCertificateData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *RelationshipDefaultTLSCertificateDefaultCertificateData) GetType() TypeTLSCertificate`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RelationshipDefaultTLSCertificateDefaultCertificateData) GetTypeOk() (*TypeTLSCertificate, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RelationshipDefaultTLSCertificateDefaultCertificateData) SetType(v TypeTLSCertificate)`

SetType sets Type field to given value.

### HasType

`func (o *RelationshipDefaultTLSCertificateDefaultCertificateData) HasType() bool`

HasType returns a boolean if a field has been set.

### GetID

`func (o *RelationshipDefaultTLSCertificateDefaultCertificateData) GetID() string`

GetID returns the ID field if non-nil, zero value otherwise.

### GetIDOk

`func (o *RelationshipDefaultTLSCertificateDefaultCertificateData) GetIDOk() (*string, bool)`

GetIDOk returns a tuple with the ID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetID

`func (o *RelationshipDefaultTLSCertificateDefaultCertificateData) SetID(v string)`

SetID sets ID field to given value.

### HasID

`func (o *RelationshipDefaultTLSCertificateDefaultCertificateData) HasID() bool`

HasID returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
