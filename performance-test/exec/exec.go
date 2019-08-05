package main

import (
	"github.com/docker/docker/pkg/jsonmessage"
	"github.com/sirupsen/logrus"
	"os"
	"os/exec"
	"runtime/trace"
)

func main() {
	michaelf, mierr := os.Create("/go/src/github.com/docker/docker/trace.out")
	if mierr != nil {
		panic(mierr)
	}
	defer michaelf.Close()

	mierr = trace.Start(michaelf)
	if mierr != nil {
		panic(mierr)
	}
	defer trace.Stop()

	logrus.SetLevel(logrus.DebugLevel)
	logrus.SetFormatter(&logrus.TextFormatter{
		TimestampFormat: jsonmessage.RFC3339NanoFixed,
		DisableColors:   false,
		FullTimestamp:   true,
	})

	//path := "/sbin/iptables"
	//params := []string{"--wait", "-t", "nat", "-C", "POSTROUTING", "-s", "172.19.0.0/16", "!", "-o", "br-10bf3866bfb9", "-j", "MASQUERADE"}
	//logrus.Debugf("before: %s, %v", path, params)
	//output, err := exec.Command(path, params...).CombinedOutput()
	logrus.Debugf("before: ")
	output, err := exec.Command("/go/src/github.com/docker/docker/a.out").CombinedOutput()
	logrus.Debug("after")
	logrus.Debug(string(output), err)

}
