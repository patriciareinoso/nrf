// Copyright 2019 free5GC.org
//
// SPDX-License-Identifier: Apache-2.0

package management

import (
	"reflect"
	"strconv"
	"time"

	"github.com/mitchellh/mapstructure"
	"github.com/omec-project/nrf/dbadapter"
	"github.com/omec-project/nrf/logger"
	"github.com/omec-project/nrf/util"
	"github.com/omec-project/openapi/models"
	"go.mongodb.org/mongo-driver/bson"
)

func GetNrfInfo() *models.NrfInfo {
	// init
	var nrfinfo models.NrfInfo

	nrfinfo.ServedUdrInfo = getUdrInfo()
	nrfinfo.ServedUdmInfo = getUdmInfo()
	nrfinfo.ServedAusfInfo = getAusfInfo()
	nrfinfo.ServedAmfInfo = getAmfInfo()
	nrfinfo.ServedSmfInfo = getSmfInfo()
	nrfinfo.ServedUpfInfo = getUpfInfo()
	nrfinfo.ServedPcfInfo = getPcfInfo()
	nrfinfo.ServedBsfInfo = getBsfInfo()
	nrfinfo.ServedChfInfo = getChfInfo()

	return &nrfinfo
}

func getUdrInfo() map[string]models.UdrInfo {
	servedUdrInfo := make(map[string]models.UdrInfo)
	var UDRProfile models.NfProfile

	collName := "NfProfile"
	filter := bson.M{"nfType": "UDR"}

	UDR, _ := dbadapter.DBClient.RestfulAPIGetMany(collName, filter)
	UDRStruct, err := util.Decode(UDR, time.RFC3339)
	if err != nil {
		logger.ManagementLog.Error(err)
	}

	for i := 0; i < len(UDRStruct); i++ {
		err := mapstructure.Decode(UDRStruct[i], &UDRProfile)
		if err != nil {
			panic(err)
		}
		index := strconv.Itoa(i)
		servedUdrInfo[index] = *UDRProfile.UdrInfo
	}
	return servedUdrInfo
}

func getUdmInfo() map[string]models.UdmInfo {
	servedUdmInfo := make(map[string]models.UdmInfo)
	var UDMProfile models.NfProfile

	collName := "NfProfile"
	filter := bson.M{"nfType": "UDM"}

	UDM, _ := dbadapter.DBClient.RestfulAPIGetMany(collName, filter)
	UDMStruct, err := util.Decode(UDM, time.RFC3339)
	if err != nil {
		logger.ManagementLog.Error(err)
	}

	for i := 0; i < len(UDMStruct); i++ {
		err := mapstructure.Decode(UDMStruct[i], &UDMProfile)
		if err != nil {
			panic(err)
		}
		index := strconv.Itoa(i)
		servedUdmInfo[index] = *UDMProfile.UdmInfo
	}
	return servedUdmInfo
}

func getAusfInfo() map[string]models.AusfInfo {
	servedAusfInfo := make(map[string]models.AusfInfo)
	var AUSFProfile models.NfProfile

	collName := "NfProfile"
	filter := bson.M{"nfType": "AUSF"}

	AUSF, _ := dbadapter.DBClient.RestfulAPIGetMany(collName, filter)
	AUSFStruct, err := util.Decode(AUSF, time.RFC3339)
	if err != nil {
		logger.ManagementLog.Error(err)
	}
	for i := 0; i < len(AUSFStruct); i++ {
		err := mapstructure.Decode(AUSFStruct[i], &AUSFProfile)
		if err != nil {
			panic(err)
		}
		index := strconv.Itoa(i)
		servedAusfInfo[index] = *AUSFProfile.AusfInfo
	}
	return servedAusfInfo
}

func getAmfInfo() map[string]models.AmfInfo {
	servedAmfinfo := make(map[string]models.AmfInfo)
	var AMFProfile models.NfProfile

	collName := "NfProfile"
	filter := bson.M{"nfType": "AMF"}

	AMF, _ := dbadapter.DBClient.RestfulAPIGetMany(collName, filter)
	AMFStruct, err := util.Decode(AMF, time.RFC3339)
	if err != nil {
		logger.ManagementLog.Error(err)
	}
	for i := 0; i < len(AMFStruct); i++ {
		err := mapstructure.Decode(AMFStruct[i], &AMFProfile)
		if err != nil {
			panic(err)
		}
		index := strconv.Itoa(i)
		servedAmfinfo[index] = *AMFProfile.AmfInfo
	}
	return servedAmfinfo
}

