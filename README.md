# cstc
cloud storage transfer cli
云存储传递命令

目前已经完成阿里云oss 的上传和下载
后续将补充 腾讯云 AWS GCP 等云存储的上传递下载和列出某目录下的文件

编译
```bash
go build 
#或者通过-o 参数指定编译输出可执行文件名
go build -o cstcli 
```
跨平台编译可以参考我在[CSDN的go博客](https://blog.csdn.net/JadeJones/article/details/118096463#2.4%20%E8%B7%A8%E5%B9%B3%E5%8F%B0%E7%BC%96%E8%AF%91) 这篇文章下面的跨平台编译
https://blog.csdn.net/JadeJones/article/details/118096463#2.4%20%E8%B7%A8%E5%B9%B3%E5%8F%B0%E7%BC%96%E8%AF%91

# 用法
```bash
$ go run main.go --help
cloud-storage-transfer-cli ...

Usage:
  cstcli [flags]
  cstcli [command]

Available Commands:
  get         download cloud storage to local
  help        Help about any command
  upload      upload file to cloud storage

Flags:
  -i, --access_id string      the cloud storage access id
  -k, --access_key string     the cloud storage access key
  -h, --help                  help for cstcli
  -p, --oss_provider string   cloud storage Provider [ali/tx] (default "ali")

Use "cstcli [command] --help" for more information about a command.
```

## 上传
```bash
$ go run main.go upload --help
upload file to cloud storage

Usage:
  cstcli upload [flags]

Flags:
  -e, --bucket_endpoint string   cloud storage endpoint
  -b, --bucket_name string       cloud storage bucket name
  -f, --file_path string         upload file path
  -h, --help                     help for upload
  -u, --public string            provide public download link (default "false")
```
例子1 不提供下载地址仅上传
```bash
cstcli upload -f go.sum -e bucket_endpoint -i access_id -k access_key -b bucket_name
```
例子2 提供临时3天的下载地址
```bash
cstcli upload -f go.sum -e bucket_endpoint -i access_id -k access_key -b bucket_name -u true
```

## 下载
```bash
download cloud storage to local

Usage:
  cstcli get [flags]

Flags:
  -e, --bucket_endpoint string     cloud storage endpoint
  -b, --bucket_name string         cloud storage bucket name
  -f, --file_path string           download file path
  -h, --help                       help for get
  -o, --out_path_filename string   save file path and filename
```
例子1 下载到当前目录下
```bash
cstcli get -f go.sum -e bucket_endpoint -i access_id -k access_key -b bucket_name
```

例子2 下载导指定目录(可以是相对路径和绝对路径)
```bash
cstcli get -f go.sum  -o ../../go2.sum -e bucket_endpoint -i access_id -k access_key -b bucket_name
```

## 查看

```
watch cloud storage file

Usage:
  cstcli watch [flags]

Flags:
  -e, --bucket_endpoint string   cloud storage endpoint
  -b, --bucket_name string       cloud storage bucket name (default "rsto")
  -h, --help                     help for watch
  -m, --watch_Max int            watch cloud storage file max (default 100)
  -x, --watch_prefix string      watch cloud storage file prefix
```

例子1 默认查看存储桶中所有文件(分页展示,每页默认100可通过-m修改)

```
cstcli watch -x csts -e bucket_endpoint -i access_id -k access_key -b bucket_name
```

例子2 分页查看csts目录下的所有文件(包括该目录下目录)

```
cstcli watch -x csts -e bucket_endpoint -i access_id -k access_key -b bucket_name
```

例子3 分页查看csts目录下包含指定前缀(golang)的文件

```
cstcli watch -x csts/golang -e bucket_endpoint -i access_id -k access_key -b bucket_name
```



# 提示

1. 默认上传地址为 bucket_name/cstc 有公开下载地址的文件保存在bucket_name/cstc_tmp
2. access_key与access_id 在对应云上请参考最佳实践, 尽量将权限控制到足够安全。例如下面阿里云的策略。建议创建编程用户并授权
https://help.aliyun.com/document_detail/141923.html
3. 如果你想设置default的access_key与access_id等私密信息 以便减少命令行导致的隐私泄露请参考下面内容
打开该项目(cstc)目录下的cmd/root.go 文件找到下面内容 进行修改即可(大概在第10行左右) 

![cstcli_default](https://51k8s.oss-cn-shenzhen.aliyuncs.com/golang/images/cstcli_default.png)

