## /用户模块/用户登录
```text
渲染用户登录界面
```
#### 接口状态
> 开发中

#### 接口URL
> localhost:3000/api/v1/login

#### 请求方式
> GET

#### Content-Type
> json

#### 请求Body参数
```javascript
{
}
```
#### 认证方式
```text
noauth
```
#### 预执行脚本
```javascript
暂无预执行脚本
```
#### 后执行脚本
```javascript
暂无后执行脚本
```
#### 成功响应示例
```javascript
{
	"msg": "OK",
	"status": 200,
}
```
## /用户模块/用户登录

```text
用户登录模块信息验证
```

#### 接口状态

> 开发中

#### 接口URL

> localhost:3000/api/v1/login

#### 请求方式

> POST

#### Content-Type

> json

#### 请求Body参数

```javascript
{
    "email": "123@qq.com",
    "password": "12345"
}
```

**请求参数**:

| 参数名称 | 参数说明 | 是否必须 | 数据类型 |
| :------- | -------- | -------- | -------- |
| email    | 用户邮箱 | true     | string   |
| password | 密码     | true     | string   |

#### 认证方式

```text
noauth
```

#### 预执行脚本

```javascript
暂无预执行脚本
```

#### 后执行脚本

```javascript
暂无后执行脚本
```

#### 成功响应示例

```javascript
{
	"msg": "OK",
	"status": 200,
}
```

#### 错误响应示例

```javascript
{
	"msg": "密码错误",
	"status": 1002,
}
```



## /用户模块/用户注册

```text
渲染用户注册界面
```

#### 接口状态

> 开发中

#### 接口URL

> localhost:3000/api/v1/register

#### 请求方式

> GET

#### Content-Type

> json

#### 请求Body参数

```javascript
{
}
```

#### 认证方式

```text
noauth
```

#### 预执行脚本

```javascript
暂无预执行脚本
```

#### 后执行脚本

```javascript
暂无后执行脚本
```

#### 成功响应示例

```javascript
{
	"msg": "OK",
	"status": 200,
}
```

## /用户模块/用户注册

```text
用户注册信息验证
```

#### 接口状态

> 开发中

#### 接口URL

> localhost:3000/api/v1/register

#### 请求方式

> POST

#### Content-Type

> json

#### 请求Body参数

```javascript
{
    "email": "123@qq.com",
    "username": "test",
    "password": "12345",
}
```

**请求参数**:

| 参数名称 | 参数说明      | 是否必须 | 数据类型 |
| :------- | ------------- | -------- | -------- |
| email    | 用户邮箱,唯一 | true     | string   |
| username | 用户名        | true     | string   |
| password | 密码          | true     | string   |

#### 认证方式

```text
noauth
```

#### 预执行脚本

```javascript
暂无预执行脚本
```

#### 后执行脚本

```javascript
暂无后执行脚本
```

#### 成功响应示例

```javascript
{
	"msg": "OK",
	"status": 200,
}
```

#### 错误响应示例

```javascript
{
	"msg": "FAIL",
	"status": 500,
}
```



## /Model模块/ModelZoo

```text
展示已有的模型信息
```

#### 接口状态

> 开发中

#### 接口URL

> localhost:3000/api/v1/modelzoo

#### 请求方式

> GET

#### Content-Type

> json

#### 请求Body参数

```javascript
{
}
```

#### 认证方式

```text
noauth
```

#### 预执行脚本

```javascript
暂无预执行脚本
```

#### 后执行脚本

```javascript
暂无后执行脚本
```

#### 成功响应示例

```javascript
{
	"msg": "OK",
	"status": 200,
    "data": [
        {
            "ID": 1,
            "UpdatedAt": "2023-08-15T11:15:34+08:00",
            "model_name": "solcv7",
            "dataset_name": "whu-opt-sar",
            "model_path": "solcv7_whu-opt-sar_path",
            "type": "segmentation"
        },
        {
            "ID": 2,
            "UpdatedAt": "2023-08-15T11:16:22+08:00",
            "model_name": "solcv7",
            "dataset_name": "potsdam",
            "model_path": "solcv7_potsdam_path",
            "type": "segmentation"
        }
    ]
}
```

#### 错误响应示例

```javascript
{
	"msg": "FAIL",
	"status": 500,
    "data": []
}
```

**响应参数**:


| 参数名称     | 参数说明             | 类型   |
| ------------ | -------------------- | ------ |
| type         | 任务类型             | string |
| model_name   | 模型名称             | string |
| dataset_name | 训练模型使用的数据集 | string |
| model_path   | 模型存储路径         | string |
| UpdatedAt    | 最近一次更新模型时间 | date   |

## /Model模块/添加模型

```text
添加已训练好的模型信息
```

#### 接口状态

> 开发中

#### 接口URL

> localhost:3000/api/v1/modelzoo/add

#### 请求方式

> POST

#### Content-Type

> json

#### 请求Body参数

```javascript
{
   "model_name": "solcv7",
   "dataset_name": "whu-opt-sar",
   "model_path": "solcv7_whu-opt-sar_path",
   "type": "segmentation"
}
```

**请求参数**:

| 参数名称     | 参数说明             | 是否必须 | 数据类型 |
| :----------- | -------------------- | -------- | -------- |
| type         | 任务类型             | true     | string   |
| model_name   | 模型名称             | true     | string   |
| dataset_name | 训练模型使用的数据集 | true     | string   |
| model_path   | 模型存储路径         | true     | string   |

#### 认证方式

```text
noauth
```

