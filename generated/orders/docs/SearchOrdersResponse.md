# SearchOrdersResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Orders** | [**[]Order**](Order.md) | An array containing all orders that match the search criteria. | 
**Pagination** | Pointer to [**Pagination**](Pagination.md) |  | [optional] 
**LastUpdatedBefore** | Pointer to **time.Time** | Only orders updated before the specified time are returned. The date must be in [ISO 8601](https://developer-docs.amazon.com/sp-api/docs/iso-8601) format. | [optional] 
**CreatedBefore** | Pointer to **time.Time** | Only orders placed before the specified time are returned. The date must be in [ISO 8601](https://developer-docs.amazon.com/sp-api/docs/iso-8601) format. | [optional] 

## Methods

### NewSearchOrdersResponse

`func NewSearchOrdersResponse(orders []Order, ) *SearchOrdersResponse`

NewSearchOrdersResponse instantiates a new SearchOrdersResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchOrdersResponseWithDefaults

`func NewSearchOrdersResponseWithDefaults() *SearchOrdersResponse`

NewSearchOrdersResponseWithDefaults instantiates a new SearchOrdersResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrders

`func (o *SearchOrdersResponse) GetOrders() []Order`

GetOrders returns the Orders field if non-nil, zero value otherwise.

### GetOrdersOk

`func (o *SearchOrdersResponse) GetOrdersOk() (*[]Order, bool)`

GetOrdersOk returns a tuple with the Orders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrders

`func (o *SearchOrdersResponse) SetOrders(v []Order)`

SetOrders sets Orders field to given value.


### GetPagination

`func (o *SearchOrdersResponse) GetPagination() Pagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *SearchOrdersResponse) GetPaginationOk() (*Pagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *SearchOrdersResponse) SetPagination(v Pagination)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *SearchOrdersResponse) HasPagination() bool`

HasPagination returns a boolean if a field has been set.

### GetLastUpdatedBefore

`func (o *SearchOrdersResponse) GetLastUpdatedBefore() time.Time`

GetLastUpdatedBefore returns the LastUpdatedBefore field if non-nil, zero value otherwise.

### GetLastUpdatedBeforeOk

`func (o *SearchOrdersResponse) GetLastUpdatedBeforeOk() (*time.Time, bool)`

GetLastUpdatedBeforeOk returns a tuple with the LastUpdatedBefore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedBefore

`func (o *SearchOrdersResponse) SetLastUpdatedBefore(v time.Time)`

SetLastUpdatedBefore sets LastUpdatedBefore field to given value.

### HasLastUpdatedBefore

`func (o *SearchOrdersResponse) HasLastUpdatedBefore() bool`

HasLastUpdatedBefore returns a boolean if a field has been set.

### GetCreatedBefore

`func (o *SearchOrdersResponse) GetCreatedBefore() time.Time`

GetCreatedBefore returns the CreatedBefore field if non-nil, zero value otherwise.

### GetCreatedBeforeOk

`func (o *SearchOrdersResponse) GetCreatedBeforeOk() (*time.Time, bool)`

GetCreatedBeforeOk returns a tuple with the CreatedBefore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBefore

`func (o *SearchOrdersResponse) SetCreatedBefore(v time.Time)`

SetCreatedBefore sets CreatedBefore field to given value.

### HasCreatedBefore

`func (o *SearchOrdersResponse) HasCreatedBefore() bool`

HasCreatedBefore returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


