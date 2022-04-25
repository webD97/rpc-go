/*********************************************************************
 * Copyright (c) Intel Corporation 2021
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/
package main

import (
	"fmt"
	"os"
	"rpc"
	"rpc/internal/client"
	"rpc/internal/rps"

	log "github.com/sirupsen/logrus"
)

func runRPC(args []string) {
	// process cli flags/env vars
	flags := handleFlags(args)

	startMessage := rps.PrepareInitialMessage(flags)

	rpc := client.NewExecutor(*flags)
	rpc.MakeItSo(startMessage)
}

func handleFlags(args []string) *rpc.Flags {
	//process flags
	flags := rpc.NewFlags(args)
	_, result := flags.ParseFlags()
	if !result {
		os.Exit(1)
	}
	if flags.SyncClock {
		fmt.Println("Time to sync the clock")
	}
	if flags.Verbose {
		log.SetLevel(log.TraceLevel)
	} else {
		log.SetLevel(log.InfoLevel)
	}
	if flags.JsonOutput {
		log.SetFormatter(&log.JSONFormatter{})
	}
	return flags
}

func main() {
	// ensure we are admin/sudo
	checkAccess()
	runRPC(os.Args)
}
