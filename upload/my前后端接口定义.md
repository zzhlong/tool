请求Http协议Head统一参数 
- Head
    ```json
    uuid:xxxx-xxxxx-xxxxx-xxxx-xxxx      //用于识别每个浏览器标示
    token:xxxxxxxxxxxxxxxxxxxxxxxx       //如果含有token请每次请求携带token
    ```

返回参数统一JSON格式 
- 返回值
    ```json
    {
        "success":true, //接口是否处理成功
        "code":0,       //错误代码(以后会根据code直识别错误类型)
        "msg":"",       //错误提示消息
        "data":object   //根据具体业务处理的返回结构体
    }
    ```

## <span id="user_manager">用户管理</span>
---


用途 | URL | Method | 参数类型 |参数/返回值|
---|---|---|---|---|---|---|
获取图形验证码 | /verifyc/code| GET |URL|[传送门](#VerifycCode)
获取短信验证码 | /user/code?phone=x | GET |URL| [传送门](#UserCode)
用户注册| /user/register| POST|JSON|[传送门](#UserRegister)
用户修改密码| /user/register| PUT|JSON|[传送门](#UserRegister1)
用户登录| /user/login| POST|JSON|[传送门](#UserLogin)  | {token:访问令牌}|














## <span id="group">分组管理</span>
---

用途 | URL | Method | 参数类型 | 参数/返回值|
---|---|---|------|---|---|
获取分组列表 | /group| GET |URL|[传送门](#group1) 
添加分组| /group| POST|JSON| [传送门](#group2) 
修改分组| /group| PUT|JSON|[传送门](#group3)
删除分组| /group/{group_id}| DELETE|URL|[传送门](#group4)














## <span id="device">设备管理</span>
---

用途 | URL | Method | 参数类型 | 参数/返回值|
---|---|---|------|---|---|
查询设备分页列表 | /device/list| GET |URL|[传送门](#device1) 
查询设备详细信息 | /device| GET |URL|[传送门](#device2) 
新建设备|/device|POST|JSON|[传送门](#device3) 
修改设备|/device|PUT|JSON|[传送门](#device4) 
连接设备|/device/connect|POST|JSON|[传送门](#device5) 
批量重启设备|/device/reboots|POST|JSON|[传送门](#device6) 
批量恢复出场设置|/device/resets|POST|JSON|[传送门](#device7) 
批量获取设备快照|/device/snapshots|POST|JSON|[传送门](#device8) 
设备音视流+操作流|


















## <span id="app">应用管理</span>
---

用途 | URL | Method | 参数类型 | 参数/返回值|
---|---|---|------|---|---|
获取设备APP列表 | /device/apps| GET |URL|[传送门](#app1)
批量安装应用|/device/apps|POST|JSON|[传送门](#app2)
批量卸载应用|/device/apps/uninstall|POST|JSON|[传送门](#app3)
批量启动应用|/device/app/starts|POST|JSON|[传送门](#app4)



















## <span id="cmd">设备批量管理</span>
---

用途 | URL | Method | 参数类型 | 参数/返回值|
---|---|---|------|---|---|
批量一键home| /device/cmd/homes| POST |JSON|[传送门](#cmd1)
批量一键清除后台|/device/cmd/clear|POST|JSON|[传送门](#cmd1)
批量移动分组|/device/cmd/groups|POST|JSON|[传送门](#cmd1)










## 参数与返回值

### 用户管理

#### 获取图形验证码	<span id="VerifycCode">`/verifyc/code`</span> [GET] [Top](#user_manager)

- 返回值
    ```json
    {
        "success":true,
        "code":0,
        "msg":null,
        "data":"http://www.baid.com/a.png"
    }
    ```
#### 获取短信验证码 <span id="UserCode">`/user/code?phone=x`</span> [GET] [Top](#user_manager)
- 传入参数
    ```json
    phone=17600603219
    ```
- 返回值
    ```json
    {
        "success":true,
        "code":0,
        "msg":"发送成功",
        "data":"ok"
    }
    ```
    
    
#### 注册用户 <span id="UserRegister">`/user/register`</span>[POST][Top](#user_manager)
- 传入参数
    ```json
    {
        "phone":"17892923838",
        "pwd":"xxxxxxx",
        "code":"xxxx"             //短信验证码
        
    }
    ```
- 返回值
    ```json
    {
        "success":true,
        "code":0,
        "msg":"登录成功",
        "data":"Token"
    }
    ```
    
    
#### 修改密码 <span id="UserRegister1">`/user/register`</span>[PUT] [Top](#user_manager)
- 传入参数
    ```json
    {
        "phone":"17892923838",
        "pwd":"xxxxxxx",
        "code":"xxxx"             //短信验证码
    }
    ```
- 返回值
    ```json
    {
        "success":true,
        "code":0,
        "msg":"修改成功",
        "data":"Token"
    }
    ```
    
#### 用户登录 <span id="UserLogin">`/user/login`</span>[POST] [Top](#user_manager) 
- 传入参数
    ```json
    {
        "phone":"17892923838",
        "pwd":"xxxxxxx",
        "code":"xxxx"             //短信验证码
        "verifycode":"xxxx",      //图形验证码
        "login_type":10           //10代表pc 20代表移动端登录
    }
    ```
- 返回值
    ```json
    {
        "success":true,
        "code":0,
        "msg":"登录成功",
        "data":"Token"
    }
    ```
### 分组管理

#### 获取分组列表<span id="group1">`/group`</span> [GET] [Top](#group)

- 返回值
    ```json
    {
        "success":true,
        "code":0,
        "msg":null,
        "data":[
            {
                "group_name":"分组名称",
                "group_id":"分组id",
                "group_count":1                 //分组中设备的数量
                
            }
        ]
    }
    ```
#### 添加分组<span id="group2">`/group`</span> [POST] [Top](#group)
- 传入参数
    ```json
    {
        "group_name":"xxx"
    }
    ```
- 返回值
    ```json
    {
        "success":true,
        "code":0,
        "msg":null,
        "data":{
            "group_id":"分组id"
        }
    }
    ```
#### 修改分组分组<span id="group3">`/group`</span> [PUT] [Top](#group)
- 传入参数
    ```json
    {
        "group_id":"xxx",
        "group_name":"分组名称"
    }
    ```
- 返回值
    ```json
    {
        "success":true,
        "code":0,
        "msg":"修改成功",
        "data":"ok"
    }
    ```
#### 删除分组<span id="group4">`/group/{group_id}`</span> [DELETE] [Top](#group)
- 传入参数
    ```json
       group_id=xxx
    ```
- 返回值
    ```json
    {
        "success":true,
        "code":0,
        "msg":"删除成功",
        "data":"ok"
    }
    ```
    
## 分组管理
### 查询设备分页列表<span id="device1">`/device/list`</span> [GET] [Top](#device)
- 传入参数
    ```json
       group_id=xxx             //分组id
       page_size=x              //每页大小
       page_index=x             //第几页
    ```
- 返回值
    ```json
    {
        "success":true,
        "code":0,
        "msg":"查询成功",
        "data":{
            "page_index":"0",
            "page_size":0,
            "total_count":0,
            "devices":[
                {
                    "device_id":"设备id",
                    "device_name":"设备名称",
                    "expire_time":"设备过期时间",
                    "state":"设备状态"
                }
            ]
        }
    }
    ```
### 查询设备详细信息<span id="device2">`/device`</span> [GET] [Top](#device)
- 传入参数
    ```json
       device_id="设备id"
    ```
- 返回值
    ```json
    {
        "success":true,
        "code":0,
        "msg":"查询成功",
        "data":{
            "device_id":"设备id",
            "device_name":"设备名称",
            "expire_time":2020-12-02 10:10,
            "state":"设备状态",
            "create_time":2020-01-02 10:10,
            "info":"设备相关配置",
            "agent_id":"代理id",
            "code":"对应物理设备编码"
        }
    }
    ```
### 新建设备<span id="device3">`/device`</span> [POST] [Top](#device)
- 传入参数
    ```json
       {
            "secretkey":"密钥"
       }
    ```
- 返回值
    ```json
    {
        "success":true,
        "code":0,
        "msg":"新建成功",
        "data":{
            "device_id":"设备id",
        }
    }
    ```
    
### 修改设备<span id="device4">`/device`</span> [PUT] [Top](#device)
- 传入参数
    ```json
       {
           "device_name":"设备名称",
           "device_id":"设备id"
       }
    ```
- 返回值
    ```json
    {
        "success":true,
        "code":0,
        "msg":"修改成功",
        "data":"ok"
    }
    ```
### 连接设备<span id="device5">`/device/connect`</span> [POST] [Top](#device)
- 传入参数
    ```json
        [
           {
               "device_id":"设备id"
           }
        ]
    ```
- 返回值
    ```json
    {
        "success":true,
        "code":0,
        "msg":"连接成功",
        "data":"ok"
    }
    ```
### 批量重启设备<span id="device6">`/device/reboots`</span> [POST] [Top](#device)
- 传入参数
    ```json
        [
           {
               "device_id":"设备id"
           }
        ]
    ```
- 返回值
    ```json
    {
        "success":true,
        "code":0,
        "msg":"重启成功",
        "data":"ok"
    }
    ```
### 批量恢复出场设置<span id="device7">`/device/resets`</span> [POST] [Top](#device)
- 传入参数
    ```json
        [
           {
               "device_id":"设备id"
           }
        ]
    ```
- 返回值
    ```json
    {
        "success":true,
        "code":0,
        "msg":"设置成功",
        "data":"ok"
    }
    ```
### 批量获取设备快照<span id="device8">`/device/snapshots`</span> [POST] [Top](#device)
- 传入参数
    ```json
        [
           {
               "device_id":"设备id"
           }
        ]
    ```
- 返回值
    ```json
    {
        "success":true,
        "code":0,
        "msg":"快照信息",
        "data":"ok"
    }
    ```
### 获取设备APP列表<span id="app1">`/device/apps`</span> [GET] [Top](#qpp)
- 传入参数
    ```json
       page_size=x              //每页大小
       page_index=x             //第几页
    ```
- 返回值
    ```json
    {
        "success":true,
        "code":0,
        "msg":"查询成功",
        "data":{
            "page_index":"0",
            "page_size":0,
            "total_count":0,
            "apps":[
                {
                    "app_id":"appid",
                    "app_name":"app名称",
                    "app_ack_url":"app包路径",
                }
            ]
        }
    }
    ``` 
### 批量安装应用<span id="app2">`/device/apps`</span> [POST] [Top](#app)
- 传入参数
    ```json
        {
            "app":["app_id"],
            "device":["device_id"]
        }
    ```
- 返回值
    ```json
    {
        "success":true,
        "code":0,
        "msg":"安装完成",
        "data":"ok"
    }
    ```
    
### 批量卸载应用<span id="app3">`/device/apps/uninstall`</span> [POST] [Top](#app)
- 传入参数
    ```json
        {
            "app":["app_id"],
            "device":["device_id"]
        }
    ```
- 返回值
    ```json
    {
        "success":true,
        "code":0,
        "msg":"卸载完成",
        "data":"ok"
    }
    ```
### 批量启动应用<span id="app4">`/device/app/starts`</span> [POST] [Top](#app)
- 传入参数
    ```json
        {
            "app":["app_id"],
            "device":["device_id"]
        }
    ```
- 返回值
    ```json
    {
        "success":true,
        "code":0,
        "msg":"启动成功",
        "data":"ok"
    }
    ```

## 设备批量管理
### 批量一键home<span id="cmd1">`/device/cmd/homes`</span> [POST] [Top](cmd)
- 传入参数
    ```json
        [
            "device_id"
        ]
    ```
- 返回值
    ```json
    {
        "success":true,
        "code":0,
        "msg":"执行完成",
        "data":"ok"
    }
    ```
### 批量一键清除后台<span id="cmd2">`/device/cmd/clear`</span> [POST] [Top](cmd)
- 传入参数
    ```json
        [
            "device_id"
        ]
    ```
- 返回值
    ```json
    {
        "success":true,
        "code":0,
        "msg":"执行完成",
        "data":"ok"
    }
    ```
### 批量移动分组<span id="cmd3">`/device/cmd/groups`</span> [POST] [Top](cmd)
- 传入参数
    ```json
        {
            "group_id":"xxxx",
            "devices":["device_id"]
        }
    ```
- 返回值
    ```json
    {
        "success":true,
        "code":0,
        "msg":"分组完成",
        "data":"ok"
    }
    ```