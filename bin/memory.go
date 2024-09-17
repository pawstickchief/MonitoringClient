package bin

import (
	"awesomeProject/datetype"
	"github.com/shirou/gopsutil/v3/mem"
)

func GetMemoryData() (datetype.MemoryData, error) {
	v, err := mem.VirtualMemory()
	if err != nil {
		return datetype.MemoryData{}, err
	}

	memData := datetype.MemoryData{
		Total:       v.Total,
		Used:        v.Used,
		Free:        v.Free,
		UsedPercent: v.UsedPercent,
	}

	return memData, nil
}
