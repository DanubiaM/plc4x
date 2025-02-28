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

// AdsMultiRequestItemRead is the corresponding interface of AdsMultiRequestItemRead
type AdsMultiRequestItemRead interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	AdsMultiRequestItem
	// GetItemIndexGroup returns ItemIndexGroup (property field)
	GetItemIndexGroup() uint32
	// GetItemIndexOffset returns ItemIndexOffset (property field)
	GetItemIndexOffset() uint32
	// GetItemReadLength returns ItemReadLength (property field)
	GetItemReadLength() uint32
}

// AdsMultiRequestItemReadExactly can be used when we want exactly this type and not a type which fulfills AdsMultiRequestItemRead.
// This is useful for switch cases.
type AdsMultiRequestItemReadExactly interface {
	AdsMultiRequestItemRead
	isAdsMultiRequestItemRead() bool
}

// _AdsMultiRequestItemRead is the data-structure of this message
type _AdsMultiRequestItemRead struct {
	*_AdsMultiRequestItem
	ItemIndexGroup  uint32
	ItemIndexOffset uint32
	ItemReadLength  uint32
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_AdsMultiRequestItemRead) GetIndexGroup() uint32 {
	return uint32(61568)
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_AdsMultiRequestItemRead) InitializeParent(parent AdsMultiRequestItem) {}

func (m *_AdsMultiRequestItemRead) GetParent() AdsMultiRequestItem {
	return m._AdsMultiRequestItem
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_AdsMultiRequestItemRead) GetItemIndexGroup() uint32 {
	return m.ItemIndexGroup
}

func (m *_AdsMultiRequestItemRead) GetItemIndexOffset() uint32 {
	return m.ItemIndexOffset
}

func (m *_AdsMultiRequestItemRead) GetItemReadLength() uint32 {
	return m.ItemReadLength
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewAdsMultiRequestItemRead factory function for _AdsMultiRequestItemRead
func NewAdsMultiRequestItemRead(itemIndexGroup uint32, itemIndexOffset uint32, itemReadLength uint32) *_AdsMultiRequestItemRead {
	_result := &_AdsMultiRequestItemRead{
		ItemIndexGroup:       itemIndexGroup,
		ItemIndexOffset:      itemIndexOffset,
		ItemReadLength:       itemReadLength,
		_AdsMultiRequestItem: NewAdsMultiRequestItem(),
	}
	_result._AdsMultiRequestItem._AdsMultiRequestItemChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastAdsMultiRequestItemRead(structType interface{}) AdsMultiRequestItemRead {
	if casted, ok := structType.(AdsMultiRequestItemRead); ok {
		return casted
	}
	if casted, ok := structType.(*AdsMultiRequestItemRead); ok {
		return *casted
	}
	return nil
}

func (m *_AdsMultiRequestItemRead) GetTypeName() string {
	return "AdsMultiRequestItemRead"
}

func (m *_AdsMultiRequestItemRead) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (itemIndexGroup)
	lengthInBits += 32

	// Simple field (itemIndexOffset)
	lengthInBits += 32

	// Simple field (itemReadLength)
	lengthInBits += 32

	return lengthInBits
}

func (m *_AdsMultiRequestItemRead) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func AdsMultiRequestItemReadParse(theBytes []byte, indexGroup uint32) (AdsMultiRequestItemRead, error) {
	return AdsMultiRequestItemReadParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes), indexGroup)
}

func AdsMultiRequestItemReadParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, indexGroup uint32) (AdsMultiRequestItemRead, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("AdsMultiRequestItemRead"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for AdsMultiRequestItemRead")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (itemIndexGroup)
	_itemIndexGroup, _itemIndexGroupErr := readBuffer.ReadUint32("itemIndexGroup", 32)
	if _itemIndexGroupErr != nil {
		return nil, errors.Wrap(_itemIndexGroupErr, "Error parsing 'itemIndexGroup' field of AdsMultiRequestItemRead")
	}
	itemIndexGroup := _itemIndexGroup

	// Simple Field (itemIndexOffset)
	_itemIndexOffset, _itemIndexOffsetErr := readBuffer.ReadUint32("itemIndexOffset", 32)
	if _itemIndexOffsetErr != nil {
		return nil, errors.Wrap(_itemIndexOffsetErr, "Error parsing 'itemIndexOffset' field of AdsMultiRequestItemRead")
	}
	itemIndexOffset := _itemIndexOffset

	// Simple Field (itemReadLength)
	_itemReadLength, _itemReadLengthErr := readBuffer.ReadUint32("itemReadLength", 32)
	if _itemReadLengthErr != nil {
		return nil, errors.Wrap(_itemReadLengthErr, "Error parsing 'itemReadLength' field of AdsMultiRequestItemRead")
	}
	itemReadLength := _itemReadLength

	if closeErr := readBuffer.CloseContext("AdsMultiRequestItemRead"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for AdsMultiRequestItemRead")
	}

	// Create a partially initialized instance
	_child := &_AdsMultiRequestItemRead{
		_AdsMultiRequestItem: &_AdsMultiRequestItem{},
		ItemIndexGroup:       itemIndexGroup,
		ItemIndexOffset:      itemIndexOffset,
		ItemReadLength:       itemReadLength,
	}
	_child._AdsMultiRequestItem._AdsMultiRequestItemChildRequirements = _child
	return _child, nil
}

func (m *_AdsMultiRequestItemRead) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_AdsMultiRequestItemRead) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("AdsMultiRequestItemRead"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for AdsMultiRequestItemRead")
		}

		// Simple Field (itemIndexGroup)
		itemIndexGroup := uint32(m.GetItemIndexGroup())
		_itemIndexGroupErr := writeBuffer.WriteUint32("itemIndexGroup", 32, (itemIndexGroup))
		if _itemIndexGroupErr != nil {
			return errors.Wrap(_itemIndexGroupErr, "Error serializing 'itemIndexGroup' field")
		}

		// Simple Field (itemIndexOffset)
		itemIndexOffset := uint32(m.GetItemIndexOffset())
		_itemIndexOffsetErr := writeBuffer.WriteUint32("itemIndexOffset", 32, (itemIndexOffset))
		if _itemIndexOffsetErr != nil {
			return errors.Wrap(_itemIndexOffsetErr, "Error serializing 'itemIndexOffset' field")
		}

		// Simple Field (itemReadLength)
		itemReadLength := uint32(m.GetItemReadLength())
		_itemReadLengthErr := writeBuffer.WriteUint32("itemReadLength", 32, (itemReadLength))
		if _itemReadLengthErr != nil {
			return errors.Wrap(_itemReadLengthErr, "Error serializing 'itemReadLength' field")
		}

		if popErr := writeBuffer.PopContext("AdsMultiRequestItemRead"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for AdsMultiRequestItemRead")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_AdsMultiRequestItemRead) isAdsMultiRequestItemRead() bool {
	return true
}

func (m *_AdsMultiRequestItemRead) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
