package vue3kratos

import (
	"os"
	"strings"

	"github.com/yylego/kratos-vue3/internal/utils"
	"github.com/yylego/done"
	"github.com/yylego/must"
	"github.com/yylego/osexistpath/osmustexist"
	"github.com/yylego/zaplog"
	"go.uber.org/zap"
)

// GenGrpcViaHttpInRoot converts gRPC TS clients to HTTP clients in a DIR
// 在整个 DIR 中查找 ts grpc client 代码，把它们转换成使用 http 请求
// This modifies files ending with ".client.ts" in place, use with caution
// 这会修改以 ".client.ts" 结尾的文件内容，使用时需谨慎
func GenGrpcViaHttpInRoot(grpcTsRoot string) {
	zaplog.LOG.Info("gen-grpc-via-http", zap.String("grpc-ts-root", grpcTsRoot))
	osmustexist.MustRoot(grpcTsRoot)

	done.Done(utils.WalkFiles(grpcTsRoot, func(path string, info os.FileInfo) error {
		zaplog.LOG.Info("walk", zap.String("path", path))
		if strings.HasSuffix(path, ".client.ts") {
			zaplog.LOG.Info("walk", zap.String("name", info.Name()))
			GenGrpcViaHttpInPath(path)
		}
		return nil
	}))
}

// GenGrpcViaHttpInPath converts gRPC TS client to HTTP client in a file
// 在文件中查找 ts grpc client 代码，把它们转换成使用 http 请求
// Writes the transformed content back to the same file
// 转换完成把新内容写回到同一个文件中
func GenGrpcViaHttpInPath(grpcTsPath string) {
	zaplog.LOG.Info("gen-grpc-via-http", zap.String("code_path", grpcTsPath))
	must.Nice(grpcTsPath)

	// Read file content
	// 读取文件内容
	srcContent := string(done.VAE(os.ReadFile(grpcTsPath)).Nice())

	// Transform content
	// 转换内容
	newContent := GenGrpcViaHttpInCode(srcContent)

	// Write back to file
	// 写回到文件
	done.Done(os.WriteFile(grpcTsPath, []byte(newContent), 0644))
	zaplog.LOG.Info("content replace success!!!")
}

// GenGrpcViaHttpInCode converts gRPC TS client code to HTTP client code
// 在代码中查找 ts grpc client 代码，把它们转换成使用 http 请求
// Returns the transformed code, idempotent when running multiple times
// 返回转换后的代码，多次运行时保持幂等性
func GenGrpcViaHttpInCode(srcContent string) string {
	newContent := srcContent

	// Replace gRPC invocations with HTTP equivalents
	// 把代码中调用 grpc 的地方改成调用 http
	newContent = strings.ReplaceAll(newContent, "stackIntercept<", "executeGrpcToHttp<")
	newContent = strings.ReplaceAll(newContent, "UnaryCall<", "GrpcToHttpPromise<")

	// Check if target imports exist
	// 判断是否存在目标引用
	targetImport := `import { executeGrpcToHttp } from '@yylego/grpc-to-http';` +
		"\n" +
		`import type { GrpcToHttpPromise } from '@yylego/grpc-to-http';`
	searchImport := `import type { RpcOptions } from "@protobuf-ts/runtime-rpc";`

	if !strings.Contains(newContent, targetImport) {
		// Find the position to insert new imports
		// 找到插入新引用的位置
		targetIndex := strings.Index(newContent, searchImport)
		if targetIndex != -1 {
			// Insert new imports following the target import
			// 在目标引用之后添加新引用
			insertIndex := targetIndex + len(searchImport)

			// Concatenate the outcome
			// 拼接最终结果
			newContent = newContent[:insertIndex] + "\n" + targetImport + newContent[insertIndex:]
		}
	}
	return newContent
}

// CloneFilesToDestRoot copies files from source DIR to target DIR
// 把源 DIR 的文件克隆到目标 DIR 中
func CloneFilesToDestRoot(sourceRoot string, targetRoot string) {
	utils.CopyFiles(sourceRoot, targetRoot)
}
