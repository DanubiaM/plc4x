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

// PortSegment is the corresponding interface of PortSegment
type PortSegment interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	PathSegment
	// GetSegmentType returns SegmentType (property field)
	GetSegmentType() PortSegmentType
}

// PortSegmentExactly can be used when we want exactly this type and not a type which fulfills PortSegment.
// This is useful for switch cases.
type PortSegmentExactly interface {
	PortSegment
	isPortSegment() bool
}

// _PortSegment is the data-structure of this message
type _PortSegment struct {
	*_PathSegment
	SegmentType PortSegmentType
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_PortSegment) GetPathSegment() uint8 {
	return 0x00
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_PortSegment) InitializeParent(parent PathSegment) {}

func (m *_PortSegment) GetParent() PathSegment {
	return m._PathSegment
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_PortSegment) GetSegmentType() PortSegmentType {
	return m.SegmentType
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewPortSegment factory function for _PortSegment
func NewPortSegment(segmentType PortSegmentType) *_PortSegment {
	_result := &_PortSegment{
		SegmentType:  segmentType,
		_PathSegment: NewPathSegment(),
	}
	_result._PathSegment._PathSegmentChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastPortSegment(structType interface{}) PortSegment {
	if casted, ok := structType.(PortSegment); ok {
		return casted
	}
	if casted, ok := structType.(*PortSegment); ok {
		return *casted
	}
	return nil
}

func (m *_PortSegment) GetTypeName() string {
	return "PortSegment"
}

func (m *_PortSegment) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (segmentType)
	lengthInBits += m.SegmentType.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_PortSegment) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func PortSegmentParse(theBytes []byte) (PortSegment, error) {
	return PortSegmentParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes))
}

func PortSegmentParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (PortSegment, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("PortSegment"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for PortSegment")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (segmentType)
	if pullErr := readBuffer.PullContext("segmentType"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for segmentType")
	}
	_segmentType, _segmentTypeErr := PortSegmentTypeParseWithBuffer(ctx, readBuffer)
	if _segmentTypeErr != nil {
		return nil, errors.Wrap(_segmentTypeErr, "Error parsing 'segmentType' field of PortSegment")
	}
	segmentType := _segmentType.(PortSegmentType)
	if closeErr := readBuffer.CloseContext("segmentType"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for segmentType")
	}

	if closeErr := readBuffer.CloseContext("PortSegment"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for PortSegment")
	}

	// Create a partially initialized instance
	_child := &_PortSegment{
		_PathSegment: &_PathSegment{},
		SegmentType:  segmentType,
	}
	_child._PathSegment._PathSegmentChildRequirements = _child
	return _child, nil
}

func (m *_PortSegment) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_PortSegment) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("PortSegment"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for PortSegment")
		}

		// Simple Field (segmentType)
		if pushErr := writeBuffer.PushContext("segmentType"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for segmentType")
		}
		_segmentTypeErr := writeBuffer.WriteSerializable(ctx, m.GetSegmentType())
		if popErr := writeBuffer.PopContext("segmentType"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for segmentType")
		}
		if _segmentTypeErr != nil {
			return errors.Wrap(_segmentTypeErr, "Error serializing 'segmentType' field")
		}

		if popErr := writeBuffer.PopContext("PortSegment"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for PortSegment")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_PortSegment) isPortSegment() bool {
	return true
}

func (m *_PortSegment) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
