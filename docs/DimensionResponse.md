# DimensionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Response** | Pointer to **string** | The HTTP reason phrase for this dimension. | [optional] 

## Methods

### NewDimensionResponse

`func NewDimensionResponse() *DimensionResponse`

NewDimensionResponse instantiates a new DimensionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDimensionResponseWithDefaults

`func NewDimensionResponseWithDefaults() *DimensionResponse`

NewDimensionResponseWithDefaults instantiates a new DimensionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResponse

`func (o *DimensionResponse) GetResponse() string`

GetResponse returns the Response field if non-nil, zero value otherwise.

### GetResponseOk

`func (o *DimensionResponse) GetResponseOk() (*string, bool)`

GetResponseOk returns a tuple with the Response field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponse

`func (o *DimensionResponse) SetResponse(v string)`

SetResponse sets Response field to given value.

### HasResponse

`func (o *DimensionResponse) HasResponse() bool`

HasResponse returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
