package example2

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/yylego/kratos-vue3/vue3kratos"
	"github.com/stretchr/testify/require"
	"github.com/yylego/done"
	"github.com/yylego/must"
	"github.com/yylego/runpath"
)

func TestGenGrpcViaHttpInCode(t *testing.T) {
	// Read gRPC client code as input
	// 读取 gRPC 客户端代码作为输入
	sourcePath := runpath.PARENT.Join("testdata", "echo.client.ts-source.txt")
	srcContent := done.VAE(os.ReadFile(sourcePath)).Nice()
	t.Log("Source content length:", len(srcContent))

	// Read HTTP client code as expected output
	// 读取 HTTP 客户端代码作为期望输出
	targetPath := runpath.PARENT.Join("testdata", "echo.client.ts-target.txt")
	resContent := done.VAE(os.ReadFile(targetPath)).Nice()
	t.Log("Target content length:", len(resContent))

	// Test conversion
	// 测试转换功能
	newContent := vue3kratos.GenGrpcViaHttpInCode(string(srcContent))

	// Check output matches expected
	// 检查输出匹配期望
	require.Equal(t, string(resContent), newContent, "Converted content should match target output")
}

func TestGenGrpcViaHttpInPath(t *testing.T) {
	// Read test data
	// 读取测试数据
	sourcePath := runpath.PARENT.Join("testdata", "echo.client.ts-source.txt")
	srcContent := done.VAE(os.ReadFile(sourcePath)).Nice()

	targetPath := runpath.PARENT.Join("testdata", "echo.client.ts-target.txt")
	resContent := done.VAE(os.ReadFile(targetPath)).Nice()

	// Create temp file using modern API
	// 使用现代 API 创建临时文件
	tempRoot, err := os.MkdirTemp("", "vue3kratos_test_*")
	require.NoError(t, err)
	defer func() {
		must.Done(os.RemoveAll(tempRoot))
	}()

	testPath := filepath.Join(tempRoot, "test.client.ts")
	done.Done(os.WriteFile(testPath, srcContent, 0644))

	// Execute conversion on file
	// 在文件上执行转换
	vue3kratos.GenGrpcViaHttpInPath(testPath)

	// Read conversion output
	// 读取转换输出
	newContent := done.VAE(os.ReadFile(testPath)).Nice()

	// Check output matches expected
	// 检查输出匹配期望
	require.Equal(t, string(resContent), string(newContent), "Converted file should match target output")
}

func TestGenGrpcViaHttpInRoot(t *testing.T) {
	// Read test data
	// 读取测试数据
	sourcePath := runpath.PARENT.Join("testdata", "echo.client.ts-source.txt")
	srcContent := done.VAE(os.ReadFile(sourcePath)).Nice()

	targetPath := runpath.PARENT.Join("testdata", "echo.client.ts-target.txt")
	resContent := done.VAE(os.ReadFile(targetPath)).Nice()

	// Create temp DIR with test files
	// 创建临时 DIR 包含测试文件
	tempRoot, err := os.MkdirTemp("", "vue3kratos_test_*")
	require.NoError(t, err)
	defer func() {
		must.Done(os.RemoveAll(tempRoot))
	}()

	// Create test file
	// 创建测试文件
	testPath := filepath.Join(tempRoot, "test.client.ts")
	done.Done(os.WriteFile(testPath, srcContent, 0644))

	// Execute conversion on entire DIR
	// 在整个 DIR 上执行转换
	vue3kratos.GenGrpcViaHttpInRoot(tempRoot)

	// Read conversion output
	// 读取转换输出
	newContent := done.VAE(os.ReadFile(testPath)).Nice()

	// Check output matches expected
	// 检查输出匹配期望
	require.Equal(t, string(resContent), string(newContent), "Converted file should match target output")
}
