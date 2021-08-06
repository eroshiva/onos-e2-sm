// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package kpmctypes

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "RIC-Format-Type.h"
import "C"
import (
	"fmt"
	e2sm_kpm_ies "github.com/sdran/onos-e2-sm/servicemodels/e2sm_kpm/v1beta1/e2sm-kpm-ies"
	"unsafe"
)

func xerEncodeRicFormatType(ricFormatType *e2sm_kpm_ies.RicFormatType) ([]byte, error) {
	ricFormatTypeCP := newRicFormatType(ricFormatType)

	bytes, err := encodeXer(&C.asn_DEF_RIC_Format_Type, unsafe.Pointer(ricFormatTypeCP))
	if err != nil {
		return nil, fmt.Errorf("xerEncodeRicFormatType() %s", err.Error())
	}
	return bytes, nil
}

func perEncodeRicFormatType(ricFormatType *e2sm_kpm_ies.RicFormatType) ([]byte, error) {
	ricFormatTypeCP := newRicFormatType(ricFormatType)

	bytes, err := encodePerBuffer(&C.asn_DEF_RIC_Format_Type, unsafe.Pointer(ricFormatTypeCP))
	if err != nil {
		return nil, fmt.Errorf("perEncodeRicFormatType() %s", err.Error())
	}
	return bytes, nil
}

func xerDecodeRicFormatType(bytes []byte) (*e2sm_kpm_ies.RicFormatType, error) {
	unsafePtr, err := decodeXer(bytes, &C.asn_DEF_RIC_Format_Type)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from XER is nil")
	}
	return decodeRicFormatType((*C.RIC_Format_Type_t)(unsafePtr)), nil
}

func perDecodeRicFormatType(bytes []byte) (*e2sm_kpm_ies.RicFormatType, error) {
	unsafePtr, err := decodePer(bytes, len(bytes), &C.asn_DEF_RIC_Format_Type)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from PER is nil")
	}
	return decodeRicFormatType((*C.RIC_Format_Type_t)(unsafePtr)), nil
}

func newRicFormatType(ricFormatType *e2sm_kpm_ies.RicFormatType) *C.RIC_Format_Type_t {

	ricStyleTypeC := C.long(ricFormatType.Value)

	return &ricStyleTypeC
}

func decodeRicFormatType(ricFormatTypeC *C.RIC_Format_Type_t) *e2sm_kpm_ies.RicFormatType {

	ricFormatType := e2sm_kpm_ies.RicFormatType{
		Value: int32(*ricFormatTypeC),
	}
	return &ricFormatType
}
