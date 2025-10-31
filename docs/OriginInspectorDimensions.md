# OriginInspectorDimensions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Region** | Pointer to **string** | The geographic region from which the edge responses in this data entry were delivered. If unspecified, results are aggregated across regions. | [optional] 
**Datacenter** | Pointer to **string** | The POP from which the edge responses in this data entry were delivered. If unspecified, results are aggregated across POPs. | [optional] 
**Host** | Pointer to **string** | The origin host from which the edge responses in this data entry were delivered. If unspecified, results are aggregated across origin hosts. | [optional] 

## Methods

### NewOriginInspectorDimensions

`func NewOriginInspectorDimensions() *OriginInspectorDimensions`

NewOriginInspectorDimensions instantiates a new OriginInspectorDimensions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOriginInspectorDimensionsWithDefaults

`func NewOriginInspectorDimensionsWithDefaults() *OriginInspectorDimensions`

NewOriginInspectorDimensionsWithDefaults instantiates a new OriginInspectorDimensions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRegion

`func (o *OriginInspectorDimensions) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *OriginInspectorDimensions) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *OriginInspectorDimensions) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *OriginInspectorDimensions) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetDatacenter

`func (o *OriginInspectorDimensions) GetDatacenter() string`

GetDatacenter returns the Datacenter field if non-nil, zero value otherwise.

### GetDatacenterOk

`func (o *OriginInspectorDimensions) GetDatacenterOk() (*string, bool)`

GetDatacenterOk returns a tuple with the Datacenter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatacenter

`func (o *OriginInspectorDimensions) SetDatacenter(v string)`

SetDatacenter sets Datacenter field to given value.

### HasDatacenter

`func (o *OriginInspectorDimensions) HasDatacenter() bool`

HasDatacenter returns a boolean if a field has been set.

### GetHost

`func (o *OriginInspectorDimensions) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *OriginInspectorDimensions) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *OriginInspectorDimensions) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *OriginInspectorDimensions) HasHost() bool`

HasHost returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


