package vue3kratos

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

// grpcClientSnippet is a minimal protobuf-ts gRPC client, just enough to exercise
// each replacement rule (the import anchor, the UnaryCall return type, the stackIntercept call).
//
// grpcClientSnippet 是一段最小的 protobuf-ts gRPC 客户端，刚好覆盖每条替换规则
// （import 锚点、UnaryCall 返回类型、stackIntercept 调用）
const grpcClientSnippet = `import type { RpcOptions } from "@protobuf-ts/runtime-rpc";
export interface IFooClient {
    bar(input: BarRequest, options?: RpcOptions): UnaryCall<BarRequest, BarReply>;
}
export class FooClient implements IFooClient {
    bar(input: BarRequest, options?: RpcOptions): UnaryCall<BarRequest, BarReply> {
        return stackIntercept<BarRequest, BarReply>("unary", this._transport, method, opt, input);
    }
}
`

// requireConverted asserts the content is a converted HTTP client.
//
// requireConverted 断言内容已转换成 HTTP 客户端
func requireConverted(t *testing.T, out string) {
	t.Helper()
	require.Contains(t, out, "import { executeGrpcToHttp } from '@yylego/grpc-to-http';")
	require.Contains(t, out, "executeGrpcToHttp<")
	require.Contains(t, out, "GrpcToHttpPromise<")
	require.NotContains(t, out, "stackIntercept<")
	require.NotContains(t, out, "UnaryCall<")
}

func TestGenGrpcViaHttpInCode(t *testing.T) {
	out := GenGrpcViaHttpInCode(grpcClientSnippet)
	requireConverted(t, out)

	// Running again keeps the same result (idempotent).
	// 再跑一次结果不变（幂等）
	require.Equal(t, out, GenGrpcViaHttpInCode(out))
}

func TestGenGrpcViaHttpInPath(t *testing.T) {
	path := filepath.Join(t.TempDir(), "foo.client.ts")
	require.NoError(t, os.WriteFile(path, []byte(grpcClientSnippet), 0644))

	GenGrpcViaHttpInPath(path)

	out, err := os.ReadFile(path)
	require.NoError(t, err)
	requireConverted(t, string(out))
}

func TestGenGrpcViaHttpInRoot(t *testing.T) {
	root := t.TempDir()

	clientPath := filepath.Join(root, "sub", "foo.client.ts")
	require.NoError(t, os.MkdirAll(filepath.Dir(clientPath), 0755))
	require.NoError(t, os.WriteFile(clientPath, []byte(grpcClientSnippet), 0644))

	// A file that does not end with .client.ts must be left untouched.
	// 不以 .client.ts 结尾的文件必须保持不变
	keepPath := filepath.Join(root, "keep.ts")
	require.NoError(t, os.WriteFile(keepPath, []byte("untouched"), 0644))

	GenGrpcViaHttpInRoot(root)

	got, err := os.ReadFile(clientPath)
	require.NoError(t, err)
	requireConverted(t, string(got))

	kept, err := os.ReadFile(keepPath)
	require.NoError(t, err)
	require.Equal(t, "untouched", string(kept))
}
