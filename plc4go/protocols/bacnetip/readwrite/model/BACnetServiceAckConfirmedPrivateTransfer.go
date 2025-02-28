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
	"io"
)

// Code generated by code-generation. DO NOT EDIT.

// BACnetServiceAckConfirmedPrivateTransfer is the corresponding interface of BACnetServiceAckConfirmedPrivateTransfer
type BACnetServiceAckConfirmedPrivateTransfer interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetServiceAck
	// GetVendorId returns VendorId (property field)
	GetVendorId() BACnetVendorIdTagged
	// GetServiceNumber returns ServiceNumber (property field)
	GetServiceNumber() BACnetContextTagUnsignedInteger
	// GetResultBlock returns ResultBlock (property field)
	GetResultBlock() BACnetConstructedData
}

// BACnetServiceAckConfirmedPrivateTransferExactly can be used when we want exactly this type and not a type which fulfills BACnetServiceAckConfirmedPrivateTransfer.
// This is useful for switch cases.
type BACnetServiceAckConfirmedPrivateTransferExactly interface {
	BACnetServiceAckConfirmedPrivateTransfer
	isBACnetServiceAckConfirmedPrivateTransfer() bool
}

// _BACnetServiceAckConfirmedPrivateTransfer is the data-structure of this message
type _BACnetServiceAckConfirmedPrivateTransfer struct {
	*_BACnetServiceAck
	VendorId      BACnetVendorIdTagged
	ServiceNumber BACnetContextTagUnsignedInteger
	ResultBlock   BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetServiceAckConfirmedPrivateTransfer) GetServiceChoice() BACnetConfirmedServiceChoice {
	return BACnetConfirmedServiceChoice_CONFIRMED_PRIVATE_TRANSFER
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetServiceAckConfirmedPrivateTransfer) InitializeParent(parent BACnetServiceAck) {}

func (m *_BACnetServiceAckConfirmedPrivateTransfer) GetParent() BACnetServiceAck {
	return m._BACnetServiceAck
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetServiceAckConfirmedPrivateTransfer) GetVendorId() BACnetVendorIdTagged {
	return m.VendorId
}

func (m *_BACnetServiceAckConfirmedPrivateTransfer) GetServiceNumber() BACnetContextTagUnsignedInteger {
	return m.ServiceNumber
}

func (m *_BACnetServiceAckConfirmedPrivateTransfer) GetResultBlock() BACnetConstructedData {
	return m.ResultBlock
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetServiceAckConfirmedPrivateTransfer factory function for _BACnetServiceAckConfirmedPrivateTransfer
func NewBACnetServiceAckConfirmedPrivateTransfer(vendorId BACnetVendorIdTagged, serviceNumber BACnetContextTagUnsignedInteger, resultBlock BACnetConstructedData, serviceAckLength uint32) *_BACnetServiceAckConfirmedPrivateTransfer {
	_result := &_BACnetServiceAckConfirmedPrivateTransfer{
		VendorId:          vendorId,
		ServiceNumber:     serviceNumber,
		ResultBlock:       resultBlock,
		_BACnetServiceAck: NewBACnetServiceAck(serviceAckLength),
	}
	_result._BACnetServiceAck._BACnetServiceAckChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetServiceAckConfirmedPrivateTransfer(structType interface{}) BACnetServiceAckConfirmedPrivateTransfer {
	if casted, ok := structType.(BACnetServiceAckConfirmedPrivateTransfer); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetServiceAckConfirmedPrivateTransfer); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetServiceAckConfirmedPrivateTransfer) GetTypeName() string {
	return "BACnetServiceAckConfirmedPrivateTransfer"
}

func (m *_BACnetServiceAckConfirmedPrivateTransfer) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (vendorId)
	lengthInBits += m.VendorId.GetLengthInBits(ctx)

	// Simple field (serviceNumber)
	lengthInBits += m.ServiceNumber.GetLengthInBits(ctx)

	// Optional Field (resultBlock)
	if m.ResultBlock != nil {
		lengthInBits += m.ResultBlock.GetLengthInBits(ctx)
	}

	return lengthInBits
}

func (m *_BACnetServiceAckConfirmedPrivateTransfer) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetServiceAckConfirmedPrivateTransferParse(theBytes []byte, serviceAckLength uint32) (BACnetServiceAckConfirmedPrivateTransfer, error) {
	return BACnetServiceAckConfirmedPrivateTransferParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes), serviceAckLength)
}

