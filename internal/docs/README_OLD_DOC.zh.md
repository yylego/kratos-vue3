# vue3kratos

Vue 3 前端与 Kratos 后端集成工具包。实现无缝通信和高效端到端开发。

---

## 英文文档

[ENGLISH README](README_OLD_DOC.en.md)

---

## 项目背景

### 开发动机

做后端做久了偶尔也要做做前端。

在学前端时发现vue3用起来还行，就想着用vue3做前端连接kratos的后端服务。

因此做了个中间的胶水包，让两个语言能够更顺畅的对接。

---

## 第一步：Proto → TypeScript gRPC Client

把 kratos 的 proto 接口定义转换为 typescript 语言的 grpc 的 client 客户端代码。

### 工具链安装

#### 安装 @protobuf-ts/plugin

首先需要安装正确的工具链，访问 NPM 包页面 [@protobuf-ts/plugin](https://www.npmjs.com/package/@protobuf-ts/plugin)。

里面有详细的 Installation 指导（虽然推荐安装在临时目录里，但这里为了演示方便，我们设置在系统范围）:
```
npm install -g @protobuf-ts/plugin
```

#### 验证安装

安装完毕以后检查可执行文件路径：
```
# 确认没有安装
admin@lele-de-MacBook-Pro ~ % which protoc-gen-ts                 
protoc-gen-ts not found

# 安装到环境里
admin@lele-de-MacBook-Pro ~ % npm install -g @protobuf-ts/plugin

added 6 packages in 659ms

# 查看安装路径
admin@lele-de-MacBook-Pro ~ % which protoc-gen-ts               
/Users/admin/.nvm/versions/node/v18.17.0/bin/protoc-gen-ts
```
#### 重要提醒

⚠️ 注意：请不要使用 [protoc-gen-ts](https://www.npmjs.com/package/protoc-gen-ts) 这个同名工具。

本文档是基于 `@protobuf-ts/plugin` 编写的，使用不同的工具会导致后续步骤出错。

### 工具链环境样例

#### 样例1：MacOS 环境
```
~ % npm list -g --depth=0
~ % 输出:
/Users/admin/.nvm/versions/node/v18.17.0/lib
├── @protobuf-ts/plugin@2.9.4
├── corepack@0.18.0
├── npm@10.8.1
├── ts-node@10.9.2
└── typescript@5.4.5
```

#### 样例2：Linux 环境
```
# npm list -g --depth=0
# 输出:
/home/yangyile/.nvm/versions/node/v20.14.0/lib
├── @protobuf-ts/plugin@2.9.4
├── corepack@0.28.1
├── npm-check-updates@17.1.3
├── npm@10.9.0
└── vite@5.4.8
```

---

### 使用工具生成代码

#### 生成流程概览

现在我们要根据 proto 文件生成 typescript 的 grpc client 客户端代码。

由于不同开发者的工具路径、三方包路径和版本都可能不同，这里只能提供 Makefile 的参考配置。

#### 完整样例参考

完整的配置样例可以参考以下两个演示项目的 Makefile：
- [demo1](https://github.com/yylego/kratos-vue3-demos/blob/main/demo1kratos/Makefile)
- [demo2](https://github.com/yylego/kratos-vue3-demos/blob/main/demo2kratos/Makefile)

---

### Makefile 配置

#### 通用配置模板

下面是通用的配置模板，需要在你的 Kratos 项目 Makefile 中增加以下内容：
``` makefile
web_api_grpc_ts:
	mkdir -p ./bin/web_api_grpc_ts.out
	PROTOC_GEN_TS=$$(which protoc-gen-ts) && \
	protoc \
	--plugin=protoc-gen-ts=$$PROTOC_GEN_TS \
	--ts_out=./bin/web_api_grpc_ts.out \
	--proto_path=./api \
	--proto_path=./third_party \
	$(API_PROTO_FILES)

	PROTOC_GEN_TS=$$(which protoc-gen-ts) && \
	protoc \
	--plugin=protoc-gen-ts=$$PROTOC_GEN_TS \
	--ts_out=./bin/web_api_grpc_ts.out \
	--proto_path=./third_party \
	$(THIRD_PARTY_GOOGLE_API_PROTO_FILES)
```
##### 路径变量配置

你需要在`Makefile`找路径的逻辑里增加(这还只是mac/ubuntu系统的，windows系统的你自己也写写吧):
``` makefile
THIRD_PARTY_GOOGLE_API_PROTO_FILES=$(shell find third_party/google/api -name *.proto)
```

---

#### 配置样例1：MacOS Node v18 环境

这是某个环境的逻辑，添加到 Kratos 项目的 Makefile 中：
``` makefile
web_api_grpc_ts:
	mkdir -p ./bin/web_api_grpc_ts.out
	protoc \
	--plugin=protoc-gen-ts=/Users/admin/.nvm/versions/node/v18.17.0/bin/protoc-gen-ts \
	--ts_out=./bin/web_api_grpc_ts.out \
	--proto_path=./api \
	--proto_path=./third_party \
	$(API_PROTO_FILES)

	protoc \
	--plugin=protoc-gen-ts=/Users/admin/.nvm/versions/node/v18.17.0/bin/protoc-gen-ts \
	--ts_out=./bin/web_api_grpc_ts.out \
	--proto_path=./third_party \
	$(THIRD_PARTY_GOOGLE_API_PROTO_FILES)
```
##### 路径变量配置

你需要在`Makefile`找路径的逻辑里增加(这还只是mac/ubuntu系统的，windows系统的你自己也写写吧):
``` makefile
THIRD_PARTY_GOOGLE_API_PROTO_FILES=$(shell find third_party/google/api -name *.proto)
```

---

#### 配置样例2：Linux Node v20 环境

这是另一个环境的逻辑，添加到 Kratos 项目的 Makefile 中：
``` makefile
web_api_grpc_ts:
	mkdir -p ./bin/web_api_grpc_ts.out
	protoc \
	--plugin=protoc-gen-ts=/home/yangyile/.nvm/versions/node/v20.14.0/bin/protoc-gen-ts \
	--ts_out=./bin/web_api_grpc_ts.out \
	--proto_path=./api \
	--proto_path=./proto3ps \
	$(API_PROTO_FILES)
	
	protoc \
	--plugin=protoc-gen-ts=/home/yangyile/.nvm/versions/node/v20.14.0/bin/protoc-gen-ts \
	--ts_out=./bin/web_api_grpc_ts.out \
	--proto_path=./proto3ps \
	$(THIRD_PARTY_GOOGLE_API_PROTO_FILES)
```
##### 路径变量配置

你需要在`Makefile`找路径的逻辑里增加(这还只是mac/ubuntu系统的，windows系统的你自己也写写吧):
``` makefile
THIRD_PARTY_GOOGLE_API_PROTO_FILES=$(shell find proto3ps/google/api -name *.proto)
```

##### 配置说明

具体使用时请自己根据实际编写吧，确保工具链相同。

---

### 工具链升级

#### 版本升级操作

当需要升级 Node.js 版本时，只需更新 Makefile 中的工具路径。例如从：
```
/Users/admin/.nvm/versions/node/v18.17.0/bin/protoc-gen-ts
```
替换为新版本：
```
/Users/admin/.nvm/versions/node/v22.11.0/bin/protoc-gen-ts
```

---

### 执行代码生成

#### 第一步：创建项目

首先创建一个 Kratos 项目（这里以 helloworld 为例）：
```
kratos new helloworld
```

#### 第二步：配置 Makefile

然后修改 Makefile，按照前面的模板增加代码生成命令。

#### 第三步：执行生成命令

配置完成后，执行生成命令：
```
make web_api_grpc_ts
```

#### 第四步：验证生成结果

如果命令执行成功，就会在 `bin/web_api_grpc_ts.out` 目录下生成对应的 TypeScript 代码：
```
admin@lele-de-MacBook-Pro helloworld % make web_api_grpc_ts
admin@lele-de-MacBook-Pro helloworld %
admin@lele-de-MacBook-Pro helloworld % cd bin/web_api_grpc_ts.out
admin@lele-de-MacBook-Pro web_api_grpc_ts.out % ls
google          helloworld
admin@lele-de-MacBook-Pro web_api_grpc_ts.out % cd helloworld/v1
admin@lele-de-MacBook-Pro v1 % ls
error_reason.ts         greeter.client.ts       greeter.ts
admin@lele-de-MacBook-Pro v1 % cat greeter.client.ts

# 在生成的文件头部会看到 // @generated by protobuf-ts 2.9.4
# 这表示代码生成成功
```

#### 关于转换的说明

到这一步，我们得到的代码是基于 gRPC 协议的。如果需要在 Web 浏览器中使用（改为 HTTP 协议传输），有多种解决方案。

##### 多种方案对比

比如使用 web-grpc 代理等，但是需要额外的配置。

##### 推荐方案

本项目采用的方案是：修改客户端代码，将 gRPC 调用转换为 HTTP 请求，底层使用 `axios` 发送。

这样做的好处是：
- 保持相同的函数调用模式
- 参数和返回值都带有 TypeScript 类型
- 无需额外的桥接配置

接下来就是使用 `vue3kratos` 程序完成这个转换。

---

## 第二步：gRPC Client → HTTP Client

现在我们要把生成的 gRPC 客户端代码转换为 HTTP 请求逻辑。

### 两种转换方式

转换操作可以在两个地方进行：
- **Golang 侧**（推荐）：在后端项目的构建流程中完成转换
- **Vue 侧**：在前端项目中完成转换

---

### Golang 侧转换（推荐）

推荐在 Golang 侧完成转换，这样每次修改 proto 并重新生成代码时，转换都能自动完成。

#### 安装 vue3kratos

首先设置转换程序：
```
go install github.com/yylego/kratos-vue3/cmd/vue3kratos@latest
```

#### 使用转换工具

指定要转换的 TypeScript 客户端文件路径：
```
vue3kratos gen-grpc-via-http-in-path --grpc-ts-path=/xxx/src/rpc/rpc_admin_login/admin_login.client.ts
```

⚠️ 注意：该命令会原地修改目标文件，请确保传入正确的绝对路径。

---

### Vue 侧转换

如果需要在前端项目中完成转换，可以参考以下演示项目。

#### 演示项目参考

以下两个演示项目的 Makefile 展示了从 Kratos 项目生成 TS 客户端代码的完整流程：
- [demo1](https://github.com/yylego/kratos-vue3-demos/blob/main/demo1kratos/Makefile)
- [demo2](https://github.com/yylego/kratos-vue3-demos/blob/main/demo2kratos/Makefile)

---

### 安装运行时依赖

#### 安装 @yylego/grpt

无论使用哪种转换方式，前端项目都需要安装运行时依赖包 [@yylego/grpt](https://www.npmjs.com/package/@yylego/grpt)。

这个包负责在运行时将 gRPC 调用转换为 HTTP 请求。

安装命令：
```bash
npm install @yylego/grpt
```

#### 相关链接

- NPM：[@yylego/grpt](https://www.npmjs.com/package/@yylego/grpt)
- 源码：[yylego/grpt](https://github.com/yylego/grpt)

---

## 总结与建议

### 推荐工作流

建议采用 Golang 侧转换的方式：
1. 修改 proto 文件
2. 运行 `make web_api_grpc_ts` 生成 TypeScript gRPC 客户端
3. 运行 `vue3kratos` 程序转换为 HTTP 客户端
4. 将生成的代码提供给前端使用

这样可以确保每次 proto 变更后，前端代码都能自动同步和更新。
