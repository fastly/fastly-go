# AsyncResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 

## Methods

### NewAsyncResponse

`func NewAsyncResponse() *AsyncResponse`

NewAsyncResponse instantiates a new AsyncResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAsyncResponseWithDefaults

`func NewAsyncResponseWithDefaults() *AsyncResponse`

NewAsyncResponseWithDefaults instantiates a new AsyncResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *AsyncResponse) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *AsyncResponse) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *AsyncResponse) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *AsyncResponse) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetStatus

`func (o *AsyncResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AsyncResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AsyncResponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AsyncResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


