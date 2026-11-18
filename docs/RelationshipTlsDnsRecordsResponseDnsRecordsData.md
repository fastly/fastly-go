# RelationshipTLSDNSRecordsResponseDNSRecordsData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to [**TypeTLSDNSRecord**](TypeTLSDNSRecord.md) |  | [optional] [default to TYPETLSDNSRECORD_DNS_RECORD]
**ID** | Pointer to **string** |  | [optional] [readonly] 

## Methods

### NewRelationshipTLSDNSRecordsResponseDNSRecordsData

`func NewRelationshipTLSDNSRecordsResponseDNSRecordsData() *RelationshipTLSDNSRecordsResponseDNSRecordsData`

NewRelationshipTLSDNSRecordsResponseDNSRecordsData instantiates a new RelationshipTLSDNSRecordsResponseDNSRecordsData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRelationshipTLSDNSRecordsResponseDNSRecordsDataWithDefaults

`func NewRelationshipTLSDNSRecordsResponseDNSRecordsDataWithDefaults() *RelationshipTLSDNSRecordsResponseDNSRecordsData`

NewRelationshipTLSDNSRecordsResponseDNSRecordsDataWithDefaults instantiates a new RelationshipTLSDNSRecordsResponseDNSRecordsData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *RelationshipTLSDNSRecordsResponseDNSRecordsData) GetType() TypeTLSDNSRecord`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RelationshipTLSDNSRecordsResponseDNSRecordsData) GetTypeOk() (*TypeTLSDNSRecord, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RelationshipTLSDNSRecordsResponseDNSRecordsData) SetType(v TypeTLSDNSRecord)`

SetType sets Type field to given value.

### HasType

`func (o *RelationshipTLSDNSRecordsResponseDNSRecordsData) HasType() bool`

HasType returns a boolean if a field has been set.

### GetID

`func (o *RelationshipTLSDNSRecordsResponseDNSRecordsData) GetID() string`

GetID returns the ID field if non-nil, zero value otherwise.

### GetIDOk

`func (o *RelationshipTLSDNSRecordsResponseDNSRecordsData) GetIDOk() (*string, bool)`

GetIDOk returns a tuple with the ID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetID

`func (o *RelationshipTLSDNSRecordsResponseDNSRecordsData) SetID(v string)`

SetID sets ID field to given value.

### HasID

`func (o *RelationshipTLSDNSRecordsResponseDNSRecordsData) HasID() bool`

HasID returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
