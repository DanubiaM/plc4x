/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   https://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

package _default

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"net/url"

	"github.com/apache/plc4x/plc4go/pkg/api"
	apiModel "github.com/apache/plc4x/plc4go/pkg/api/model"
	"github.com/apache/plc4x/plc4go/spi"
	"github.com/apache/plc4x/plc4go/spi/options"
	"github.com/apache/plc4x/plc4go/spi/transports"
)

type DefaultDriverRequirements interface {
	GetConnectionWithContext(ctx context.Context, transportUrl url.URL, transports map[string]transports.Transport, options map[string][]string) <-chan plc4go.PlcConnectionConnectResult
	DiscoverWithContext(callback context.Context, event func(event apiModel.PlcDiscoveryItem), discoveryOptions ...options.WithDiscoveryOption) error
}

type DefaultDriver interface {
	fmt.Stringer
	plc4go.PlcDriver
	spi.PlcDiscoverer
	GetPlcTagHandler() spi.PlcTagHandler
}

func NewDefaultDriver(defaultDriverRequirements DefaultDriverRequirements, protocolCode string, protocolName string, defaultTransport string, plcTagHandler spi.PlcTagHandler) DefaultDriver {
	return &defaultDriver{
		DefaultDriverRequirements: defaultDriverRequirements,
		protocolCode:              protocolCode,
		protocolName:              protocolName,
		defaultTransport:          defaultTransport,
		plcTagHandler:             plcTagHandler,
	}
}

///////////////////////////////////////
///////////////////////////////////////
//
// Internal section
//

type defaultDriver struct {
	DefaultDriverRequirements
	protocolCode     string
	protocolName     string
	defaultTransport string
	plcTagHandler    spi.PlcTagHandler
}

//
// Internal section
//
///////////////////////////////////////
///////////////////////////////////////

func (d *defaultDriver) GetProtocolCode() string {
	return d.protocolCode
}

func (d *defaultDriver) GetProtocolName() string {
	return d.protocolName
}

func (d *defaultDriver) GetDefaultTransport() string {
	return d.defaultTransport
}

func (d *defaultDriver) CheckTagAddress(query string) error {
	_, err := d.plcTagHandler.ParseTag(query)
	return err
}

func (d *defaultDriver) CheckQuery(query string) error {
	_, err := d.plcTagHandler.ParseQuery(query)
	return err
}

func (d *defaultDriver) GetConnection(transportUrl url.URL, transports map[string]transports.Transport, options map[string][]string) <-chan plc4go.PlcConnectionConnectResult {
	return d.GetConnectionWithContext(context.Background(), transportUrl, transports, options)
}

func (d *defaultDriver) SupportsDiscovery() bool {
	return false
}

func (d *defaultDriver) Discover(callback func(event apiModel.PlcDiscoveryItem), discoveryOptions ...options.WithDiscoveryOption) error {
	return d.DiscoverWithContext(context.Background(), callback, discoveryOptions...)
}

func (d *defaultDriver) DiscoverWithContext(_ context.Context, _ func(event apiModel.PlcDiscoveryItem), _ ...options.WithDiscoveryOption) error {
	return errors.New("not available")
}

func (d *defaultDriver) GetPlcTagHandler() spi.PlcTagHandler {
	return d.plcTagHandler
}

func (d *defaultDriver) String() string {
	return fmt.Sprintf("%s (%s) [%s]", d.protocolName, d.protocolCode, d.defaultTransport)
}
