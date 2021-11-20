# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [transaction.proto](#transaction.proto)
    - [Transaction](#transactionService.Transaction)
    - [TransactionRequest](#transactionService.TransactionRequest)
    - [TransactionResponse](#transactionService.TransactionResponse)
    - [TransactionsRequest](#transactionService.TransactionsRequest)
    - [TransactionsResponse](#transactionService.TransactionsResponse)
  
    - [balanceService](#transactionService.balanceService)
  
- [Scalar Value Types](#scalar-value-types)



<a name="transaction.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## transaction.proto



<a name="transactionService.Transaction"></a>

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






<a name="transactionService.TransactionRequest"></a>

### TransactionRequest
Represents Transaction Request.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| transaction | [Transaction](#transactionService.Transaction) |  | User balance change request. |






<a name="transactionService.TransactionResponse"></a>

### TransactionResponse
Represents Transaction Response.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| transaction_id | [string](#string) |  | User balance change response. |






<a name="transactionService.TransactionsRequest"></a>

### TransactionsRequest
Represents request for accrual or reduction of users balance.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| user_id | [string](#string) |  | Unique user ID(UUID v4). |
| currency | [string](#string) |  | Currency type. Three capital letters are used. Ex. RUB, USD etc. |
| page | [sint64](#sint64) |  | Page number. |
| size | [sint64](#sint64) |  | Page size. |






<a name="transactionService.TransactionsResponse"></a>

### TransactionsResponse
Represents sorted user transactions


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| total_count | [sint64](#sint64) |  | Total count. |
| total_pages | [sint64](#sint64) |  | Total pages. |
| page | [sint64](#sint64) |  | Page number. |
| size | [sint64](#sint64) |  | Page size. |
| has_more | [bool](#bool) |  | End of message flag. |
| transactions | [Transaction](#transactionService.Transaction) | repeated | List of transactions. |





 

 

 


<a name="transactionService.balanceService"></a>

### balanceService


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| CreateTransaction | [TransactionRequest](#transactionService.TransactionRequest) | [TransactionResponse](#transactionService.TransactionResponse) | Used for crediting or debiting funds from the user&#39;s account. |
| GetTransactions | [TransactionsRequest](#transactionService.TransactionsRequest) | [TransactionsResponse](#transactionService.TransactionsResponse) stream | Used to view the history of transactions |

 



## Scalar Value Types

| .proto Type | Notes | C++ | Java | Python | Go | C# | PHP | Ruby |
| ----------- | ----- | --- | ---- | ------ | -- | -- | --- | ---- |
| <a name="double" /> double |  | double | double | float | float64 | double | float | Float |
| <a name="float" /> float |  | float | float | float | float32 | float | float | Float |
| <a name="int32" /> int32 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint32 instead. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="int64" /> int64 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint64 instead. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="uint32" /> uint32 | Uses variable-length encoding. | uint32 | int | int/long | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="uint64" /> uint64 | Uses variable-length encoding. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum or Fixnum (as required) |
| <a name="sint32" /> sint32 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int32s. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sint64" /> sint64 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int64s. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="fixed32" /> fixed32 | Always four bytes. More efficient than uint32 if values are often greater than 2^28. | uint32 | int | int | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="fixed64" /> fixed64 | Always eight bytes. More efficient than uint64 if values are often greater than 2^56. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum |
| <a name="sfixed32" /> sfixed32 | Always four bytes. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sfixed64" /> sfixed64 | Always eight bytes. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="bool" /> bool |  | bool | boolean | boolean | bool | bool | boolean | TrueClass/FalseClass |
| <a name="string" /> string | A string must always contain UTF-8 encoded or 7-bit ASCII text. | string | String | str/unicode | string | string | string | String (UTF-8) |
| <a name="bytes" /> bytes | May contain any arbitrary sequence of bytes. | string | ByteString | str | []byte | ByteString | string | String (ASCII-8BIT) |

