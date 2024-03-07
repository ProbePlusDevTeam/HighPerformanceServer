package tlvparser

import (
	"encoding/binary"
	//"encoding/hex"
	"encoding/json"
	"fmt"
	"strconv"
)

type TLVData map[string]interface{}
type SensorData struct {
	Accel                 []int `json:"Accel"`
	Checksum              int   `json:"Checksum"`
	ECG0                  []int `json:"ECG0"`
	ECG1                  []int `json:"ECG1"`
	ExtAmbientTemperature []int `json:"ExtAmbientTemperature"`
	IAGain                int   `json:"IAGain"`
	LeadStatus            int   `json:"LeadStatus"`
	PatchId               []int `json:"PatchId"`
	RLDInformation        int   `json:"RLDInformation"`
	Respiration           []int `json:"Respiration"`
	Seq                   int   `json:"Seq"`
	Temperature           int   `json:"Temperature"`
	TsECG                 int   `json:"TsECG"`
	VBat                  int   `json:"vBat"`
}

type Data struct {
	Commit           map[string]interface{} `json:"Commit"`
	ConfigureSSID    map[string]interface{} `json:"ConfigureSSID"`
	DataRequest      map[string]interface{} `json:"DataRequest"`
	Fail             int                    `json:"Fail"`
	Identify         int                    `json:"Identify"`
	OTAUpdateCommand map[string]interface{} `json:"OTAUpdateCommand"`
	SensorData       []SensorData           `json:"SensorData"`
	TurnOff          map[string]interface{} `json:"TurnOff"`
}

func readTLV(tlv []byte) (int, int, []byte, []byte) {
	T := int(tlv[0]) // T of TLV
	L := int(tlv[1]) // L of TLV
	dp := 2          // Offset to first V byte]
	if L == 0 {      // Extended length
		L = int(tlv[3])<<0x8 + int(tlv[2])
		dp += 2
	}
	V := tlv[dp : dp+int(L)]
	dp += int(L)
	if dp%2 != 0 {
		dp++
	}
	if dp > len(tlv) {
		return T, int(L), V, []byte{}

	} else {
		return T, int(L), V, tlv[dp:]

	}

}

func intFromBytes(bytes []byte, nBytes int, signed bool) int {
	var v int
	switch nBytes {
	case 1:
		v = int(bytes[0])
	case 2:
		v = int(binary.LittleEndian.Uint16(bytes))
	case 4:
		v = int(binary.LittleEndian.Uint32(bytes))
	case 6:
		// Prepend with zero bytes and then convert to uint32
		var intValue uint64
		for i := 0; i < len(bytes); i++ {
			intValue |= uint64(bytes[i]) << uint(8*i)
		}
		v = int(intValue)
	case 8:
		v = int(binary.LittleEndian.Uint64(bytes))
	}

	if signed {
		// Calculate the mask for the most significant bit
		mask := 1 << ((nBytes * 8) - 1)
		// Calculate the maximum value for the number of bytes
		imax := 1 << (nBytes * 8)
		if (v & mask) != 0 { // MSBit set?
			// Perform two's complement
			v = 0 - imax + v
		}
	}
	return v
}

func readJson() TLVData {

	var tlvdef TLVData
	if err := json.Unmarshal([]byte(JsonString), &tlvdef); err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return TLVData{}
	}
	return tlvdef

}

func parseTlv(tlvData []byte, tlvDefn map[string]interface{}) (map[string]interface{}, string) {
	result := make(map[string]interface{})
	var nestedResult map[string]interface{}
	var signed bool

	for len(tlvData) > 4 {
		T, L, Vdat, newTlvData := readTLV(tlvData)
		tlvData = newTlvData

		if defn, ok := tlvDefn[strconv.Itoa(T)].(map[string]interface{}); ok {
			VDefn, err := getData(defn, "V")
			if err != "" {
				return nil, err
			}
			if VDefn == nil {
				VDefn = "U"
			}
			key := defn["N"].(string)
			if VDefn == "A" || VDefn == "L" {
				elemSize := 2
				if size, ok := defn["S"].(float64); ok {
					elemSize = int(size)
				}
				array := make([]int, 0)
				if VDefn == "L" {
					signed = true
				} else {
					signed = false
				}
				for indx := 0; indx < L/elemSize; indx++ {
					array = append(array, intFromBytes(Vdat, elemSize, signed))
					Vdat = Vdat[elemSize:]
				}
				result[key] = array
			} else if VDefn == "U" || VDefn == "I" {
				result[key] = intFromBytes(Vdat, L, VDefn == "I")
			} else {
				// Handle nested TLV case
				if VDefnm, ok := VDefn.(map[string]interface{}); ok {
					// Use the map
					nestedResult, err = parseTlv(Vdat, VDefnm)
					if err != "" {
						return nil, err
					}
				}

				if existing, exists := result[key]; exists {
					if existingMap, ok := existing.(map[string]interface{}); ok {
						result[key] = []interface{}{existingMap, nestedResult}
					} else if existingSlice, ok := existing.([]interface{}); ok {
						result[key] = append(existingSlice, nestedResult)
					}
				} else {
					result[key] = nestedResult
				}
			}
		}

	}

	return result, ""
}
func getData(data interface{}, key string) (interface{}, string) {
	var err string
	switch v := data.(type) {
	case string, int:
		return v, ""
	case map[string]interface{}:
		return v[key], ""
	default:
		err = "Not find any type"
	}
	return "", err
}
func createJson(data map[string]interface{}) interface{} {

	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Println("Error:", err)
		return Data{}
	}

	if _, ok := data["Capability"]; ok {
		var configdata ConfigData
		eerr := json.Unmarshal(jsonData, &configdata)
		// fmt.Printf("Capability: %+v\n", configdata.Capability)
		// fmt.Printf("ConfigurePatch: %+v\n", configdata.ConfigurePatch)
		// fmt.Printf("PatchInfo: %+v\n", configdata.PatchInfo)
		// fmt.Printf("SensorCalibration: %+v\n", configdata.SensorCalibration)
		// fmt.Printf("VersionInfo: %+v\n", configdata.VersionInfo)

		if err != nil {
			fmt.Println("Error:", eerr)
			return Data{}
		}
		return configdata
	} else if _, ok := data["SensorData"]; ok {
		var sensordata ConfigData
		eerr := json.Unmarshal(jsonData, &sensordata)
		if err != nil {
			fmt.Println("Error:", eerr)
			return Data{}
		}
		return sensordata
	}
	return nil

}

func ParseData(relayData []uint8) {
	tlvdef := readJson()
	result, errs := parseTlv([]byte(relayData[4:]), tlvdef)
	fmt.Println(result, errs)
	//sensordata := createJson(result)
	//fmt.Println(sensordata)

}
