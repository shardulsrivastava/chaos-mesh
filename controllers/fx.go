// Copyright 2021 Chaos Mesh Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// See the License for the specific language governing permissions and
// limitations under the License.

package controllers

import (
	"go.uber.org/fx"

	"github.com/chaos-mesh/chaos-mesh/controllers/chaosimpl"
	"github.com/chaos-mesh/chaos-mesh/controllers/common"
	"github.com/chaos-mesh/chaos-mesh/controllers/desiredphase"
	"github.com/chaos-mesh/chaos-mesh/controllers/finalizers"
)

var Module = fx.Options(
	fx.Provide(
		fx.Annotated{
			Group:  "controller",
			Target: common.NewController,
		},
		fx.Annotated{
			Group:  "controller",
			Target: finalizers.NewController,
		},
		fx.Annotated{
			Group:  "controller",
			Target: desiredphase.NewController,
		},
	), chaosimpl.AllImpl)