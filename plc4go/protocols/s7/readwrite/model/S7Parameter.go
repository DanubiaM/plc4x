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

// S7Parameter is the corresponding interface of S7Parameter
type S7Parameter interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetMessageType returns MessageType (discriminator field)
	GetMessageType() uint8
	// GetParameterType returns ParameterType (discriminator field)
	GetParameterType() uint8
}

// S7ParameterExactly can be used when we want exactly this type and not a type which fulfills S7Parameter.
// This is useful for switch cases.
type S7ParameterExactly interface {
	S7Parameter
	isS7Parameter() bool
}

// _S7Parameter is the data-structure of this message
type _S7Parameter struct {
	_S7ParameterChildRequirements
}

type _S7ParameterChildRequirements interface {
	utils.Serializable
	GetLengthInBits(ctx context.Context) uint16
	GetParameterType() uint8
	GetMessageType() uint8
}

type S7ParameterParent interface {
	SerializeParent(ctx context.Context, writeBuffer utils.WriteBuffer, child S7Parameter, serializeChildFunction func() error) error
	GetTypeName() string
}

type S7ParameterChild interface {
	utils.Serializable
	InitializeParent(parent S7Parameter)
	GetParent() *S7Parameter

	GetTypeName() string
	S7Parameter
}

// NewS7Parameter factory function for _S7Parameter
func NewS7Parameter() *_S7Parameter {
	return &_S7Parameter{}
}

// Deprecated: use the interface for direct cast
func CastS7Parameter(structType interface{}) S7Parameter {
	if casted, ok := structType.(S7Parameter); ok {
		return casted
	}
	if casted, ok := structType.(*S7Parameter); ok {
		return *casted
	}
	return nil
}

func (m *_S7Parameter) GetTypeName() string {
	return "S7Parameter"
}

func (m *_S7Parameter) GetParentLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)
	// Discriminator Field (parameterType)
	lengthInBits += 8

	return lengthInBits
}

func (m *_S7Parameter) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func S7ParameterParse(theBytes []byte, messageType uint8) (S7Parameter, error) {
	return S7ParameterParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes), messageType)
}

func S7ParameterParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, messageType uint8) (S7Parameter, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("S7Parameter"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for S7Parameter")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Discriminator Field (parameterType) (Used as input to a switch field)
	parameterType, _parameterTypeErr := readBuffer.ReadUint8("parameterType", 8)
	if _parameterTypeErr != nil {
		return nil, errors.Wrap(_parameterTypeErr, "Error parsing 'parameterType' field of S7Parameter")
	}

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	type S7ParameterChildSerializeRequirement interface {
		S7Parameter
		InitializeParent(S7Parameter)
		GetParent() S7Parameter
	}
	var _childTemp interface{}
	var _child S7ParameterChildSerializeRequirement
	var typeSwitchError error
	switch {
	case parameterType == 0xF0: // S7ParameterSetupCommunication
		_childTemp, typeSwitchError = S7ParameterSetupCommunicationParseWithBuffer(ctx, readBuffer, messageType)
	case parameterType == 0x04 && messageType == 0x01: // S7ParameterReadVarRequest
		_childTemp, typeSwitchError = S7ParameterReadVarRequestParseWithBuffer(ctx, readBuffer, messageType)
	case parameterType == 0x04 && messageType == 0x03: // S7ParameterReadVarResponse
		_childTemp, typeSwitchError = S7ParameterReadVarResponseParseWithBuffer(ctx, readBuffer, messageType)
	case parameterType == 0x05 && messageType == 0x01: // S7ParameterWriteVarRequest
		_childTemp, typeSwitchError = S7ParameterWriteVarRequestParseWithBuffer(ctx, readBuffer, messageType)
	case parameterType == 0x05 && messageType == 0x03: // S7ParameterWriteVarResponse
		_childTemp, typeSwitchError = S7ParameterWriteVarResponseParseWithBuffer(ctx, readBuffer, messageType)
	case parameterType == 0x00 && messageType == 0x07: // S7ParameterUserData
		_childTemp, typeSwitchError = S7ParameterUserDataParseWithBuffer(ctx, readBuffer, messageType)
	case parameterType == 0x01 && messageType == 0x07: // S7ParameterModeTransition
		_childTemp, typeSwitchError = S7ParameterModeTransitionParseWithBuffer(ctx, readBuffer, messageType)
	default:
		typeSwitchError = errors.Errorf("Unmapped type for parameters [parameterType=%v, messageType=%v]", parameterType, messageType)
	}
	if typeSwitchError != nil {
		return nil, errors.Wrap(typeSwitchError, "Error parsing sub-type for type-switch of S7Parameter")
	}
	_child = _childTemp.(S7ParameterChildSerializeRequirement)

	if closeErr := readBuffer.CloseContext("S7Parameter"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for S7Parameter")
	}

	// Finish initializing
	_child.InitializeParent(_child)
	return _child, nil
}

func (pm *_S7Parameter) SerializeParent(ctx context.Context, writeBuffer utils.WriteBuffer, child S7Parameter, serializeChildFunction func() error) error {
	// We redirect all calls through client as some methods are only implemented there
	m := child
	_ = m
	positionAware := writeBuffer
	_ = positionAware
	if pushErr := writeBuffer.PushContext("S7Parameter"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for S7Parameter")
	}

	// Discriminator Field (parameterType) (Used as input to a switch field)
	parameterType := uint8(child.GetParameterType())
	_parameterTypeErr := writeBuffer.WriteUint8("parameterType", 8, (parameterType))

	if _parameterTypeErr != nil {
		return errors.Wrap(_parameterTypeErr, "Error serializing 'parameterType' field")
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	if _typeSwitchErr := serializeChildFunction(); _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	if popErr := writeBuffer.PopContext("S7Parameter"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for S7Parameter")
	}
	return nil
}

func (m *_S7Parameter) isS7Parameter() bool {
	return true
}

func (m *_S7Parameter) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
