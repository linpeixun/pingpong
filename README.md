# pingpong
go study

## 消息结构
```
Message
|-Header（[]byte）
|-ServiceId(string)
|-ServiceMethod(string)
|-MetaInfo(map[string]string)
|-payload([]byte)
```
## 消息byte数组结构
Header+(后续数据总长度-4位)+(ServiceId长度-4位)+ServiceId+（ServiceMethod长度-4位)+ServiceMethod+(MetaInfo长度-4位)+MetaInfo+(payload长度-4位)+payload

## HEADER 
12位

|0|1|2|3|4|5|6|7|8|9|10|11|
|-----|-----|-----|-----|-----|-----|-----|-----|-----|-----|-----|-----|
|magic(0x80)|Version|MessageType|SerializeType|Seq|Seq|Seq|Seq|Seq|Seq|Seq|Seq|

### MessageType
单字节8位

|7|6|5|4|3|2|1|0|
 |-----|-----|-----|-----|-----|-----|-----|-----|
 |request or response|1为心跳|1为需要返回值|压缩类型|压缩类型|压缩类型|消息状态|消息状态|
 
 
## ServiceId
服务标识

## ServiceMethod
服务方法名

## Meta Info

Meta Info为Map，存储的key、value均为String

Meta Info转为byte数组的结果如下：

Key长度（4位）|Key内容|Value长度(4位)|Value内容|

## payload
消息体