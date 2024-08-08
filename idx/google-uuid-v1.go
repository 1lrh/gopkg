package idx

import (
	"fmt"
	"github.com/google/uuid"
	"net"
	"sync"
)

var once = sync.Once{}

type GoogleUUIDv1Generator struct {
}

func (g *GoogleUUIDv1Generator) GetMac() {
	interfaces, err := net.Interfaces()
	if err != nil {
		fmt.Println("Error fetching network interfaces:", err)
		return
	}

	// 遍历所有网络接口
	for _, iface := range interfaces {
		// 检查接口是否启用并且不是回环接口
		if iface.Flags&net.FlagUp == 0 || iface.Flags&net.FlagLoopback != 0 {
			continue // 忽略未启用或回环的接口
		}

		// 获取接口的硬件地址（MAC地址）
		hwAddr := iface.HardwareAddr

		// 打印接口名称和MAC地址
		fmt.Printf("Interface Name: %s, MAC Address: %s\n", iface.Name, hwAddr)
	}
}

func NewGoogleUUIDv1Generator() *GoogleUUIDv1Generator {
	once.Do(func() {
		uuid.SetNodeID([]byte{byte('1')})
	})

	return &GoogleUUIDv1Generator{}
}

func (g *GoogleUUIDv1Generator) GenerateToken() string {
	uid, _ := uuid.NewUUID()
	return uid.String()
}
