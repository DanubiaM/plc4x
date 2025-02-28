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

// BACnetConfirmedServiceRequestReadPropertyConditional is the corresponding interface of BACnetConfirmedServiceRequestReadPropertyConditional
type BACnetConfirmedServiceRequestReadPropertyConditional interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetConfirmedServiceRequest
	// GetBytesOfRemovedService returns BytesOfRemovedService (property field)
	GetBytesOfRemovedService() []byte
}

// BACnetConfirmedServiceRequestReadPropertyConditionalExactly can be used when we want exactly this type and not a type which fulfills BACnetConfirmedServiceRequestReadPropertyConditional.
// This is useful for switch cases.
type BACnetConfirmedServiceRequestReadPropertyConditionalExactly interface {
	BACnetConfirmedServiceRequestReadPropertyConditional
	isBACnetConfirmedServiceRequestReadPropertyConditional() bool
}

// _BACnetConfirmedServiceRequestReadPropertyConditional is the data-structure of this message
type _BACnetConfirmedServiceRequestReadPropertyConditional struct {
	*_BACnetConfirmedServiceRequest
	BytesOfRemovedService []byte

	// Arguments.
	ServiceRequestPayloadLength uint32
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConfirmedServiceRequestReadPropertyConditional) GetServiceChoice() BACnetConfirmedServiceChoice {
	return BACnetConfirmedServiceChoice_READ_PROPERTY_CONDITIONAL
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConfirmedServiceRequestReadPropertyConditional) InitializeParent(parent BACnetConfirmedServiceRequest) {
}

func (m *_BACnetConfirmedServiceRequestReadPropertyConditional) GetParent() BACnetConfirmedServiceRequest {
	return m._BACnetConfirmedServiceRequest
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConfirmedServiceRequestReadPropertyConditional) GetBytesOfRemovedService() []byte {
	return m.BytesOfRemovedService
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConfirmedServiceRequestReadPropertyConditional factory function for _BACnetConfirmedServiceRequestReadPropertyConditional
func NewBACnetConfirmedServiceRequestReadPropertyConditional(bytesOfRemovedService []byte, serviceRequestPayloadLength uint32, serviceRequestLength uint32) *_BACnetConfirmedServiceRequestReadPropertyConditional {
	_result := &_BACnetConfirmedServiceRequestReadPropertyConditional{
		BytesOfRemovedService:          bytesOfRemovedService,
		_BACnetConfirmedServiceRequest: NewBACnetConfirmedServiceRequest(serviceRequestLength),
	}
	_result._BACnetConfirmedServiceRequest._BACnetConfirmedServiceRequestChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConfirmedServiceRequestReadPropertyConditional(structType interface{}) BACnetConfirmedServiceRequestReadPropertyConditional {
	if casted, ok := structType.(BACnetConfirmedServiceRequestReadPropertyConditional); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConfirmedServiceRequestReadPropertyConditional); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConfirmedServiceRequestReadPropertyConditional) GetTypeName() string {
	return "BACnetConfirmedServiceRequestReadPropertyConditional"
}

func (m *_BACnetConfirmedServiceRequestReadPropertyConditional) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Array field
	if len(m.BytesOfRemovedService) > 0 {
		lengthInBits += 8 * uint16(len(m.BytesOfRemovedService))
	}

	return lengthInBits
}

func (m *_BACnetConfirmedServiceRequestReadPropertyConditional) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetConfirmedServiceRequestReadPropertyConditionalParse(theBytes []byte, serviceRequestPayloadLength uint32, serviceRequestLength uint32) (BACnetConfirmedServiceRequestReadPropertyConditional, error) {
	return BACnetConfirmedServiceRequestReadPropertyConditionalParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes), serviceRequestPayloadLength, serviceRequestLength)
}

func BACnetConfirmedServiceRequestReadPropertyConditionalParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, serviceRequestPayloadLength uint32, serviceRequestLength uint32) (BACnetConfirmedServiceRequestReadPropertyConditional, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConfirmedServiceRequestReadPropertyConditional"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConfirmedServiceRequestReadPropertyConditional")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos
	// Byte Array field (bytesOfRemovedService)
	numberOfBytesbytesOfRemovedService := int(serviceRequestPayloadLength)
	bytesOfRemovedService, _readArrayErr := readBuffer.ReadByteArray("bytesOfRemovedService", numberOfBytesbytesOfRemovedService)
	if _readArrayErr != nil {
		return nil, errors.Wrap(_readArrayErr, "Error parsing 'bytesOfRemovedService' field of BACnetConfirmedServiceRequestReadPropertyConditional")
	}

	if closeErr := readBuffer.CloseContext("BACnetConfirmedServiceRequestReadPropertyConditional"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConfirmedServiceRequestReadPropertyConditional")
	}

	// Create a partially initialized instance
	_child := &_BACnetConfirmedServiceRequestReadPropertyConditional{
		_BACnetConfirmedServiceRequest: &_BACnetConfirmedServiceRequest{
			ServiceRequestLength: serviceRequestLength,
		},
		BytesOfRemovedService: bytesOfRemovedService,
	}
	_child._BACnetConfirmedServiceRequest._BACnetConfirmedServiceRequestChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConfirmedServiceRequestReadPropertyConditional) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConfirmedServiceRequestReadPropertyConditional) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConfirmedServiceRequestReadPropertyConditional"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConfirmedServiceRequestReadPropertyConditional")
		}

		// Array Field (bytesOfRemovedService)
		// Byte Array field (bytesOfRemovedService)
		if err := writeBuffer.WriteByteArray("bytesOfRemovedService", m.GetBytesOfRemovedService()); err != nil {
			return errors.Wrap(err, "Error serializing 'bytesOfRemovedService' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConfirmedServiceRequestReadPropertyConditional"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConfirmedServiceRequestReadPropertyConditional")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

////
// Arguments Getter

func (m *_BACnetConfirmedServiceRequestReadPropertyConditional) GetServiceRequestPayloadLength() uint32 {
	return m.ServiceRequestPayloadLength
}

//
////

func (m *_BACnetConfirmedServiceRequestReadPropertyConditional) isBACnetConfirmedServiceRequestReadPropertyConditional() bool {
	return true
}

func (m *_BACnetConfirmedServiceRequestReadPropertyConditional) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
