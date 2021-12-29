# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [balance.proto](#balance.proto)
    - [Balance](#protoBalance.Balance)
    - [CreateBalanceRequest](#protoBalance.CreateBalanceRequest)
    - [CreateBalanceResponse](#protoBalance.CreateBalanceResponse)
    - [GetBalanceRequest](#protoBalance.GetBalanceRequest)
    - [GetBalanceResponse](#protoBalance.GetBalanceResponse)
  
    - [balanceService](#protoBalance.balanceService)
  
- [Scalar Value Types](#scalar-value-types)



<a name="balance.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## balance.proto



<a name="protoBalance.Balance"></a>

### Balance
Represents the user balance


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| user_id | [string](#string) |  | Unique user ID(UUID v4). |
| currency | [string](#string) |  | Currency type. Three capital letters are used. Ex. RUB, USD etc. |
| amount | [uint64](#uint64) |  | Change user balance. A negative number means a decrease in the user&#39;s balance. |
| updated_at | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  | Time of processing. |






<a name="protoBalance.CreateBalanceRequest"></a>

### CreateBalanceRequest
Represents create user balance request.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| balance | [Balance](#protoBalance.Balance) |  |  |






<a name="protoBalance.CreateBalanceResponse"></a>

### CreateBalanceResponse
Represents create user balance response.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| user_id | [string](#string) |  | Unique user ID(UUID v4). |






<a name="protoBalance.GetBalanceRequest"></a>

### GetBalanceRequest
Represents user balance request.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| user_id | [string](#string) |  | Unique user ID(UUID v4). |
| currency | [string](#string) |  | Currency type. Three capital letters are used. Ex. RUB, USD etc. |






<a name="protoBalance.GetBalanceResponse"></a>

### GetBalanceResponse
Represents user balance response.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| balance | [Balance](#protoBalance.Balance) |  | User balance response. |





 

 

 


<a name="protoBalance.balanceService"></a>

### balanceService


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| CreateBalance | [CreateBalanceRequest](#protoBalance.CreateBalanceRequest) | [CreateBalanceResponse](#protoBalance.CreateBalanceResponse) | Create new users balance. |
| GetBalance | [GetBalanceRequest](#protoBalance.GetBalanceRequest) | [GetBalanceResponse](#protoBalance.GetBalanceResponse) | Used to check your balance. |

 



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

