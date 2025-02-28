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
package org.apache.plc4x.java.profinet.readwrite;

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

public class PnIoCm_IoCrBlockReqApi implements Message {

  // Constant values.
  public static final Long API = 0x00000000L;

  // Properties.
  protected final List<PnIoCm_IoDataObject> ioDataObjects;
  protected final List<PnIoCm_IoCs> ioCss;

  public PnIoCm_IoCrBlockReqApi(List<PnIoCm_IoDataObject> ioDataObjects, List<PnIoCm_IoCs> ioCss) {
    super();
    this.ioDataObjects = ioDataObjects;
    this.ioCss = ioCss;
  }

  public List<PnIoCm_IoDataObject> getIoDataObjects() {
    return ioDataObjects;
  }

  public List<PnIoCm_IoCs> getIoCss() {
    return ioCss;
  }

  public long getApi() {
    return API;
  }

  public void serialize(WriteBuffer writeBuffer) throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    int startPos = positionAware.getPos();
    writeBuffer.pushContext("PnIoCm_IoCrBlockReqApi");

    // Const Field (api)
    writeConstField(
        "api",
        API,
        writeUnsignedLong(writeBuffer, 32),
        WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    // Implicit Field (numIoDataObjects) (Used for parsing, but its value is not stored as it's
    // implicitly given by the objects content)
    int numIoDataObjects = (int) (COUNT(getIoDataObjects()));
    writeImplicitField(
        "numIoDataObjects",
        numIoDataObjects,
        writeUnsignedInt(writeBuffer, 16),
        WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    // Array Field (ioDataObjects)
    writeComplexTypeArrayField(
        "ioDataObjects",
        ioDataObjects,
        writeBuffer,
        WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    // Implicit Field (numIoCss) (Used for parsing, but its value is not stored as it's implicitly
    // given by the objects content)
    int numIoCss = (int) (COUNT(getIoCss()));
    writeImplicitField(
        "numIoCss",
        numIoCss,
        writeUnsignedInt(writeBuffer, 16),
        WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    // Array Field (ioCss)
    writeComplexTypeArrayField(
        "ioCss", ioCss, writeBuffer, WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    writeBuffer.popContext("PnIoCm_IoCrBlockReqApi");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = 0;
    PnIoCm_IoCrBlockReqApi _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Const Field (api)
    lengthInBits += 32;

    // Implicit Field (numIoDataObjects)
    lengthInBits += 16;

    // Array field
    if (ioDataObjects != null) {
      int i = 0;
      for (PnIoCm_IoDataObject element : ioDataObjects) {
        ThreadLocalHelper.lastItemThreadLocal.set(++i >= ioDataObjects.size());
        lengthInBits += element.getLengthInBits();
      }
    }

    // Implicit Field (numIoCss)
    lengthInBits += 16;

    // Array field
    if (ioCss != null) {
      int i = 0;
      for (PnIoCm_IoCs element : ioCss) {
        ThreadLocalHelper.lastItemThreadLocal.set(++i >= ioCss.size());
        lengthInBits += element.getLengthInBits();
      }
    }

    return lengthInBits;
  }

  public static PnIoCm_IoCrBlockReqApi staticParse(ReadBuffer readBuffer, Object... args)
      throws ParseException {
    PositionAware positionAware = readBuffer;
    return staticParse(readBuffer);
  }

  public static PnIoCm_IoCrBlockReqApi staticParse(ReadBuffer readBuffer) throws ParseException {
    readBuffer.pullContext("PnIoCm_IoCrBlockReqApi");
    PositionAware positionAware = readBuffer;
    int startPos = positionAware.getPos();
    int curPos;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    long api =
        readConstField(
            "api",
            readUnsignedLong(readBuffer, 32),
            PnIoCm_IoCrBlockReqApi.API,
            WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    int numIoDataObjects =
        readImplicitField(
            "numIoDataObjects",
            readUnsignedInt(readBuffer, 16),
            WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    List<PnIoCm_IoDataObject> ioDataObjects =
        readCountArrayField(
            "ioDataObjects",
            new DataReaderComplexDefault<>(
                () -> PnIoCm_IoDataObject.staticParse(readBuffer), readBuffer),
            numIoDataObjects,
            WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    int numIoCss =
        readImplicitField(
            "numIoCss",
            readUnsignedInt(readBuffer, 16),
            WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    List<PnIoCm_IoCs> ioCss =
        readCountArrayField(
            "ioCss",
            new DataReaderComplexDefault<>(() -> PnIoCm_IoCs.staticParse(readBuffer), readBuffer),
            numIoCss,
            WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    readBuffer.closeContext("PnIoCm_IoCrBlockReqApi");
    // Create the instance
    PnIoCm_IoCrBlockReqApi _pnIoCm_IoCrBlockReqApi;
    _pnIoCm_IoCrBlockReqApi = new PnIoCm_IoCrBlockReqApi(ioDataObjects, ioCss);
    return _pnIoCm_IoCrBlockReqApi;
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof PnIoCm_IoCrBlockReqApi)) {
      return false;
    }
    PnIoCm_IoCrBlockReqApi that = (PnIoCm_IoCrBlockReqApi) o;
    return (getIoDataObjects() == that.getIoDataObjects())
        && (getIoCss() == that.getIoCss())
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(getIoDataObjects(), getIoCss());
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
