# DomainInspectorEntryDimensions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Region** | Pointer to **string** | The geographic region from which the edge responses in this data entry were delivered. If unspecified, results are aggregated across regions. | [optional] 
**Datacenter** | Pointer to **string** | The POP from which the edge responses in this data entry were delivered. If unspecified, results are aggregated across POPs. | [optional] 
**Domain** | Pointer to **string** | The domain from which the edge responses in this data entry were delivered. If unspecified, results are aggregated across domains. | [optional] 

## Methods

### NewDomainInspectorEntryDimensions

`func NewDomainInspectorEntryDimensions() *DomainInspectorEntryDimensions`

NewDomainInspectorEntryDimensions instantiates a new DomainInspectorEntryDimensions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDomainInspectorEntryDimensionsWithDefaults

`func NewDomainInspectorEntryDimensionsWithDefaults() *DomainInspectorEntryDimensions`

NewDomainInspectorEntryDimensionsWithDefaults instantiates a new DomainInspectorEntryDimensions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRegion

`func (o *DomainInspectorEntryDimensions) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *DomainInspectorEntryDimensions) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *DomainInspectorEntryDimensions) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *DomainInspectorEntryDimensions) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetDatacenter

`func (o *DomainInspectorEntryDimensions) GetDatacenter() string`

GetDatacenter returns the Datacenter field if non-nil, zero value otherwise.

### GetDatacenterOk

`func (o *DomainInspectorEntryDimensions) GetDatacenterOk() (*string, bool)`

GetDatacenterOk returns a tuple with the Datacenter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatacenter

`func (o *DomainInspectorEntryDimensions) SetDatacenter(v string)`

SetDatacenter sets Datacenter field to given value.

### HasDatacenter

`func (o *DomainInspectorEntryDimensions) HasDatacenter() bool`

HasDatacenter returns a boolean if a field has been set.

### GetDomain

`func (o *DomainInspectorEntryDimensions) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *DomainInspectorEntryDimensions) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *DomainInspectorEntryDimensions) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *DomainInspectorEntryDimensions) HasDomain() bool`

HasDomain returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
