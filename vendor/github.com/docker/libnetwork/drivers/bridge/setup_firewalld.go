package bridge

import (
	"time"

	"github.com/docker/libnetwork/iptables"
	"github.com/sirupsen/logrus"
)

func (n *bridgeNetwork) setupFirewalld(config *networkConfiguration, i *bridgeInterface) error {
	logrus.Info("step: 9")
	logrus.Info(time.Now())
	d := n.driver
	d.Lock()
	driverConfig := d.config
	d.Unlock()

	// Sanity check.
	if !driverConfig.EnableIPTables {
		return IPTableCfgError(config.BridgeName)
	}

	iptables.OnReloaded(func() { n.setupIPTables(config, i) })
	iptables.OnReloaded(n.portMapper.ReMapAll)

	return nil
}
