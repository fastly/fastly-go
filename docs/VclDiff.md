# VclDiff

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**From** | Pointer to **int32** | The version number of the service to which changes in the generated VCL are being compared. | [optional] 
**To** | Pointer to **int32** | The version number of the service from which changes in the generated VCL are being compared. | [optional] 
**Format** | Pointer to **string** | The format in which compared VCL changes are being returned in. | [optional] 
**Diff** | Pointer to **string** | The differences between two specified versions. | [optional] 

## Methods

### NewVclDiff

`func NewVclDiff() *VclDiff`

NewVclDiff instantiates a new VclDiff object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVclDiffWithDefaults

`func NewVclDiffWithDefaults() *VclDiff`

NewVclDiffWithDefaults instantiates a new VclDiff object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFrom

`func (o *VclDiff) GetFrom() int32`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *VclDiff) GetFromOk() (*int32, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *VclDiff) SetFrom(v int32)`

SetFrom sets From field to given value.

### HasFrom

`func (o *VclDiff) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetTo

`func (o *VclDiff) GetTo() int32`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *VclDiff) GetToOk() (*int32, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *VclDiff) SetTo(v int32)`

SetTo sets To field to given value.

### HasTo

`func (o *VclDiff) HasTo() bool`

HasTo returns a boolean if a field has been set.

### GetFormat

`func (o *VclDiff) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *VclDiff) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *VclDiff) SetFormat(v string)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *VclDiff) HasFormat() bool`

HasFormat returns a boolean if a field has been set.

### GetDiff

`func (o *VclDiff) GetDiff() string`

GetDiff returns the Diff field if non-nil, zero value otherwise.

### GetDiffOk

`func (o *VclDiff) GetDiffOk() (*string, bool)`

GetDiffOk returns a tuple with the Diff field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiff

`func (o *VclDiff) SetDiff(v string)`

SetDiff sets Diff field to given value.

### HasDiff

`func (o *VclDiff) HasDiff() bool`

HasDiff returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


