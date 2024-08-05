# TLSDNSRecord

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Region** | Pointer to **string** | Specifies the regions that will be used to route traffic. Select DNS records with a `global` region to route traffic to the most performant point of presence (POP) worldwide (global pricing will apply). Select DNS records with a `na/eu` region to exclusively land traffic on North American and European POPs. | [optional] 
**RecordType** | Pointer to **string** | The type of the DNS record. `A` specifies an IPv4 address to be used for an A record to be used for apex domains (e.g., `example.com`). `AAAA` specifies an IPv6 address for use in an A record for apex domains. `CNAME` specifies the hostname to be used for a CNAME record for subdomains or wildcard domains (e.g., `www.example.com` or `*.example.com`). | [optional] 

## Methods

### NewTLSDNSRecord

`func NewTLSDNSRecord() *TLSDNSRecord`

NewTLSDNSRecord instantiates a new TLSDNSRecord object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTLSDNSRecordWithDefaults

`func NewTLSDNSRecordWithDefaults() *TLSDNSRecord`

NewTLSDNSRecordWithDefaults instantiates a new TLSDNSRecord object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRegion

`func (o *TLSDNSRecord) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *TLSDNSRecord) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *TLSDNSRecord) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *TLSDNSRecord) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetRecordType

`func (o *TLSDNSRecord) GetRecordType() string`

GetRecordType returns the RecordType field if non-nil, zero value otherwise.

### GetRecordTypeOk

`func (o *TLSDNSRecord) GetRecordTypeOk() (*string, bool)`

GetRecordTypeOk returns a tuple with the RecordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordType

`func (o *TLSDNSRecord) SetRecordType(v string)`

SetRecordType sets RecordType field to given value.

### HasRecordType

`func (o *TLSDNSRecord) HasRecordType() bool`

HasRecordType returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
