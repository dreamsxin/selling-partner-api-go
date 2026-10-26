# ListBalancesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Balances** | Pointer to [**[]Balance**](Balance.md) | A list of balances. | [optional] 
**NextToken** | Pointer to **string** | A token to retrieve the next page of results. If null, no more results are available. | [optional] 

## Methods

### NewListBalancesResponse

`func NewListBalancesResponse() *ListBalancesResponse`

NewListBalancesResponse instantiates a new ListBalancesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListBalancesResponseWithDefaults

`func NewListBalancesResponseWithDefaults() *ListBalancesResponse`

NewListBalancesResponseWithDefaults instantiates a new ListBalancesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBalances

`func (o *ListBalancesResponse) GetBalances() []Balance`

GetBalances returns the Balances field if non-nil, zero value otherwise.

### GetBalancesOk

`func (o *ListBalancesResponse) GetBalancesOk() (*[]Balance, bool)`

GetBalancesOk returns a tuple with the Balances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalances

`func (o *ListBalancesResponse) SetBalances(v []Balance)`

SetBalances sets Balances field to given value.

### HasBalances

`func (o *ListBalancesResponse) HasBalances() bool`

HasBalances returns a boolean if a field has been set.

### GetNextToken

`func (o *ListBalancesResponse) GetNextToken() string`

GetNextToken returns the NextToken field if non-nil, zero value otherwise.

### GetNextTokenOk

`func (o *ListBalancesResponse) GetNextTokenOk() (*string, bool)`

GetNextTokenOk returns a tuple with the NextToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextToken

`func (o *ListBalancesResponse) SetNextToken(v string)`

SetNextToken sets NextToken field to given value.

### HasNextToken

`func (o *ListBalancesResponse) HasNextToken() bool`

HasNextToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


