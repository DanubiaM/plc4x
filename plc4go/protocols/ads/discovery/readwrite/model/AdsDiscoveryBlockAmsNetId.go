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

// Constant values.
const AdsDiscoveryBlockAmsNetId_AMSNETIDLENGTH uint16 = 0x0006

// AdsDiscoveryBlockAmsNetId is the corresponding interface of AdsDiscoveryBlockAmsNetId
type AdsDiscoveryBlockAmsNetId interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	AdsDiscoveryBlock
	// GetAmsNetId returns AmsNetId (property field)
	GetAmsNetId() AmsNetId
}

// AdsDiscoveryBlockAmsNetIdExactly can be used when we want exactly this type and not a type which fulfills AdsDiscoveryBlockAmsNetId.
// This is useful for switch cases.
type AdsDiscoveryBlockAmsNetIdExactly interface {
	AdsDiscoveryBlockAmsNetId
	isAdsDiscoveryBlockAmsNetId() bool
}

// _AdsDiscoveryBlockAmsNetId is the data-structure of this message
type _AdsDiscoveryBlockAmsNetId struct {
	*_AdsDiscoveryBlock
	AmsNetId AmsNetId
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_AdsDiscoveryBlockAmsNetId) GetBlockType() AdsDiscoveryBlockType {
	return AdsDiscoveryBlockType_AMS_NET_ID
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_AdsDiscoveryBlockAmsNetId) InitializeParent(parent AdsDiscoveryBlock) {}

func (m *_AdsDiscoveryBlockAmsNetId) GetParent() AdsDiscoveryBlock {
	return m._AdsDiscoveryBlock
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_AdsDiscoveryBlockAmsNetId) GetAmsNetId() AmsNetId {
	return m.AmsNetId
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for const fields.
///////////////////////

func (m *_AdsDiscoveryBlockAmsNetId) GetAmsNetIdLength() uint16 {
	return AdsDiscoveryBlockAmsNetId_AMSNETIDLENGTH
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewAdsDiscoveryBlockAmsNetId factory function for _AdsDiscoveryBlockAmsNetId
func NewAdsDiscoveryBlockAmsNetId(amsNetId AmsNetId) *_AdsDiscoveryBlockAmsNetId {
	_result := &_AdsDiscoveryBlockAmsNetId{
		AmsNetId:           amsNetId,
		_AdsDiscoveryBlock: NewAdsDiscoveryBlock(),
	}
	_result._AdsDiscoveryBlock._AdsDiscoveryBlockChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastAdsDiscoveryBlockAmsNetId(structType interface{}) AdsDiscoveryBlockAmsNetId {
	if casted, ok := structType.(AdsDiscoveryBlockAmsNetId); ok {
		return casted
	}
	if casted, ok := structType.(*AdsDiscoveryBlockAmsNetId); ok {
		return *casted
	}
	return nil
}

func (m *_AdsDiscoveryBlockAmsNetId) GetTypeName() string {
	return "AdsDiscoveryBlockAmsNetId"
}

func (m *_AdsDiscoveryBlockAmsNetId) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Const Field (amsNetIdLength)
	lengthInBits += 16

	// Simple field (amsNetId)
	lengthInBits += m.AmsNetId.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_AdsDiscoveryBlockAmsNetId) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func AdsDiscoveryBlockAmsNetIdParse(theBytes []byte) (AdsDiscoveryBlockAmsNetId, error) {
	return AdsDiscoveryBlockAmsNetIdParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes))
}

func AdsDiscoveryBlockAmsNetIdParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (AdsDiscoveryBlockAmsNetId, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("AdsDiscoveryBlockAmsNetId"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for AdsDiscoveryBlockAmsNetId")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Const Field (amsNetIdLength)
	amsNetIdLength, _amsNetIdLengthErr := readBuffer.ReadUint16("amsNetIdLength", 16)
	if _amsNetIdLengthErr != nil {
		return nil, errors.Wrap(_amsNetIdLengthErr, "Error parsing 'amsNetIdLength' field of AdsDiscoveryBlockAmsNetId")
	}
	if amsNetIdLength != AdsDiscoveryBlockAmsNetId_AMSNETIDLENGTH {
		return nil, errors.New("Expected constant value " + fmt.Sprintf("%d", AdsDiscoveryBlockAmsNetId_AMSNETIDLENGTH) + " but got " + fmt.Sprintf("%d", amsNetIdLength))
	}

	// Simple Field (amsNetId)
	if pullErr := readBuffer.PullContext("amsNetId"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for amsNetId")
	}
	_amsNetId, _amsNetIdErr := AmsNetIdParseWithBuffer(ctx, readBuffer)
	if _amsNetIdErr != nil {
		return nil, errors.Wrap(_amsNetIdErr, "Error parsing 'amsNetId' field of AdsDiscoveryBlockAmsNetId")
	}
	amsNetId := _amsNetId.(AmsNetId)
	if closeErr := readBuffer.CloseContext("amsNetId"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for amsNetId")
	}

	if closeErr := readBuffer.CloseContext("AdsDiscoveryBlockAmsNetId"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for AdsDiscoveryBlockAmsNetId")
	}

	// Create a partially initialized instance
	_child := &_AdsDiscoveryBlockAmsNetId{
		_AdsDiscoveryBlock: &_AdsDiscoveryBlock{},
		AmsNetId:           amsNetId,
	}
	_child._AdsDiscoveryBlock._AdsDiscoveryBlockChildRequirements = _child
	return _child, nil
}

func (m *_AdsDiscoveryBlockAmsNetId) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_AdsDiscoveryBlockAmsNetId) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("AdsDiscoveryBlockAmsNetId"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for AdsDiscoveryBlockAmsNetId")
		}

		// Const Field (amsNetIdLength)
		_amsNetIdLengthErr := writeBuffer.WriteUint16("amsNetIdLength", 16, 0x0006)
		if _amsNetIdLengthErr != nil {
			return errors.Wrap(_amsNetIdLengthErr, "Error serializing 'amsNetIdLength' field")
		}

		// Simple Field (amsNetId)
		if pushErr := writeBuffer.PushContext("amsNetId"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for amsNetId")
		}
		_amsNetIdErr := writeBuffer.WriteSerializable(ctx, m.GetAmsNetId())
		if popErr := writeBuffer.PopContext("amsNetId"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for amsNetId")
		}
		if _amsNetIdErr != nil {
			return errors.Wrap(_amsNetIdErr, "Error serializing 'amsNetId' field")
		}

		if popErr := writeBuffer.PopContext("AdsDiscoveryBlockAmsNetId"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for AdsDiscoveryBlockAmsNetId")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_AdsDiscoveryBlockAmsNetId) isAdsDiscoveryBlockAmsNetId() bool {
	return true
}

func (m *_AdsDiscoveryBlockAmsNetId) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
