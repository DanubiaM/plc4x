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
package org.apache.plc4x.java.bacnetip.readwrite;

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

public class BACnetStatusFlagsTagged implements Message {

  // Properties.
  protected final BACnetTagHeader header;
  protected final BACnetTagPayloadBitString payload;

  // Arguments.
  protected final Short tagNumber;
  protected final TagClass tagClass;

  public BACnetStatusFlagsTagged(
      BACnetTagHeader header,
      BACnetTagPayloadBitString payload,
      Short tagNumber,
      TagClass tagClass) {
    super();
    this.header = header;
    this.payload = payload;
    this.tagNumber = tagNumber;
    this.tagClass = tagClass;
  }

  public BACnetTagHeader getHeader() {
    return header;
  }

  public BACnetTagPayloadBitString getPayload() {
    return payload;
  }

  public boolean getInAlarm() {
    return (boolean)
        (((((COUNT(getPayload().getData())) > (0))) ? getPayload().getData().get(0) : false));
  }

  public boolean getFault() {
    return (boolean)
        (((((COUNT(getPayload().getData())) > (1))) ? getPayload().getData().get(1) : false));
  }

  public boolean getOverridden() {
    return (boolean)
        (((((COUNT(getPayload().getData())) > (2))) ? getPayload().getData().get(2) : false));
  }

  public boolean getOutOfService() {
    return (boolean)
        (((((COUNT(getPayload().getData())) > (3))) ? getPayload().getData().get(3) : false));
  }

  public void serialize(WriteBuffer writeBuffer) throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    int startPos = positionAware.getPos();
    writeBuffer.pushContext("BACnetStatusFlagsTagged");

    // Simple Field (header)
    writeSimpleField("header", header, new DataWriterComplexDefault<>(writeBuffer));

    // Simple Field (payload)
    writeSimpleField("payload", payload, new DataWriterComplexDefault<>(writeBuffer));

    // Virtual field (doesn't actually serialize anything, just makes the value available)
    boolean inAlarm = getInAlarm();
    writeBuffer.writeVirtual("inAlarm", inAlarm);

    // Virtual field (doesn't actually serialize anything, just makes the value available)
    boolean fault = getFault();
    writeBuffer.writeVirtual("fault", fault);

    // Virtual field (doesn't actually serialize anything, just makes the value available)
    boolean overridden = getOverridden();
    writeBuffer.writeVirtual("overridden", overridden);

    // Virtual field (doesn't actually serialize anything, just makes the value available)
    boolean outOfService = getOutOfService();
    writeBuffer.writeVirtual("outOfService", outOfService);

    writeBuffer.popContext("BACnetStatusFlagsTagged");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = 0;
    BACnetStatusFlagsTagged _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (header)
    lengthInBits += header.getLengthInBits();

    // Simple field (payload)
    lengthInBits += payload.getLengthInBits();

    // A virtual field doesn't have any in- or output.

    // A virtual field doesn't have any in- or output.

    // A virtual field doesn't have any in- or output.

    // A virtual field doesn't have any in- or output.

    return lengthInBits;
  }

  public static BACnetStatusFlagsTagged staticParse(ReadBuffer readBuffer, Object... args)
      throws ParseException {
    PositionAware positionAware = readBuffer;
    if ((args == null) || (args.length != 2)) {
      throw new PlcRuntimeException(
          "Wrong number of arguments, expected 2, but got " + args.length);
    }
    Short tagNumber;
    if (args[0] instanceof Short) {
      tagNumber = (Short) args[0];
    } else if (args[0] instanceof String) {
      tagNumber = Short.valueOf((String) args[0]);
    } else {
      throw new PlcRuntimeException(
          "Argument 0 expected to be of type Short or a string which is parseable but was "
              + args[0].getClass().getName());
    }
    TagClass tagClass;
    if (args[1] instanceof TagClass) {
      tagClass = (TagClass) args[1];
    } else if (args[1] instanceof String) {
      tagClass = TagClass.valueOf((String) args[1]);
    } else {
      throw new PlcRuntimeException(
          "Argument 1 expected to be of type TagClass or a string which is parseable but was "
              + args[1].getClass().getName());
    }
    return staticParse(readBuffer, tagNumber, tagClass);
  }

  public static BACnetStatusFlagsTagged staticParse(
      ReadBuffer readBuffer, Short tagNumber, TagClass tagClass) throws ParseException {
    readBuffer.pullContext("BACnetStatusFlagsTagged");
    PositionAware positionAware = readBuffer;
    int startPos = positionAware.getPos();
    int curPos;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    BACnetTagHeader header =
        readSimpleField(
            "header",
            new DataReaderComplexDefault<>(
                () -> BACnetTagHeader.staticParse(readBuffer), readBuffer));
    // Validation
    if (!((header.getTagClass()) == (tagClass))) {
      throw new ParseValidationException("tag class doesn't match");
    }
    // Validation
    if (!((((header.getTagClass()) == (TagClass.APPLICATION_TAGS)))
        || (((header.getActualTagNumber()) == (tagNumber))))) {
      throw new ParseAssertException("tagnumber doesn't match");
    }

    BACnetTagPayloadBitString payload =
        readSimpleField(
            "payload",
            new DataReaderComplexDefault<>(
                () ->
                    BACnetTagPayloadBitString.staticParse(
                        readBuffer, (long) (header.getActualLength())),
                readBuffer));
    boolean inAlarm =
        readVirtualField(
            "inAlarm",
            boolean.class,
            ((((COUNT(payload.getData())) > (0))) ? payload.getData().get(0) : false));
    boolean fault =
        readVirtualField(
            "fault",
            boolean.class,
            ((((COUNT(payload.getData())) > (1))) ? payload.getData().get(1) : false));
    boolean overridden =
        readVirtualField(
            "overridden",
            boolean.class,
            ((((COUNT(payload.getData())) > (2))) ? payload.getData().get(2) : false));
    boolean outOfService =
        readVirtualField(
            "outOfService",
            boolean.class,
            ((((COUNT(payload.getData())) > (3))) ? payload.getData().get(3) : false));

    readBuffer.closeContext("BACnetStatusFlagsTagged");
    // Create the instance
    BACnetStatusFlagsTagged _bACnetStatusFlagsTagged;
    _bACnetStatusFlagsTagged = new BACnetStatusFlagsTagged(header, payload, tagNumber, tagClass);
    return _bACnetStatusFlagsTagged;
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof BACnetStatusFlagsTagged)) {
      return false;
    }
    BACnetStatusFlagsTagged that = (BACnetStatusFlagsTagged) o;
    return (getHeader() == that.getHeader()) && (getPayload() == that.getPayload()) && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(getHeader(), getPayload());
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
