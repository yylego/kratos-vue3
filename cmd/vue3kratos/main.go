package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/yylego/done"
	"github.com/yylego/kratos-vue3/vue3kratos"
	"github.com/yylego/osexistpath/osmustexist"
	"github.com/yylego/zaplog"
	"go.uber.org/zap"
)

func main() {
	var rootCmd = &cobra.Command{
		Use:   "vue3kratos",
		Short: "A CLI app to the vue3kratos chain.",
		Long:  `vue3kratos is a command-line app that helps bridge the gap between Kratos gRPC backends and Vue 3 frontends through converting generated TypeScript clients to be web-safe.`,
	}

	// example: vue3kratos gen-grpc-via-http-in-root --grpc-ts-root=/xxx/src/rpc
	rootCmd.AddCommand(makeRootCmd())

	// example: vue3kratos gen-grpc-via-http-in-path --grpc-ts-path=/xxx/src/rpc/rpc_admin_login/admin_login.client.ts
	rootCmd.AddCommand(makePathCmd())

	// example: vue3kratos gen-grpc-via-http-in-code --grpc-ts-code=/xxx/src/rpc/rpc_admin_login/admin_login.client.ts
	rootCmd.AddCommand(makeCodeCmd())

	// Execute root command
	done.Done(rootCmd.Execute())
}

// makeRootCmd creates the root command that converts each .client.ts file in a DIR
// Searches through the DIR and subdirectories to find and convert each gRPC client file
//
// makeRootCmd 创建 root 命令递归转换 DIR 中的所有 .client.ts 文件
// 搜索 DIR 和子目录查找并转换每个 gRPC 客户端文件
func makeRootCmd() *cobra.Command {
	var grpcTsRoot string
	var genCmd = &cobra.Command{
		Use:   "gen-grpc-via-http-in-root",
		Short: "Finds and converts each .client.ts file in a DIR.",
		Long:  `Searches files ending with .client.ts within the specified root DIR and its subdirectories, then converts each file from a gRPC client to an HTTP client.`,
		Run: func(cmd *cobra.Command, args []string) {
			vue3kratos.GenGrpcViaHttpInRoot(grpcTsRoot)
			zaplog.LOG.Info("✅ Success! Each found .client.ts file has been converted.", zap.String("root", grpcTsRoot))
		},
	}
	genCmd.Flags().StringVarP(&grpcTsRoot, "grpc-ts-root", "r", "", "Root DIR with .client.ts files.")
	done.Done(genCmd.MarkFlagRequired("grpc-ts-root"))
	return genCmd
}

// makePathCmd creates the path command that converts a single .client.ts file
// Modifies the specified file in place, replacing gRPC code with HTTP code
//
// makePathCmd 创建 path 命令转换单个 .client.ts 文件
// 就地修改指定文件，将 gRPC 代码替换为 HTTP 代码
func makePathCmd() *cobra.Command {
	var grpcTsPath string
	var genCmd = &cobra.Command{
		Use:   "gen-grpc-via-http-in-path",
		Short: "Converts one .client.ts file.",
		Long:  `Converts one specified .client.ts file from a gRPC client to an HTTP client, modifying it in place.`,
		Run: func(cmd *cobra.Command, args []string) {
			vue3kratos.GenGrpcViaHttpInPath(grpcTsPath)
			zaplog.LOG.Info("✅ Success! File has been converted.", zap.String("path", grpcTsPath))
		},
	}
	genCmd.Flags().StringVarP(&grpcTsPath, "grpc-ts-path", "p", "", "Path to the target .client.ts file.")
	done.Done(genCmd.MarkFlagRequired("grpc-ts-path"))
	return genCmd
}

// makeCodeCmd creates the code command that converts and prints to stdout
// Reads the file, performs conversion in RAM, and outputs without modifying the source
//
// makeCodeCmd 创建 code 命令转换并打印到标准输出
// 读取文件，在内存中转换，输出结果但不修改源文件
func makeCodeCmd() *cobra.Command {
	var grpcTsCodePath string
	var genCmd = &cobra.Command{
		Use:   "gen-grpc-via-http-in-code",
		Short: "Converts and prints the content of a .client.ts file to stdout.",
		Long:  `Reads a .client.ts file, performs the conversion in RAM, and prints the resulting code to standard output without modifying the source file.`,
		Run: func(cmd *cobra.Command, args []string) {
			clientPath := osmustexist.FILE(grpcTsCodePath)
			srcContent := done.VAE(os.ReadFile(clientPath)).Nice()
			newContent := vue3kratos.GenGrpcViaHttpInCode(string(srcContent))
			fmt.Println(newContent)
		},
	}
	genCmd.Flags().StringVarP(&grpcTsCodePath, "grpc-ts-code", "c", "", "Path to the .client.ts file to read and convert.")
	done.Done(genCmd.MarkFlagRequired("grpc-ts-code"))
	return genCmd
}
