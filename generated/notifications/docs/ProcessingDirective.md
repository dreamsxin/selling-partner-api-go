# ProcessingDirective

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventFilter** | Pointer to [**EventFilter**](EventFilter.md) |  | [optional] 
**FilterExpression** | Pointer to **string** | An expression for filtering events before delivery to destination based on the notification payload (example: FulfillmentOrderStatusNotification.FulfillmentOrderStatus &#x3D;&#x3D; &#x60;SHIPPED&#x60; ). The &#x60;filterExpression&#x60; is a string that follows the CEL expression syntax (https://github.com/google/cel-spec) excluding arithmetic operators (+, -, *, /, %) and list/map indexing ([]). Refer to Notification Type Values to determine if filter Expression is supported for a Notification Type. Refer to CEL Operators (https://developer-docs.amazon.com/sp-api/docs/filter-notification-subscriptions) to see if a CEL operator is supported.   Note: eventFilter and filterExpression are mutually exclusive. You can use filterExpression to replace existing eventFilter configurations. | [optional] 

## Methods

### NewProcessingDirective

`func NewProcessingDirective() *ProcessingDirective`

NewProcessingDirective instantiates a new ProcessingDirective object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProcessingDirectiveWithDefaults

`func NewProcessingDirectiveWithDefaults() *ProcessingDirective`

NewProcessingDirectiveWithDefaults instantiates a new ProcessingDirective object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventFilter

`func (o *ProcessingDirective) GetEventFilter() EventFilter`

GetEventFilter returns the EventFilter field if non-nil, zero value otherwise.

### GetEventFilterOk

`func (o *ProcessingDirective) GetEventFilterOk() (*EventFilter, bool)`

GetEventFilterOk returns a tuple with the EventFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventFilter

`func (o *ProcessingDirective) SetEventFilter(v EventFilter)`

SetEventFilter sets EventFilter field to given value.

### HasEventFilter

`func (o *ProcessingDirective) HasEventFilter() bool`

HasEventFilter returns a boolean if a field has been set.

### GetFilterExpression

`func (o *ProcessingDirective) GetFilterExpression() string`

GetFilterExpression returns the FilterExpression field if non-nil, zero value otherwise.

### GetFilterExpressionOk

`func (o *ProcessingDirective) GetFilterExpressionOk() (*string, bool)`

GetFilterExpressionOk returns a tuple with the FilterExpression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterExpression

`func (o *ProcessingDirective) SetFilterExpression(v string)`

SetFilterExpression sets FilterExpression field to given value.

### HasFilterExpression

`func (o *ProcessingDirective) HasFilterExpression() bool`

HasFilterExpression returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


