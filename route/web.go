package route

import (
	"awesomeProject/bin"
	"awesomeProject/datetype"
	"awesomeProject/logger"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func Setup() *gin.Engine {
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))
	r.GET("/api/monitor", MonitorHandler)
	return r
}
func MonitorHandler(c *gin.Context) {
	// 获取内存数据
	memData, err := bin.GetMemoryData()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "获取内存数据出错",
		})
		log.Println("获取内存数据出错：", err)
		return
	}

	// 获取网络数据
	netData, err := bin.GetNetworkData()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "获取网络数据出错",
		})
		log.Println("获取网络数据出错：", err)
		return
	}

	// 获取 CPU 数据
	cpuData, err := bin.GetCPUInfo()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "获取 CPU 数据出错",
		})
		log.Println("获取 CPU 数据出错：", err)
		return
	}

	// 获取磁盘数据
	diskData, err := bin.GetDiskInfo()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "获取磁盘数据出错",
		})
		log.Println("获取磁盘数据出错：", err)
		return
	}

	// 构造最终返回的数据结构
	monitorData := datetype.MonitorData{
		Memory:  memData,
		Network: netData,
		CPU:     cpuData,
		Disk:    diskData,
	}

	// 返回 JSON 数据
	c.JSON(http.StatusOK, monitorData)
}
