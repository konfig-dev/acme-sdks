

# Pet

A pet for sale in the pet store

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
|**tags** | [**List&lt;Tag&gt;**](Tag.md) |  |  [optional] |
|**id** | **Long** |  |  [optional] |
|**category** | [**Category**](Category.md) |  |  [optional] |
|**name** | **String** |  |  |
|**photoUrls** | **List&lt;String&gt;** |  |  |
|**status** | [**StatusEnum**](#StatusEnum) | pet status in the store |  [optional] |



## Enum: StatusEnum

| Name | Value |
|---- | -----|
| AVAILABLE | &quot;available&quot; |
| PENDING | &quot;pending&quot; |
| SOLD | &quot;sold&quot; |



