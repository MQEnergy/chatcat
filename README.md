<p align="center" style="text-align: center">
<img src="assets/chatcat.png" width="180" height="180" />
</p>

# chatcat
ChatCat ChatGPT Client Supports Mac (Intel, M1, M2), win7 and above systems

The ChatGPT client aims to be the best assistant, providing support for learning and work, and allowing users to deeply understand the charm of artificial intelligence. In the process of using it, you will unconsciously progress and learn in the AI era.

As an efficient chat tool, Chatcat has always been committed to providing the best chat experience for users. We always maintain an attitude of continuous evolution and improvement, striving to make the most optimized and outstanding products.

To this end, we will constantly improve and enhance the efficiency and quality of the chat tool, striving to meet the needs and expectations of users. Through continuous technological upgrades and optimization, we will provide users with faster, smarter and more convenient chat experience.

At the same time, we will continue to expand and extend the features and characteristics of the chat tool. We will actively listen to user feedback and suggestions, as much as possible to meet the needs and usage scenarios of different users, and strive to make the most comprehensive, diverse and flexible tool.

We believe that as long as we continue to work hard and pursue, Chatcat will surely become the best chat efficiency tool, creating more value for users.

English | [中文文档](README-zh_CN.md)

<p align="center" style="text-align: center">
<img src="screenshot/use.gif" />
<img src="screenshot/home.jpg" />
<img src="screenshot/setting_general.png" />
<img src="screenshot/setting_prompt.png" />
</p>

## Technology Stack
Vue3 + arco.design + golang + wails + sqlite

## Instructions
Install Wails
```
go install github.com/wailsapp/wails/v2/cmd/wails@latest
```
Note:
The official Wails packaging cannot run on win7. To support win7, you need to download it from the following link and compile and install it locally.
```
https://github.com/MQEnergy/wails
```

### Mac Packaging
```shell
# 
wails build
```

### Windows Packaging
```shell
# md64 refers to X86 architecture CPU, 64-bit version. X86_64 is also known as X86_64. Mainstream desktop PCs, laptops, servers (including virtual machines) are all using X86_64 CPUs.
wails build -platform=windows/amd64

# Arm64 refers to ARM architecture CPU, 64-bit version. Apple's new computers are using ARM architecture CPUs. Some routers and embedded devices use Arm64 CPUs. The CPU most commonly used in mobile phones and Android tablets is also ARM architecture.
wails build -platform=windows/arm64
```