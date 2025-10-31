# Star

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**StarData**](StarData.md) |  | [optional] 

## Methods

### NewStar

`func NewStar() *Star`

NewStar instantiates a new Star object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStarWithDefaults

`func NewStarWithDefaults() *Star`

NewStarWithDefaults instantiates a new Star object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *Star) GetData() StarData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *Star) GetDataOk() (*StarData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *Star) SetData(v StarData)`

SetData sets Data field to given value.

### HasData

`func (o *Star) HasData() bool`

HasData returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


