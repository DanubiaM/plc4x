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

// BACnetUnconfirmedServiceRequestUTCTimeSynchronization is the corresponding interface of BACnetUnconfirmedServiceRequestUTCTimeSynchronization
type BACnetUnconfirmedServiceRequestUTCTimeSynchronization interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetUnconfirmedServiceRequest
	// GetSynchronizedDate returns SynchronizedDate (property field)
	GetSynchronizedDate() BACnetApplicationTagDate
	// GetSynchronizedTime returns SynchronizedTime (property field)
	GetSynchronizedTime() BACnetApplicationTagTime
}

// BACnetUnconfirmedServiceRequestUTCTimeSynchronizationExactly can be used when we want exactly this type and not a type which fulfills BACnetUnconfirmedServiceRequestUTCTimeSynchronization.
// This is useful for switch cases.
type BACnetUnconfirmedServiceRequestUTCTimeSynchronizationExactly interface {
	BACnetUnconfirmedServiceRequestUTCTimeSynchronization
	isBACnetUnconfirmedServiceRequestUTCTimeSynchronization() bool
}

// _BACnetUnconfirmedServiceRequestUTCTimeSynchronization is the data-structure of this message
type _BACnetUnconfirmedServiceRequestUTCTimeSynchronization struct {
	*_BACnetUnconfirmedServiceRequest
	SynchronizedDate BACnetApplicationTagDate
	SynchronizedTime BACnetApplicationTagTime
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetUnconfirmedServiceRequestUTCTimeSynchronization) GetServiceChoice() BACnetUnconfirmedServiceChoice {
	return BACnetUnconfirmedServiceChoice_UTC_TIME_SYNCHRONIZATION
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetUnconfirmedServiceRequestUTCTimeSynchronization) InitializeParent(parent BACnetUnconfirmedServiceRequest) {
}

func (m *_BACnetUnconfirmedServiceRequestUTCTimeSynchronization) GetParent() BACnetUnconfirmedServiceRequest {
	return m._BACnetUnconfirmedServiceRequest
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetUnconfirmedServiceRequestUTCTimeSynchronization) GetSynchronizedDate() BACnetApplicationTagDate {
	return m.SynchronizedDate
}

func (m *_BACnetUnconfirmedServiceRequestUTCTimeSynchronization) GetSynchronizedTime() BACnetApplicationTagTime {
	return m.SynchronizedTime
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetUnconfirmedServiceRequestUTCTimeSynchronization factory function for _BACnetUnconfirmedServiceRequestUTCTimeSynchronization
func NewBACnetUnconfirmedServiceRequestUTCTimeSynchronization(synchronizedDate BACnetApplicationTagDate, synchronizedTime BACnetApplicationTagTime, serviceRequestLength uint16) *_BACnetUnconfirmedServiceRequestUTCTimeSynchronization {
	_result := &_BACnetUnconfirmedServiceRequestUTCTimeSynchronization{
		SynchronizedDate:                 synchronizedDate,
		SynchronizedTime:                 synchronizedTime,
		_BACnetUnconfirmedServiceRequest: NewBACnetUnconfirmedServiceRequest(serviceRequestLength),
	}
	_result._BACnetUnconfirmedServiceRequest._BACnetUnconfirmedServiceRequestChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetUnconfirmedServiceRequestUTCTimeSynchronization(structType interface{}) BACnetUnconfirmedServiceRequestUTCTimeSynchronization {
	if casted, ok := structType.(BACnetUnconfirmedServiceRequestUTCTimeSynchronization); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetUnconfirmedServiceRequestUTCTimeSynchronization); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetUnconfirmedServiceRequestUTCTimeSynchronization) GetTypeName() string {
	return "BACnetUnconfirmedServiceRequestUTCTimeSynchronization"
}

func (m *_BACnetUnconfirmedServiceRequestUTCTimeSynchronization) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (synchronizedDate)
	lengthInBits += m.SynchronizedDate.GetLengthInBits(ctx)

	// Simple field (synchronizedTime)
	lengthInBits += m.SynchronizedTime.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetUnconfirmedServiceRequestUTCTimeSynchronization) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetUnconfirmedServiceRequestUTCTimeSynchronizationParse(theBytes []byte, serviceRequestLength uint16) (BACnetUnconfirmedServiceRequestUTCTimeSynchronization, error) {
	return BACnetUnconfirmedServiceRequestUTCTimeSynchronizationParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes), serviceRequestLength)
}