func getSmfInfo() map[string]models.SmfInfo {
	servedSmfInfo := make(map[string]models.SmfInfo)
	var SMFProfile models.NfProfile

	collName := "NfProfile"
	filter := bson.M{"nfType": "SMF"}

	SMF, _ := dbadapter.DBClient.RestfulAPIGetMany(collName, filter)
	SMFStruct, err := util.Decode(SMF, time.RFC3339)
	if err != nil {
		logger.ManagementLog.Error(err)
	}
	for i := 0; i < len(SMFStruct); i++ {
		err := mapstructure.Decode(SMFStruct[i], &SMFProfile)
		if err != nil {
			panic(err)
		}
		index := strconv.Itoa(i)
		servedSmfInfo[index] = *SMFProfile.SmfInfo
	}
	return servedSmfInfo
}

func getUpfInfo() map[string]models.UpfInfo {
	servedUpfInfo := make(map[string]models.UpfInfo)
	var UPFProfile models.NfProfile

	collName := "NfProfile"
	filter := bson.M{"nfType": "UPF"}

	UPF, _ := dbadapter.DBClient.RestfulAPIGetMany(collName, filter)
	UPFStruct, err := util.Decode(UPF, time.RFC3339)
	if err != nil {
		logger.ManagementLog.Error(err)
	}
	for i := 0; i < len(UPFStruct); i++ {
		err := mapstructure.Decode(UPFStruct[i], &UPFProfile)
		if err != nil {
			panic(err)
		}
		index := strconv.Itoa(i)
		servedUpfInfo[index] = *UPFProfile.UpfInfo
	}
	return servedUpfInfo
}

func getPcfInfo() map[string]models.PcfInfo {
	servedPcfInfo := make(map[string]models.PcfInfo)
	var PCFProfile models.NfProfile

	collName := "NfProfile"
	filter := bson.M{"nfType": "PCF"}

	PCF, _ := dbadapter.DBClient.RestfulAPIGetMany(collName, filter)
	PCFStruct, err := util.Decode(PCF, time.RFC3339)
	if err != nil {
		logger.ManagementLog.Error(err)
	}
	for i := 0; i < len(PCFStruct); i++ {
		err := mapstructure.Decode(PCFStruct[i], &PCFProfile)
		if err != nil {
			panic(err)
		}
		index := strconv.Itoa(i)
		servedPcfInfo[index] = *PCFProfile.PcfInfo
	}
	return servedPcfInfo
}

func getBsfInfo() map[string]models.BsfInfo {
	servedBsfInfo := make(map[string]models.BsfInfo)
	var BSFProfile models.NfProfile

	collName := "NfProfile"
	filter := bson.M{"nfType": "BSF"}

	BSF, _ := dbadapter.DBClient.RestfulAPIGetMany(collName, filter)
	BSFStruct, err := util.Decode(BSF, time.RFC3339)
	if err != nil {
		logger.ManagementLog.Error(err)
	}
	for i := 0; i < len(BSFStruct); i++ {
		err := mapstructure.Decode(BSFStruct[i], &BSFProfile)
		if err != nil {
			panic(err)
		}
		index := strconv.Itoa(i)
		servedBsfInfo[index] = *BSFProfile.BsfInfo
	}
	return servedBsfInfo
}

func getChfInfo() map[string]models.ChfInfo {
	servedChfInfo := make(map[string]models.ChfInfo)
	var CHFProfile models.NfProfile

	collName := "NfProfile"
	filter := bson.M{"nfType": "CHF"}

	CHF, _ := dbadapter.DBClient.RestfulAPIGetMany(collName, filter)
	CHFStruct, err := util.Decode(CHF, time.RFC3339)
	if err != nil {
		logger.ManagementLog.Error(err)
	}
	for i := 0; i < len(CHFStruct); i++ {
		err := mapstructure.Decode(CHFStruct[i], &CHFProfile)
		if err != nil {
			panic(err)
		}
		index := strconv.Itoa(i)
		servedChfInfo[index] = *CHFProfile.ChfInfo
	}
	return servedChfInfo
}

// DecodeNfProfile - Only support []map[string]interface to []models.NfProfile
func DecodeNfProfile(source interface{}, format string) (models.NfProfile, error) {
	var target models.NfProfile

	// config mapstruct
	stringToDateTimeHook := func(
		f reflect.Type,
		t reflect.Type,
		data interface{},
	) (interface{}, error) {
		if t == reflect.TypeOf(time.Time{}) && f == reflect.TypeOf("") {
			return time.Parse(format, data.(string))
		}
		return data, nil
	}

	config := mapstructure.DecoderConfig{
		DecodeHook: stringToDateTimeHook,
		Result:     &target,
	}

	decoder, err := mapstructure.NewDecoder(&config)
	if err != nil {
		return target, err
	}

	// Decode result to NfProfile structure
	err = decoder.Decode(source)
	if err != nil {
		return target, err
	}
	return target, nil
}
