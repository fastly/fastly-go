# Server

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

## Methods

### NewServer

`func NewServer() *Server`

NewServer instantiates a new Server object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerWithDefaults

`func NewServerWithDefaults() *Server`

NewServerWithDefaults instantiates a new Server object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWeight

`func (o *Server) GetWeight() int32`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *Server) GetWeightOk() (*int32, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *Server) SetWeight(v int32)`

SetWeight sets Weight field to given value.

### HasWeight

`func (o *Server) HasWeight() bool`

HasWeight returns a boolean if a field has been set.

### GetMaxConn

`func (o *Server) GetMaxConn() int32`

GetMaxConn returns the MaxConn field if non-nil, zero value otherwise.

### GetMaxConnOk

`func (o *Server) GetMaxConnOk() (*int32, bool)`

GetMaxConnOk returns a tuple with the MaxConn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxConn

`func (o *Server) SetMaxConn(v int32)`

SetMaxConn sets MaxConn field to given value.

### HasMaxConn

`func (o *Server) HasMaxConn() bool`

HasMaxConn returns a boolean if a field has been set.

### GetPort

`func (o *Server) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *Server) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *Server) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *Server) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetAddress

`func (o *Server) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *Server) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *Server) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *Server) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetComment

`func (o *Server) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *Server) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *Server) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *Server) HasComment() bool`

HasComment returns a boolean if a field has been set.

### SetCommentNil

`func (o *Server) SetCommentNil(b bool)`

 SetCommentNil sets the value for Comment to be an explicit nil

### UnsetComment
`func (o *Server) UnsetComment()`

UnsetComment ensures that no value is present for Comment, not even an explicit nil
### GetDisabled

`func (o *Server) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *Server) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *Server) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.

### HasDisabled

`func (o *Server) HasDisabled() bool`

HasDisabled returns a boolean if a field has been set.

### GetOverrideHost

`func (o *Server) GetOverrideHost() string`

GetOverrideHost returns the OverrideHost field if non-nil, zero value otherwise.

### GetOverrideHostOk

`func (o *Server) GetOverrideHostOk() (*string, bool)`

GetOverrideHostOk returns a tuple with the OverrideHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverrideHost

`func (o *Server) SetOverrideHost(v string)`

SetOverrideHost sets OverrideHost field to given value.

### HasOverrideHost

`func (o *Server) HasOverrideHost() bool`

HasOverrideHost returns a boolean if a field has been set.

### SetOverrideHostNil

`func (o *Server) SetOverrideHostNil(b bool)`

 SetOverrideHostNil sets the value for OverrideHost to be an explicit nil

### UnsetOverrideHost
`func (o *Server) UnsetOverrideHost()`

UnsetOverrideHost ensures that no value is present for OverrideHost, not even an explicit nil

[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
