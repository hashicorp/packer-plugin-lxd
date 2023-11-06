// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package lxd

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/hashicorp/packer-plugin-sdk/multistep"
	packersdk "github.com/hashicorp/packer-plugin-sdk/packer"
)

type stepLxdLaunch struct{}

func (s *stepLxdLaunch) Run(ctx context.Context, state multistep.StateBag) multistep.StepAction {
	config := state.Get("config").(*Config)
	ui := state.Get("ui").(packersdk.Ui)

	name := config.ContainerName
	image := config.Image
	profile := fmt.Sprintf("--profile=%s", config.Profile)
	vm := fmt.Sprintf("--vm=%s", strconv.FormatBool(config.VirtualMachine))

	launch_args := []string{
		"launch", "--ephemeral=false", vm, profile, image, name,
	}

	for k, v := range config.LaunchConfig {
		launch_args = append(launch_args, "--config", fmt.Sprintf("%s=%s", k, v))
	}

	sleep_seconds, err := strconv.Atoi(config.InitSleep)
	if err != nil {
		err := fmt.Errorf("Error parsing InitSleep into int: %s", err)
		state.Put("error", err)
		ui.Error(err.Error())
		return multistep.ActionHalt
	}

	ui.Say("Creating container...")
	_, err = LXDCommand(launch_args...)
	if err != nil {
		err := fmt.Errorf("Error creating container: %s", err)
		state.Put("error", err)
		ui.Error(err.Error())
		return multistep.ActionHalt
	}

	// TODO: Should we check `lxc info <container>` for "Running"?
	// We have to do this so /tmp doesn't get cleared and lose our provisioner scripts.

	log.Printf("Sleeping for %d seconds...", sleep_seconds)
	time.Sleep(time.Duration(sleep_seconds) * time.Second)

	// instance_id is the generic term used so that users can have access to the
	// instance id inside the provisioners, used in step_provision.
	state.Put("instance_id", name)
	ui.Message(fmt.Sprintf("Instance Name: %s", name))
	return multistep.ActionContinue
}

func (s *stepLxdLaunch) Cleanup(state multistep.StateBag) {
	config := state.Get("config").(*Config)
	ui := state.Get("ui").(packersdk.Ui)

	if config.SkipPublish {
		return
	}

	cleanup_args := []string{
		"delete", "--force", config.ContainerName,
	}

	ui.Say("Unregistering and deleting deleting container...")
	if _, err := LXDCommand(cleanup_args...); err != nil {
		ui.Error(fmt.Sprintf("Error deleting container: %s", err))
	}
}
