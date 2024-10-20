
# 
[TOC]

## 整体说明
1.	字符串都为utf8格式;
1.	HTTP Headers:
	1.	Content-Type设置为：application/json
1.	DataTime格式参考RFC3339标准

## 错误处理
错误的具体信息将在error字段中返回。

### 错误码示例
```json
{
    "code": "400",
    "message": "Param Error"
}
```


### 状态码列表
| 状态码 | 说明 |
|---|---|
| 200 | 返回正常 |
| 400 | 参数错误 |
| 401 | 无access<br> key或key无效 |
| 500 | 服务器内部错误 |


## 查询cat列表

### 请求路径
```http
GET /tnr/v1/cats
```


### 请求参数

#### Query 参数
| 参数名 | 参数类型 | 格式类型 | 是否必须 | 默认值 | 说明 |
|---|---|---|---|---|---|
| `page_size` | `integer` | `Int32` | 否 |  |  |
| `page_token` | `string` |  | 否 |  |  |
| `skip` | `integer` | `Int32` | 否 |  |  |


### 返回值

#### 返回对象
| type | description |
|---|---|
| `Array<cat.Cat>` |  |


#### `cat.Cat`
| field | type | format | required | default | description |
|---|---|---|---|---|---|
| `id` | `string` |  | N |  |
| `name` | `string` |  | N |  | 名称 |
| `nickName` | `string` |  | N |  | 昵称 |
| `sex` | `string` |  | N |  | 性别, male(Tomcat), female(Queen) |
| `brithDay` | `string` | `Timestamp` | N |  | 出生日期 |
| `location` | `string` |  | N |  | 地理位置 |
| `area` | `string` |  | N |  | 地理范围 |
| `lastFoundTime` | `string` | `Timestamp` | N |  | 最后一次发现的时间 |
| `status` | `string` |  | N |  | 状态, trap, neuter, return |
| `state` | `string` |  | N |  | cat的物理或情感状态 |
| `createTime` | `string` | `Timestamp` | N |  | create time |
| `updateTime` | `string` | `Timestamp` | N |  | update time |
| `deleteTime` | `string` | `Timestamp` | N |  | delete time |


## 创建cat信息

### 请求路径
```http
POST /tnr/v1/cats
```


### 请求参数

#### Body 请求对象
| field | type | format | required | default | description |
|---|---|---|---|---|---|
| `id` | `string` |  | N |  |
| `name` | `string` |  | N |  | 名称 |
| `nickName` | `string` |  | N |  | 昵称 |
| `sex` | `string` |  | N |  | 性别, male(Tomcat), female(Queen) |
| `brithDay` | `string` | `Timestamp` | N |  | 出生日期 |
| `location` | `string` |  | N |  | 地理位置 |
| `area` | `string` |  | N |  | 地理范围 |
| `lastFoundTime` | `string` | `Timestamp` | N |  | 最后一次发现的时间 |
| `status` | `string` |  | N |  | 状态, trap, neuter, return |
| `state` | `string` |  | N |  | cat的物理或情感状态 |
| `createTime` | `string` | `Timestamp` | N |  | create time |
| `updateTime` | `string` | `Timestamp` | N |  | update time |
| `deleteTime` | `string` | `Timestamp` | N |  | delete time |


### 返回值

#### 返回对象
| field | type | format | required | default | description |
|---|---|---|---|---|---|
| `id` | `string` |  | N |  |
| `name` | `string` |  | N |  | 名称 |
| `nickName` | `string` |  | N |  | 昵称 |
| `sex` | `string` |  | N |  | 性别, male(Tomcat), female(Queen) |
| `brithDay` | `string` | `Timestamp` | N |  | 出生日期 |
| `location` | `string` |  | N |  | 地理位置 |
| `area` | `string` |  | N |  | 地理范围 |
| `lastFoundTime` | `string` | `Timestamp` | N |  | 最后一次发现的时间 |
| `status` | `string` |  | N |  | 状态, trap, neuter, return |
| `state` | `string` |  | N |  | cat的物理或情感状态 |
| `createTime` | `string` | `Timestamp` | N |  | create time |
| `updateTime` | `string` | `Timestamp` | N |  | update time |
| `deleteTime` | `string` | `Timestamp` | N |  | delete time |


## 更新cat信息

### 请求路径
```http
PUT /tnr/v1/cats/{cat_id}
```


### 请求参数

