package main

import (
	"github.com/josh-diamond/rancher-terratest/config"
	"github.com/josh-diamond/rancher-terratest/functions"
)


func main() {
	config.BuildConfig1()
	functions.SetConfigTF(config.K3s, config.Config1)
}
