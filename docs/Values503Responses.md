# Values503Responses

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | Pointer to **string** | The HTTP request path. | [optional] 
**RatePerUrl** | Pointer to **float32** | The rate at which the reason in this dimension occurs among responses to this URL with a 503 status code. | [optional] 
**Var503RatePerUrl** | Pointer to **float32** | The rate at which 503 status codes are returned for this URL. | [optional] 

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

### GetUrl

`func (o *Values503Responses) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *Values503Responses) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *Values503Responses) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *Values503Responses) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetRatePerUrl

`func (o *Values503Responses) GetRatePerUrl() float32`

GetRatePerUrl returns the RatePerUrl field if non-nil, zero value otherwise.

### GetRatePerUrlOk

`func (o *Values503Responses) GetRatePerUrlOk() (*float32, bool)`

GetRatePerUrlOk returns a tuple with the RatePerUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRatePerUrl

`func (o *Values503Responses) SetRatePerUrl(v float32)`

SetRatePerUrl sets RatePerUrl field to given value.

### HasRatePerUrl

`func (o *Values503Responses) HasRatePerUrl() bool`

HasRatePerUrl returns a boolean if a field has been set.

### GetVar503RatePerUrl

`func (o *Values503Responses) GetVar503RatePerUrl() float32`

GetVar503RatePerUrl returns the Var503RatePerUrl field if non-nil, zero value otherwise.

### GetVar503RatePerUrlOk

`func (o *Values503Responses) GetVar503RatePerUrlOk() (*float32, bool)`

GetVar503RatePerUrlOk returns a tuple with the Var503RatePerUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar503RatePerUrl

`func (o *Values503Responses) SetVar503RatePerUrl(v float32)`

SetVar503RatePerUrl sets Var503RatePerUrl field to given value.

### HasVar503RatePerUrl

`func (o *Values503Responses) HasVar503RatePerUrl() bool`

HasVar503RatePerUrl returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


