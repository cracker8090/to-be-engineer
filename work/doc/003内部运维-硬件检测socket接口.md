

hos/go-lib-common/socket_struct/cloud_device_up.go

固件上行数据包结构体
连接上报，验证权限



hos/go-lib-common/socket_struct/cloud_device_down.go

固件下行数据包结构体
上报成功后必回

云端下发





# 云端负责上行下行接口

## hos/go-lib-common/socket_struct/cloud_device_down.go

```go
/////////////////// 检测类 //////////////////////
//act——set_hdcheck_info
//硬件检测
type DevDownSetHardwareCheckStruct struct {
	Act          string `json:"act"`
	Reqtype      string `json:"reqtype"`
	Act_id       string `json:"act_id"`
	Client_id    string `json:"client_id"`
	Act_code     string `json:"act_code"`
	hardwareDetection struct{
		RootDirDetection struct{
			Size int64 `json:"Size"`  //B
		} `json:"root_dir_detection"`
		MemDetection struct{
			MemorySize int `json:"MemorySize"` //GB
		} `json:"mem_detection"`
		DiskDetection []struct{
			Size int64 `json:"Size"`  //B
			Num int `json:"Num"`
			Type int `json:"Type"` //0 = any  1 = HDD   2 =  SSD
		} `json:"disk_detection"`
		GPUDetection struct{
			DisplayName string `json:"DisplayName"`   //* any
			Num int `json:"Num"`
		} `json:"gpu_detection"`
	}
}

//act——get_hardware_check_result
//获取检测结果
type DevDownGetHardwareCheckResultStruct struct {
	Act            string `json:"act"`
	Reqtype        string `json:"reqtype"`
	Act_id         string `json:"act_id"`
	Client_id      string `json:"client_id"`
	Act_code       string `json:"act_code"`
	Passed			bool  `json:"passed"`
	ResultList     []struct{
		Name string `json:"name"`
		Passed bool `json:"passed"`
		Msg string `json:"msg"`
	} `json:"result_list"`
}
```

## hos/go-lib-common/socket_struct/cloud_device_up.go

```go
///////////////////////////检类///////////////////////////////
//获取检测结果
type DevDownGetHardwareCheckResultStruct struct {
	Act            string `json:"act"`
	Reqtype        string `json:"reqtype"`
	Act_id         string `json:"act_id"`
	Client_id      string `json:"client_id"`
	Act_code       string `json:"act_code"`
	Passed			bool `json:"passed"`
	ResultList     []struct{
		Name string `json:"name"`
		Status int `json:"status"`
		Msg string `json:"msg"`
	} `json:"result_list"`
}
```







## go-srv-cloud-worker/srv/handler/device.go

检测操作

```go
device_get_system_time











```









# 示例

## 矿机-上行接口



每次建立连接上报 act-start

基本信息上报（开机） act

心跳 act-heartbeat

set操作返回 act set_return code/desc

报警 act-set_warning_info

查系统时间 act-get_system_time



收到后上报 act-get system_info

系统显示信息 act-get_system_showinfo

磁盘信息 act_system_diskinfo



软件列表 act-get_app_list

软件详情 act-get_app_info



网卡信息 act-get_net_baseinfo

磁盘信息 act-get_disk_info

矿机类



## 矿机-下行接口



通用回复 上报成功后必返回

act =return

action 矿机-上行接口

code 错误码，0代表成功，其他失败

description 出错原因



查询类

查询基本信息 act-get_system_baseinfo





































































