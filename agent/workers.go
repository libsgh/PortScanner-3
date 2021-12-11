package agent

import (
	"github.com/libsgh/PortScanner-3/common/lib/go-masscan"
	"github.com/libsgh/PortScanner-3/common/lib/goworker"
)

func resultToNmapQueue(results []masscan.Result, queue, class string) error {
	pushList := make([]*goworker.Payload, len(results))
	for i, v := range results {
		pushList[i] = &goworker.Payload{
			Class: class,
			Args:  []interface{}{v.IP, v.Port, v.IsPublic},
		}
	}
	return goworker.EnqueueMoreOne(queue, pushList...)
}
