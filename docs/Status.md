# Status

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Scope** | Pointer to **string** | The scope provided in the status request. | [optional] 
**Domain** | Pointer to **string** | The domain provided in the status request. | [optional] 
**Zone** | Pointer to **string** | The zone of the domain provided of the status request. | [optional] 
**Status** | Pointer to **string** | A space-delimited string of the varying statuses associated with the domain provided. | [optional] 
**Tags** | Pointer to **string** | A space-delimited string of the varying tags associated with the domain provided. | [optional] 
**Offers** | Pointer to [**[]Offer**](Offer.md) |  | [optional] 

## Methods

### NewStatus

`func NewStatus() *Status`

NewStatus instantiates a new Status object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStatusWithDefaults

`func NewStatusWithDefaults() *Status`

NewStatusWithDefaults instantiates a new Status object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetScope

`func (o *Status) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *Status) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *Status) SetScope(v string)`

SetScope sets Scope field to given value.

### HasScope

`func (o *Status) HasScope() bool`

HasScope returns a boolean if a field has been set.

### GetDomain

`func (o *Status) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *Status) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *Status) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *Status) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetZone

`func (o *Status) GetZone() string`

GetZone returns the Zone field if non-nil, zero value otherwise.

### GetZoneOk

`func (o *Status) GetZoneOk() (*string, bool)`

GetZoneOk returns a tuple with the Zone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZone

`func (o *Status) SetZone(v string)`

SetZone sets Zone field to given value.

### HasZone

`func (o *Status) HasZone() bool`

HasZone returns a boolean if a field has been set.

### GetStatus

`func (o *Status) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Status) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Status) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Status) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTags

`func (o *Status) GetTags() string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *Status) GetTagsOk() (*string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *Status) SetTags(v string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *Status) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetOffers

`func (o *Status) GetOffers() []Offer`

GetOffers returns the Offers field if non-nil, zero value otherwise.

### GetOffersOk

`func (o *Status) GetOffersOk() (*[]Offer, bool)`

GetOffersOk returns a tuple with the Offers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffers

`func (o *Status) SetOffers(v []Offer)`

SetOffers sets Offers field to given value.

### HasOffers

`func (o *Status) HasOffers() bool`

HasOffers returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


