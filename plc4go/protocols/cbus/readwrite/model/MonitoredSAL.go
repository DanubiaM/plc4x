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

// MonitoredSAL is the corresponding interface of MonitoredSAL
type MonitoredSAL interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetSalType returns SalType (property field)
	GetSalType() byte
}

// MonitoredSALExactly can be used when we want exactly this type and not a type which fulfills MonitoredSAL.
// This is useful for switch cases.
type MonitoredSALExactly interface {
	MonitoredSAL
	isMonitoredSAL() bool
}

// _MonitoredSAL is the data-structure of this message
type _MonitoredSAL struct {
	_MonitoredSALChildRequirements
	SalType byte

	// Arguments.
	CBusOptions CBusOptions
}

type _MonitoredSALChildRequirements interface {
	utils.Serializable
	GetLengthInBits(ctx context.Context) uint16
	GetSalType() byte
}

type MonitoredSALParent interface {
	SerializeParent(ctx context.Context, writeBuffer utils.WriteBuffer, child MonitoredSAL, serializeChildFunction func() error) error
	GetTypeName() string
}

type MonitoredSALChild interface {
	utils.Serializable
	InitializeParent(parent MonitoredSAL, salType byte)
	GetParent() *MonitoredSAL

	GetTypeName() string
	MonitoredSAL
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_MonitoredSAL) GetSalType() byte {
	return m.SalType
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewMonitoredSAL factory function for _MonitoredSAL
func NewMonitoredSAL(salType byte, cBusOptions CBusOptions) *_MonitoredSAL {
	return &_MonitoredSAL{SalType: salType, CBusOptions: cBusOptions}
}

// Deprecated: use the interface for direct cast
func CastMonitoredSAL(structType interface{}) MonitoredSAL {
	if casted, ok := structType.(MonitoredSAL); ok {
		return casted
	}
	if casted, ok := structType.(*MonitoredSAL); ok {
		return *casted
	}
	return nil
}

func (m *_MonitoredSAL) GetTypeName() string {
	return "MonitoredSAL"
}

func (m *_MonitoredSAL) GetParentLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	return lengthInBits
}

func (m *_MonitoredSAL) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func MonitoredSALParse(theBytes []byte, cBusOptions CBusOptions) (MonitoredSAL, error) {
	return MonitoredSALParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes), cBusOptions)
}

func MonitoredSALParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, cBusOptions CBusOptions) (MonitoredSAL, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("MonitoredSAL"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for MonitoredSAL")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Peek Field (salType)
	currentPos = positionAware.GetPos()
	salType, _err := readBuffer.ReadByte("salType")
	if _err != nil {
		return nil, errors.Wrap(_err, "Error parsing 'salType' field of MonitoredSAL")
	}

	readBuffer.Reset(currentPos)

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	type MonitoredSALChildSerializeRequirement interface {
		MonitoredSAL
		InitializeParent(MonitoredSAL, byte)
		GetParent() MonitoredSAL
	}
	var _childTemp interface{}
	var _child MonitoredSALChildSerializeRequirement
	var typeSwitchError error
	switch {
	case salType == 0x05: // MonitoredSALLongFormSmartMode
		_childTemp, typeSwitchError = MonitoredSALLongFormSmartModeParseWithBuffer(ctx, readBuffer, cBusOptions)
	case 0 == 0: // MonitoredSALShortFormBasicMode
		_childTemp, typeSwitchError = MonitoredSALShortFormBasicModeParseWithBuffer(ctx, readBuffer, cBusOptions)
	default:
		typeSwitchError = errors.Errorf("Unmapped type for parameters [salType=%v]", salType)
	}
	if typeSwitchError != nil {
		return nil, errors.Wrap(typeSwitchError, "Error parsing sub-type for type-switch of MonitoredSAL")
	}
	_child = _childTemp.(MonitoredSALChildSerializeRequirement)

	if closeErr := readBuffer.CloseContext("MonitoredSAL"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for MonitoredSAL")
	}

	// Finish initializing
	_child.InitializeParent(_child, salType)
	return _child, nil
}

func (pm *_MonitoredSAL) SerializeParent(ctx context.Context, writeBuffer utils.WriteBuffer, child MonitoredSAL, serializeChildFunction func() error) error {
	// We redirect all calls through client as some methods are only implemented there
	m := child
	_ = m
	positionAware := writeBuffer
	_ = positionAware
	if pushErr := writeBuffer.PushContext("MonitoredSAL"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for MonitoredSAL")
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	if _typeSwitchErr := serializeChildFunction(); _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	if popErr := writeBuffer.PopContext("MonitoredSAL"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for MonitoredSAL")
	}
	return nil
}

////
// Arguments Getter

func (m *_MonitoredSAL) GetCBusOptions() CBusOptions {
	return m.CBusOptions
}

//
////

func (m *_MonitoredSAL) isMonitoredSAL() bool {
	return true
}

func (m *_MonitoredSAL) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
