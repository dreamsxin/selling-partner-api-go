# EventFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AggregationSettings** | Pointer to [**AggregationSettings**](AggregationSettings.md) |  | [optional] 
**MarketplaceIds** | Pointer to **[]string** | A list of marketplace identifiers you can subscribe to (for example, &#x60;ATVPDKIKX0DER&#x60;). To receive notifications in every marketplace, do not provide this list. | [optional] 
**OrderChangeTypes** | Pointer to [**[]OrderChangeTypeEnum**](OrderChangeTypeEnum.md) | A list of order change types you can subscribe to (for example, &#x60;BuyerRequestedChange&#x60;). To receive notifications of all change types, do not provide this list. | [optional] 
**TrackingIdentifier** | Pointer to [**TrackingIdentifier**](TrackingIdentifier.md) |  | [optional] 
**EventFilterType** | **string** | An &#x60;eventFilterType&#x60; value that the &#x60;notificationType&#x60; supports. The subscription service uses the &#x60;eventFilterType&#x60; to determine the type of event filter. To determine if a specific &#x60;notificationType&#x60; supports an &#x60;eventFilterType&#x60;, refer to [Notification Type Values]( https://developer-docs.amazon.com/sp-api/docs/notification-type-values). | 

## Methods

### NewEventFilter

`func NewEventFilter(eventFilterType string, ) *EventFilter`

NewEventFilter instantiates a new EventFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventFilterWithDefaults

`func NewEventFilterWithDefaults() *EventFilter`

NewEventFilterWithDefaults instantiates a new EventFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAggregationSettings

`func (o *EventFilter) GetAggregationSettings() AggregationSettings`

GetAggregationSettings returns the AggregationSettings field if non-nil, zero value otherwise.

### GetAggregationSettingsOk

`func (o *EventFilter) GetAggregationSettingsOk() (*AggregationSettings, bool)`

GetAggregationSettingsOk returns a tuple with the AggregationSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregationSettings

`func (o *EventFilter) SetAggregationSettings(v AggregationSettings)`

SetAggregationSettings sets AggregationSettings field to given value.

### HasAggregationSettings

`func (o *EventFilter) HasAggregationSettings() bool`

HasAggregationSettings returns a boolean if a field has been set.

### GetMarketplaceIds

`func (o *EventFilter) GetMarketplaceIds() []string`

GetMarketplaceIds returns the MarketplaceIds field if non-nil, zero value otherwise.

### GetMarketplaceIdsOk

`func (o *EventFilter) GetMarketplaceIdsOk() (*[]string, bool)`

GetMarketplaceIdsOk returns a tuple with the MarketplaceIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarketplaceIds

`func (o *EventFilter) SetMarketplaceIds(v []string)`

SetMarketplaceIds sets MarketplaceIds field to given value.

### HasMarketplaceIds

`func (o *EventFilter) HasMarketplaceIds() bool`

HasMarketplaceIds returns a boolean if a field has been set.

### GetOrderChangeTypes

`func (o *EventFilter) GetOrderChangeTypes() []OrderChangeTypeEnum`

GetOrderChangeTypes returns the OrderChangeTypes field if non-nil, zero value otherwise.

### GetOrderChangeTypesOk

`func (o *EventFilter) GetOrderChangeTypesOk() (*[]OrderChangeTypeEnum, bool)`

GetOrderChangeTypesOk returns a tuple with the OrderChangeTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderChangeTypes

`func (o *EventFilter) SetOrderChangeTypes(v []OrderChangeTypeEnum)`

SetOrderChangeTypes sets OrderChangeTypes field to given value.

### HasOrderChangeTypes

`func (o *EventFilter) HasOrderChangeTypes() bool`

HasOrderChangeTypes returns a boolean if a field has been set.

### GetTrackingIdentifier

`func (o *EventFilter) GetTrackingIdentifier() TrackingIdentifier`

GetTrackingIdentifier returns the TrackingIdentifier field if non-nil, zero value otherwise.

### GetTrackingIdentifierOk

`func (o *EventFilter) GetTrackingIdentifierOk() (*TrackingIdentifier, bool)`

GetTrackingIdentifierOk returns a tuple with the TrackingIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackingIdentifier

`func (o *EventFilter) SetTrackingIdentifier(v TrackingIdentifier)`

SetTrackingIdentifier sets TrackingIdentifier field to given value.

### HasTrackingIdentifier

`func (o *EventFilter) HasTrackingIdentifier() bool`

HasTrackingIdentifier returns a boolean if a field has been set.

### GetEventFilterType

`func (o *EventFilter) GetEventFilterType() string`

GetEventFilterType returns the EventFilterType field if non-nil, zero value otherwise.

### GetEventFilterTypeOk

`func (o *EventFilter) GetEventFilterTypeOk() (*string, bool)`

GetEventFilterTypeOk returns a tuple with the EventFilterType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventFilterType

`func (o *EventFilter) SetEventFilterType(v string)`

SetEventFilterType sets EventFilterType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


