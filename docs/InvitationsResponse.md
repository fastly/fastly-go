# InvitationsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | Pointer to [**PaginationLinks**](PaginationLinks.md) |  | [optional] 
**Meta** | Pointer to [**PaginationMeta**](PaginationMeta.md) |  | [optional] 
**Data** | Pointer to [**[]InvitationResponseData**](InvitationResponseData.md) |  | [optional] 

## Methods

### NewInvitationsResponse

`func NewInvitationsResponse() *InvitationsResponse`

NewInvitationsResponse instantiates a new InvitationsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvitationsResponseWithDefaults

`func NewInvitationsResponseWithDefaults() *InvitationsResponse`

NewInvitationsResponseWithDefaults instantiates a new InvitationsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *InvitationsResponse) GetLinks() PaginationLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *InvitationsResponse) GetLinksOk() (*PaginationLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *InvitationsResponse) SetLinks(v PaginationLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *InvitationsResponse) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetMeta

`func (o *InvitationsResponse) GetMeta() PaginationMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *InvitationsResponse) GetMetaOk() (*PaginationMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *InvitationsResponse) SetMeta(v PaginationMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *InvitationsResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetData

`func (o *InvitationsResponse) GetData() []InvitationResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *InvitationsResponse) GetDataOk() (*[]InvitationResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *InvitationsResponse) SetData(v []InvitationResponseData)`

SetData sets Data field to given value.

### HasData

`func (o *InvitationsResponse) HasData() bool`

HasData returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
