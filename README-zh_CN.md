# chatcat
聊天猫 ChatGPT客户端 支持 Mac（Intel、M1、M2）、win7及以上系统

Chatcat的ChatGPT客户端旨在成为最佳的助手，为学习和工作提供支持，并让用户深入了解人工智能的魅力。在使用过程中，您将不知不觉地在AI时代中进步学习。

作为一款高效的聊天工具，Chatcat一直致力于为用户提供最佳的聊天体验。我们始终保持着不断进化和改进的态度，力争做出最优化和最出色的产品。

为此，我们将不断完善和提高聊天工具的效率和质量，努力满足用户的需求和期望。我们将通过不断地技术升级和优化，为用户提供更快速、更智能、更便捷的聊天体验。

同时，我们还将不断拓展和扩展聊天工具的功能和特性。我们将积极倾听用户的建议和反馈，尽可能地满足各种不同用户的需求和使用场景，力求做到最全面、最多样化、最灵活性的工具。

我们相信，只要我们不断努力和追求，Chatcat必将成为最好的聊天效率工具，为用户创造更多的价值。

中文文档 | [English](README.md)

<p align="center" style="text-align: center">
<img src="screenshot/use.gif" />
<img src="screenshot/home.jpg" />
<img src="screenshot/setting_general.png" />
<img src="screenshot/setting_prompt.png" />
</p>

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