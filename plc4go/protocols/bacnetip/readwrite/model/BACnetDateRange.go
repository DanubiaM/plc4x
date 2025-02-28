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

// BACnetDateRange is the corresponding interface of BACnetDateRange
type BACnetDateRange interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetStartDate returns StartDate (property field)
	GetStartDate() BACnetApplicationTagDate
	// GetEndDate returns EndDate (property field)
	GetEndDate() BACnetApplicationTagDate
}

// BACnetDateRangeExactly can be used when we want exactly this type and not a type which fulfills BACnetDateRange.
// This is useful for switch cases.
type BACnetDateRangeExactly interface {
	BACnetDateRange
	isBACnetDateRange() bool
}

// _BACnetDateRange is the data-structure of this message
type _BACnetDateRange struct {
	StartDate BACnetApplicationTagDate
	EndDate   BACnetApplicationTagDate
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetDateRange) GetStartDate() BACnetApplicationTagDate {
	return m.StartDate
}

func (m *_BACnetDateRange) GetEndDate() BACnetApplicationTagDate {
	return m.EndDate
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetDateRange factory function for _BACnetDateRange
func NewBACnetDateRange(startDate BACnetApplicationTagDate, endDate BACnetApplicationTagDate) *_BACnetDateRange {
	return &_BACnetDateRange{StartDate: startDate, EndDate: endDate}
}

// Deprecated: use the interface for direct cast
func CastBACnetDateRange(structType interface{}) BACnetDateRange {
	if casted, ok := structType.(BACnetDateRange); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetDateRange); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetDateRange) GetTypeName() string {
	return "BACnetDateRange"
}

func (m *_BACnetDateRange) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (startDate)
	lengthInBits += m.StartDate.GetLengthInBits(ctx)

	// Simple field (endDate)
	lengthInBits += m.EndDate.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetDateRange) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetDateRangeParse(theBytes []byte) (BACnetDateRange, error) {
	return BACnetDateRangeParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes))
}

func BACnetDateRangeParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetDateRange, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetDateRange"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetDateRange")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (startDate)
	if pullErr := readBuffer.PullContext("startDate"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for startDate")
	}
	_startDate, _startDateErr := BACnetApplicationTagParseWithBuffer(ctx, readBuffer)
	if _startDateErr != nil {
		return nil, errors.Wrap(_startDateErr, "Error parsing 'startDate' field of BACnetDateRange")
	}
	startDate := _startDate.(BACnetApplicationTagDate)
	if closeErr := readBuffer.CloseContext("startDate"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for startDate")
	}

	// Simple Field (endDate)
	if pullErr := readBuffer.PullContext("endDate"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for endDate")
	}
	_endDate, _endDateErr := BACnetApplicationTagParseWithBuffer(ctx, readBuffer)
	if _endDateErr != nil {
		return nil, errors.Wrap(_endDateErr, "Error parsing 'endDate' field of BACnetDateRange")
	}
	endDate := _endDate.(BACnetApplicationTagDate)
	if closeErr := readBuffer.CloseContext("endDate"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for endDate")
	}

	if closeErr := readBuffer.CloseContext("BACnetDateRange"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetDateRange")
	}

	// Create the instance
	return &_BACnetDateRange{
		StartDate: startDate,
		EndDate:   endDate,
	}, nil
}

func (m *_BACnetDateRange) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetDateRange) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	if pushErr := writeBuffer.PushContext("BACnetDateRange"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetDateRange")
	}

	// Simple Field (startDate)
	if pushErr := writeBuffer.PushContext("startDate"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for startDate")
	}
	_startDateErr := writeBuffer.WriteSerializable(ctx, m.GetStartDate())
	if popErr := writeBuffer.PopContext("startDate"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for startDate")
	}
	if _startDateErr != nil {
		return errors.Wrap(_startDateErr, "Error serializing 'startDate' field")
	}

	// Simple Field (endDate)
	if pushErr := writeBuffer.PushContext("endDate"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for endDate")
	}
	_endDateErr := writeBuffer.WriteSerializable(ctx, m.GetEndDate())
	if popErr := writeBuffer.PopContext("endDate"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for endDate")
	}
	if _endDateErr != nil {
		return errors.Wrap(_endDateErr, "Error serializing 'endDate' field")
	}

	if popErr := writeBuffer.PopContext("BACnetDateRange"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetDateRange")
	}
	return nil
}

func (m *_BACnetDateRange) isBACnetDateRange() bool {
	return true
}

func (m *_BACnetDateRange) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
