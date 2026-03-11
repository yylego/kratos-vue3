# vue3kratos

Vue 3 frontends and Kratos backends integration toolkit. Enabling seamless communication and efficient end-to-end development.

---

## CHINESE README

[中文说明](README_OLD_DOC.zh.md)

---

## Project Background

### Motivation

When doing backend development long enough, sometimes need to work on frontend.

While learning frontend, I found Vue3 quite nice, so I wanted to use Vue3 frontend to connect with Kratos backend services.

Therefore, I created this intermediate glue package to enable smooth integration between the two languages.

---

## Step 1: Proto → TypeScript gRPC Client

Convert Kratos proto interface definitions into TypeScript gRPC client code.

### Toolchain Installation

#### Install @protobuf-ts/plugin

First, you need to install the correct toolchain. Visit the NPM package page [@protobuf-ts/plugin](https://www.npmjs.com/package/@protobuf-ts/plugin).

It contains detailed Installation instructions (although installing in a temp DIR is recommended, we will set it up at system-wide scope here as demo):
```
npm install -g @protobuf-ts/plugin
```

#### Verify Installation

After installation, check the executable path:
```
# Check if not installed
admin@lele-de-MacBook-Pro ~ % which protoc-gen-ts
protoc-gen-ts not found

# Install to environment
admin@lele-de-MacBook-Pro ~ % npm install -g @protobuf-ts/plugin

added 6 packages in 659ms

# Check installation path
admin@lele-de-MacBook-Pro ~ % which protoc-gen-ts
/Users/admin/.nvm/versions/node/v18.17.0/bin/protoc-gen-ts
```

#### Important Notice

⚠️ Note: Do not use [protoc-gen-ts](https://www.npmjs.com/package/protoc-gen-ts), the tool with the same name.

This document is based on `@protobuf-ts/plugin`. Using different ones can break subsequent steps.

### Toolchain Environment Examples

#### Example 1: MacOS Environment
```
~ % npm list -g --depth=0
~ % Output:
/Users/admin/.nvm/versions/node/v18.17.0/lib
├── @protobuf-ts/plugin@2.9.4
├── corepack@0.18.0
├── npm@10.8.1
├── ts-node@10.9.2
└── typescript@5.4.5
```

#### Example 2: Linux Environment
```
# npm list -g --depth=0
# Output:
/home/yangyile/.nvm/versions/node/v20.14.0/lib
├── @protobuf-ts/plugin@2.9.4
├── corepack@0.28.1
├── npm-check-updates@17.1.3
├── npm@10.9.0
└── vite@5.4.8
```

---

### Generate Code with Tools

#### Generation Flow Overview

Now we need to generate TypeScript gRPC client code based on proto files.

Since different devs have different paths, third-party package paths and versions, we can just provide reference Makefile configurations here.

#### Complete Example Reference

Complete configuration examples can be found in the Makefile of the following two demo projects:
- [demo1](https://github.com/yylego/kratos-vue3-demos/blob/main/demo1kratos/Makefile)
- [demo2](https://github.com/yylego/kratos-vue3-demos/blob/main/demo2kratos/Makefile)

---

### Makefile Configuration

#### Generic Configuration Template

Below is the generic configuration template. Add the following to the Kratos project Makefile:
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

##### Path Variable Configuration

Add the following path logic to the `Makefile` (this is mac/ubuntu, write the config on Windows):
``` makefile
THIRD_PARTY_GOOGLE_API_PROTO_FILES=$(shell find third_party/google/api -name *.proto)
```

---

#### Configuration Example 1: MacOS Node v18 Environment

This is the logic in one environment. Add this to the Kratos project Makefile:
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

##### Path Variable Configuration

Add the following path logic to the `Makefile` (this is mac/ubuntu, write the config on Windows):
``` makefile
THIRD_PARTY_GOOGLE_API_PROTO_FILES=$(shell find third_party/google/api -name *.proto)
```

---

#### Configuration Example 2: Linux Node v20 Environment

This is the logic in another environment. Add this to the Kratos project Makefile:
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

##### Path Variable Configuration

Add the following path logic to the `Makefile` (this is mac/ubuntu, write the config on Windows):
``` makefile
THIRD_PARTY_GOOGLE_API_PROTO_FILES=$(shell find proto3ps/google/api -name *.proto)
```

##### Configuration Notes

Adapt these configurations based on the current setup, ensuring the toolchain matches.

---

### Toolchain Upgrade

#### Version Upgrade Operation

When you need to upgrade Node.js version, just update the executable path in Makefile. Change from:
```
/Users/admin/.nvm/versions/node/v18.17.0/bin/protoc-gen-ts
```
To the new version:
```
/Users/admin/.nvm/versions/node/v22.11.0/bin/protoc-gen-ts
```

---

### Execute Code Generation

#### Step 1: Create Project

First, create a Kratos project (using helloworld as demo):
```
kratos new helloworld
```

#### Step 2: Configure Makefile

Then update the Makefile, adding the code generation commands based on the previous templates.

#### Step 3: Execute Generation Command

After configuration, run the generation command:
```
make web_api_grpc_ts
```

#### Step 4: Verify Generation Result

If the command executes without errors, TypeScript code gets generated in the `bin/web_api_grpc_ts.out` DIR:
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

# You will see // @generated by protobuf-ts 2.9.4 at the file head
# This indicates successful code generation
```

#### About Conversion

At this step, the generated code is based on gRPC protocol. If you need to use it in web (convert to HTTP protocol transmission), there are multiple solutions.

##### Multiple Solutions Comparison

Such as using web-grpc bridge, but extra configuration is needed.

##### Recommended Solution

This project adopts the solution: change client code to convert gRPC invocations into HTTP requests, using `axios` at the base.

Benefits of this approach:
- Keep the same function invocation pattern
- Parameters and return values have TypeScript types
- No extra bridge configuration needed

Next, we use the `vue3kratos` program to complete this conversion.

---

## Step 2: gRPC Client → HTTP Client

Now we need to convert the generated gRPC client code into HTTP request logic.

### Two Conversion Methods

Conversion can be done in two places:
- **Golang Side** (Recommended): Complete conversion in backend project build flow
- **Vue Side**: Complete conversion in frontend project

---

### Golang Side Conversion (Recommended)

It's recommended to complete conversion on Golang side, so that each time proto is modified and code is regenerated, conversion is done with automation.

#### Install vue3kratos

First set up the conversion program:
```
go install github.com/yylego/kratos-vue3/cmd/vue3kratos@latest
```

#### Use Conversion Tool

Specify the TypeScript client file path to convert:
```
vue3kratos gen-grpc-via-http-in-path --grpc-ts-path=/xxx/src/rpc/rpc_admin_login/admin_login.client.ts
```

⚠️ Note: This command changes the target file in-place. Make sure to pass the correct absolute path.

---

### Vue Side Conversion

If you need to complete conversion in the frontend project, see the following demo projects.

#### Demo Project Reference

The Makefiles of the following two demo projects show the complete flow as generating TS client code from Kratos projects:
- [demo1](https://github.com/yylego/kratos-vue3-demos/blob/main/demo1kratos/Makefile)
- [demo2](https://github.com/yylego/kratos-vue3-demos/blob/main/demo2kratos/Makefile)

---

### Install Runtime Dependencies

#### Install @yylego/grpt

Regardless of which conversion method is used, the frontend project needs to install the runtime dependency package [@yylego/grpt](https://www.npmjs.com/package/@yylego/grpt).

This package is responsible to convert gRPC calls into HTTP requests at runtime.

Installation command:
```bash
npm install @yylego/grpt
```

#### Related Links

- NPM: [@yylego/grpt](https://www.npmjs.com/package/@yylego/grpt)
- Source: [yylego/grpt](https://github.com/yylego/grpt)

---

## Summary and Recommendations

### Recommended Workflow

It's recommended to use the Golang side conversion approach:
1. Modify proto files
2. Run `make web_api_grpc_ts` to generate TypeScript gRPC client
3. Run `vue3kratos` program to convert to HTTP client
4. Provide the generated code to frontend

This ensures that each time proto changes, the frontend code can be synced and updated with automation.
