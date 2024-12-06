# Values503Responses

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**URL** | Pointer to **string** | The HTTP request path. | [optional] 
**RatePerURL** | Pointer to **float32** | The rate at which the reason in this dimension occurs among responses to this URL with a 503 status code. | [optional] 
**Var503RatePerURL** | Pointer to **float32** | The rate at which 503 status codes are returned for this URL. | [optional] 

## Methods

### NewValues503Responses

`func NewValues503Responses() *Values503Responses`

NewValues503Responses instantiates a new Values503Responses object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewValues503ResponsesWithDefaults

`func NewValues503ResponsesWithDefaults() *Values503Responses`

NewValues503ResponsesWithDefaults instantiates a new Values503Responses object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetURL

`func (o *Values503Responses) GetURL() string`

GetURL returns the URL field if non-nil, zero value otherwise.

### GetURLOk

`func (o *Values503Responses) GetURLOk() (*string, bool)`

GetURLOk returns a tuple with the URL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetURL

`func (o *Values503Responses) SetURL(v string)`

SetURL sets URL field to given value.

### HasURL

`func (o *Values503Responses) HasURL() bool`

HasURL returns a boolean if a field has been set.

### GetRatePerURL

`func (o *Values503Responses) GetRatePerURL() float32`

GetRatePerURL returns the RatePerURL field if non-nil, zero value otherwise.

### GetRatePerURLOk

`func (o *Values503Responses) GetRatePerURLOk() (*float32, bool)`

GetRatePerURLOk returns a tuple with the RatePerURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRatePerURL

`func (o *Values503Responses) SetRatePerURL(v float32)`

SetRatePerURL sets RatePerURL field to given value.

### HasRatePerURL

`func (o *Values503Responses) HasRatePerURL() bool`

HasRatePerURL returns a boolean if a field has been set.

### GetVar503RatePerURL

`func (o *Values503Responses) GetVar503RatePerURL() float32`

GetVar503RatePerURL returns the Var503RatePerURL field if non-nil, zero value otherwise.

### GetVar503RatePerURLOk

`func (o *Values503Responses) GetVar503RatePerURLOk() (*float32, bool)`

GetVar503RatePerURLOk returns a tuple with the Var503RatePerURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar503RatePerURL

`func (o *Values503Responses) SetVar503RatePerURL(v float32)`

SetVar503RatePerURL sets Var503RatePerURL field to given value.

### HasVar503RatePerURL

`func (o *Values503Responses) HasVar503RatePerURL() bool`

HasVar503RatePerURL returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