#### 预执行脚本

```javascript
暂无预执行脚本
```

#### 后执行脚本

```javascript
暂无后执行脚本
```

#### 成功响应示例

```javascript
{
	"msg": "OK",
	"status": 200
}
```

#### 错误响应示例

```javascript
{
	"msg": "FAIL",
	"status": 500,
}
```



## /推理模块/InferencePage

```text
渲染推理页面
```

#### 接口状态

> 开发中

#### 接口URL

> localhost:3000/api/v1/inference

#### 请求方式

> GET

#### Content-Type

> json

#### 请求Body参数

```javascript
{
}
```

#### 认证方式

```text
noauth
```

#### 预执行脚本

```javascript
暂无预执行脚本
```

#### 后执行脚本

```javascript
暂无后执行脚本
```

#### 成功响应示例

```javascript
{
	"msg": "OK",
	"status": 200
}
```

#### 错误响应示例

```javascript
{
	"msg": "FAIL",
	"status": 500,
}
```

## /推理模块/推理任务

```text
执行推理任务
```

#### 接口状态

> 开发中

#### 接口URL

> localhost:3000/api/v1/inference

#### 请求方式

> POST

#### Content-Type

> json

#### 请求Body参数

```javascript
{
   "type": "solcv7",
   "model_name": "whu-opt-sar",
   "dataset_name": "solcv7_whu-opt-sar_path",
   "image": image
}
```

**请求参数**:

| 参数名称     | 参数说明             | 是否必须 | 数据类型 |
| :----------- | -------------------- | -------- | -------- |
| type         | 任务类型             | true     | string   |
| model_name   | 要使用的模型名称     | true     | string   |
| dataset_name | 训练模型使用的数据集 | true     | string   |
| image        | 要推理的图像         | true     | file     |

#### 认证方式

```text
noauth
```

#### 预执行脚本

```javascript
暂无预执行脚本
```

#### 后执行脚本

```javascript
暂无后执行脚本
```

#### 成功响应示例

```javascript
{
	"msg": "OK",
	"status": 200
    "data":{
   		"image_name": "image1.jpg",
        "image_path": "user/inference/images/image1.jpg"
        "result_name": "result_image1.jpg"
        "result_path": "user/inference/results/result_image1.jpg"
    }
}
```

#### 错误响应示例

```javascript
{
	"msg": "FAIL",
	"status": 500,
    "data":{}
}
```

| 参数名称    | 参数说明           | 类型   |
| ----------- | ------------------ | ------ |
| image_name  | 前端图像的名称     | string |
| image_path  | 前端图像保存路径   | string |
| result_name | 推理结果图像的名称 | string |
| result_path | 推理结果保存路径   | string |



## /Gee模块/GeePage

```text
渲染Gee页面
```

#### 接口状态

> 开发中

#### 接口URL

> localhost:3000/api/v1/gee

#### 请求方式

> GET

#### Content-Type

> json

#### 请求Body参数

```javascript
{
}
```

#### 认证方式

```text
noauth
```

#### 预执行脚本

```javascript
暂无预执行脚本
```

#### 后执行脚本

```javascript
暂无后执行脚本
```

#### 成功响应示例

```javascript
{
	"msg": "OK",
	"status": 200
}
```

#### 错误响应示例

```javascript
{
	"msg": "FAIL",
	"status": 500,
}
```

## /Gee模块/任务

```text
执行Gee采集图像任务
```

#### 接口状态

> 开发中

#### 接口URL

> localhost:3000/api/v1/gee

#### 请求方式

> POST

#### Content-Type

> json

#### 请求Body参数

```javascript
{
    "start_date" "2023-08-19",
	"end_date"   "2023-08-20",
	"max_pixel"  255,   
	"min_pixel"  0,   
	"point_lon1" 123.123,   
	"point_lat1" 123.456,   
	"point_lon2" 456.123,   
	"point_lat2" 456.456,   
	"dataset"    "landsat",    
	"image_name" "gee_image"    
}
```

**请求参数**:

| 参数名称   | 参数说明               | 是否必须 | 数据类型 |
| :--------- | ---------------------- | -------- | -------- |
| start_date | 起始时间               | true     | date     |
| end_date   | 终止时间               | true     | date     |
| max_pixel  | 图像像素值最大值       | true     | float32  |
| min_pixel  | 图像像素值最小值       | true     | float32  |
| point_lon1 | 采集区域左上点的经度值 | true     | float32  |
| point_lat1 | 采集区域左上点的纬度值 | true     | float32  |
| point_lon2 | 采集区域右下点的经度值 | true     | float32  |
| point_lat2 | 采集区域右下点的纬度值 | true     | float32  |
| dataset    | 采集使用的数据集名称   | true     | string   |
| image_name | 保存图像时的名称       | true     | string   |

#### 认证方式

```text
noauth
```

#### 预执行脚本

```javascript
暂无预执行脚本
```

#### 后执行脚本

```javascript
暂无后执行脚本
```

#### 成功响应示例

```javascript
{
	"msg": "OK",
	"status": 200
    "data":{
   		"image_name": "save_image_name",
        "image_path": "save_image_path"
    }
}
```

#### 错误响应示例

```javascript
{
	"msg": "FAIL",
	"status": 500,
    "data":{}
}
```


**响应参数**:


| 参数名称   | 参数说明              | 类型   |
| ---------- | --------------------- | ------ |
| image_name | Gee采集图像保存的名称 | string |
| image_path | Gee采集图像保存的路径 | string |