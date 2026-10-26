# Balance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PartnerMetadata** | Pointer to [**PartnerMetadata**](PartnerMetadata.md) |  | [optional] 
**BalanceType** | Pointer to **string** | The type of balance. | [optional] 
**Amount** | Pointer to [**Currency**](Currency.md) |  | [optional] 
**Balances** | Pointer to [**[]Balance**](Balance.md) | A list of balances that comprise this balance. | [optional] 
**AsOfDate** | Pointer to **time.Time** | A date in [ISO 8601](https://developer-docs.amazon.com/sp-api/docs/iso-8601) date-time format. | [optional] 
**LastUpdatedTime** | Pointer to **time.Time** | The date and time when the balance was last updated, in [ISO 8601](https://developer-docs.amazon.com/sp-api/docs/iso-8601) date-time format. The value is in the marketplace&#39;s local time zone. This field is present for point-in-time balance queries (when no &#x60;asOfDate&#x60; query parameter is provided). Null for ending balance queries. | [optional] 

## Methods

### NewBalance

`func NewBalance() *Balance`

NewBalance instantiates a new Balance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBalanceWithDefaults

`func NewBalanceWithDefaults() *Balance`

NewBalanceWithDefaults instantiates a new Balance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPartnerMetadata

`func (o *Balance) GetPartnerMetadata() PartnerMetadata`

GetPartnerMetadata returns the PartnerMetadata field if non-nil, zero value otherwise.

### GetPartnerMetadataOk

`func (o *Balance) GetPartnerMetadataOk() (*PartnerMetadata, bool)`

GetPartnerMetadataOk returns a tuple with the PartnerMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerMetadata

`func (o *Balance) SetPartnerMetadata(v PartnerMetadata)`

SetPartnerMetadata sets PartnerMetadata field to given value.

### HasPartnerMetadata

`func (o *Balance) HasPartnerMetadata() bool`

HasPartnerMetadata returns a boolean if a field has been set.

### GetBalanceType

`func (o *Balance) GetBalanceType() string`

GetBalanceType returns the BalanceType field if non-nil, zero value otherwise.

### GetBalanceTypeOk

`func (o *Balance) GetBalanceTypeOk() (*string, bool)`

GetBalanceTypeOk returns a tuple with the BalanceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalanceType

`func (o *Balance) SetBalanceType(v string)`

SetBalanceType sets BalanceType field to given value.

### HasBalanceType

`func (o *Balance) HasBalanceType() bool`

HasBalanceType returns a boolean if a field has been set.

### GetAmount

`func (o *Balance) GetAmount() Currency`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *Balance) GetAmountOk() (*Currency, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *Balance) SetAmount(v Currency)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *Balance) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetBalances

`func (o *Balance) GetBalances() []Balance`

GetBalances returns the Balances field if non-nil, zero value otherwise.

### GetBalancesOk

`func (o *Balance) GetBalancesOk() (*[]Balance, bool)`

GetBalancesOk returns a tuple with the Balances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalances

`func (o *Balance) SetBalances(v []Balance)`

SetBalances sets Balances field to given value.

### HasBalances

`func (o *Balance) HasBalances() bool`

HasBalances returns a boolean if a field has been set.

### GetAsOfDate

`func (o *Balance) GetAsOfDate() time.Time`

GetAsOfDate returns the AsOfDate field if non-nil, zero value otherwise.

### GetAsOfDateOk

`func (o *Balance) GetAsOfDateOk() (*time.Time, bool)`

GetAsOfDateOk returns a tuple with the AsOfDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsOfDate

`func (o *Balance) SetAsOfDate(v time.Time)`

SetAsOfDate sets AsOfDate field to given value.

### HasAsOfDate

`func (o *Balance) HasAsOfDate() bool`

HasAsOfDate returns a boolean if a field has been set.

### GetLastUpdatedTime

`func (o *Balance) GetLastUpdatedTime() time.Time`

GetLastUpdatedTime returns the LastUpdatedTime field if non-nil, zero value otherwise.

### GetLastUpdatedTimeOk

`func (o *Balance) GetLastUpdatedTimeOk() (*time.Time, bool)`

GetLastUpdatedTimeOk returns a tuple with the LastUpdatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedTime

`func (o *Balance) SetLastUpdatedTime(v time.Time)`

SetLastUpdatedTime sets LastUpdatedTime field to given value.

### HasLastUpdatedTime

`func (o *Balance) HasLastUpdatedTime() bool`

HasLastUpdatedTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


