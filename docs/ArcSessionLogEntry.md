# ArcSessionLogEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Request** | Pointer to **string** | The raw request data. | [optional] 
**Response** | Pointer to **string** | The raw response data. | [optional] 
**Attrs** | Pointer to **map[string]map[string]any** | Additional attributes recorded for this request/response pair. | [optional] 

## Methods

### NewArcSessionLogEntry

`func NewArcSessionLogEntry() *ArcSessionLogEntry`

NewArcSessionLogEntry instantiates a new ArcSessionLogEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArcSessionLogEntryWithDefaults

`func NewArcSessionLogEntryWithDefaults() *ArcSessionLogEntry`

NewArcSessionLogEntryWithDefaults instantiates a new ArcSessionLogEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequest

`func (o *ArcSessionLogEntry) GetRequest() string`

GetRequest returns the Request field if non-nil, zero value otherwise.

### GetRequestOk

`func (o *ArcSessionLogEntry) GetRequestOk() (*string, bool)`

GetRequestOk returns a tuple with the Request field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequest

`func (o *ArcSessionLogEntry) SetRequest(v string)`

SetRequest sets Request field to given value.

### HasRequest

`func (o *ArcSessionLogEntry) HasRequest() bool`

HasRequest returns a boolean if a field has been set.

### GetResponse

`func (o *ArcSessionLogEntry) GetResponse() string`

GetResponse returns the Response field if non-nil, zero value otherwise.

### GetResponseOk

`func (o *ArcSessionLogEntry) GetResponseOk() (*string, bool)`

GetResponseOk returns a tuple with the Response field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponse

`func (o *ArcSessionLogEntry) SetResponse(v string)`

SetResponse sets Response field to given value.

### HasResponse

`func (o *ArcSessionLogEntry) HasResponse() bool`

HasResponse returns a boolean if a field has been set.

### GetAttrs

`func (o *ArcSessionLogEntry) GetAttrs() map[string]map[string]any`

GetAttrs returns the Attrs field if non-nil, zero value otherwise.

### GetAttrsOk

`func (o *ArcSessionLogEntry) GetAttrsOk() (*map[string]map[string]any, bool)`

GetAttrsOk returns a tuple with the Attrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttrs

`func (o *ArcSessionLogEntry) SetAttrs(v map[string]map[string]any)`

SetAttrs sets Attrs field to given value.

### HasAttrs

`func (o *ArcSessionLogEntry) HasAttrs() bool`

HasAttrs returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
