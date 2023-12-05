//go:build !linux

package device

import (
	"github.com/realglebivanov/amnezia-wg/conn"
	"github.com/realglebivanov/amnezia-wg/rwcancel"
)

func (device *Device) startRouteListener(bind conn.Bind) (*rwcancel.RWCancel, error) {
	return nil, nil
}