func BACnetUnconfirmedServiceRequestUTCTimeSynchronizationParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, serviceRequestLength uint16) (BACnetUnconfirmedServiceRequestUTCTimeSynchronization, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetUnconfirmedServiceRequestUTCTimeSynchronization"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetUnconfirmedServiceRequestUTCTimeSynchronization")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (synchronizedDate)
	if pullErr := readBuffer.PullContext("synchronizedDate"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for synchronizedDate")
	}
	_synchronizedDate, _synchronizedDateErr := BACnetApplicationTagParseWithBuffer(ctx, readBuffer)
	if _synchronizedDateErr != nil {
		return nil, errors.Wrap(_synchronizedDateErr, "Error parsing 'synchronizedDate' field of BACnetUnconfirmedServiceRequestUTCTimeSynchronization")
	}
	synchronizedDate := _synchronizedDate.(BACnetApplicationTagDate)
	if closeErr := readBuffer.CloseContext("synchronizedDate"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for synchronizedDate")
	}

	// Simple Field (synchronizedTime)
	if pullErr := readBuffer.PullContext("synchronizedTime"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for synchronizedTime")
	}
	_synchronizedTime, _synchronizedTimeErr := BACnetApplicationTagParseWithBuffer(ctx, readBuffer)
	if _synchronizedTimeErr != nil {
		return nil, errors.Wrap(_synchronizedTimeErr, "Error parsing 'synchronizedTime' field of BACnetUnconfirmedServiceRequestUTCTimeSynchronization")
	}
	synchronizedTime := _synchronizedTime.(BACnetApplicationTagTime)
	if closeErr := readBuffer.CloseContext("synchronizedTime"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for synchronizedTime")
	}

	if closeErr := readBuffer.CloseContext("BACnetUnconfirmedServiceRequestUTCTimeSynchronization"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetUnconfirmedServiceRequestUTCTimeSynchronization")
	}

	// Create a partially initialized instance
	_child := &_BACnetUnconfirmedServiceRequestUTCTimeSynchronization{
		_BACnetUnconfirmedServiceRequest: &_BACnetUnconfirmedServiceRequest{
			ServiceRequestLength: serviceRequestLength,
		},
		SynchronizedDate: synchronizedDate,
		SynchronizedTime: synchronizedTime,
	}
	_child._BACnetUnconfirmedServiceRequest._BACnetUnconfirmedServiceRequestChildRequirements = _child
	return _child, nil
}

func (m *_BACnetUnconfirmedServiceRequestUTCTimeSynchronization) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetUnconfirmedServiceRequestUTCTimeSynchronization) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetUnconfirmedServiceRequestUTCTimeSynchronization"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetUnconfirmedServiceRequestUTCTimeSynchronization")
		}

		// Simple Field (synchronizedDate)
		if pushErr := writeBuffer.PushContext("synchronizedDate"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for synchronizedDate")
		}
		_synchronizedDateErr := writeBuffer.WriteSerializable(ctx, m.GetSynchronizedDate())
		if popErr := writeBuffer.PopContext("synchronizedDate"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for synchronizedDate")
		}
		if _synchronizedDateErr != nil {
			return errors.Wrap(_synchronizedDateErr, "Error serializing 'synchronizedDate' field")
		}

		// Simple Field (synchronizedTime)
		if pushErr := writeBuffer.PushContext("synchronizedTime"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for synchronizedTime")
		}
		_synchronizedTimeErr := writeBuffer.WriteSerializable(ctx, m.GetSynchronizedTime())
		if popErr := writeBuffer.PopContext("synchronizedTime"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for synchronizedTime")
		}
		if _synchronizedTimeErr != nil {
			return errors.Wrap(_synchronizedTimeErr, "Error serializing 'synchronizedTime' field")
		}

		if popErr := writeBuffer.PopContext("BACnetUnconfirmedServiceRequestUTCTimeSynchronization"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetUnconfirmedServiceRequestUTCTimeSynchronization")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetUnconfirmedServiceRequestUTCTimeSynchronization) isBACnetUnconfirmedServiceRequestUTCTimeSynchronization() bool {
	return true
}

func (m *_BACnetUnconfirmedServiceRequestUTCTimeSynchronization) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
