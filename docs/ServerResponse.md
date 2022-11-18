# ServerResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Weight** | Pointer to **int32** | Weight (`1-100`) used to load balance this server against others. | [optional] [default to 100]
**MaxConn** | Pointer to **int32** | Maximum number of connections. If the value is `0`, it inherits the value from pool&#39;s `max_conn_default`. | [optional] [default to 0]
**Port** | Pointer to **int32** | Port number. Setting port `443` does not force TLS. Set `use_tls` in pool to force TLS. | [optional] [default to 80]
**Address** | Pointer to **string** | A hostname, IPv4, or IPv6 address for the server. Required. | [optional] 
**Comment** | Pointer to **NullableString** | A freeform descriptive note. | [optional] 
**Disabled** | Pointer to **bool** | Allows servers to be enabled and disabled in a pool. | [optional] [default to false]
**OverrideHost** | Pointer to **NullableString** | The hostname to override the Host header. Defaults to `null` meaning no override of the Host header if not set. This setting can also be added to a Pool definition. However, the server setting will override the Pool setting. | [optional] [default to "null"]
**CreatedAt** | Pointer to **NullableTime** | Date and time in ISO 8601 format. | [optional] [readonly] 
**DeletedAt** | Pointer to **NullableTime** | Date and time in ISO 8601 format. | [optional] [readonly] 
**UpdatedAt** | Pointer to **NullableTime** | Date and time in ISO 8601 format. | [optional] [readonly] 
**ServiceID** | Pointer to **string** |  | [optional] [readonly] 
**ID** | Pointer to **string** |  | [optional] [readonly] 
**PoolID** | Pointer to **string** |  | [optional] [readonly] 

## Methods

### NewServerResponse

`func NewServerResponse() *ServerResponse`

NewServerResponse instantiates a new ServerResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerResponseWithDefaults

`func NewServerResponseWithDefaults() *ServerResponse`

NewServerResponseWithDefaults instantiates a new ServerResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWeight

`func (o *ServerResponse) GetWeight() int32`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *ServerResponse) GetWeightOk() (*int32, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *ServerResponse) SetWeight(v int32)`

SetWeight sets Weight field to given value.

### HasWeight

`func (o *ServerResponse) HasWeight() bool`

HasWeight returns a boolean if a field has been set.

### GetMaxConn

`func (o *ServerResponse) GetMaxConn() int32`

GetMaxConn returns the MaxConn field if non-nil, zero value otherwise.

### GetMaxConnOk

`func (o *ServerResponse) GetMaxConnOk() (*int32, bool)`

GetMaxConnOk returns a tuple with the MaxConn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxConn

`func (o *ServerResponse) SetMaxConn(v int32)`

SetMaxConn sets MaxConn field to given value.

### HasMaxConn

`func (o *ServerResponse) HasMaxConn() bool`

HasMaxConn returns a boolean if a field has been set.

### GetPort

`func (o *ServerResponse) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *ServerResponse) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *ServerResponse) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *ServerResponse) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetAddress

`func (o *ServerResponse) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *ServerResponse) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *ServerResponse) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *ServerResponse) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetComment

`func (o *ServerResponse) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *ServerResponse) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *ServerResponse) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *ServerResponse) HasComment() bool`

HasComment returns a boolean if a field has been set.

### SetCommentNil

`func (o *ServerResponse) SetCommentNil(b bool)`

 SetCommentNil sets the value for Comment to be an explicit nil

### UnsetComment
`func (o *ServerResponse) UnsetComment()`

UnsetComment ensures that no value is present for Comment, not even an explicit nil
### GetDisabled

`func (o *ServerResponse) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *ServerResponse) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *ServerResponse) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.

### HasDisabled

`func (o *ServerResponse) HasDisabled() bool`

HasDisabled returns a boolean if a field has been set.

### GetOverrideHost

`func (o *ServerResponse) GetOverrideHost() string`

GetOverrideHost returns the OverrideHost field if non-nil, zero value otherwise.

### GetOverrideHostOk

`func (o *ServerResponse) GetOverrideHostOk() (*string, bool)`

GetOverrideHostOk returns a tuple with the OverrideHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverrideHost

`func (o *ServerResponse) SetOverrideHost(v string)`

SetOverrideHost sets OverrideHost field to given value.

### HasOverrideHost

`func (o *ServerResponse) HasOverrideHost() bool`

HasOverrideHost returns a boolean if a field has been set.

### SetOverrideHostNil

`func (o *ServerResponse) SetOverrideHostNil(b bool)`

 SetOverrideHostNil sets the value for OverrideHost to be an explicit nil

### UnsetOverrideHost
`func (o *ServerResponse) UnsetOverrideHost()`

UnsetOverrideHost ensures that no value is present for OverrideHost, not even an explicit nil
### GetCreatedAt

`func (o *ServerResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ServerResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ServerResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ServerResponse) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *ServerResponse) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *ServerResponse) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetDeletedAt

`func (o *ServerResponse) GetDeletedAt() time.Time`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *ServerResponse) GetDeletedAtOk() (*time.Time, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *ServerResponse) SetDeletedAt(v time.Time)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *ServerResponse) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### SetDeletedAtNil

`func (o *ServerResponse) SetDeletedAtNil(b bool)`

 SetDeletedAtNil sets the value for DeletedAt to be an explicit nil

### UnsetDeletedAt
`func (o *ServerResponse) UnsetDeletedAt()`

UnsetDeletedAt ensures that no value is present for DeletedAt, not even an explicit nil
### GetUpdatedAt

`func (o *ServerResponse) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ServerResponse) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ServerResponse) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ServerResponse) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *ServerResponse) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *ServerResponse) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
### GetServiceID

`func (o *ServerResponse) GetServiceID() string`

GetServiceID returns the ServiceID field if non-nil, zero value otherwise.

### GetServiceIDOk

`func (o *ServerResponse) GetServiceIDOk() (*string, bool)`

GetServiceIDOk returns a tuple with the ServiceID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceID

`func (o *ServerResponse) SetServiceID(v string)`

SetServiceID sets ServiceID field to given value.

### HasServiceID

`func (o *ServerResponse) HasServiceID() bool`

HasServiceID returns a boolean if a field has been set.

### GetID

`func (o *ServerResponse) GetID() string`

GetID returns the ID field if non-nil, zero value otherwise.

### GetIDOk

`func (o *ServerResponse) GetIDOk() (*string, bool)`

GetIDOk returns a tuple with the ID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetID

`func (o *ServerResponse) SetID(v string)`

SetID sets ID field to given value.

### HasID

`func (o *ServerResponse) HasID() bool`

HasID returns a boolean if a field has been set.

### GetPoolID

`func (o *ServerResponse) GetPoolID() string`

GetPoolID returns the PoolID field if non-nil, zero value otherwise.

### GetPoolIDOk

`func (o *ServerResponse) GetPoolIDOk() (*string, bool)`

GetPoolIDOk returns a tuple with the PoolID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolID

`func (o *ServerResponse) SetPoolID(v string)`

SetPoolID sets PoolID field to given value.

### HasPoolID

`func (o *ServerResponse) HasPoolID() bool`

HasPoolID returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
