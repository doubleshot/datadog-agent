// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2016-present Datadog, Inc.

//go:build kubeapiserver
// +build kubeapiserver

package app

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/DataDog/datadog-agent/pkg/flare"
)

func init() {
	ClusterAgentCmd.AddCommand(telemetryCmd)
}

var telemetryCmd = &cobra.Command{
	Use:   "telemetry",
	Short: "Print the telemetry metrics exposed by the cluster agent",
	Long:  ``,
	RunE: func(cmd *cobra.Command, args []string) error {
		payload, err := flare.QueryDCAMetrics()
		if err != nil {
			return err
		}
		fmt.Print(string(payload))
		return nil
	},
}
