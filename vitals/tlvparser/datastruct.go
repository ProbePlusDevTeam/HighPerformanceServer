package tlvparser

type ConfigData struct {
    Capability         CapabilityData      `json:"Capability"`
    ConfigurePatch     ConfigurePatchData  `json:"ConfigurePatch"`
    PatchInfo          PatchInfoData       `json:"PatchInfo"`
    SensorCalibration  SensorCalibrationData `json:"SensorCalibration"`
    VersionInfo        VersionInfoData     `json:"VersionInfo"`
}

type CapabilityData struct {
    AccelInfo             []int   `json:"AccelInfo"`
    ClockDiv              int     `json:"ClockDiv"`
    DestIP                int     `json:"DestIP"`
    ECGChSps              []int   `json:"ECGChSps"`
    // Add more fields as needed
}

type ConfigurePatchData struct {
    AccelInfo             []int   `json:"AccelInfo"`
    ClockDiv              int     `json:"ClockDiv"`
    ECGChSps              []int   `json:"ECGChSps"`
    // Add more fields as needed
}

type PatchInfoData struct {
    MacId                []int   `json:"MacId"`
    PartNumber           []int   `json:"PartNumber"`
    PatchDebug           []int   `json:"PatchDebug"`
    // Add more fields as needed
}

type SensorCalibrationData struct {
    AccelCalib           []int   `json:"AccelCalib"`
    IAConvHi1mv          []int   `json:"IAConvHi1mv"`
    IAConvLo1mv          []int   `json:"IAConvLo1mv"`
    // Add more fields as needed
}

type VersionInfoData struct {
    FWVersion            int     `json:"FWVersion"`
    Hash                 int     `json:"Hash"`
    // Add more fields as needed
}

