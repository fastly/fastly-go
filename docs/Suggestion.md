# Suggestion

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Domain** | Pointer to **string** | The suggested domain, consisting of a subdomain and zone. | [optional] 
**Subdomain** | Pointer to **string** | The subdomain of the suggested domain. | [optional] 
**Zone** | Pointer to **string** | The zone of the suggested domain. | [optional] 
**Path** | Pointer to **string** | If present, the path is to be appended to the domain to complete the suggestion. | [optional] 

## Methods

### NewSuggestion

`func NewSuggestion() *Suggestion`

NewSuggestion instantiates a new Suggestion object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSuggestionWithDefaults

`func NewSuggestionWithDefaults() *Suggestion`

NewSuggestionWithDefaults instantiates a new Suggestion object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDomain

`func (o *Suggestion) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *Suggestion) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *Suggestion) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *Suggestion) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetSubdomain

`func (o *Suggestion) GetSubdomain() string`

GetSubdomain returns the Subdomain field if non-nil, zero value otherwise.

### GetSubdomainOk

`func (o *Suggestion) GetSubdomainOk() (*string, bool)`

GetSubdomainOk returns a tuple with the Subdomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubdomain

`func (o *Suggestion) SetSubdomain(v string)`

SetSubdomain sets Subdomain field to given value.

### HasSubdomain

`func (o *Suggestion) HasSubdomain() bool`

HasSubdomain returns a boolean if a field has been set.

### GetZone

`func (o *Suggestion) GetZone() string`

GetZone returns the Zone field if non-nil, zero value otherwise.

### GetZoneOk

`func (o *Suggestion) GetZoneOk() (*string, bool)`

GetZoneOk returns a tuple with the Zone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZone

`func (o *Suggestion) SetZone(v string)`

SetZone sets Zone field to given value.

### HasZone

`func (o *Suggestion) HasZone() bool`

HasZone returns a boolean if a field has been set.

### GetPath

`func (o *Suggestion) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *Suggestion) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *Suggestion) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *Suggestion) HasPath() bool`

HasPath returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