func BACnetServiceAckConfirmedPrivateTransferParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, serviceAckLength uint32) (BACnetServiceAckConfirmedPrivateTransfer, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetServiceAckConfirmedPrivateTransfer"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetServiceAckConfirmedPrivateTransfer")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (vendorId)
	if pullErr := readBuffer.PullContext("vendorId"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for vendorId")
	}
	_vendorId, _vendorIdErr := BACnetVendorIdTaggedParseWithBuffer(ctx, readBuffer, uint8(uint8(0)), TagClass(TagClass_CONTEXT_SPECIFIC_TAGS))
	if _vendorIdErr != nil {
		return nil, errors.Wrap(_vendorIdErr, "Error parsing 'vendorId' field of BACnetServiceAckConfirmedPrivateTransfer")
	}
	vendorId := _vendorId.(BACnetVendorIdTagged)
	if closeErr := readBuffer.CloseContext("vendorId"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for vendorId")
	}

	// Simple Field (serviceNumber)
	if pullErr := readBuffer.PullContext("serviceNumber"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for serviceNumber")
	}
	_serviceNumber, _serviceNumberErr := BACnetContextTagParseWithBuffer(ctx, readBuffer, uint8(uint8(1)), BACnetDataType(BACnetDataType_UNSIGNED_INTEGER))
	if _serviceNumberErr != nil {
		return nil, errors.Wrap(_serviceNumberErr, "Error parsing 'serviceNumber' field of BACnetServiceAckConfirmedPrivateTransfer")
	}
	serviceNumber := _serviceNumber.(BACnetContextTagUnsignedInteger)
	if closeErr := readBuffer.CloseContext("serviceNumber"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for serviceNumber")
	}

	// Optional Field (resultBlock) (Can be skipped, if a given expression evaluates to false)
	var resultBlock BACnetConstructedData = nil
	{
		currentPos = positionAware.GetPos()
		if pullErr := readBuffer.PullContext("resultBlock"); pullErr != nil {
			return nil, errors.Wrap(pullErr, "Error pulling for resultBlock")
		}
		_val, _err := BACnetConstructedDataParseWithBuffer(ctx, readBuffer, uint8(2), BACnetObjectType_VENDOR_PROPRIETARY_VALUE, BACnetPropertyIdentifier_VENDOR_PROPRIETARY_VALUE, nil)
		switch {
		case errors.Is(_err, utils.ParseAssertError{}) || errors.Is(_err, io.EOF):
			Plc4xModelLog.Debug().Err(_err).Msg("Resetting position because optional threw an error")
			readBuffer.Reset(currentPos)
		case _err != nil:
			return nil, errors.Wrap(_err, "Error parsing 'resultBlock' field of BACnetServiceAckConfirmedPrivateTransfer")
		default:
			resultBlock = _val.(BACnetConstructedData)
			if closeErr := readBuffer.CloseContext("resultBlock"); closeErr != nil {
				return nil, errors.Wrap(closeErr, "Error closing for resultBlock")
			}
		}
	}

	if closeErr := readBuffer.CloseContext("BACnetServiceAckConfirmedPrivateTransfer"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetServiceAckConfirmedPrivateTransfer")
	}

	// Create a partially initialized instance
	_child := &_BACnetServiceAckConfirmedPrivateTransfer{
		_BACnetServiceAck: &_BACnetServiceAck{
			ServiceAckLength: serviceAckLength,
		},
		VendorId:      vendorId,
		ServiceNumber: serviceNumber,
		ResultBlock:   resultBlock,
	}
	_child._BACnetServiceAck._BACnetServiceAckChildRequirements = _child
	return _child, nil
}

func (m *_BACnetServiceAckConfirmedPrivateTransfer) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetServiceAckConfirmedPrivateTransfer) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetServiceAckConfirmedPrivateTransfer"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetServiceAckConfirmedPrivateTransfer")
		}

		// Simple Field (vendorId)
		if pushErr := writeBuffer.PushContext("vendorId"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for vendorId")
		}
		_vendorIdErr := writeBuffer.WriteSerializable(ctx, m.GetVendorId())
		if popErr := writeBuffer.PopContext("vendorId"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for vendorId")
		}
		if _vendorIdErr != nil {
			return errors.Wrap(_vendorIdErr, "Error serializing 'vendorId' field")
		}

		// Simple Field (serviceNumber)
		if pushErr := writeBuffer.PushContext("serviceNumber"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for serviceNumber")
		}
		_serviceNumberErr := writeBuffer.WriteSerializable(ctx, m.GetServiceNumber())
		if popErr := writeBuffer.PopContext("serviceNumber"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for serviceNumber")
		}
		if _serviceNumberErr != nil {
			return errors.Wrap(_serviceNumberErr, "Error serializing 'serviceNumber' field")
		}

		// Optional Field (resultBlock) (Can be skipped, if the value is null)
		var resultBlock BACnetConstructedData = nil
		if m.GetResultBlock() != nil {
			if pushErr := writeBuffer.PushContext("resultBlock"); pushErr != nil {
				return errors.Wrap(pushErr, "Error pushing for resultBlock")
			}
			resultBlock = m.GetResultBlock()
			_resultBlockErr := writeBuffer.WriteSerializable(ctx, resultBlock)
			if popErr := writeBuffer.PopContext("resultBlock"); popErr != nil {
				return errors.Wrap(popErr, "Error popping for resultBlock")
			}
			if _resultBlockErr != nil {
				return errors.Wrap(_resultBlockErr, "Error serializing 'resultBlock' field")
			}
		}

		if popErr := writeBuffer.PopContext("BACnetServiceAckConfirmedPrivateTransfer"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetServiceAckConfirmedPrivateTransfer")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetServiceAckConfirmedPrivateTransfer) isBACnetServiceAckConfirmedPrivateTransfer() bool {
	return true
}

func (m *_BACnetServiceAckConfirmedPrivateTransfer) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
