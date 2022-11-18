# HTTPStreamFormat

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Content** | Pointer to **string** | A fragment of body data as a string. | [optional] 
**ContentBin** | Pointer to **string** | A fragment of body data as a base64-encoded binary blob. | [optional] 

## Methods

### NewHTTPStreamFormat

`func NewHTTPStreamFormat() *HTTPStreamFormat`

NewHTTPStreamFormat instantiates a new HTTPStreamFormat object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHTTPStreamFormatWithDefaults

`func NewHTTPStreamFormatWithDefaults() *HTTPStreamFormat`

NewHTTPStreamFormatWithDefaults instantiates a new HTTPStreamFormat object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContent

`func (o *HTTPStreamFormat) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *HTTPStreamFormat) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *HTTPStreamFormat) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *HTTPStreamFormat) HasContent() bool`

HasContent returns a boolean if a field has been set.

### GetContentBin

`func (o *HTTPStreamFormat) GetContentBin() string`

GetContentBin returns the ContentBin field if non-nil, zero value otherwise.

### GetContentBinOk

`func (o *HTTPStreamFormat) GetContentBinOk() (*string, bool)`

GetContentBinOk returns a tuple with the ContentBin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentBin

`func (o *HTTPStreamFormat) SetContentBin(v string)`

SetContentBin sets ContentBin field to given value.

### HasContentBin

`func (o *HTTPStreamFormat) HasContentBin() bool`

HasContentBin returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
