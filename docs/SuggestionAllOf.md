# SuggestionAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Domain** | Pointer to **string** | The suggested domain, consisting of a subdomain and zone. | [optional] 
**Subdomain** | Pointer to **string** | The subdomain of the suggested domain. | [optional] 
**Zone** | Pointer to **string** | The zone of the suggested domain. | [optional] 
**Path** | Pointer to **string** | If present, the path is to be appended to the domain to complete the suggestion. | [optional] 

## Methods

### NewSuggestionAllOf

`func NewSuggestionAllOf() *SuggestionAllOf`

NewSuggestionAllOf instantiates a new SuggestionAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSuggestionAllOfWithDefaults

`func NewSuggestionAllOfWithDefaults() *SuggestionAllOf`

NewSuggestionAllOfWithDefaults instantiates a new SuggestionAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDomain

`func (o *SuggestionAllOf) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *SuggestionAllOf) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *SuggestionAllOf) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *SuggestionAllOf) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetSubdomain

`func (o *SuggestionAllOf) GetSubdomain() string`

GetSubdomain returns the Subdomain field if non-nil, zero value otherwise.

### GetSubdomainOk

`func (o *SuggestionAllOf) GetSubdomainOk() (*string, bool)`

GetSubdomainOk returns a tuple with the Subdomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubdomain

`func (o *SuggestionAllOf) SetSubdomain(v string)`

SetSubdomain sets Subdomain field to given value.

### HasSubdomain

`func (o *SuggestionAllOf) HasSubdomain() bool`

HasSubdomain returns a boolean if a field has been set.

### GetZone

`func (o *SuggestionAllOf) GetZone() string`

GetZone returns the Zone field if non-nil, zero value otherwise.

### GetZoneOk

`func (o *SuggestionAllOf) GetZoneOk() (*string, bool)`

GetZoneOk returns a tuple with the Zone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZone

`func (o *SuggestionAllOf) SetZone(v string)`

SetZone sets Zone field to given value.

### HasZone

`func (o *SuggestionAllOf) HasZone() bool`

HasZone returns a boolean if a field has been set.

### GetPath

`func (o *SuggestionAllOf) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *SuggestionAllOf) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *SuggestionAllOf) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *SuggestionAllOf) HasPath() bool`

HasPath returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


