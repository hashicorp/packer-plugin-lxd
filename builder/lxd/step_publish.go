// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package lxd

import (
	"context"
	"fmt"
	"regexp"

	"github.com/hashicorp/packer-plugin-sdk/multistep"
	packersdk "github.com/hashicorp/packer-plugin-sdk/packer"
)

type stepPublish struct{}

func (s *stepPublish) Run(ctx context.Context, state multistep.StateBag) multistep.StepAction {

	var remote string = ""

	config := state.Get("config").(*Config)
	ui := state.Get("ui").(packersdk.Ui)

	if config.SkipPublish {
		ui.Say("skip_publish is true. Skipping step for publish")
		return multistep.ActionContinue
	}

	name := config.ContainerName
	stop_args := []string{
		// We created the container with "--ephemeral=false" so we know it is safe to stop.
		"stop", name,
	}

	ui.Say("Stopping container...")
	_, err := LXDCommand(stop_args...)
	if err != nil {
		err := fmt.Errorf("Error stopping container: %s", err)
		state.Put("error", err)
		ui.Error(err.Error())
		return multistep.ActionHalt
	}

	if config.PublishRemoteName != "" {
		remote = config.PublishRemoteName + ":"
	}

	publish_args := []string{
		"publish", name, remote, "--alias", config.OutputImage,
	}

	if len(config.OutputAliases) != 0 {
		for _, v := range config.OutputAliases {
			publish_args = append(publish_args, fmt.Sprintf("--alias"))
			publish_args = append(publish_args, fmt.Sprintf("%s", v))
		}
	}
	
	for k, v := range config.PublishProperties {
		publish_args = append(publish_args, fmt.Sprintf("%s=%s", k, v))
	}

	ui.Say("Publishing container...")
	stdoutString, err := LXDCommand(publish_args...)
	if err != nil {
		err := fmt.Errorf("Error publishing container: %s", err)
		state.Put("error", err)
		ui.Error(err.Error())
		return multistep.ActionHalt
	}

	r := regexp.MustCompile("([0-9a-fA-F]+)$")
	fingerprint := r.FindAllStringSubmatch(stdoutString, -1)[0][0]

	ui.Say(fmt.Sprintf("Created image: %s", fingerprint))

	state.Put("imageFingerprint", fingerprint)

	return multistep.ActionContinue
}

func (s *stepPublish) Cleanup(state multistep.StateBag) {}
