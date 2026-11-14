package utils

import (
	"os"
	"path/filepath"

	"github.com/yylego/erero"
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
