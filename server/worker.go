package server

import "github.com/libsgh/PortScanner-3/common/lib/goworker"

func taskToMasscan(queue, class, ipRange, scanPort, scanRate string, isPublic bool) error {
	return goworker.Enqueue(&goworker.Job{
		Queue: queue,
		Payload: goworker.Payload{
			Class: class,
			Args:  []interface{}{ipRange, scanPort, scanRate, isPublic},
		},
	})
}
