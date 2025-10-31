# DiffResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**From** | Pointer to **int32** | The version number being diffed from. | [optional] 
**To** | Pointer to **int32** | The version number being diffed to. | [optional] 
**Format** | Pointer to **string** | The format the diff is being returned in (`text`, `html` or `html_simple`). | [optional] 
**Diff** | Pointer to **string** | The differences between two specified service versions. Returns the full config if the version configurations are identical. | [optional] 

## Methods

### NewDiffResponse

`func NewDiffResponse() *DiffResponse`

NewDiffResponse instantiates a new DiffResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDiffResponseWithDefaults

`func NewDiffResponseWithDefaults() *DiffResponse`

NewDiffResponseWithDefaults instantiates a new DiffResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFrom

`func (o *DiffResponse) GetFrom() int32`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *DiffResponse) GetFromOk() (*int32, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *DiffResponse) SetFrom(v int32)`

SetFrom sets From field to given value.

### HasFrom

`func (o *DiffResponse) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetTo

`func (o *DiffResponse) GetTo() int32`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *DiffResponse) GetToOk() (*int32, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *DiffResponse) SetTo(v int32)`

SetTo sets To field to given value.

### HasTo

`func (o *DiffResponse) HasTo() bool`

HasTo returns a boolean if a field has been set.

### GetFormat

`func (o *DiffResponse) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *DiffResponse) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *DiffResponse) SetFormat(v string)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *DiffResponse) HasFormat() bool`

HasFormat returns a boolean if a field has been set.

### GetDiff

`func (o *DiffResponse) GetDiff() string`

GetDiff returns the Diff field if non-nil, zero value otherwise.

### GetDiffOk

`func (o *DiffResponse) GetDiffOk() (*string, bool)`

GetDiffOk returns a tuple with the Diff field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiff

`func (o *DiffResponse) SetDiff(v string)`

SetDiff sets Diff field to given value.

### HasDiff

`func (o *DiffResponse) HasDiff() bool`

HasDiff returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


