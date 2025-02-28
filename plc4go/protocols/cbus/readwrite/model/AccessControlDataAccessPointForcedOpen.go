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

package model

import (
	"context"
	"fmt"
	"github.com/apache/plc4x/plc4go/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// AccessControlDataAccessPointForcedOpen is the corresponding interface of AccessControlDataAccessPointForcedOpen
type AccessControlDataAccessPointForcedOpen interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	AccessControlData
}

// AccessControlDataAccessPointForcedOpenExactly can be used when we want exactly this type and not a type which fulfills AccessControlDataAccessPointForcedOpen.
// This is useful for switch cases.
type AccessControlDataAccessPointForcedOpenExactly interface {
	AccessControlDataAccessPointForcedOpen
	isAccessControlDataAccessPointForcedOpen() bool
}

// _AccessControlDataAccessPointForcedOpen is the data-structure of this message
type _AccessControlDataAccessPointForcedOpen struct {
	*_AccessControlData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_AccessControlDataAccessPointForcedOpen) InitializeParent(parent AccessControlData, commandTypeContainer AccessControlCommandTypeContainer, networkId byte, accessPointId byte) {
	m.CommandTypeContainer = commandTypeContainer
	m.NetworkId = networkId
	m.AccessPointId = accessPointId
}

func (m *_AccessControlDataAccessPointForcedOpen) GetParent() AccessControlData {
	return m._AccessControlData
}

// NewAccessControlDataAccessPointForcedOpen factory function for _AccessControlDataAccessPointForcedOpen
func NewAccessControlDataAccessPointForcedOpen(commandTypeContainer AccessControlCommandTypeContainer, networkId byte, accessPointId byte) *_AccessControlDataAccessPointForcedOpen {
	_result := &_AccessControlDataAccessPointForcedOpen{
		_AccessControlData: NewAccessControlData(commandTypeContainer, networkId, accessPointId),
	}
	_result._AccessControlData._AccessControlDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastAccessControlDataAccessPointForcedOpen(structType interface{}) AccessControlDataAccessPointForcedOpen {
	if casted, ok := structType.(AccessControlDataAccessPointForcedOpen); ok {
		return casted
	}
	if casted, ok := structType.(*AccessControlDataAccessPointForcedOpen); ok {
		return *casted
	}
	return nil
}

func (m *_AccessControlDataAccessPointForcedOpen) GetTypeName() string {
	return "AccessControlDataAccessPointForcedOpen"
}

func (m *_AccessControlDataAccessPointForcedOpen) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	return lengthInBits
}

func (m *_AccessControlDataAccessPointForcedOpen) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func AccessControlDataAccessPointForcedOpenParse(theBytes []byte) (AccessControlDataAccessPointForcedOpen, error) {
	return AccessControlDataAccessPointForcedOpenParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes))
}

func AccessControlDataAccessPointForcedOpenParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (AccessControlDataAccessPointForcedOpen, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("AccessControlDataAccessPointForcedOpen"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for AccessControlDataAccessPointForcedOpen")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("AccessControlDataAccessPointForcedOpen"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for AccessControlDataAccessPointForcedOpen")
	}

	// Create a partially initialized instance
	_child := &_AccessControlDataAccessPointForcedOpen{
		_AccessControlData: &_AccessControlData{},
	}
	_child._AccessControlData._AccessControlDataChildRequirements = _child
	return _child, nil
}

func (m *_AccessControlDataAccessPointForcedOpen) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_AccessControlDataAccessPointForcedOpen) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("AccessControlDataAccessPointForcedOpen"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for AccessControlDataAccessPointForcedOpen")
		}

		if popErr := writeBuffer.PopContext("AccessControlDataAccessPointForcedOpen"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for AccessControlDataAccessPointForcedOpen")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_AccessControlDataAccessPointForcedOpen) isAccessControlDataAccessPointForcedOpen() bool {
	return true
}

func (m *_AccessControlDataAccessPointForcedOpen) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
