# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [balance.proto](#balance.proto)
    - [Balance](#balanceService.Balance)
    - [GetBalanceRequest](#balanceService.GetBalanceRequest)
    - [GetBalanceResponse](#balanceService.GetBalanceResponse)
    - [OperationRequest](#balanceService.OperationRequest)
    - [OperationResponse](#balanceService.OperationResponse)
    - [Transaction](#balanceService.Transaction)
    - [TransactionRequest](#balanceService.TransactionRequest)
    - [TransactionResponse](#balanceService.TransactionResponse)
    - [TransferRequest](#balanceService.TransferRequest)
    - [TransferResponse](#balanceService.TransferResponse)
  
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
| request_id | [string](#string) |  | Unique request operation ID(UUID v4). |
| user_id | [string](#string) |  | Unique user ID(UUID v4). |
| currency | [string](#string) |  | Currency type. Three capital letters are used. Ex. RUB, USD etc. |
| amount | [int64](#int64) |  | Change user balance. A negative number means a decrease in the user&#39;s balance. |
| created_at | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  | Time of processing. Format will be determined later. |






<a name="balanceService.GetBalanceRequest"></a>

### GetBalanceRequest
Represents user balance request


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| request_id | [string](#string) |  | Unique request operation ID(UUID v4). |
| user_id | [string](#string) |  | Unique user ID(UUID v4). |






<a name="balanceService.GetBalanceResponse"></a>

### GetBalanceResponse
Represents user balance response


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| request_id | [string](#string) |  | Unique request operation ID(UUID v4). |
| balance | [Balance](#balanceService.Balance) |  | The users balance. |
| status | [string](#string) |  | State change status. |






<a name="balanceService.OperationRequest"></a>

### OperationRequest
Represents request for accrual or reduction of money


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| request_id | [string](#string) |  | Unique request operation ID(UUID v4). |
| user_id | [string](#string) |  | Unique user ID(UUID v4). |
| currency | [string](#string) |  | Currency type. Three capital letters are used. Ex. RUB, USD etc. |
| amount | [int64](#int64) |  | Change user balance. A negative number means a decrease in the user&#39;s balance. |
| created_at | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  | Time of processing. Format will be determined later. |






<a name="balanceService.OperationResponse"></a>

### OperationResponse
Represents operation response


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| request_id | [string](#string) |  | Unique request operation ID(UUID v4). |
| status | [string](#string) |  | State change status |






<a name="balanceService.Transaction"></a>

### Transaction
Represents the transaction


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| request_id | [string](#string) |  | Unique request operation ID(UUID v4). |
| from | [string](#string) |  | Source of account change. |
| description | [string](#string) |  | Transaction description. |
| amount | [int64](#int64) |  | Change user balance. A negative number means a decrease in the user&#39;s balance. |
| created_at | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  | Time of processing. Format will be determined later. |






<a name="balanceService.TransactionRequest"></a>

### TransactionRequest
Its a request for a list of transactions with pagination and sorting


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| request_id | [string](#string) |  | Unique request operation ID(UUID v4). |
| user_id | [string](#string) |  | Unique user ID(UUID v4). |
| currency | [string](#string) |  | Currency type. Three capital letters are used. Ex. RUB, USD etc. |
| page | [uint64](#uint64) |  | Page number. |
| size | [uint64](#uint64) |  | Page size. |






<a name="balanceService.TransactionResponse"></a>

### TransactionResponse
Represents sorted user transactions


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| request_id | [string](#string) |  | Unique request operation ID(UUID v4). |
| user_id | [string](#string) |  | Unique user ID(UUID v4). |
| currency | [string](#string) |  | Currency type. Three capital letters are used. Ex. RUB, USD etc. |
| transaction | [Transaction](#balanceService.Transaction) | repeated | List of transactions. |
| status | [string](#string) |  | State change status |






<a name="balanceService.TransferRequest"></a>

### TransferRequest
Represents transfer request from one user to another


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| request_id | [string](#string) |  | Unique request operation ID(UUID v4). |
| sender_id | [string](#string) |  | Unique sender ID(UUID v4). |
| recipient_id | [string](#string) |  | Unique receiver ID(UUID v4). |
| currency | [string](#string) |  | Currency type. Three capital letters are used. Ex. RUB, USD etc. |
| amount | [uint64](#uint64) |  | Change user balance. A negative number means a decrease in the user&#39;s balance. |
| created_at | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  | Time of processing. Format will be determined later. |






<a name="balanceService.TransferResponse"></a>

### TransferResponse
Represents transfer response


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| request_id | [string](#string) |  | Unique request operation ID(UUID v4). |
| status | [string](#string) |  | State change status |





 

 

 


<a name="balanceService.balanceService"></a>

### balanceService


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| Operation | [OperationRequest](#balanceService.OperationRequest) | [OperationResponse](#balanceService.OperationResponse) | Used for crediting or debiting funds from the user&#39;s account |
| Transfer | [TransferRequest](#balanceService.TransferRequest) | [TransferResponse](#balanceService.TransferResponse) | Used to transfer funds between users |
| GetBalance | [GetBalanceRequest](#balanceService.GetBalanceRequest) | [GetBalanceResponse](#balanceService.GetBalanceResponse) | Used to check your balance |
| Transaction | [TransactionRequest](#balanceService.TransactionRequest) | [TransactionResponse](#balanceService.TransactionResponse) | Used to view the history of transactions |

 


