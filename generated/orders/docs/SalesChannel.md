# SalesChannel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChannelName** | **string** | The name of the sales platform or channel where the customer placed this order.  **Possible values**: &#x60;AMAZON&#x60;, &#x60;NON_AMAZON&#x60; | 
**MarketplaceId** | Pointer to **string** | The unique identifier for the specific marketplace within the sales channel where this order was placed. | [optional] 
**MarketplaceName** | Pointer to **string** | The human-readable name of the marketplace where this order was placed. | [optional] 

## Methods

### NewSalesChannel

`func NewSalesChannel(channelName string, ) *SalesChannel`

NewSalesChannel instantiates a new SalesChannel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSalesChannelWithDefaults

`func NewSalesChannelWithDefaults() *SalesChannel`

NewSalesChannelWithDefaults instantiates a new SalesChannel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChannelName

`func (o *SalesChannel) GetChannelName() string`

GetChannelName returns the ChannelName field if non-nil, zero value otherwise.

### GetChannelNameOk

`func (o *SalesChannel) GetChannelNameOk() (*string, bool)`

GetChannelNameOk returns a tuple with the ChannelName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelName

`func (o *SalesChannel) SetChannelName(v string)`

SetChannelName sets ChannelName field to given value.


### GetMarketplaceId

`func (o *SalesChannel) GetMarketplaceId() string`

GetMarketplaceId returns the MarketplaceId field if non-nil, zero value otherwise.

### GetMarketplaceIdOk

`func (o *SalesChannel) GetMarketplaceIdOk() (*string, bool)`

GetMarketplaceIdOk returns a tuple with the MarketplaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarketplaceId

`func (o *SalesChannel) SetMarketplaceId(v string)`

SetMarketplaceId sets MarketplaceId field to given value.

### HasMarketplaceId

`func (o *SalesChannel) HasMarketplaceId() bool`

HasMarketplaceId returns a boolean if a field has been set.

### GetMarketplaceName

`func (o *SalesChannel) GetMarketplaceName() string`

GetMarketplaceName returns the MarketplaceName field if non-nil, zero value otherwise.

### GetMarketplaceNameOk

`func (o *SalesChannel) GetMarketplaceNameOk() (*string, bool)`

GetMarketplaceNameOk returns a tuple with the MarketplaceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarketplaceName

`func (o *SalesChannel) SetMarketplaceName(v string)`

SetMarketplaceName sets MarketplaceName field to given value.

### HasMarketplaceName

`func (o *SalesChannel) HasMarketplaceName() bool`

HasMarketplaceName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


