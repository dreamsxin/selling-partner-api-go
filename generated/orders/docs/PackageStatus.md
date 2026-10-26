# PackageStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | **string** | Primary status classification of the package in the shipping workflow. | 
**DetailedStatus** | Pointer to **string** | Granular status information providing specific details about the package&#39;s current location and handling stage.   **Possible values**: - &#x60;PENDING_SCHEDULE&#x60; (Package awaiting pickup scheduling) - &#x60;PENDING_PICK_UP&#x60; (Package ready for carrier collection from seller) - &#x60;PENDING_DROP_OFF&#x60; (Package awaiting seller delivery to carrier) - &#x60;LABEL_CANCELLED&#x60; (Shipping label canceled by seller) - &#x60;PICKED_UP&#x60; (Package collected by carrier from seller location) - &#x60;DROPPED_OFF&#x60; (Package delivered to carrier by seller) - &#x60;AT_ORIGIN_FC&#x60; (Package at originating fulfillment center) - &#x60;AT_DESTINATION_FC&#x60; (Package at destination fulfillment center) - &#x60;DELIVERED&#x60; (Package successfully delivered to recipient) - &#x60;REJECTED_BY_BUYER&#x60; (Package refused by intended recipient) - &#x60;UNDELIVERABLE&#x60; (Package cannot be delivered due to address or access issues) - &#x60;RETURNING_TO_SELLER&#x60; (Package in transit back to seller) - &#x60;RETURNED_TO_SELLER&#x60; (Package successfully returned to seller) - &#x60;LOST&#x60; (Package location unknown or confirmed lost) - &#x60;OUT_FOR_DELIVERY&#x60; (Package on delivery vehicle for final delivery) - &#x60;DAMAGED&#x60; (Package damaged during transit)  | [optional] 

## Methods

### NewPackageStatus

`func NewPackageStatus(status string, ) *PackageStatus`

NewPackageStatus instantiates a new PackageStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPackageStatusWithDefaults

`func NewPackageStatusWithDefaults() *PackageStatus`

NewPackageStatusWithDefaults instantiates a new PackageStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *PackageStatus) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PackageStatus) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PackageStatus) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetDetailedStatus

`func (o *PackageStatus) GetDetailedStatus() string`

GetDetailedStatus returns the DetailedStatus field if non-nil, zero value otherwise.

### GetDetailedStatusOk

`func (o *PackageStatus) GetDetailedStatusOk() (*string, bool)`

GetDetailedStatusOk returns a tuple with the DetailedStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetailedStatus

`func (o *PackageStatus) SetDetailedStatus(v string)`

SetDetailedStatus sets DetailedStatus field to given value.

### HasDetailedStatus

`func (o *PackageStatus) HasDetailedStatus() bool`

HasDetailedStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


