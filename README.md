<!-- TEMPLATE (EN) BEGIN: BADGES -->

[![GitHub Workflow Status (branch)](https://img.shields.io/github/actions/workflow/status/yylego/kratos-vue3/release.yml?branch=main&label=BUILD)](https://github.com/yylego/kratos-vue3/actions/workflows/release.yml?query=branch%3Amain)
[![GoDoc](https://pkg.go.dev/badge/github.com/yylego/kratos-vue3)](https://pkg.go.dev/github.com/yylego/kratos-vue3)
[![Coverage Status](https://img.shields.io/coveralls/github/yylego/kratos-vue3/main.svg)](https://coveralls.io/github/yylego/kratos-vue3?branch=main)
[![Supported Go Versions](https://img.shields.io/badge/Go-1.24%2B-lightgrey.svg)](https://go.dev/)
[![GitHub Release](https://img.shields.io/github/release/yylego/kratos-vue3.svg)](https://github.com/yylego/kratos-vue3/releases)
[![Go Report Card](https://goreportcard.com/badge/github.com/yylego/kratos-vue3)](https://goreportcard.com/report/github.com/yylego/kratos-vue3)

<!-- TEMPLATE (EN) CLOSE: BADGES -->

# kratos-vue3

> Vue 3 frontends and Kratos backends integration toolkit
> Generate type-safe TypeScript clients from Go Kratos proto files
> Seamless TypeScript + Go RPC integration, use `@protobuf-ts/plugin`
> Complete type-safe code from backend to frontend
> Invoke backend APIs just as native functions

⚠️ **Replaced with [kratos-ts-http](https://github.com/yylego/kratos-ts-http).** This package rewrites the clients in a second pass and gets no more updates; it does not handle stream methods. New and existing projects should move to kratos-ts-http, which folds the rewrite into generation. The code here stays as it is.

---

![grpc-to-http-overview](https://raw.githubusercontent.com/yylego/grpc-to-http/main/assets/grpc-to-http-overview.svg)

---

<!-- TEMPLATE (EN) BEGIN: LANGUAGE NAVIGATION -->

## CHINESE README

[中文说明](README.zh.md)

<!-- TEMPLATE (EN) CLOSE: LANGUAGE NAVIGATION -->

## Core Architecture

`vue3kratos` bridges Go backends and Vue 3 frontends with TypeScript client.

### Development Toolchain

```
+-------------+    +----------+    +---------------+    +--------------+    +---------------+
| .proto files| -> | protoc   | -> | gRPC TS Client| -> | vue3kratos   | -> | HTTP TS Client|
|             |    | + plugin |    |               |    | CLI Convert  |    |               |
+-------------+    +----------+    +---------------+    +--------------+    +---------------+
```

## 🌟 Highlights

- **Auto Code Generation**: Generate clean TypeScript clients from proto files
- **No Handwriting**: Forget about handwriting API clients
- **Single Command Convert**: Transform gRPC clients to HTTP with one command
- **Web Compatible**: Works in web without gRPC issues
- **Complete Type-Safe**: Type checking all the way from backend to frontend
- **IDE Autocompletion**: Rich development experience with smart suggestions
- **Makefile Integration**: Simple integration into existing build process
- **CI/CD Pipeline**: Smooth workflow automation support
- **Axios HTTP Clients**: Modern HTTP client implementation
- **Native Function Experience**: Invoke APIs just as native functions

## Related Projects

- [grpc-to-http](https://github.com/yylego/grpc-to-http) — npm package [`@yylego/grpc-to-http`](https://www.npmjs.com/package/@yylego/grpc-to-http), converts protobuf-ts gRPC invocations to HTTP/REST requests via Axios
- [kratos-examples/vue3](https://github.com/kratos-examples/vue3) — Complete demo projects with backend and frontend integration

## 🚀 Quick Start

Experience the magic of `vue3kratos` in minutes with the demo projects.

Visit the separate demo repo: **[kratos-examples/vue3](https://github.com/kratos-examples/vue3)**

Follow the instructions in the demo repo to run the complete examples.

## The `vue3kratos` Flow Explained

Using `vue3kratos` in the project involves just a few steps.

### Step 1: Setup the Chain

Ensure you have `@protobuf-ts/plugin` installed, and the `vue3kratos` Go CLI app.

```bash
# Setup the protobuf-ts plugin (use this one, NOT others)
npm install -g @protobuf-ts/plugin

# Check the installation
which protoc-gen-ts

# Setup the vue3kratos CLI
go install github.com/yylego/kratos-vue3/cmd/vue3kratos@latest
```

#### Example Development Environment

```bash
# Example dependencies in system-wide
npm list -g --depth=0
├── @protobuf-ts/plugin@2.11.1
├── typescript@5.4.5
├── vite@5.4.8
```

### Step 2: Generate the gRPC Client

A Kratos v3 project ships the `buf` toolchain, so generate the TypeScript client with `buf` plus `protoc-gen-ts`. Add a `buf.gen.ts.yaml` template next to the project's `buf.yaml`:

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

Then drive it from a `Makefile` target:

```makefile
web_api_grpc_ts:
	mkdir -p ./bin/web_api_grpc_ts.out
	buf generate --template buf.gen.ts.yaml --include-imports
```

**`opt` options explained:**

| Option             | Effect                                                                                                |
| ------------------ | ----------------------------------------------------------------------------------------------------- |
| `ts_nocheck`       | Adds `// @ts-nocheck` to generated files, suppresses TypeScript type checking on auto-generated code  |
| `eslint_disable`   | Adds `/* eslint-disable */` to generated files, prevents ESLint from flagging auto-generated code     |
| `long_type_string` | Uses `string` instead of `bigint` on int64/uint64 fields, avoids `bigint` issues in some environments |

`--include-imports` makes `buf` emit the imported google well-known types too, so the generated client's imports resolve. `buf` pulls the google protos from its registry (declare `buf.build/googleapis/googleapis` under `deps` in `buf.yaml`), so no local `third_party` proto tree is needed.

### Step 3: Convert to HTTP Client

Use the `vue3kratos` CLI to convert the gRPC client file generated in the previous step into an HTTP client.

```bash
vue3kratos gen-grpc-via-http-in-path  --grpc-ts-path=/path/to/the/generated.client.ts
```

This command modifies the target file and transforms gRPC invocations into Axios HTTP requests.

### Step 4: Use in Vue

In the Vue project, setup the `@yylego/grpc-to-http` runtime module. Then you can invoke APIs just as native functions.

```bash
npm install @yylego/grpc-to-http
```

> npm module URL: [@yylego/grpc-to-http](https://www.npmjs.com/package/@yylego/grpc-to-http)

```typescript
// Demo example from kratos-examples/vue3 repo
import { GrpcWebFetchTransport } from "@protobuf-ts/grpcweb-transport";
import { RpcpingClient } from "./rpc/rpcping/rpcping.client";
import { StringValue } from "./rpc/google/protobuf/wrappers";

// Create transport instance
const demoTransport = new GrpcWebFetchTransport({
  baseUrl: "http://127.0.0.1:28000",
  meta: {
    Authorization: "TOKEN-888",
  },
});

const rpcpingClient = new RpcpingClient(demoTransport);

// Invoke API example
async function demoPing() {
  const request = StringValue.create({
    value: "Hello from Vue3 Kratos!",
  });

  const response = await rpcpingClient.ping(request, {});
  console.log("Ping success:", response.data.value);
  return response.data.value;
}
```

---

## 🔁 Demo Projects

See the complete working examples in the separate repo:

**[kratos-examples/vue3](https://github.com/kratos-examples/vue3)** - Complete demo projects with backend and frontend integration

The Makefiles in the demo projects show the complete flow:

- [demo1kratos Makefile](https://github.com/kratos-examples/vue3/blob/main/demo1kratos/Makefile)
- [demo2kratos Makefile](https://github.com/kratos-examples/vue3/blob/main/demo2kratos/Makefile)

The [`examples`](internal/examples) DIR contains these examples:

- [`example1`](internal/examples/example1) - Service demo
- [`example2`](internal/examples/example2) - Service demo

These examples show proto-based test data generation.

---

## ✅ Feature Overview

- Generate type-safe TypeScript gRPC clients from Kratos proto files
- Support automatic conversion to HTTP requests (Axios-based)
- Complete type-safe with IDE autocomplete and checking
- Simple integration into Makefiles and CI/CD pipelines
- Web-compatible HTTP clients with direct frontend use

---

## 💡 Target Audience

- Developers using Kratos as the backend
- Frontend developers using Vue 3 to invoke backend services
- Developers who want complete type-safe client-backend integration

---

## Setup Guide

[SETUP GUIDE](internal/docs/SETUP_GUIDE.md)

---

<!-- TEMPLATE (EN) BEGIN: STANDARD PROJECT FOOTER -->
<!-- VERSION 2025-11-25 03:52:28.131064 +0000 UTC -->

## 📄 License

MIT License - see [LICENSE](LICENSE).

---

## 💬 Contact & Feedback

Contributions are welcome! Report bugs, suggest features, and contribute code:

- 🐛 **Mistake reports?** Open an issue on GitHub with reproduction steps
- 💡 **Fresh ideas?** Create an issue to discuss
- 📖 **Documentation confusing?** Report it so we can improve
- 🚀 **Need new features?** Share the use cases to help us understand requirements
- ⚡ **Performance issue?** Help us optimize through reporting slow operations
- 🔧 **Configuration problem?** Ask questions about complex setups
- 📢 **Follow project progress?** Watch the repo to get new releases and features
- 🌟 **Success stories?** Share how this package improved the workflow
- 💬 **Feedback?** We welcome suggestions and comments

---

## 🔧 Development

New code contributions, follow this process:

1. **Fork**: Fork the repo on GitHub (using the webpage UI).
2. **Clone**: Clone the forked project (`git clone https://github.com/yourname/repo-name.git`).
3. **Navigate**: Navigate to the cloned project (`cd repo-name`)
4. **Branch**: Create a feature branch (`git checkout -b feature/xxx`).
5. **Code**: Implement the changes with comprehensive tests
6. **Testing**: (Golang project) Ensure tests pass (`go test ./...`) and follow Go code style conventions
7. **Documentation**: Update documentation to support client-facing changes
8. **Stage**: Stage changes (`git add .`)
9. **Commit**: Commit changes (`git commit -m "Add feature xxx"`) ensuring backward compatible code
10. **Push**: Push to the branch (`git push origin feature/xxx`).
11. **PR**: Open a merge request on GitHub (on the GitHub webpage) with detailed description.

Please ensure tests pass and include relevant documentation updates.

---

## 🌟 Support

Welcome to contribute to this project via submitting merge requests and reporting issues.

**Project Support:**

- ⭐ **Give GitHub stars** if this project helps you
- 🤝 **Share with teammates** and (golang) programming friends
- 📝 **Write tech blogs** about development tools and workflows - we provide content writing support
- 🌟 **Join the ecosystem** - committed to supporting open source and the (golang) development scene

**Have Fun Coding with this package!** 🎉🎉🎉

<!-- TEMPLATE (EN) CLOSE: STANDARD PROJECT FOOTER -->