#### Body 请求对象
| field | type | format | required | default | description |
|---|---|---|---|---|---|
| `id` | `string` |  | N |  |
| `name` | `string` |  | N |  | 名称 |
| `nickName` | `string` |  | N |  | 昵称 |
| `sex` | `string` |  | N |  | 性别, male(Tomcat), female(Queen) |
| `brithDay` | `string` | `Timestamp` | N |  | 出生日期 |
| `location` | `string` |  | N |  | 地理位置 |
| `area` | `string` |  | N |  | 地理范围 |
| `lastFoundTime` | `string` | `Timestamp` | N |  | 最后一次发现的时间 |
| `status` | `string` |  | N |  | 状态, trap, neuter, return |
| `state` | `string` |  | N |  | cat的物理或情感状态 |
| `createTime` | `string` | `Timestamp` | N |  | create time |
| `updateTime` | `string` | `Timestamp` | N |  | update time |
| `deleteTime` | `string` | `Timestamp` | N |  | delete time |


### 返回值

#### 返回对象
| field | type | format | required | default | description |
|---|---|---|---|---|---|
| `id` | `string` |  | N |  |
| `name` | `string` |  | N |  | 名称 |
| `nickName` | `string` |  | N |  | 昵称 |
| `sex` | `string` |  | N |  | 性别, male(Tomcat), female(Queen) |
| `brithDay` | `string` | `Timestamp` | N |  | 出生日期 |
| `location` | `string` |  | N |  | 地理位置 |
| `area` | `string` |  | N |  | 地理范围 |
| `lastFoundTime` | `string` | `Timestamp` | N |  | 最后一次发现的时间 |
| `status` | `string` |  | N |  | 状态, trap, neuter, return |
| `state` | `string` |  | N |  | cat的物理或情感状态 |
| `createTime` | `string` | `Timestamp` | N |  | create time |
| `updateTime` | `string` | `Timestamp` | N |  | update time |
| `deleteTime` | `string` | `Timestamp` | N |  | delete time |


## 查询cat信息

### 请求路径
```http
GET /tnr/v1/cats/{id}
```


### 请求参数

#### Path 参数
| 参数名 | 参数类型 | 格式类型 | 说明 |
|---|---|---|---|
| `id` | `string` |  |  |


### 返回值

#### 返回对象
| field | type | format | required | default | description |
|---|---|---|---|---|---|
| `id` | `string` |  | N |  |
| `name` | `string` |  | N |  | 名称 |
| `nickName` | `string` |  | N |  | 昵称 |
| `sex` | `string` |  | N |  | 性别, male(Tomcat), female(Queen) |
| `brithDay` | `string` | `Timestamp` | N |  | 出生日期 |
| `location` | `string` |  | N |  | 地理位置 |
| `area` | `string` |  | N |  | 地理范围 |
| `lastFoundTime` | `string` | `Timestamp` | N |  | 最后一次发现的时间 |
| `status` | `string` |  | N |  | 状态, trap, neuter, return |
| `state` | `string` |  | N |  | cat的物理或情感状态 |
| `createTime` | `string` | `Timestamp` | N |  | create time |
| `updateTime` | `string` | `Timestamp` | N |  | update time |
| `deleteTime` | `string` | `Timestamp` | N |  | delete time |


## 删除cat信息

### 请求路径
```http
DELETE /tnr/v1/cats/{id}
```


### 请求参数

#### Path 参数
| 参数名 | 参数类型 | 格式类型 | 说明 |
|---|---|---|---|
| `id` | `string` |  |  |


### 返回值

#### 返回对象
对象为空

## 批量获取cat信息

### 请求路径
```http
GET /tnr/v1/cats:batch-get
```


### 请求参数

#### Query 参数
| 参数名 | 参数类型 | 格式类型 | 是否必须 | 默认值 | 说明 |
|---|---|---|---|---|---|
| `ids` | `Array<string>` |  | 否 |  |  |


### 返回值

#### 返回对象
| type | description |
|---|---|
| `Array<cat.Cat>` |  |


#### `cat.Cat`
| field | type | format | required | default | description |
|---|---|---|---|---|---|
| `id` | `string` |  | N |  |
| `name` | `string` |  | N |  | 名称 |
| `nickName` | `string` |  | N |  | 昵称 |
| `sex` | `string` |  | N |  | 性别, male(Tomcat), female(Queen) |
| `brithDay` | `string` | `Timestamp` | N |  | 出生日期 |
| `location` | `string` |  | N |  | 地理位置 |
| `area` | `string` |  | N |  | 地理范围 |
| `lastFoundTime` | `string` | `Timestamp` | N |  | 最后一次发现的时间 |
| `status` | `string` |  | N |  | 状态, trap, neuter, return |
| `state` | `string` |  | N |  | cat的物理或情感状态 |
| `createTime` | `string` | `Timestamp` | N |  | create time |
| `updateTime` | `string` | `Timestamp` | N |  | update time |
| `deleteTime` | `string` | `Timestamp` | N |  | delete time |
