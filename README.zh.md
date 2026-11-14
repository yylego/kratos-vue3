<!-- TEMPLATE (ZH) BEGIN: BADGES -->

[![GitHub Workflow Status (branch)](https://img.shields.io/github/actions/workflow/status/yylego/kratos-vue3/release.yml?branch=main&label=BUILD)](https://github.com/yylego/kratos-vue3/actions/workflows/release.yml?query=branch%3Amain)
[![GoDoc](https://pkg.go.dev/badge/github.com/yylego/kratos-vue3)](https://pkg.go.dev/github.com/yylego/kratos-vue3)
[![Coverage Status](https://img.shields.io/coveralls/github/yylego/kratos-vue3/main.svg)](https://coveralls.io/github/yylego/kratos-vue3?branch=main)
[![Supported Go Versions](https://img.shields.io/badge/Go-1.24%2B-lightgrey.svg)](https://go.dev/)
[![GitHub Release](https://img.shields.io/github/release/yylego/kratos-vue3.svg)](https://github.com/yylego/kratos-vue3/releases)
[![Go Report Card](https://goreportcard.com/badge/github.com/yylego/kratos-vue3)](https://goreportcard.com/report/github.com/yylego/kratos-vue3)

<!-- TEMPLATE (ZH) CLOSE: BADGES -->

# kratos-vue3

> Vue 3 前端 + Kratos 后端集成开发工具包
> 从 Go Kratos proto 文件生成类型安全的 TypeScript 客户端
> 无缝的 TypeScript + Go RPC 集成，使用 `@protobuf-ts/plugin`
> 从后端到前端的完整类型安全代码
> 像调用原生函数一样调用后端 API

⚠️ **已被 [kratos-ts-http](https://github.com/yylego/kratos-ts-http) 取代。** 本包要分两步改写客户端，已不再更新，也不支持流式方法。新老项目都请换用 kratos-ts-http，它把改写直接做进了生成插件。这里的代码照原样保留。

---

![grpc-to-http-overview](https://raw.githubusercontent.com/yylego/grpc-to-http/main/assets/grpc-to-http-overview.svg)

---

<!-- TEMPLATE (ZH) BEGIN: LANGUAGE NAVIGATION -->

## 英文文档

[ENGLISH README](README.md)

<!-- TEMPLATE (ZH) CLOSE: LANGUAGE NAVIGATION -->

## 核心架构

`vue3kratos` 连接 Go 后端与 Vue 3 前端，生成 TypeScript 客户端。

### 开发工具链

```
+-------------+    +----------+    +---------------+    +--------------+    +---------------+
| .proto 文件 | -> | protoc   | -> | gRPC TS 客户端| -> | vue3kratos   | -> | HTTP TS 客户端|
|             |    | + 插件   |    |               |    | CLI 转换工具 |    |               |
+-------------+    +----------+    +---------------+    +--------------+    +---------------+
```

## 🌟 核心特性

- **自动代码生成**: 从 proto 文件生成简洁的 TypeScript 客户端
- **告别手写**: 忘掉手动编写 API 客户端代码
- **一键转换**: 单条命令将 gRPC 客户端转换成 HTTP 客户端
- **Web 兼容**: 直接在 Web 中使用，无需 gRPC 问题
- **完整类型安全**: 从后端到前端全程类型检查
- **IDE 智能提示**: 丰富的开发体验，智能代码补全
- **Makefile 集成**: 简单集成到现有构建流程
- **CI/CD 管线**: 平滑的工作流自动化支持
- **Axios HTTP 客户端**: 现代化的 HTTP 客户端实现
- **原生函数体验**: 像调用原生函数一样调用 API

## 关联项目

- [grpc-to-http](https://github.com/yylego/grpc-to-http) — npm 包 [`@yylego/grpc-to-http`](https://www.npmjs.com/package/@yylego/grpc-to-http)，将 protobuf-ts gRPC 调用转换为基于 Axios 的 HTTP/REST 请求
- [kratos-examples/vue3](https://github.com/kratos-examples/vue3) — 完整的演示项目，包含后端和前端集成

## 🚀 快速上手

在几分钟内，通过演示项目亲身体验 `vue3kratos` 的强大之处。

访问独立的演示代码：**[kratos-examples/vue3](https://github.com/kratos-examples/vue3)**

按照演示代码中的说明运行完整示例。

## `vue3kratos` 工作流详解

在项目中使用 `vue3kratos` 只需以下几步。

### 第1步: 配置链

确保已在系统安装 `@protobuf-ts/plugin`，并已设置 `vue3kratos` 的 Go CLI 应用。

```bash
# 配置 protobuf-ts 插件（使用这个，而不是其他插件）
npm install -g @protobuf-ts/plugin

# 检查是否配置成功
which protoc-gen-ts

# 配置 vue3kratos CLI
go install github.com/yylego/kratos-vue3/cmd/vue3kratos@latest
```

#### 基础开发环境示例

```bash
# 示例系统范围依赖
npm list -g --depth=0
├── @protobuf-ts/plugin@2.11.1
├── typescript@5.4.5
├── vite@5.4.8
```

### 第2步: 生成 gRPC 客户端

Kratos v3 项目自带 `buf` 工具链，所以用 `buf` 配合 `protoc-gen-ts` 生成 TypeScript 客户端。在项目 `buf.yaml` 旁边加一个 `buf.gen.ts.yaml` 模板：

```yaml
version: v2
inputs:
  - directory: api
plugins:
  - local: protoc-gen-ts
    out: bin/web_api_grpc_ts.out
    opt:
      - ts_nocheck
      - eslint_disable
      - long_type_string
```

再用 `Makefile` target 驱动它：

```makefile
web_api_grpc_ts:
	mkdir -p ./bin/web_api_grpc_ts.out
	buf generate --template buf.gen.ts.yaml --include-imports
```

**`opt` 各项配置说明：**

| 选项               | 作用                                                                                    |
| ------------------ | --------------------------------------------------------------------------------------- |
| `ts_nocheck`       | 在生成文件顶部添加 `// @ts-nocheck`，跳过 TypeScript 类型检查（自动生成的代码无需检查） |
| `eslint_disable`   | 在生成文件顶部添加 `/* eslint-disable */`，防止 ESLint 对自动生成的代码报错             |
| `long_type_string` | int64/uint64 字段使用 `string` 代替 `bigint`，避免部分环境下 `bigint` 兼容问题          |

`--include-imports` 让 `buf` 把 import 进来的 google 标准类型也一并生成，这样生成的客户端 import 才能解析。google protos 由 `buf` 从它的 registry 拉取（在 `buf.yaml` 的 `deps` 里声明 `buf.build/googleapis/googleapis`），所以**不再需要本地的 `third_party` proto 目录**。

### 第3步: 转换为 HTTP 客户端

使用 `vue3kratos` CLI 将上一步生成的 gRPC 客户端文件转换为 HTTP 客户端。

```bash
vue3kratos gen-grpc-via-http-in-path --grpc-ts-path=/path/to/the/generated.client.ts
```

该命令会修改目标文件并将 gRPC 调用转换成基于 Axios 的 HTTP 请求。

### 第4步: 在 Vue 中使用

在 Vue 项目中，配置 `@yylego/grpc-to-http` 运行时模块，然后就可以像调用原生函数一样，调用所有 API。

```bash
npm install @yylego/grpc-to-http
```

> npm 模块地址：[@yylego/grpc-to-http](https://www.npmjs.com/package/@yylego/grpc-to-http)

```typescript
// 来自 kratos-examples/vue3 演示代码
import { GrpcWebFetchTransport } from "@protobuf-ts/grpcweb-transport";
import { RpcpingClient } from "./rpc/rpcping/rpcping.client";
import { StringValue } from "./rpc/google/protobuf/wrappers";

// 创建传输实例
const demoTransport = new GrpcWebFetchTransport({
  baseUrl: "http://127.0.0.1:28000",
  meta: {
    Authorization: "TOKEN-888",
  },
});

const rpcpingClient = new RpcpingClient(demoTransport);

// API 调用示例
async function demoPing() {
  const request = StringValue.create({
    value: "Hello from Vue3 Kratos!",
  });

  const response = await rpcpingClient.ping(request, {});
  console.log("Ping 成功:", response.data.value);
  return response.data.value;
}
```

---

## 🔁 示例项目

在独立代码中查看完整的工作示例：

**[kratos-examples/vue3](https://github.com/kratos-examples/vue3)** - 完整的演示项目，包含后端和前端集成

演示项目的 Makefile 展示了完整流程：

- [demo1kratos Makefile](https://github.com/kratos-examples/vue3/blob/main/demo1kratos/Makefile)
- [demo2kratos Makefile](https://github.com/kratos-examples/vue3/blob/main/demo2kratos/Makefile)

在仓库的 [`examples`](internal/examples) 中还包含这些示例：

- [`example1`](internal/examples/example1) - 服务示例
- [`example2`](internal/examples/example2) - 服务示例

这些示例展示基于 proto 的测试数据生成。

---

## ✅ 特性概览

- 将 Kratos proto 文件生成成类型安全的 TypeScript gRPC 客户端
- 支持自动转换成 HTTP 请求形式（Axios 封装）
- 支持类型提示与自动补全，前端集成体验优秀
- 可集成至 Makefile 和 CI/CD 中
- Web 兼容的 HTTP 客户端，支持直接前端调用

---

## 💡 适合人群

- 使用 Kratos 作为后端的开发者
- 想通过 Vue 3 编写前端，调用后端服务
- 希望实现"完整类型安全"客户端和后端集成的开发者

---

## 配置指南

[配置指南](internal/docs/SETUP_GUIDE.zh.md)

---

<!-- TEMPLATE (ZH) BEGIN: STANDARD PROJECT FOOTER -->
<!-- VERSION 2025-11-25 03:52:28.131064 +0000 UTC -->

## 📄 许可证类型

MIT 许可证 - 详见 [LICENSE](LICENSE)。

---

## 💬 联系与反馈

非常欢迎贡献代码！报告 BUG、建议功能、贡献代码：

- 🐛 **问题报告？** 在 GitHub 上提交问题并附上重现步骤
- 💡 **新颖思路？** 创建 issue 讨论
- 📖 **文档疑惑？** 报告问题，帮助我们完善文档
- 🚀 **需要功能？** 分享使用场景，帮助理解需求
- ⚡ **性能瓶颈？** 报告慢操作，协助解决性能问题
- 🔧 **配置困扰？** 询问复杂设置的相关问题
- 📢 **关注进展？** 关注仓库以获取新版本和功能
- 🌟 **成功案例？** 分享这个包如何改善工作流程
- 💬 **反馈意见？** 欢迎提出建议和意见

---

## 🔧 代码贡献

新代码贡献，请遵循此流程：

1. **Fork**：在 GitHub 上 Fork 仓库（使用网页界面）
2. **克隆**：克隆 Fork 的项目（`git clone https://github.com/yourname/repo-name.git`）
3. **导航**：进入克隆的项目（`cd repo-name`）
4. **分支**：创建功能分支（`git checkout -b feature/xxx`）
5. **编码**：实现您的更改并编写全面的测试
6. **测试**：（Golang 项目）确保测试通过（`go test ./...`）并遵循 Go 代码风格约定
7. **文档**：面向用户的更改需要更新文档
8. **暂存**：暂存更改（`git add .`）
9. **提交**：提交更改（`git commit -m "Add feature xxx"`）确保向后兼容的代码
10. **推送**：推送到分支（`git push origin feature/xxx`）
11. **PR**：在 GitHub 上打开 Merge Request（在 GitHub 网页上）并提供详细描述

请确保测试通过并包含相关的文档更新。

---

## 🌟 项目支持

非常欢迎通过提交 Merge Request 和报告问题来贡献此项目。

**项目支持：**

- ⭐ **给予星标**如果项目对您有帮助
- 🤝 **分享项目**给团队成员和（golang）编程朋友
- 📝 **撰写博客**关于开发工具和工作流程 - 我们提供写作支持
- 🌟 **加入生态** - 致力于支持开源和（golang）开发场景

**祝你用这个包编程愉快！** 🎉🎉🎉

<!-- TEMPLATE (ZH) CLOSE: STANDARD PROJECT FOOTER -->
