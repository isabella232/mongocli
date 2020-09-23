// Copyright 2020 MongoDB Inc
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// +build unit

package aws

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/mongodb/mongocli/internal/cli"
	"github.com/mongodb/mongocli/internal/flag"
	"github.com/mongodb/mongocli/internal/mocks"
	atlas "go.mongodb.org/atlas/mongodbatlas"
)

func TestEnableBuilder(t *testing.T) {
	cli.CmdValidator(
		t,
		EnableBuilder(),
		0,
		[]string{flag.ProjectID},
	)
}

func TestEnableOpts_Run(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := mocks.NewMockCustomDNSEnabler(ctrl)
	defer ctrl.Finish()

	expected := &atlas.AWSCustomDNSSetting{
		Enabled: true,
	}

	opts := &EnableOpts{
		store: mockStore,
	}

	mockStore.
		EXPECT().
		EnableCustomDNS(opts.ProjectID).
		Return(expected, nil).
		Times(1)

	err := opts.Run()
	if err != nil {
		t.Fatalf("Run() unexpected error: %v", err)
	}
}