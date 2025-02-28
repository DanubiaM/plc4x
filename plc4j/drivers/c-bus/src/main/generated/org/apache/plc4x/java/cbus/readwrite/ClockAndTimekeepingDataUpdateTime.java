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
package org.apache.plc4x.java.cbus.readwrite;

import static org.apache.plc4x.java.spi.codegen.fields.FieldReaderFactory.*;
import static org.apache.plc4x.java.spi.codegen.fields.FieldWriterFactory.*;
import static org.apache.plc4x.java.spi.codegen.io.DataReaderFactory.*;
import static org.apache.plc4x.java.spi.codegen.io.DataWriterFactory.*;
import static org.apache.plc4x.java.spi.generation.StaticHelper.*;

import java.time.*;
import java.util.*;
import org.apache.plc4x.java.api.exceptions.*;
import org.apache.plc4x.java.api.value.*;
import org.apache.plc4x.java.spi.codegen.*;
import org.apache.plc4x.java.spi.codegen.fields.*;
import org.apache.plc4x.java.spi.codegen.io.*;
import org.apache.plc4x.java.spi.generation.*;

// Code generated by code-generation. DO NOT EDIT.

public class ClockAndTimekeepingDataUpdateTime extends ClockAndTimekeepingData implements Message {

  // Accessors for discriminator values.

  // Properties.
  protected final short hours;
  protected final short minute;
  protected final short second;
  protected final byte daylightSaving;

  public ClockAndTimekeepingDataUpdateTime(
      ClockAndTimekeepingCommandTypeContainer commandTypeContainer,
      byte argument,
      short hours,
      short minute,
      short second,
      byte daylightSaving) {
    super(commandTypeContainer, argument);
    this.hours = hours;
    this.minute = minute;
    this.second = second;
    this.daylightSaving = daylightSaving;
  }

  public short getHours() {
    return hours;
  }

  public short getMinute() {
    return minute;
  }

  public short getSecond() {
    return second;
  }

  public byte getDaylightSaving() {
    return daylightSaving;
  }

  public boolean getIsNoDaylightSavings() {
    return (boolean) ((getDaylightSaving()) == (0x00));
  }

  public boolean getIsAdvancedBy1Hour() {
    return (boolean) ((getDaylightSaving()) == (0x01));
  }

  public boolean getIsReserved() {
    return (boolean) (((getDaylightSaving()) > (0x01)) && ((getDaylightSaving()) <= (0xFE)));
  }

  public boolean getIsUnknown() {
    return (boolean) ((getDaylightSaving()) > (0xFE));
  }

  @Override
  protected void serializeClockAndTimekeepingDataChild(WriteBuffer writeBuffer)
      throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    int startPos = positionAware.getPos();
    writeBuffer.pushContext("ClockAndTimekeepingDataUpdateTime");

    // Simple Field (hours)
    writeSimpleField("hours", hours, writeUnsignedShort(writeBuffer, 8));

    // Simple Field (minute)
    writeSimpleField("minute", minute, writeUnsignedShort(writeBuffer, 8));

    // Simple Field (second)
    writeSimpleField("second", second, writeUnsignedShort(writeBuffer, 8));

    // Simple Field (daylightSaving)
    writeSimpleField("daylightSaving", daylightSaving, writeByte(writeBuffer, 8));

    // Virtual field (doesn't actually serialize anything, just makes the value available)
    boolean isNoDaylightSavings = getIsNoDaylightSavings();
    writeBuffer.writeVirtual("isNoDaylightSavings", isNoDaylightSavings);

    // Virtual field (doesn't actually serialize anything, just makes the value available)
    boolean isAdvancedBy1Hour = getIsAdvancedBy1Hour();
    writeBuffer.writeVirtual("isAdvancedBy1Hour", isAdvancedBy1Hour);

    // Virtual field (doesn't actually serialize anything, just makes the value available)
    boolean isReserved = getIsReserved();
    writeBuffer.writeVirtual("isReserved", isReserved);

