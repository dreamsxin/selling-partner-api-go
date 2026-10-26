# MarketplaceParticipation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Marketplace** | [**Marketplace**](Marketplace.md) |  | 
**Participation** | [**Participation**](Participation.md) |  | 
**StoreName** | **string** | The name of the seller&#39;s store as displayed in the marketplace. | 

## Methods

### NewMarketplaceParticipation

`func NewMarketplaceParticipation(marketplace Marketplace, participation Participation, storeName string, ) *MarketplaceParticipation`

NewMarketplaceParticipation instantiates a new MarketplaceParticipation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMarketplaceParticipationWithDefaults

`func NewMarketplaceParticipationWithDefaults() *MarketplaceParticipation`

NewMarketplaceParticipationWithDefaults instantiates a new MarketplaceParticipation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMarketplace

`func (o *MarketplaceParticipation) GetMarketplace() Marketplace`

GetMarketplace returns the Marketplace field if non-nil, zero value otherwise.

### GetMarketplaceOk

`func (o *MarketplaceParticipation) GetMarketplaceOk() (*Marketplace, bool)`

GetMarketplaceOk returns a tuple with the Marketplace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarketplace

`func (o *MarketplaceParticipation) SetMarketplace(v Marketplace)`

SetMarketplace sets Marketplace field to given value.


### GetParticipation

`func (o *MarketplaceParticipation) GetParticipation() Participation`

GetParticipation returns the Participation field if non-nil, zero value otherwise.

### GetParticipationOk

`func (o *MarketplaceParticipation) GetParticipationOk() (*Participation, bool)`

GetParticipationOk returns a tuple with the Participation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParticipation

`func (o *MarketplaceParticipation) SetParticipation(v Participation)`

SetParticipation sets Participation field to given value.


### GetStoreName

`func (o *MarketplaceParticipation) GetStoreName() string`

GetStoreName returns the StoreName field if non-nil, zero value otherwise.

### GetStoreNameOk

`func (o *MarketplaceParticipation) GetStoreNameOk() (*string, bool)`

GetStoreNameOk returns a tuple with the StoreName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoreName

`func (o *MarketplaceParticipation) SetStoreName(v string)`

SetStoreName sets StoreName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


