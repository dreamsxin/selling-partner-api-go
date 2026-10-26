# Participation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsParticipating** | **bool** | If &#x60;true&#x60;, the seller participates in the marketplace. Otherwise &#x60;false&#x60;. | 
**HasSuspendedListings** | **bool** | Specifies if the seller has suspended listings. &#x60;true&#x60; if the seller Listing Status is set to Inactive, otherwise &#x60;false&#x60;. | 

## Methods

### NewParticipation

`func NewParticipation(isParticipating bool, hasSuspendedListings bool, ) *Participation`

NewParticipation instantiates a new Participation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewParticipationWithDefaults

`func NewParticipationWithDefaults() *Participation`

NewParticipationWithDefaults instantiates a new Participation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsParticipating

`func (o *Participation) GetIsParticipating() bool`

GetIsParticipating returns the IsParticipating field if non-nil, zero value otherwise.

### GetIsParticipatingOk

`func (o *Participation) GetIsParticipatingOk() (*bool, bool)`

GetIsParticipatingOk returns a tuple with the IsParticipating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsParticipating

`func (o *Participation) SetIsParticipating(v bool)`

SetIsParticipating sets IsParticipating field to given value.


### GetHasSuspendedListings

`func (o *Participation) GetHasSuspendedListings() bool`

GetHasSuspendedListings returns the HasSuspendedListings field if non-nil, zero value otherwise.

### GetHasSuspendedListingsOk

`func (o *Participation) GetHasSuspendedListingsOk() (*bool, bool)`

GetHasSuspendedListingsOk returns a tuple with the HasSuspendedListings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasSuspendedListings

`func (o *Participation) SetHasSuspendedListings(v bool)`

SetHasSuspendedListings sets HasSuspendedListings field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