    // Virtual field (doesn't actually serialize anything, just makes the value available)
    boolean isUnknown = getIsUnknown();
    writeBuffer.writeVirtual("isUnknown", isUnknown);

    writeBuffer.popContext("ClockAndTimekeepingDataUpdateTime");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    ClockAndTimekeepingDataUpdateTime _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (hours)
    lengthInBits += 8;

    // Simple field (minute)
    lengthInBits += 8;

    // Simple field (second)
    lengthInBits += 8;

    // Simple field (daylightSaving)
    lengthInBits += 8;

    // A virtual field doesn't have any in- or output.

    // A virtual field doesn't have any in- or output.

    // A virtual field doesn't have any in- or output.

    // A virtual field doesn't have any in- or output.

    return lengthInBits;
  }

  public static ClockAndTimekeepingDataBuilder staticParseClockAndTimekeepingDataBuilder(
      ReadBuffer readBuffer) throws ParseException {
    readBuffer.pullContext("ClockAndTimekeepingDataUpdateTime");
    PositionAware positionAware = readBuffer;
    int startPos = positionAware.getPos();
    int curPos;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    short hours = readSimpleField("hours", readUnsignedShort(readBuffer, 8));

    short minute = readSimpleField("minute", readUnsignedShort(readBuffer, 8));

    short second = readSimpleField("second", readUnsignedShort(readBuffer, 8));

    byte daylightSaving = readSimpleField("daylightSaving", readByte(readBuffer, 8));
    boolean isNoDaylightSavings =
        readVirtualField("isNoDaylightSavings", boolean.class, (daylightSaving) == (0x00));
    boolean isAdvancedBy1Hour =
        readVirtualField("isAdvancedBy1Hour", boolean.class, (daylightSaving) == (0x01));
    boolean isReserved =
        readVirtualField(
            "isReserved",
            boolean.class,
            ((daylightSaving) > (0x01)) && ((daylightSaving) <= (0xFE)));
    boolean isUnknown = readVirtualField("isUnknown", boolean.class, (daylightSaving) > (0xFE));

    readBuffer.closeContext("ClockAndTimekeepingDataUpdateTime");
    // Create the instance
    return new ClockAndTimekeepingDataUpdateTimeBuilderImpl(hours, minute, second, daylightSaving);
  }

  public static class ClockAndTimekeepingDataUpdateTimeBuilderImpl
      implements ClockAndTimekeepingData.ClockAndTimekeepingDataBuilder {
    private final short hours;
    private final short minute;
    private final short second;
    private final byte daylightSaving;

    public ClockAndTimekeepingDataUpdateTimeBuilderImpl(
        short hours, short minute, short second, byte daylightSaving) {
      this.hours = hours;
      this.minute = minute;
      this.second = second;
      this.daylightSaving = daylightSaving;
    }

    public ClockAndTimekeepingDataUpdateTime build(
        ClockAndTimekeepingCommandTypeContainer commandTypeContainer, byte argument) {
      ClockAndTimekeepingDataUpdateTime clockAndTimekeepingDataUpdateTime =
          new ClockAndTimekeepingDataUpdateTime(
              commandTypeContainer, argument, hours, minute, second, daylightSaving);
      return clockAndTimekeepingDataUpdateTime;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof ClockAndTimekeepingDataUpdateTime)) {
      return false;
    }
    ClockAndTimekeepingDataUpdateTime that = (ClockAndTimekeepingDataUpdateTime) o;
    return (getHours() == that.getHours())
        && (getMinute() == that.getMinute())
        && (getSecond() == that.getSecond())
        && (getDaylightSaving() == that.getDaylightSaving())
        && super.equals(that)
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(
        super.hashCode(), getHours(), getMinute(), getSecond(), getDaylightSaving());
  }

  @Override
  public String toString() {
    WriteBufferBoxBased writeBufferBoxBased = new WriteBufferBoxBased(true, true);
    try {
      writeBufferBoxBased.writeSerializable(this);
    } catch (SerializationException e) {
      throw new RuntimeException(e);
    }
    return "\n" + writeBufferBoxBased.getBox().toString() + "\n";
  }
}
