package utils

import (
	"encoding/json"
	"os"
	"path/filepath"
)

func CurrentDir(subpath ...string) string {
	// 获取可执行程序路径
	path, _ := os.Executable()
	path, _ = filepath.EvalSymlinks(filepath.Dir(path))

	// 返回路径
	return filepath.Join(append([]string{path}, subpath...)...)
}

func Any2Str(obj any) string {
	if obj == nil {return ""}
	data, _ := json.Marshal(obj)
	return string(data)
}

func Any2MapAny(str any) map[string]any {
	obj := make(map[string]any)
	json.Unmarshal([]byte(Any2Str(str)), &obj)
	return obj
}
