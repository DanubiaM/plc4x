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

// BACnetConstructedDataIPDHCPLeaseTimeRemaining is the corresponding interface of BACnetConstructedDataIPDHCPLeaseTimeRemaining
type BACnetConstructedDataIPDHCPLeaseTimeRemaining interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetIpDhcpLeaseTimeRemaining returns IpDhcpLeaseTimeRemaining (property field)
	GetIpDhcpLeaseTimeRemaining() BACnetApplicationTagUnsignedInteger
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagUnsignedInteger
}

// BACnetConstructedDataIPDHCPLeaseTimeRemainingExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataIPDHCPLeaseTimeRemaining.
// This is useful for switch cases.
type BACnetConstructedDataIPDHCPLeaseTimeRemainingExactly interface {
	BACnetConstructedDataIPDHCPLeaseTimeRemaining
	isBACnetConstructedDataIPDHCPLeaseTimeRemaining() bool
}

// _BACnetConstructedDataIPDHCPLeaseTimeRemaining is the data-structure of this message
type _BACnetConstructedDataIPDHCPLeaseTimeRemaining struct {
	*_BACnetConstructedData
	IpDhcpLeaseTimeRemaining BACnetApplicationTagUnsignedInteger
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataIPDHCPLeaseTimeRemaining) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataIPDHCPLeaseTimeRemaining) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_IP_DHCP_LEASE_TIME_REMAINING
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataIPDHCPLeaseTimeRemaining) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataIPDHCPLeaseTimeRemaining) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataIPDHCPLeaseTimeRemaining) GetIpDhcpLeaseTimeRemaining() BACnetApplicationTagUnsignedInteger {
	return m.IpDhcpLeaseTimeRemaining
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataIPDHCPLeaseTimeRemaining) GetActualValue() BACnetApplicationTagUnsignedInteger {
	ctx := context.Background()
	_ = ctx
	return CastBACnetApplicationTagUnsignedInteger(m.GetIpDhcpLeaseTimeRemaining())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataIPDHCPLeaseTimeRemaining factory function for _BACnetConstructedDataIPDHCPLeaseTimeRemaining
func NewBACnetConstructedDataIPDHCPLeaseTimeRemaining(ipDhcpLeaseTimeRemaining BACnetApplicationTagUnsignedInteger, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataIPDHCPLeaseTimeRemaining {
	_result := &_BACnetConstructedDataIPDHCPLeaseTimeRemaining{
		IpDhcpLeaseTimeRemaining: ipDhcpLeaseTimeRemaining,
		_BACnetConstructedData:   NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataIPDHCPLeaseTimeRemaining(structType interface{}) BACnetConstructedDataIPDHCPLeaseTimeRemaining {
	if casted, ok := structType.(BACnetConstructedDataIPDHCPLeaseTimeRemaining); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataIPDHCPLeaseTimeRemaining); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataIPDHCPLeaseTimeRemaining) GetTypeName() string {
	return "BACnetConstructedDataIPDHCPLeaseTimeRemaining"
}

func (m *_BACnetConstructedDataIPDHCPLeaseTimeRemaining) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (ipDhcpLeaseTimeRemaining)
	lengthInBits += m.IpDhcpLeaseTimeRemaining.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataIPDHCPLeaseTimeRemaining) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetConstructedDataIPDHCPLeaseTimeRemainingParse(theBytes []byte, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataIPDHCPLeaseTimeRemaining, error) {
	return BACnetConstructedDataIPDHCPLeaseTimeRemainingParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes), tagNumber, objectTypeArgument, propertyIdentifierArgument, arrayIndexArgument)
}

func BACnetConstructedDataIPDHCPLeaseTimeRemainingParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataIPDHCPLeaseTimeRemaining, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataIPDHCPLeaseTimeRemaining"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataIPDHCPLeaseTimeRemaining")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (ipDhcpLeaseTimeRemaining)
	if pullErr := readBuffer.PullContext("ipDhcpLeaseTimeRemaining"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ipDhcpLeaseTimeRemaining")
	}
	_ipDhcpLeaseTimeRemaining, _ipDhcpLeaseTimeRemainingErr := BACnetApplicationTagParseWithBuffer(ctx, readBuffer)
	if _ipDhcpLeaseTimeRemainingErr != nil {
		return nil, errors.Wrap(_ipDhcpLeaseTimeRemainingErr, "Error parsing 'ipDhcpLeaseTimeRemaining' field of BACnetConstructedDataIPDHCPLeaseTimeRemaining")
	}
	ipDhcpLeaseTimeRemaining := _ipDhcpLeaseTimeRemaining.(BACnetApplicationTagUnsignedInteger)
	if closeErr := readBuffer.CloseContext("ipDhcpLeaseTimeRemaining"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ipDhcpLeaseTimeRemaining")
	}

	// Virtual field
	_actualValue := ipDhcpLeaseTimeRemaining
	actualValue := _actualValue
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataIPDHCPLeaseTimeRemaining"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataIPDHCPLeaseTimeRemaining")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataIPDHCPLeaseTimeRemaining{
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber:          tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
		IpDhcpLeaseTimeRemaining: ipDhcpLeaseTimeRemaining,
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataIPDHCPLeaseTimeRemaining) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataIPDHCPLeaseTimeRemaining) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataIPDHCPLeaseTimeRemaining"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataIPDHCPLeaseTimeRemaining")
		}

		// Simple Field (ipDhcpLeaseTimeRemaining)
		if pushErr := writeBuffer.PushContext("ipDhcpLeaseTimeRemaining"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ipDhcpLeaseTimeRemaining")
		}
		_ipDhcpLeaseTimeRemainingErr := writeBuffer.WriteSerializable(ctx, m.GetIpDhcpLeaseTimeRemaining())
		if popErr := writeBuffer.PopContext("ipDhcpLeaseTimeRemaining"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ipDhcpLeaseTimeRemaining")
		}
		if _ipDhcpLeaseTimeRemainingErr != nil {
			return errors.Wrap(_ipDhcpLeaseTimeRemainingErr, "Error serializing 'ipDhcpLeaseTimeRemaining' field")
		}
		// Virtual field
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataIPDHCPLeaseTimeRemaining"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataIPDHCPLeaseTimeRemaining")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataIPDHCPLeaseTimeRemaining) isBACnetConstructedDataIPDHCPLeaseTimeRemaining() bool {
	return true
}

func (m *_BACnetConstructedDataIPDHCPLeaseTimeRemaining) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
