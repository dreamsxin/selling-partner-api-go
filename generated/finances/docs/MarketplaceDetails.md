# MarketplaceDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MarketplaceId** | Pointer to **string** | The identifier of the marketplace where the transaction occurred. The marketplace ID is the globally unique identifier of a marketplace. To find the ID for a marketplace, refer to [Marketplace IDs](https://developer-docs.amazon.com/sp-api/docs/marketplace-ids). | [optional] 
**MarketplaceName** | Pointer to **string** | The name of the marketplace where the transaction occurred. For example: &#x60;Amazon.com&#x60;,&#x60;Amazon.in&#x60; | [optional] 

## Methods

### NewMarketplaceDetails

`func NewMarketplaceDetails() *MarketplaceDetails`

NewMarketplaceDetails instantiates a new MarketplaceDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMarketplaceDetailsWithDefaults

`func NewMarketplaceDetailsWithDefaults() *MarketplaceDetails`

NewMarketplaceDetailsWithDefaults instantiates a new MarketplaceDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMarketplaceId

`func (o *MarketplaceDetails) GetMarketplaceId() string`

GetMarketplaceId returns the MarketplaceId field if non-nil, zero value otherwise.

### GetMarketplaceIdOk

`func (o *MarketplaceDetails) GetMarketplaceIdOk() (*string, bool)`

GetMarketplaceIdOk returns a tuple with the MarketplaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarketplaceId

`func (o *MarketplaceDetails) SetMarketplaceId(v string)`

SetMarketplaceId sets MarketplaceId field to given value.

### HasMarketplaceId

`func (o *MarketplaceDetails) HasMarketplaceId() bool`

HasMarketplaceId returns a boolean if a field has been set.

### GetMarketplaceName

`func (o *MarketplaceDetails) GetMarketplaceName() string`

GetMarketplaceName returns the MarketplaceName field if non-nil, zero value otherwise.

### GetMarketplaceNameOk

`func (o *MarketplaceDetails) GetMarketplaceNameOk() (*string, bool)`

GetMarketplaceNameOk returns a tuple with the MarketplaceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarketplaceName

`func (o *MarketplaceDetails) SetMarketplaceName(v string)`

SetMarketplaceName sets MarketplaceName field to given value.

### HasMarketplaceName

`func (o *MarketplaceDetails) HasMarketplaceName() bool`

HasMarketplaceName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


