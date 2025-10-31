# EventAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Admin** | Pointer to **bool** | Indicates if event was performed by Fastly. | [optional] 
**CreatedAt** | Pointer to **NullableTime** | Date and time in ISO 8601 format. | [optional] [readonly] 
**CustomerId** | Pointer to **string** |  | [optional] [readonly] 
**Description** | Pointer to **string** | Description of the event. | [optional] 
**EventType** | Pointer to **string** | Type of event. Can be used with `filter[event_type]` | [optional] 
**Ip** | Pointer to **string** | IP addresses that the event was requested from. | [optional] 
**Metadata** | Pointer to **map[string]map[string]interface{}** | Hash of key value pairs of additional information. | [optional] 
**ServiceId** | Pointer to **string** |  | [optional] [readonly] 
**UserId** | Pointer to **string** |  | [optional] [readonly] 
**TokenId** | Pointer to **string** |  | [optional] [readonly] 

## Methods

### NewEventAttributes

`func NewEventAttributes() *EventAttributes`

NewEventAttributes instantiates a new EventAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventAttributesWithDefaults

`func NewEventAttributesWithDefaults() *EventAttributes`

NewEventAttributesWithDefaults instantiates a new EventAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdmin

`func (o *EventAttributes) GetAdmin() bool`

GetAdmin returns the Admin field if non-nil, zero value otherwise.

### GetAdminOk

`func (o *EventAttributes) GetAdminOk() (*bool, bool)`

GetAdminOk returns a tuple with the Admin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdmin

`func (o *EventAttributes) SetAdmin(v bool)`

SetAdmin sets Admin field to given value.

### HasAdmin

`func (o *EventAttributes) HasAdmin() bool`

HasAdmin returns a boolean if a field has been set.

### GetCreatedAt

`func (o *EventAttributes) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *EventAttributes) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *EventAttributes) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *EventAttributes) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *EventAttributes) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *EventAttributes) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetCustomerId

`func (o *EventAttributes) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *EventAttributes) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *EventAttributes) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.

### HasCustomerId

`func (o *EventAttributes) HasCustomerId() bool`

HasCustomerId returns a boolean if a field has been set.

### GetDescription

`func (o *EventAttributes) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *EventAttributes) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *EventAttributes) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *EventAttributes) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEventType

`func (o *EventAttributes) GetEventType() string`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *EventAttributes) GetEventTypeOk() (*string, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *EventAttributes) SetEventType(v string)`

SetEventType sets EventType field to given value.

### HasEventType

`func (o *EventAttributes) HasEventType() bool`

HasEventType returns a boolean if a field has been set.

### GetIp

`func (o *EventAttributes) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *EventAttributes) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *EventAttributes) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *EventAttributes) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetMetadata

`func (o *EventAttributes) GetMetadata() map[string]map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *EventAttributes) GetMetadataOk() (*map[string]map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *EventAttributes) SetMetadata(v map[string]map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *EventAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetServiceId

`func (o *EventAttributes) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *EventAttributes) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *EventAttributes) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.

### HasServiceId

`func (o *EventAttributes) HasServiceId() bool`

HasServiceId returns a boolean if a field has been set.

### GetUserId

`func (o *EventAttributes) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *EventAttributes) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *EventAttributes) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *EventAttributes) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetTokenId

`func (o *EventAttributes) GetTokenId() string`

GetTokenId returns the TokenId field if non-nil, zero value otherwise.

### GetTokenIdOk

`func (o *EventAttributes) GetTokenIdOk() (*string, bool)`

GetTokenIdOk returns a tuple with the TokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenId

`func (o *EventAttributes) SetTokenId(v string)`

SetTokenId sets TokenId field to given value.

### HasTokenId

`func (o *EventAttributes) HasTokenId() bool`

HasTokenId returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


