package utils

import (
	"os"
	"path/filepath"

	"github.com/yylego/erero"
	"github.com/yylego/must"
	"github.com/yylego/must/muststrings"
	"github.com/yylego/osexistpath/osmustexist"
	"github.com/yylego/rese"
	"github.com/yylego/zaplog"
)

// WalkFiles walks through files in the DIR and runs the function on each file
// Skips directories and passes each file path and info to the run function
//
// WalkFiles 遍历 DIR 中的文件并对每个文件运行函数
// 跳过 DIR 并将每个文件路径和信息传递给 run 函数
func WalkFiles(root string, run func(path string, info os.FileInfo) error) error {
	if err := filepath.Walk(root,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return erero.Wro(err)
			}
			if info.IsDir() {
				return nil
			}
			return run(path, info)
		},
	); err != nil {
		return erero.Wro(err)
	}
	return nil
}

// CopyFiles copies files from source DIR to target DIR
// Preserves the same relative paths and file permissions
//
// CopyFiles 将文件从源 DIR 复制到目标 DIR
// 保持相同的相对路径和文件权限
func CopyFiles(sourceRoot string, targetRoot string) {
	must.Done(WalkFiles(sourceRoot, func(srcPath string, info os.FileInfo) error {
		zaplog.SUG.Debugln(srcPath)
		muststrings.HasPrefix(srcPath, sourceRoot)

		relPath := must.V1(filepath.Rel(sourceRoot, srcPath))
		absPath := filepath.Join(sourceRoot, relPath)
		must.Same(absPath, srcPath)

		content := rese.V1(os.ReadFile(osmustexist.FILE(absPath)))
		dstPath := filepath.Join(targetRoot, relPath)
		zaplog.SUG.Debugln(dstPath)

		must.Done(os.MkdirAll(filepath.Dir(dstPath), 0755))
		must.Done(os.WriteFile(dstPath, content, 0644))
		return nil
	}))
}
