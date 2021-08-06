// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package mhoctypes

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "RIC-ReportStyle-List.h"
import "C"
import (
	"fmt"
	e2sm_mho "github.com/sdran/onos-e2-sm/servicemodels/e2sm_mho/v1/e2sm-mho"
	"unsafe"
)

func xerEncodeRicReportStyleItem(ricReportStyleItem *e2sm_mho.RicReportStyleList) ([]byte, error) {
	ricReportStyleItemCP := newRicReportStyleListItem(ricReportStyleItem)

	bytes, err := encodeXer(&C.asn_DEF_RIC_ReportStyle_List, unsafe.Pointer(ricReportStyleItemCP))
	if err != nil {
		return nil, fmt.Errorf("xerEncodeRicReportStyleItem() %s", err.Error())
	}
	return bytes, nil
}

func perEncodeRicReportStyleItem(ricReportStyleItem *e2sm_mho.RicReportStyleList) ([]byte, error) {
	ricReportStyleItemCP := newRicReportStyleListItem(ricReportStyleItem)

	bytes, err := encodePerBuffer(&C.asn_DEF_RIC_ReportStyle_List, unsafe.Pointer(ricReportStyleItemCP))
	if err != nil {
		return nil, fmt.Errorf("perEncodeRicReportStyleItem() %s", err.Error())
	}
	return bytes, nil
}

func xerDecodeRicReportStyleItem(bytes []byte) (*e2sm_mho.RicReportStyleList, error) {
	unsafePtr, err := decodeXer(bytes, &C.asn_DEF_RIC_ReportStyle_List)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from XER is nil")
	}
	return decodeRicReportStyleListItem((*C.RIC_ReportStyle_List_t)(unsafePtr)), nil
}

func perDecodeRicReportStyleItem(bytes []byte) (*e2sm_mho.RicReportStyleList, error) {
	unsafePtr, err := decodePer(bytes, len(bytes), &C.asn_DEF_RIC_ReportStyle_List)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from PER is nil")
	}
	return decodeRicReportStyleListItem((*C.RIC_ReportStyle_List_t)(unsafePtr)), nil
}

func newRicReportStyleListItem(ricReportStyleListItem *e2sm_mho.RicReportStyleList) *C.RIC_ReportStyle_List_t {

	ricReportStyleTypeC, _ := newRicStyleType(ricReportStyleListItem.RicReportStyleType)
	ricReportStyleNameC, _ := newRicStyleName(ricReportStyleListItem.RicReportStyleName)
	ricReportStyleIndicationHeaderFormatTypeC, _ := newRicFormatType(ricReportStyleListItem.RicIndicationHeaderFormatType)
	ricReportStyleIndicationMessageFormatTypeC, _ := newRicFormatType(ricReportStyleListItem.RicIndicationMessageFormatType)

	ricReportStyleListItemC := C.RIC_ReportStyle_List_t{
		ric_ReportStyle_Type:             *ricReportStyleTypeC,
		ric_ReportStyle_Name:             *ricReportStyleNameC,
		ric_IndicationHeaderFormat_Type:  *ricReportStyleIndicationHeaderFormatTypeC,
		ric_IndicationMessageFormat_Type: *ricReportStyleIndicationMessageFormatTypeC,
	}

	return &ricReportStyleListItemC
}

func decodeRicReportStyleListItem(ricReportStyleListItemC *C.RIC_ReportStyle_List_t) *e2sm_mho.RicReportStyleList {

	ricReportStyleType, _ := decodeRicStyleType(&ricReportStyleListItemC.ric_ReportStyle_Type)
	ricReportStyleName, _ := decodeRicStyleName(&ricReportStyleListItemC.ric_ReportStyle_Name)
	ricReportStyleIndicationHeaderFormatType, _ := decodeRicFormatType(&ricReportStyleListItemC.ric_IndicationHeaderFormat_Type)
	ricReportStyleIndicationMessageFormatType, _ := decodeRicFormatType(&ricReportStyleListItemC.ric_IndicationMessageFormat_Type)

	ricReportStyleListItem := e2sm_mho.RicReportStyleList{
		RicReportStyleType:             ricReportStyleType,
		RicReportStyleName:             ricReportStyleName,
		RicIndicationHeaderFormatType:  ricReportStyleIndicationHeaderFormatType,
		RicIndicationMessageFormatType: ricReportStyleIndicationMessageFormatType,
	}
	return &ricReportStyleListItem
}
