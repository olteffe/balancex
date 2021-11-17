# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [balance.proto](#balance.proto)
    - [Balance](#balanceService.Balance)
    - [BalanceRequest](#balanceService.BalanceRequest)
    - [BalanceResponse](#balanceService.BalanceResponse)
    - [OperationsRequest](#balanceService.OperationsRequest)
    - [OperationsResponse](#balanceService.OperationsResponse)
    - [Transaction](#balanceService.Transaction)
    - [TransactionRequest](#balanceService.TransactionRequest)
    - [TransactionResponse](#balanceService.TransactionResponse)
  
    - [balanceService](#balanceService.balanceService)
  
- [Scalar Value Types](#scalar-value-types)



<a name="balance.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## balance.proto



<a name="balanceService.Balance"></a>

### Balance
Represents the user balance


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| user_id | [string](#string) |  | Unique user ID(UUID v4). |
| currency | [string](#string) |  | Currency type. Three capital letters are used. Ex. RUB, USD etc. |
| amount | [uint64](#uint64) |  | Change user balance. A negative number means a decrease in the user&#39;s balance. |
| updated_at | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  | Time of processing. |






<a name="balanceService.BalanceRequest"></a>

### BalanceRequest
Represents user balance request.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| user_id | [string](#string) |  | Unique user ID(UUID v4). |
| currency | [string](#string) |  | Currency type. Three capital letters are used. Ex. RUB, USD etc. |






<a name="balanceService.BalanceResponse"></a>

### BalanceResponse
Represents user balance response.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| balance | [Balance](#balanceService.Balance) |  | User balance response. |






<a name="balanceService.OperationsRequest"></a>

### OperationsRequest
Represents request for accrual or reduction of users balance.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| user_id | [string](#string) |  | Unique user ID(UUID v4). |
| currency | [string](#string) |  | Currency type. Three capital letters are used. Ex. RUB, USD etc. |
| page | [sint64](#sint64) |  | Page number. |
| size | [sint64](#sint64) |  | Page size. |






<a name="balanceService.OperationsResponse"></a>

### OperationsResponse
Represents sorted user transactions


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| total_count | [sint64](#sint64) |  | Total count. |
| total_pages | [sint64](#sint64) |  | Total pages. |
| page | [sint64](#sint64) |  | Page number. |
| size | [sint64](#sint64) |  | Page size. |
| has_more | [bool](#bool) |  | End of message flag. |
| transactions | [Transaction](#balanceService.Transaction) | repeated | List of transactions. |






<a name="balanceService.Transaction"></a>

### Transaction
Represents the transaction


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| transaction_id | [string](#string) |  | Unique request operation ID(UUID v4). |
| source | [string](#string) |  | Source of account change. |
| description | [string](#string) |  | Transaction description. |
| sender_id | [string](#string) |  | Unique sender ID(UUID v4). If sender is other microservice, then use static personal mc ID. |
| recipient_id | [string](#string) |  | Unique receiver user ID(UUID v4). |
| currency | [string](#string) |  | Currency type. Three capital letters are used. Ex. RUB, USD etc. |
| amount | [sint64](#sint64) |  | Change user balance. A negative number means a decrease in the user&#39;s balance. |
| created_at | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  | Time of processing. |






<a name="balanceService.TransactionRequest"></a>

### TransactionRequest
Represents Transaction Request.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| transaction | [Transaction](#balanceService.Transaction) |  | User balance change request. |






<a name="balanceService.TransactionResponse"></a>

### TransactionResponse
Represents Transaction Response.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| transaction_id | [string](#string) |  | User balance change response. |





 

 

 


<a name="balanceService.balanceService"></a>

### balanceService


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| CreateTransaction | [TransactionRequest](#balanceService.TransactionRequest) | [TransactionResponse](#balanceService.TransactionResponse) | Used for crediting or debiting funds from the user&#39;s account. |
| GetBalance | [BalanceRequest](#balanceService.BalanceRequest) | [BalanceResponse](#balanceService.BalanceResponse) | Used to check your balance. |
| GetOperations | [OperationsRequest](#balanceService.OperationsRequest) | [OperationsResponse](#balanceService.OperationsResponse) stream | Used to view the history of transactions |

