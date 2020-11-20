[toc] 



# http接口

规范

# 1.user

## 1.1user-用户模块

启动日志

开始连接 120.79.169.139
连接成功 120.79.169.139
开始连接 120.79.169.139
连接成功 120.79.169.139
INFO[0000] model 初始化成功
读到 447297 IP条数据
INFO[0000] 系统初始化成功!
INFO[0000] Transport [http] Listening on [::]:37945
INFO[0000] Broker [http] Connected to [::]:34095
INFO[0000] Registry [mdns] Registering node: ars.cloud.api.user-9524910a-8e72-4daa-8633-a433e29dbcc0
INFO[0020] User/Login
INFO[0020] user-GetByUserName
INFO[0020] token-Add
Set Redis Cache failed!
INFO[0020] userClient-Add
INFO[0020] user is {<nil> 3 1 2 13662240977 1ddf5d63cdeafcfea3ba87baa1577dad mio [mio.wang@arsyun.com](mailto:mio.wang@arsyun.com) 0 0000-00-00 00:00:00 1}



### user用户部分

#### add

```jsx
192.168.1.212:8080/user/user/add
{
    "code": 0,
    "description": "",
    "data": {
        "user_id": 7
    }
}
```

#### del

#### edit

#### editPassword

#### get

```jsx
{
    "code": 0,
    "description": "",
    "data": {
        "user_id": 1,
        "username": "15170195695",
        "name": "总管理员",
        "email": "",
        "role_id": 1,
        "role_name": "总管理员",
        "status": 9
    }
}

{
    "code": 0,
    "description": "",
    "data": {
        "user_id": 2,
        "username": "15256688999",
        "name": "lotus",
        "email": "",
        "role_id": 3,
        "role_name": "系统运维",
        "status": 1
    }
}

{
    "code": 0,
    "description": "",
    "data": {
        "user_id": 3,
        "username": "13662240977",
        "name": "mio",
        "email": "mio.wang@arsyun.com",
        "role_id": 2,
        "role_name": "管理员",
        "status": 1
    }
}

user_id 2 系统运维
```

#### list

```jsx
192.168.1.212:8080/user/user/list
{
    "code": 0,
    "description": "",
    "data": {
        "total": 6,
        "per_page": 10,
        "current_page": 1,
        "last_page": 1,
        "data": [
            {
                "user_id": 6,
                "username": "15222222222",
                "name": "admin",
                "email": "",
                "role_id": 2,
                "role_name": "管理员",
                "status": 0
            },
            {
                "user_id": 5,
                "username": "18565629925",
                "name": "运维用户",
                "email": "",
                "role_id": 3,
                "role_name": "系统运维",
                "status": 0
            },
            {
                "user_id": 4,
                "username": "17665339196",
                "name": "管理员用户",
                "email": "",
                "role_id": 2,
                "role_name": "管理员",
                "status": 0
            },
            {
                "user_id": 3,
                "username": "13662240977",
                "name": "mio",
                "email": "mio.wang@arsyun.com",
                "role_id": 2,
                "role_name": "管理员",
                "status": 0
            },
            {
                "user_id": 2,
                "username": "15256688999",
                "name": "lotus",
                "email": "",
                "role_id": 3,
                "role_name": "系统运维",
                "status": 0
            },
            {
                "user_id": 1,
                "username": "15170195695",
                "name": "总管理员",
                "email": "",
                "role_id": 1,
                "role_name": "总管理员",
                "status": 0
            }
        ]
    }
}
```

#### checkToken

```jsx
{
    "code": 100000,
    "description": "用户token有误或已过期",
    "data": null
}
{
    "code": 0,
    "description": "",
    "data": null
}
```

#### checkUserName

#### login

```jsx
192.168.1.212:8080/user/user/login
{
    "code": 0,
    "description": "",
    "data": {
        "user_id": 3,
        "token": "34aaf9925cf33ab4bc149fa328f98161"
    }
}
```

#### logout

```jsx
192.168.1.212:8080/user/user/logout
{
    "code": 0,
    "description": "",
    "data": null
}
```

#### sendSms



### role用户角色

add

del

edit

editPermission

get

list

```jsx
192.168.1.212:8080/user/role/list
{
    "code": 0,
    "description": "",
    "data": {
        "total": 3,
        "per_page": 10,
        "current_page": 1,
        "last_page": 1,
        "data": [
            {
                "Id": 3,
                "Name": "系统运维",
                "Info": "系统运维",
                "Status": 3
            },
            {
                "Id": 2,
                "Name": "管理员",
                "Info": "管理员",
                "Status": 5
            },
            {
                "Id": 1,
                "Name": "总管理员",
                "Info": "总管理员",
                "Status": 9
            }
        ]
    }
}

```

initRole

### userclient用户设备



### file用户文件



## 1.2user-管理员模块

### adminUser

#### add

#### del

#### edit

#### editPassword

#### get

#### list

#### checkToken

#### checkUserName

#### login

```
192.168.1.212:8080/user/adminUser/login

1

8b0ed838900ad046e7ba8eaa8329fff2
```





#### logout



#### editRole

#### editUserName



### adminRole





### adminClient

#### list



### adminApp













# 2.warning

## 2.1warning-报警模块



## 2.2warning-管理员报警



# 3.workorder

## 3.1workorder-工单模块



## 3.2workorder-管理员工单





# 4.pool

## 4.1pool-矿池模块

### pool矿池部分



### group矿池分组



### income收益



### device设备



### license



### miner矿机相关



## 4.2pool-管理员矿池

### adminPool



### adminDevice



### adminMiner



















# 参考

[链接](https://www.tapd.cn/51838973/documents/view/1151838973001000514?file_type=mindmap)

