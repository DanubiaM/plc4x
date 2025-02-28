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

// AdsDiscoveryBlockPassword is the corresponding interface of AdsDiscoveryBlockPassword
type AdsDiscoveryBlockPassword interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	AdsDiscoveryBlock
	// GetPassword returns Password (property field)
	GetPassword() AmsString
}

// AdsDiscoveryBlockPasswordExactly can be used when we want exactly this type and not a type which fulfills AdsDiscoveryBlockPassword.
// This is useful for switch cases.
type AdsDiscoveryBlockPasswordExactly interface {
	AdsDiscoveryBlockPassword
	isAdsDiscoveryBlockPassword() bool
}

// _AdsDiscoveryBlockPassword is the data-structure of this message
type _AdsDiscoveryBlockPassword struct {
	*_AdsDiscoveryBlock
	Password AmsString
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_AdsDiscoveryBlockPassword) GetBlockType() AdsDiscoveryBlockType {
	return AdsDiscoveryBlockType_PASSWORD
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_AdsDiscoveryBlockPassword) InitializeParent(parent AdsDiscoveryBlock) {}

func (m *_AdsDiscoveryBlockPassword) GetParent() AdsDiscoveryBlock {
	return m._AdsDiscoveryBlock
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_AdsDiscoveryBlockPassword) GetPassword() AmsString {
	return m.Password
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewAdsDiscoveryBlockPassword factory function for _AdsDiscoveryBlockPassword
func NewAdsDiscoveryBlockPassword(password AmsString) *_AdsDiscoveryBlockPassword {
	_result := &_AdsDiscoveryBlockPassword{
		Password:           password,
		_AdsDiscoveryBlock: NewAdsDiscoveryBlock(),
	}
	_result._AdsDiscoveryBlock._AdsDiscoveryBlockChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastAdsDiscoveryBlockPassword(structType interface{}) AdsDiscoveryBlockPassword {
	if casted, ok := structType.(AdsDiscoveryBlockPassword); ok {
		return casted
	}
	if casted, ok := structType.(*AdsDiscoveryBlockPassword); ok {
		return *casted
	}
	return nil
}

func (m *_AdsDiscoveryBlockPassword) GetTypeName() string {
	return "AdsDiscoveryBlockPassword"
}

func (m *_AdsDiscoveryBlockPassword) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (password)
	lengthInBits += m.Password.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_AdsDiscoveryBlockPassword) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func AdsDiscoveryBlockPasswordParse(theBytes []byte) (AdsDiscoveryBlockPassword, error) {
	return AdsDiscoveryBlockPasswordParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes))
}

func AdsDiscoveryBlockPasswordParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (AdsDiscoveryBlockPassword, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("AdsDiscoveryBlockPassword"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for AdsDiscoveryBlockPassword")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (password)
	if pullErr := readBuffer.PullContext("password"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for password")
	}
	_password, _passwordErr := AmsStringParseWithBuffer(ctx, readBuffer)
	if _passwordErr != nil {
		return nil, errors.Wrap(_passwordErr, "Error parsing 'password' field of AdsDiscoveryBlockPassword")
	}
	password := _password.(AmsString)
	if closeErr := readBuffer.CloseContext("password"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for password")
	}

	if closeErr := readBuffer.CloseContext("AdsDiscoveryBlockPassword"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for AdsDiscoveryBlockPassword")
	}

	// Create a partially initialized instance
	_child := &_AdsDiscoveryBlockPassword{
		_AdsDiscoveryBlock: &_AdsDiscoveryBlock{},
		Password:           password,
	}
	_child._AdsDiscoveryBlock._AdsDiscoveryBlockChildRequirements = _child
	return _child, nil
}

func (m *_AdsDiscoveryBlockPassword) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_AdsDiscoveryBlockPassword) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("AdsDiscoveryBlockPassword"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for AdsDiscoveryBlockPassword")
		}

		// Simple Field (password)
		if pushErr := writeBuffer.PushContext("password"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for password")
		}
		_passwordErr := writeBuffer.WriteSerializable(ctx, m.GetPassword())
		if popErr := writeBuffer.PopContext("password"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for password")
		}
		if _passwordErr != nil {
			return errors.Wrap(_passwordErr, "Error serializing 'password' field")
		}

		if popErr := writeBuffer.PopContext("AdsDiscoveryBlockPassword"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for AdsDiscoveryBlockPassword")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_AdsDiscoveryBlockPassword) isAdsDiscoveryBlockPassword() bool {
	return true
}

func (m *_AdsDiscoveryBlockPassword) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
