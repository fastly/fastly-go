# WafSimulateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Request** | **string** | The raw HTTP request in wire format to simulate through the WAF. Must include the request line, headers, and optionally a body, separated by CRLF sequences. | 
**Response** | Pointer to **string** | The raw HTTP response in wire format. The WAF engine inspects response headers during its PostRequest phase and may generate signals from them. When omitted, a default response of `HTTP/1.1 200 OK\\r\\n\\r\\n` is used. | [optional] 

## Methods

### NewWafSimulateRequest

`func NewWafSimulateRequest(request string, ) *WafSimulateRequest`

NewWafSimulateRequest instantiates a new WafSimulateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWafSimulateRequestWithDefaults

`func NewWafSimulateRequestWithDefaults() *WafSimulateRequest`

NewWafSimulateRequestWithDefaults instantiates a new WafSimulateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequest

`func (o *WafSimulateRequest) GetRequest() string`

GetRequest returns the Request field if non-nil, zero value otherwise.

### GetRequestOk

`func (o *WafSimulateRequest) GetRequestOk() (*string, bool)`

GetRequestOk returns a tuple with the Request field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequest

`func (o *WafSimulateRequest) SetRequest(v string)`

SetRequest sets Request field to given value.


### GetResponse

`func (o *WafSimulateRequest) GetResponse() string`

GetResponse returns the Response field if non-nil, zero value otherwise.

### GetResponseOk

`func (o *WafSimulateRequest) GetResponseOk() (*string, bool)`

GetResponseOk returns a tuple with the Response field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponse

`func (o *WafSimulateRequest) SetResponse(v string)`

SetResponse sets Response field to given value.

### HasResponse

`func (o *WafSimulateRequest) HasResponse() bool`

HasResponse returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


