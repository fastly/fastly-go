# ValuesBrowser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BrowserVersion** | Pointer to **string** | The version of the client&#39;s browser. | [optional] 
**Rate** | Pointer to **float32** | The percentage of requests by this version of the browser specified in the dimension. | [optional] 

## Methods

### NewValuesBrowser

`func NewValuesBrowser() *ValuesBrowser`

NewValuesBrowser instantiates a new ValuesBrowser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewValuesBrowserWithDefaults

`func NewValuesBrowserWithDefaults() *ValuesBrowser`

NewValuesBrowserWithDefaults instantiates a new ValuesBrowser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBrowserVersion

`func (o *ValuesBrowser) GetBrowserVersion() string`

GetBrowserVersion returns the BrowserVersion field if non-nil, zero value otherwise.

### GetBrowserVersionOk

`func (o *ValuesBrowser) GetBrowserVersionOk() (*string, bool)`

GetBrowserVersionOk returns a tuple with the BrowserVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrowserVersion

`func (o *ValuesBrowser) SetBrowserVersion(v string)`

SetBrowserVersion sets BrowserVersion field to given value.

### HasBrowserVersion

`func (o *ValuesBrowser) HasBrowserVersion() bool`

HasBrowserVersion returns a boolean if a field has been set.

### GetRate

`func (o *ValuesBrowser) GetRate() float32`

GetRate returns the Rate field if non-nil, zero value otherwise.

### GetRateOk

`func (o *ValuesBrowser) GetRateOk() (*float32, bool)`

GetRateOk returns a tuple with the Rate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRate

`func (o *ValuesBrowser) SetRate(v float32)`

SetRate sets Rate field to given value.

### HasRate

`func (o *ValuesBrowser) HasRate() bool`

HasRate returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


