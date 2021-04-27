package system

import "runtime"

func GetCpuNumber() int {
	return runtime.NumCPU()
}
