# RelationshipMemberTLSDNSRecord

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to [**TypeTLSDNSRecord**](TypeTLSDNSRecord.md) |  | [optional] [default to TYPETLSDNSRECORD_DNS_RECORD]
**ID** | Pointer to **string** |  | [optional] [readonly] 

## Methods

### NewRelationshipMemberTLSDNSRecord

`func NewRelationshipMemberTLSDNSRecord() *RelationshipMemberTLSDNSRecord`

NewRelationshipMemberTLSDNSRecord instantiates a new RelationshipMemberTLSDNSRecord object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRelationshipMemberTLSDNSRecordWithDefaults

`func NewRelationshipMemberTLSDNSRecordWithDefaults() *RelationshipMemberTLSDNSRecord`

NewRelationshipMemberTLSDNSRecordWithDefaults instantiates a new RelationshipMemberTLSDNSRecord object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *RelationshipMemberTLSDNSRecord) GetType() TypeTLSDNSRecord`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RelationshipMemberTLSDNSRecord) GetTypeOk() (*TypeTLSDNSRecord, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RelationshipMemberTLSDNSRecord) SetType(v TypeTLSDNSRecord)`

SetType sets Type field to given value.

### HasType

`func (o *RelationshipMemberTLSDNSRecord) HasType() bool`

HasType returns a boolean if a field has been set.

### GetID

`func (o *RelationshipMemberTLSDNSRecord) GetID() string`

GetID returns the ID field if non-nil, zero value otherwise.

### GetIDOk

`func (o *RelationshipMemberTLSDNSRecord) GetIDOk() (*string, bool)`

GetIDOk returns a tuple with the ID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetID

`func (o *RelationshipMemberTLSDNSRecord) SetID(v string)`

SetID sets ID field to given value.

### HasID

`func (o *RelationshipMemberTLSDNSRecord) HasID() bool`

HasID returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
