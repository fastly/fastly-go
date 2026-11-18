# HTTPResponseFormat

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **int32** | The HTTP status code. | [optional] [default to 200]
**Reason** | Pointer to **string** | The HTTP status string. Defaults to a string appropriate for `code`. | [optional] 
**Headers** | Pointer to **map[string]string** | A map of arbitrary HTTP response headers to include in the response. | [optional] 
**Body** | Pointer to **string** | The response body as a string. | [optional] 
**BodyBin** | Pointer to **string** | The response body as a base64-encoded binary blob. | [optional] 

## Methods

### NewHTTPResponseFormat

`func NewHTTPResponseFormat() *HTTPResponseFormat`

NewHTTPResponseFormat instantiates a new HTTPResponseFormat object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHTTPResponseFormatWithDefaults

`func NewHTTPResponseFormatWithDefaults() *HTTPResponseFormat`

NewHTTPResponseFormatWithDefaults instantiates a new HTTPResponseFormat object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *HTTPResponseFormat) GetCode() int32`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *HTTPResponseFormat) GetCodeOk() (*int32, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *HTTPResponseFormat) SetCode(v int32)`

SetCode sets Code field to given value.

### HasCode

`func (o *HTTPResponseFormat) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetReason

`func (o *HTTPResponseFormat) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *HTTPResponseFormat) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *HTTPResponseFormat) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *HTTPResponseFormat) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetHeaders

`func (o *HTTPResponseFormat) GetHeaders() map[string]string`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *HTTPResponseFormat) GetHeadersOk() (*map[string]string, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *HTTPResponseFormat) SetHeaders(v map[string]string)`

SetHeaders sets Headers field to given value.

### HasHeaders

`func (o *HTTPResponseFormat) HasHeaders() bool`

HasHeaders returns a boolean if a field has been set.

### GetBody

`func (o *HTTPResponseFormat) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *HTTPResponseFormat) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *HTTPResponseFormat) SetBody(v string)`

SetBody sets Body field to given value.

### HasBody

`func (o *HTTPResponseFormat) HasBody() bool`

HasBody returns a boolean if a field has been set.

### GetBodyBin

`func (o *HTTPResponseFormat) GetBodyBin() string`

GetBodyBin returns the BodyBin field if non-nil, zero value otherwise.

### GetBodyBinOk

`func (o *HTTPResponseFormat) GetBodyBinOk() (*string, bool)`

GetBodyBinOk returns a tuple with the BodyBin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBodyBin

`func (o *HTTPResponseFormat) SetBodyBin(v string)`

SetBodyBin sets BodyBin field to given value.

### HasBodyBin

`func (o *HTTPResponseFormat) HasBodyBin() bool`

HasBodyBin returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
