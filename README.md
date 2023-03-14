# chatcat
聊天猫 ChatGPT客户端 

支持 Mac（Intel、M1、M2）、win7及以上系统

## 技术栈
vue3 + arco.design + golang + wails + sqlite

## 使用说明
安装wails
```
go install github.com/wailsapp/wails/v2/cmd/wails@latest
```
注意：
官方的wails打包 不能在win7下运行，需要在以下链接下载，并本地编译安装才支持win7
```
https://github.com/MQEnergy/wails
```

### mac打包
```shell
# 
wails build
```

### window打包
```shell
# md64是X86架构的CPU，64位版。amd64又叫X86_64。主流的桌面PC，笔记本电脑，服务器（包括虚拟机）都在用X86_64的CPU。
wails build -platform=windows/amd64

# arm64是ARM架构的CPU，64位版。苹果新出的电脑在用ARM架构的CPU。有些路由器和嵌入式设备在用arm64的CPU。手机和安卓平板电脑最常用的CPU也是ARM架构的。
wails build -platform=windows/arm64
```