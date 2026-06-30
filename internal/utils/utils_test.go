package utils

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/yylego/runpath"
)

func TestWalkFiles(t *testing.T) {
	require.NoError(t, WalkFiles(runpath.PARENT.Path(), func(path string, info os.FileInfo) error {
		t.Log(path)
		return nil
	}))
}

func TestCopyFiles(t *testing.T) {
	sourceRoot := t.TempDir()
	targetRoot := t.TempDir()

	// A nested file in the source DIR.
	// 源 DIR 里的一个嵌套文件
	require.NoError(t, os.MkdirAll(filepath.Join(sourceRoot, "a"), 0755))
	require.NoError(t, os.WriteFile(filepath.Join(sourceRoot, "a", "f.txt"), []byte("hello"), 0644))

	CopyFiles(sourceRoot, targetRoot)

	// The same relative path lands in the target DIR with the same content.
	// 相同相对路径落到目标 DIR，内容一致
	got, err := os.ReadFile(filepath.Join(targetRoot, "a", "f.txt"))
	require.NoError(t, err)
	require.Equal(t, "hello", string(got))
}
