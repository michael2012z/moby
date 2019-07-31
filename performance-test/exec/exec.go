package main

import (
        "os/exec"
        "github.com/sirupsen/logrus"
        "github.com/docker/docker/pkg/jsonmessage"
)

func main() {
     logrus.SetLevel(logrus.DebugLevel)
    logrus.SetFormatter(&logrus.TextFormatter{
        TimestampFormat: jsonmessage.RFC3339NanoFixed,
        DisableColors:   false,
        FullTimestamp:   true,
    })

        path := "/sbin/iptables"
        params := []string{ "--wait", "-t" ,"nat", "-C", "POSTROUTING", "-s", "172.19.0.0/16", "!", "-o", "br-10bf3866bfb9", "-j", "MASQUERADE"}
        logrus.Debugf("before: %s, %v", path, params)
        output, err := exec.Command(path, params...).CombinedOutput()
        // output, err := exec.Command("/bin/ls", "--version").CombinedOutput()
        logrus.Debug("after")
        logrus.Debug(output, err)

}
