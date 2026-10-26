# OrderPackage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PackageReferenceId** | **string** | A unique identifier for this package within the context of the order. | 
**CreatedTime** | Pointer to **time.Time** | The exact time when this shipping package was created and prepared for shipment. In [ISO 8601](https://developer-docs.amazon.com/sp-api/docs/iso-8601) format. | [optional] 
**PackageStatus** | Pointer to [**PackageStatus**](PackageStatus.md) |  | [optional] 
**Carrier** | Pointer to **string** | The carrier responsible for transporting this package to the customer. | [optional] 
**ShipTime** | Pointer to **time.Time** | The exact time when this package was handed over to the carrier and began its journey to the customer. In [ISO 8601](https://developer-docs.amazon.com/sp-api/docs/iso-8601) format. | [optional] 
**ShippingService** | Pointer to **string** | The specific shipping method or service used for delivering this package. | [optional] 
**TrackingNumber** | Pointer to **string** | The carrier-provided tracking number that customers can use to monitor the package&#39;s delivery progress. | [optional] 
**ShipFromAddress** | Pointer to [**MerchantAddress**](MerchantAddress.md) |  | [optional] 
**PackageItems** | Pointer to [**[]PackageItem**](PackageItem.md) | A list of all order items included in this specific package. | [optional] 

## Methods

### NewOrderPackage

`func NewOrderPackage(packageReferenceId string, ) *OrderPackage`

NewOrderPackage instantiates a new OrderPackage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderPackageWithDefaults

`func NewOrderPackageWithDefaults() *OrderPackage`

NewOrderPackageWithDefaults instantiates a new OrderPackage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPackageReferenceId

`func (o *OrderPackage) GetPackageReferenceId() string`

GetPackageReferenceId returns the PackageReferenceId field if non-nil, zero value otherwise.

### GetPackageReferenceIdOk

`func (o *OrderPackage) GetPackageReferenceIdOk() (*string, bool)`

GetPackageReferenceIdOk returns a tuple with the PackageReferenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageReferenceId

`func (o *OrderPackage) SetPackageReferenceId(v string)`

SetPackageReferenceId sets PackageReferenceId field to given value.


### GetCreatedTime

`func (o *OrderPackage) GetCreatedTime() time.Time`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *OrderPackage) GetCreatedTimeOk() (*time.Time, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *OrderPackage) SetCreatedTime(v time.Time)`

SetCreatedTime sets CreatedTime field to given value.

### HasCreatedTime

`func (o *OrderPackage) HasCreatedTime() bool`

HasCreatedTime returns a boolean if a field has been set.

### GetPackageStatus

`func (o *OrderPackage) GetPackageStatus() PackageStatus`

GetPackageStatus returns the PackageStatus field if non-nil, zero value otherwise.

### GetPackageStatusOk

`func (o *OrderPackage) GetPackageStatusOk() (*PackageStatus, bool)`

GetPackageStatusOk returns a tuple with the PackageStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageStatus

`func (o *OrderPackage) SetPackageStatus(v PackageStatus)`

SetPackageStatus sets PackageStatus field to given value.

### HasPackageStatus

`func (o *OrderPackage) HasPackageStatus() bool`

HasPackageStatus returns a boolean if a field has been set.

### GetCarrier

`func (o *OrderPackage) GetCarrier() string`

GetCarrier returns the Carrier field if non-nil, zero value otherwise.

### GetCarrierOk

`func (o *OrderPackage) GetCarrierOk() (*string, bool)`

GetCarrierOk returns a tuple with the Carrier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarrier

`func (o *OrderPackage) SetCarrier(v string)`

SetCarrier sets Carrier field to given value.

### HasCarrier

`func (o *OrderPackage) HasCarrier() bool`

HasCarrier returns a boolean if a field has been set.

### GetShipTime

`func (o *OrderPackage) GetShipTime() time.Time`

GetShipTime returns the ShipTime field if non-nil, zero value otherwise.

### GetShipTimeOk

`func (o *OrderPackage) GetShipTimeOk() (*time.Time, bool)`

GetShipTimeOk returns a tuple with the ShipTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipTime

`func (o *OrderPackage) SetShipTime(v time.Time)`

SetShipTime sets ShipTime field to given value.

### HasShipTime

`func (o *OrderPackage) HasShipTime() bool`

HasShipTime returns a boolean if a field has been set.

### GetShippingService

`func (o *OrderPackage) GetShippingService() string`

GetShippingService returns the ShippingService field if non-nil, zero value otherwise.

### GetShippingServiceOk

`func (o *OrderPackage) GetShippingServiceOk() (*string, bool)`

GetShippingServiceOk returns a tuple with the ShippingService field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingService

`func (o *OrderPackage) SetShippingService(v string)`

SetShippingService sets ShippingService field to given value.

### HasShippingService

`func (o *OrderPackage) HasShippingService() bool`

HasShippingService returns a boolean if a field has been set.

### GetTrackingNumber

`func (o *OrderPackage) GetTrackingNumber() string`

GetTrackingNumber returns the TrackingNumber field if non-nil, zero value otherwise.

### GetTrackingNumberOk

`func (o *OrderPackage) GetTrackingNumberOk() (*string, bool)`

GetTrackingNumberOk returns a tuple with the TrackingNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackingNumber

`func (o *OrderPackage) SetTrackingNumber(v string)`

SetTrackingNumber sets TrackingNumber field to given value.

### HasTrackingNumber

`func (o *OrderPackage) HasTrackingNumber() bool`

HasTrackingNumber returns a boolean if a field has been set.

### GetShipFromAddress

`func (o *OrderPackage) GetShipFromAddress() MerchantAddress`

GetShipFromAddress returns the ShipFromAddress field if non-nil, zero value otherwise.

### GetShipFromAddressOk

`func (o *OrderPackage) GetShipFromAddressOk() (*MerchantAddress, bool)`

GetShipFromAddressOk returns a tuple with the ShipFromAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipFromAddress

`func (o *OrderPackage) SetShipFromAddress(v MerchantAddress)`

SetShipFromAddress sets ShipFromAddress field to given value.

### HasShipFromAddress

`func (o *OrderPackage) HasShipFromAddress() bool`

HasShipFromAddress returns a boolean if a field has been set.

### GetPackageItems

`func (o *OrderPackage) GetPackageItems() []PackageItem`

GetPackageItems returns the PackageItems field if non-nil, zero value otherwise.

### GetPackageItemsOk

`func (o *OrderPackage) GetPackageItemsOk() (*[]PackageItem, bool)`

GetPackageItemsOk returns a tuple with the PackageItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageItems

`func (o *OrderPackage) SetPackageItems(v []PackageItem)`

SetPackageItems sets PackageItems field to given value.

### HasPackageItems

`func (o *OrderPackage) HasPackageItems() bool`

HasPackageItems returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


