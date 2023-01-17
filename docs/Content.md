# Content

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Hash** | Pointer to **string** |  | [optional] 
**Request** | Pointer to **map[string]any** |  | [optional] 
**Response** | Pointer to **map[string]any** |  | [optional] 
**ResponseTime** | Pointer to **float32** |  | [optional] 
**Server** | Pointer to **string** |  | [optional] 
**Pop** | Pointer to **string** |  | [optional] 

## Methods

### NewContent

`func NewContent() *Content`

NewContent instantiates a new Content object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContentWithDefaults

`func NewContentWithDefaults() *Content`

NewContentWithDefaults instantiates a new Content object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHash

`func (o *Content) GetHash() string`

GetHash returns the Hash field if non-nil, zero value otherwise.

### GetHashOk

`func (o *Content) GetHashOk() (*string, bool)`

GetHashOk returns a tuple with the Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHash

`func (o *Content) SetHash(v string)`

SetHash sets Hash field to given value.

### HasHash

`func (o *Content) HasHash() bool`

HasHash returns a boolean if a field has been set.

### GetRequest

`func (o *Content) GetRequest() map[string]any`

GetRequest returns the Request field if non-nil, zero value otherwise.

### GetRequestOk

`func (o *Content) GetRequestOk() (*map[string]any, bool)`

GetRequestOk returns a tuple with the Request field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequest

`func (o *Content) SetRequest(v map[string]any)`

SetRequest sets Request field to given value.

### HasRequest

`func (o *Content) HasRequest() bool`

HasRequest returns a boolean if a field has been set.

### GetResponse

`func (o *Content) GetResponse() map[string]any`

GetResponse returns the Response field if non-nil, zero value otherwise.

### GetResponseOk

`func (o *Content) GetResponseOk() (*map[string]any, bool)`

GetResponseOk returns a tuple with the Response field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponse

`func (o *Content) SetResponse(v map[string]any)`

SetResponse sets Response field to given value.

### HasResponse

`func (o *Content) HasResponse() bool`

HasResponse returns a boolean if a field has been set.

### GetResponseTime

`func (o *Content) GetResponseTime() float32`

GetResponseTime returns the ResponseTime field if non-nil, zero value otherwise.

### GetResponseTimeOk

`func (o *Content) GetResponseTimeOk() (*float32, bool)`

GetResponseTimeOk returns a tuple with the ResponseTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseTime

`func (o *Content) SetResponseTime(v float32)`

SetResponseTime sets ResponseTime field to given value.

### HasResponseTime

`func (o *Content) HasResponseTime() bool`

HasResponseTime returns a boolean if a field has been set.

### GetServer

`func (o *Content) GetServer() string`

GetServer returns the Server field if non-nil, zero value otherwise.

### GetServerOk

`func (o *Content) GetServerOk() (*string, bool)`

GetServerOk returns a tuple with the Server field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer

`func (o *Content) SetServer(v string)`

SetServer sets Server field to given value.

### HasServer

`func (o *Content) HasServer() bool`

HasServer returns a boolean if a field has been set.

### GetPop

`func (o *Content) GetPop() string`

GetPop returns the Pop field if non-nil, zero value otherwise.

### GetPopOk

`func (o *Content) GetPopOk() (*string, bool)`

GetPopOk returns a tuple with the Pop field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPop

`func (o *Content) SetPop(v string)`

SetPop sets Pop field to given value.

### HasPop

`func (o *Content) HasPop() bool`

HasPop returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
