// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package rcprectypes

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "RIC-ReportStyle-List-RCPRE.h"
import "C"
import (
	"fmt"
	e2sm_rc_pre_v2 "github.com/sdran/onos-e2-sm/servicemodels/e2sm_rc_pre/v2/e2sm-rc-pre-v2"
	"unsafe"
)

func xerEncodeRicReportStyleItem(ricReportStyleItem *e2sm_rc_pre_v2.RicReportStyleList) ([]byte, error) {
	ricReportStyleItemCP := newRicReportStyleListItem(ricReportStyleItem)

	bytes, err := encodeXer(&C.asn_DEF_RIC_ReportStyle_List_RCPRE, unsafe.Pointer(ricReportStyleItemCP))
	if err != nil {
		return nil, fmt.Errorf("xerEncodeRicReportStyleItem() %s", err.Error())
	}
	return bytes, nil
}

func perEncodeRicReportStyleItem(ricReportStyleItem *e2sm_rc_pre_v2.RicReportStyleList) ([]byte, error) {
	ricReportStyleItemCP := newRicReportStyleListItem(ricReportStyleItem)

	bytes, err := encodePerBuffer(&C.asn_DEF_RIC_ReportStyle_List_RCPRE, unsafe.Pointer(ricReportStyleItemCP))
	if err != nil {
		return nil, fmt.Errorf("perEncodeRicReportStyleItem() %s", err.Error())
	}
	return bytes, nil
}

func xerDecodeRicReportStyleItem(bytes []byte) (*e2sm_rc_pre_v2.RicReportStyleList, error) {
	unsafePtr, err := decodeXer(bytes, &C.asn_DEF_RIC_ReportStyle_List_RCPRE)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from XER is nil")
	}
	return decodeRicReportStyleListItem((*C.RIC_ReportStyle_List_RCPRE_t)(unsafePtr)), nil
}

func perDecodeRicReportStyleItem(bytes []byte) (*e2sm_rc_pre_v2.RicReportStyleList, error) {
	unsafePtr, err := decodePer(bytes, len(bytes), &C.asn_DEF_RIC_ReportStyle_List_RCPRE)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from PER is nil")
	}
	return decodeRicReportStyleListItem((*C.RIC_ReportStyle_List_RCPRE_t)(unsafePtr)), nil
}

func newRicReportStyleListItem(ricReportStyleListItem *e2sm_rc_pre_v2.RicReportStyleList) *C.RIC_ReportStyle_List_RCPRE_t {

	ricReportStyleTypeC := newRicStyleType(ricReportStyleListItem.RicReportStyleType)
	ricReportStyleNameC := newRicStyleName(ricReportStyleListItem.RicReportStyleName)
	ricReportStyleIndicationHeaderFormatTypeC := newRicFormatType(ricReportStyleListItem.RicIndicationHeaderFormatType)
	ricReportStyleIndicationMessageFormatTypeC := newRicFormatType(ricReportStyleListItem.RicIndicationMessageFormatType)

	ricReportStyleListItemC := C.RIC_ReportStyle_List_RCPRE_t{
		ric_ReportStyle_Type:             *ricReportStyleTypeC,
		ric_ReportStyle_Name:             *ricReportStyleNameC,
		ric_IndicationHeaderFormat_Type:  *ricReportStyleIndicationHeaderFormatTypeC,
		ric_IndicationMessageFormat_Type: *ricReportStyleIndicationMessageFormatTypeC,
	}

	return &ricReportStyleListItemC
}

func decodeRicReportStyleListItem(ricReportStyleListItemC *C.RIC_ReportStyle_List_RCPRE_t) *e2sm_rc_pre_v2.RicReportStyleList {

	ricReportStyleType := decodeRicStyleType(&ricReportStyleListItemC.ric_ReportStyle_Type)
	ricReportStyleName := decodeRicStyleName(&ricReportStyleListItemC.ric_ReportStyle_Name)
	ricReportStyleIndicationHeaderFormatType := decodeRicFormatType(&ricReportStyleListItemC.ric_IndicationHeaderFormat_Type)
	ricReportStyleIndicationMessageFormatType := decodeRicFormatType(&ricReportStyleListItemC.ric_IndicationMessageFormat_Type)

	ricReportStyleListItem := e2sm_rc_pre_v2.RicReportStyleList{
		RicReportStyleType:             ricReportStyleType,
		RicReportStyleName:             ricReportStyleName,
		RicIndicationHeaderFormatType:  ricReportStyleIndicationHeaderFormatType,
		RicIndicationMessageFormatType: ricReportStyleIndicationMessageFormatType,
	}
	return &ricReportStyleListItem
}
