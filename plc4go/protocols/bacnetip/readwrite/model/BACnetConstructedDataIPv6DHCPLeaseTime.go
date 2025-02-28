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

// BACnetConstructedDataIPv6DHCPLeaseTime is the corresponding interface of BACnetConstructedDataIPv6DHCPLeaseTime
type BACnetConstructedDataIPv6DHCPLeaseTime interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetIpv6DhcpLeaseTime returns Ipv6DhcpLeaseTime (property field)
	GetIpv6DhcpLeaseTime() BACnetApplicationTagUnsignedInteger
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagUnsignedInteger
}

// BACnetConstructedDataIPv6DHCPLeaseTimeExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataIPv6DHCPLeaseTime.
// This is useful for switch cases.
type BACnetConstructedDataIPv6DHCPLeaseTimeExactly interface {
	BACnetConstructedDataIPv6DHCPLeaseTime
	isBACnetConstructedDataIPv6DHCPLeaseTime() bool
}

// _BACnetConstructedDataIPv6DHCPLeaseTime is the data-structure of this message
type _BACnetConstructedDataIPv6DHCPLeaseTime struct {
	*_BACnetConstructedData
	Ipv6DhcpLeaseTime BACnetApplicationTagUnsignedInteger
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataIPv6DHCPLeaseTime) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataIPv6DHCPLeaseTime) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_IPV6_DHCP_LEASE_TIME
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataIPv6DHCPLeaseTime) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataIPv6DHCPLeaseTime) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataIPv6DHCPLeaseTime) GetIpv6DhcpLeaseTime() BACnetApplicationTagUnsignedInteger {
	return m.Ipv6DhcpLeaseTime
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataIPv6DHCPLeaseTime) GetActualValue() BACnetApplicationTagUnsignedInteger {
	ctx := context.Background()
	_ = ctx
	return CastBACnetApplicationTagUnsignedInteger(m.GetIpv6DhcpLeaseTime())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataIPv6DHCPLeaseTime factory function for _BACnetConstructedDataIPv6DHCPLeaseTime
func NewBACnetConstructedDataIPv6DHCPLeaseTime(ipv6DhcpLeaseTime BACnetApplicationTagUnsignedInteger, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataIPv6DHCPLeaseTime {
	_result := &_BACnetConstructedDataIPv6DHCPLeaseTime{
		Ipv6DhcpLeaseTime:      ipv6DhcpLeaseTime,
		_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataIPv6DHCPLeaseTime(structType interface{}) BACnetConstructedDataIPv6DHCPLeaseTime {
	if casted, ok := structType.(BACnetConstructedDataIPv6DHCPLeaseTime); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataIPv6DHCPLeaseTime); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataIPv6DHCPLeaseTime) GetTypeName() string {
	return "BACnetConstructedDataIPv6DHCPLeaseTime"
}

func (m *_BACnetConstructedDataIPv6DHCPLeaseTime) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (ipv6DhcpLeaseTime)
	lengthInBits += m.Ipv6DhcpLeaseTime.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataIPv6DHCPLeaseTime) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetConstructedDataIPv6DHCPLeaseTimeParse(theBytes []byte, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataIPv6DHCPLeaseTime, error) {
	return BACnetConstructedDataIPv6DHCPLeaseTimeParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes), tagNumber, objectTypeArgument, propertyIdentifierArgument, arrayIndexArgument)
}

func BACnetConstructedDataIPv6DHCPLeaseTimeParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataIPv6DHCPLeaseTime, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataIPv6DHCPLeaseTime"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataIPv6DHCPLeaseTime")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (ipv6DhcpLeaseTime)
	if pullErr := readBuffer.PullContext("ipv6DhcpLeaseTime"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ipv6DhcpLeaseTime")
	}
	_ipv6DhcpLeaseTime, _ipv6DhcpLeaseTimeErr := BACnetApplicationTagParseWithBuffer(ctx, readBuffer)
	if _ipv6DhcpLeaseTimeErr != nil {
		return nil, errors.Wrap(_ipv6DhcpLeaseTimeErr, "Error parsing 'ipv6DhcpLeaseTime' field of BACnetConstructedDataIPv6DHCPLeaseTime")
	}
	ipv6DhcpLeaseTime := _ipv6DhcpLeaseTime.(BACnetApplicationTagUnsignedInteger)
	if closeErr := readBuffer.CloseContext("ipv6DhcpLeaseTime"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ipv6DhcpLeaseTime")
	}

	// Virtual field
	_actualValue := ipv6DhcpLeaseTime
	actualValue := _actualValue
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataIPv6DHCPLeaseTime"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataIPv6DHCPLeaseTime")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataIPv6DHCPLeaseTime{
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber:          tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
		Ipv6DhcpLeaseTime: ipv6DhcpLeaseTime,
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataIPv6DHCPLeaseTime) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataIPv6DHCPLeaseTime) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataIPv6DHCPLeaseTime"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataIPv6DHCPLeaseTime")
		}

		// Simple Field (ipv6DhcpLeaseTime)
		if pushErr := writeBuffer.PushContext("ipv6DhcpLeaseTime"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ipv6DhcpLeaseTime")
		}
		_ipv6DhcpLeaseTimeErr := writeBuffer.WriteSerializable(ctx, m.GetIpv6DhcpLeaseTime())
		if popErr := writeBuffer.PopContext("ipv6DhcpLeaseTime"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ipv6DhcpLeaseTime")
		}
		if _ipv6DhcpLeaseTimeErr != nil {
			return errors.Wrap(_ipv6DhcpLeaseTimeErr, "Error serializing 'ipv6DhcpLeaseTime' field")
		}
		// Virtual field
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataIPv6DHCPLeaseTime"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataIPv6DHCPLeaseTime")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataIPv6DHCPLeaseTime) isBACnetConstructedDataIPv6DHCPLeaseTime() bool {
	return true
}

func (m *_BACnetConstructedDataIPv6DHCPLeaseTime) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
