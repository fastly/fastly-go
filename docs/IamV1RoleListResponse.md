# IamV1RoleListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Limit** | Pointer to **int32** | Maximum number of results returned. | [optional] 
**NextCursor** | Pointer to **string** | Cursor for the next page. | [optional] 
**Data** | Pointer to [**[]IamV1RoleResponse**](IamV1RoleResponse.md) | Page of IAM roles (length ≤ limit). | [optional] 

## Methods

### NewIamV1RoleListResponse

`func NewIamV1RoleListResponse() *IamV1RoleListResponse`

NewIamV1RoleListResponse instantiates a new IamV1RoleListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIamV1RoleListResponseWithDefaults

`func NewIamV1RoleListResponseWithDefaults() *IamV1RoleListResponse`

NewIamV1RoleListResponseWithDefaults instantiates a new IamV1RoleListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLimit

`func (o *IamV1RoleListResponse) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *IamV1RoleListResponse) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *IamV1RoleListResponse) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *IamV1RoleListResponse) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetNextCursor

`func (o *IamV1RoleListResponse) GetNextCursor() string`

GetNextCursor returns the NextCursor field if non-nil, zero value otherwise.

### GetNextCursorOk

`func (o *IamV1RoleListResponse) GetNextCursorOk() (*string, bool)`

GetNextCursorOk returns a tuple with the NextCursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextCursor

`func (o *IamV1RoleListResponse) SetNextCursor(v string)`

SetNextCursor sets NextCursor field to given value.

### HasNextCursor

`func (o *IamV1RoleListResponse) HasNextCursor() bool`

HasNextCursor returns a boolean if a field has been set.

### GetData

`func (o *IamV1RoleListResponse) GetData() []IamV1RoleResponse`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *IamV1RoleListResponse) GetDataOk() (*[]IamV1RoleResponse, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *IamV1RoleListResponse) SetData(v []IamV1RoleResponse)`

SetData sets Data field to given value.

### HasData

`func (o *IamV1RoleListResponse) HasData() bool`

HasData returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


