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

// DateAndTime is the corresponding interface of DateAndTime
type DateAndTime interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetYear returns Year (property field)
	GetYear() uint8
	// GetMonth returns Month (property field)
	GetMonth() uint8
	// GetDay returns Day (property field)
	GetDay() uint8
	// GetHour returns Hour (property field)
	GetHour() uint8
	// GetMinutes returns Minutes (property field)
	GetMinutes() uint8
	// GetSeconds returns Seconds (property field)
	GetSeconds() uint8
	// GetMsec returns Msec (property field)
	GetMsec() uint16
	// GetDow returns Dow (property field)
	GetDow() uint8
}

// DateAndTimeExactly can be used when we want exactly this type and not a type which fulfills DateAndTime.
// This is useful for switch cases.
type DateAndTimeExactly interface {
	DateAndTime
	isDateAndTime() bool
}

// _DateAndTime is the data-structure of this message
type _DateAndTime struct {
	Year    uint8
	Month   uint8
	Day     uint8
	Hour    uint8
	Minutes uint8
	Seconds uint8
	Msec    uint16
	Dow     uint8
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_DateAndTime) GetYear() uint8 {
	return m.Year
}

func (m *_DateAndTime) GetMonth() uint8 {
	return m.Month
}

func (m *_DateAndTime) GetDay() uint8 {
	return m.Day
}

func (m *_DateAndTime) GetHour() uint8 {
	return m.Hour
}

func (m *_DateAndTime) GetMinutes() uint8 {
	return m.Minutes
}

func (m *_DateAndTime) GetSeconds() uint8 {
	return m.Seconds
}

func (m *_DateAndTime) GetMsec() uint16 {
	return m.Msec
}

func (m *_DateAndTime) GetDow() uint8 {
	return m.Dow
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewDateAndTime factory function for _DateAndTime
func NewDateAndTime(year uint8, month uint8, day uint8, hour uint8, minutes uint8, seconds uint8, msec uint16, dow uint8) *_DateAndTime {
	return &_DateAndTime{Year: year, Month: month, Day: day, Hour: hour, Minutes: minutes, Seconds: seconds, Msec: msec, Dow: dow}
}

// Deprecated: use the interface for direct cast
func CastDateAndTime(structType interface{}) DateAndTime {
	if casted, ok := structType.(DateAndTime); ok {
		return casted
	}
	if casted, ok := structType.(*DateAndTime); ok {
		return *casted
	}
	return nil
}

func (m *_DateAndTime) GetTypeName() string {
	return "DateAndTime"
}

func (m *_DateAndTime) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Manual Field (year)
	lengthInBits += uint16(int32(8))

	// Manual Field (month)
	lengthInBits += uint16(int32(8))

	// Manual Field (day)
	lengthInBits += uint16(int32(8))

	// Manual Field (hour)
	lengthInBits += uint16(int32(8))

	// Manual Field (minutes)
	lengthInBits += uint16(int32(8))

	// Manual Field (seconds)
	lengthInBits += uint16(int32(8))

	// Manual Field (msec)
	lengthInBits += uint16(int32(12))

	// Simple field (dow)
	lengthInBits += 4

	return lengthInBits
}

func (m *_DateAndTime) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func DateAndTimeParse(theBytes []byte) (DateAndTime, error) {
	return DateAndTimeParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes))
}

func DateAndTimeParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (DateAndTime, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("DateAndTime"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for DateAndTime")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Manual Field (year)
	_year, _yearErr := BcdToInt(readBuffer)
	if _yearErr != nil {
		return nil, errors.Wrap(_yearErr, "Error parsing 'year' field of DateAndTime")
	}
	var year uint8
	if _year != nil {
		year = _year.(uint8)
	}

	// Manual Field (month)
	_month, _monthErr := BcdToInt(readBuffer)
	if _monthErr != nil {
		return nil, errors.Wrap(_monthErr, "Error parsing 'month' field of DateAndTime")
	}
	var month uint8
	if _month != nil {
		month = _month.(uint8)
	}

	// Manual Field (day)
	_day, _dayErr := BcdToInt(readBuffer)
	if _dayErr != nil {
		return nil, errors.Wrap(_dayErr, "Error parsing 'day' field of DateAndTime")
	}
	var day uint8
	if _day != nil {
		day = _day.(uint8)
	}

	// Manual Field (hour)
	_hour, _hourErr := BcdToInt(readBuffer)
	if _hourErr != nil {
		return nil, errors.Wrap(_hourErr, "Error parsing 'hour' field of DateAndTime")
	}
	var hour uint8
	if _hour != nil {
		hour = _hour.(uint8)
	}

	// Manual Field (minutes)
	_minutes, _minutesErr := BcdToInt(readBuffer)
	if _minutesErr != nil {
		return nil, errors.Wrap(_minutesErr, "Error parsing 'minutes' field of DateAndTime")
	}
	var minutes uint8
	if _minutes != nil {
		minutes = _minutes.(uint8)
	}

	// Manual Field (seconds)
	_seconds, _secondsErr := BcdToInt(readBuffer)
	if _secondsErr != nil {
		return nil, errors.Wrap(_secondsErr, "Error parsing 'seconds' field of DateAndTime")
	}
	var seconds uint8
	if _seconds != nil {
		seconds = _seconds.(uint8)
	}

	// Manual Field (msec)
	_msec, _msecErr := S7msecToInt(readBuffer)
	if _msecErr != nil {
		return nil, errors.Wrap(_msecErr, "Error parsing 'msec' field of DateAndTime")
	}
	var msec uint16
	if _msec != nil {
		msec = _msec.(uint16)
	}

	// Simple Field (dow)
	_dow, _dowErr := readBuffer.ReadUint8("dow", 4)
	if _dowErr != nil {
		return nil, errors.Wrap(_dowErr, "Error parsing 'dow' field of DateAndTime")
	}
	dow := _dow

	if closeErr := readBuffer.CloseContext("DateAndTime"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for DateAndTime")
	}

	// Create the instance
	return &_DateAndTime{
		Year:    year,
		Month:   month,
		Day:     day,
		Hour:    hour,
		Minutes: minutes,
		Seconds: seconds,
		Msec:    msec,
		Dow:     dow,
	}, nil
}

func (m *_DateAndTime) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_DateAndTime) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	if pushErr := writeBuffer.PushContext("DateAndTime"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for DateAndTime")
	}

	// Manual Field (year)
	_yearErr := ByteToBcd(writeBuffer, m.GetYear())
	if _yearErr != nil {
		return errors.Wrap(_yearErr, "Error serializing 'year' field")
	}

	// Manual Field (month)
	_monthErr := ByteToBcd(writeBuffer, m.GetMonth())
	if _monthErr != nil {
		return errors.Wrap(_monthErr, "Error serializing 'month' field")
	}

	// Manual Field (day)
	_dayErr := ByteToBcd(writeBuffer, m.GetDay())
	if _dayErr != nil {
		return errors.Wrap(_dayErr, "Error serializing 'day' field")
	}

	// Manual Field (hour)
	_hourErr := ByteToBcd(writeBuffer, m.GetHour())
	if _hourErr != nil {
		return errors.Wrap(_hourErr, "Error serializing 'hour' field")
	}

	// Manual Field (minutes)
	_minutesErr := ByteToBcd(writeBuffer, m.GetMinutes())
	if _minutesErr != nil {
		return errors.Wrap(_minutesErr, "Error serializing 'minutes' field")
	}

	// Manual Field (seconds)
	_secondsErr := ByteToBcd(writeBuffer, m.GetSeconds())
	if _secondsErr != nil {
		return errors.Wrap(_secondsErr, "Error serializing 'seconds' field")
	}

	// Manual Field (msec)
	_msecErr := IntToS7msec(writeBuffer, m.GetMsec())
	if _msecErr != nil {
		return errors.Wrap(_msecErr, "Error serializing 'msec' field")
	}

	// Simple Field (dow)
	dow := uint8(m.GetDow())
	_dowErr := writeBuffer.WriteUint8("dow", 4, (dow))
	if _dowErr != nil {
		return errors.Wrap(_dowErr, "Error serializing 'dow' field")
	}

	if popErr := writeBuffer.PopContext("DateAndTime"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for DateAndTime")
	}
	return nil
}

func (m *_DateAndTime) isDateAndTime() bool {
	return true
}

func (m *_DateAndTime) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
