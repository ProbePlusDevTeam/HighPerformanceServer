package tlvparser

var Ash string = "asdrft"
var JsonString string = `{
    "1": {
        "N": "SensorData",
        "V": {
            "1": {
                "N": "Seq",
                "S": 4
            },
            "2": {
                "N": "ECG0",
                "V": "L"
            },
            "3": {
                "N": "ECG1",
                "V": "L"
            },
            "4": {
                "N": "ECG2",
                "V": "L"
            },
            "5": {
                "N": "ECG3",
                "V": "L"
            },
            "6": {
                "N": "ECG4",
                "V": "L"
            },
            "7": {
                "N": "ECG5",
                "V": "L"
            },
            "8": {
                "N": "ECG6",
                "V": "L"
            },
            "9": {
                "N": "ECG7",
                "V": "L"
            },
            "10": {
                "N": "Respiration",
                "V": "L",
                "S": 4
            },
            "11": {
                "N": "IAGain"
            },
            "12": {
                "N": "LeadStatus"
            },
            "13": {
                "N": "Accel",
                "V": "L"
            },
            "14": {
                "N": "RLDInformation"
            },
            "16": {
                "N": "Temperature",
                "V": "I"
            },
            "17": {
                "N": "vBat"
            },
            "18": {
                "N": "PatchId",
                "V": "A",
                "S": 1
            },
            "19": {
                "N": "SpO2",
                "V": "A",
                "S": 1
            },
            "20": {
                "N": "SpO2Raw",
                "V": "A",
                "S": 1
            },
            "21": {
                "N": "AmbientTemperature",
                "V": "L",
                "S": 4
            },
            "22": {
                "N": "SpO2WhiteData",
                "V": "A",
                "S": 1
            },
            "23": {
                "N": "SpO2SlowData",
                "V": "A",
                "S": 1
            },
            "25": {
                "N": "RespirationIQData",
                "V": "L",
                "S": 4
            },
            "32": {
                "N": "TsECG",
                "S": 6
            },
            "33": {
                "N": "TsResp",
                "S": 6
            },
            "34": {
                "N": "TsTemp",
                "S": 6
            },
            "35": {
                "N": "TsAccel",
                "S": 6
            },
            "36": {
                "N": "TsSpO2",
                "S": 6
            },
            "37": {
                "N": "TsECG0",
                "S": 6
            },
            "38": {
                "N": "TsECG1",
                "S": 6
            },
            "39": {
                "N": "TsECG2",
                "S": 6
            },
            "40": {
                "N": "TsECG3",
                "S": 6
            },
            "41": {
                "N": "TsECG4",
                "S": 6
            },
            "42": {
                "N": "TsECG5",
                "S": 6
            },
            "43": {
                "N": "TsECG6",
                "S": 6
            },
            "44": {
                "N": "TsECG7",
                "S": 6
            },
            "48": {
                "N": "FuelGauge",
                "V": "A"
            },
            "49": {
                "N": "ExtAmbientTemperature",
                "V": "L",
                "S": 4
            },
            "66": {
                "N": "CECG0",
                "V": "L"
            },
            "67": {
                "N": "CECG1",
                "V": "L"
            },
            "68": {
                "N": "CECG2",
                "V": "L"
            },
            "69": {
                "N": "CECG3",
                "V": "L"
            },
            "70": {
                "N": "CECG4",
                "V": "L"
            },
            "71": {
                "N": "CECG5",
                "V": "L"
            },
            "72": {
                "N": "CECG6",
                "V": "L"
            },
            "73": {
                "N": "CECG7",
                "V": "L"
            },
            "247": {
                "N": "Checksum",
                "S": 4
            }
        }
    },
    "129": {
        "N": "PlacementData",
        "V": {
            "1": {
                "N": "Seq",
                "S": 4
            },
            "2": {
                "N": "ECG0",
                "V": "L"
            },
            "3": {
                "N": "ECG1",
                "V": "L"
            },
            "4": {
                "N": "ECG2",
                "V": "L"
            },
            "5": {
                "N": "ECG3",
                "V": "L"
            },
            "6": {
                "N": "ECG4",
                "V": "L"
            },
            "7": {
                "N": "ECG5",
                "V": "L"
            },
            "8": {
                "N": "ECG6",
                "V": "L"
            },
            "9": {
                "N": "ECG7",
                "V": "L"
            },
            "10": {
                "N": "Respiration",
                "V": "L",
                "S": 4
            },
            "11": {
                "N": "IAGain"
            },
            "12": {
                "N": "LeadStatus"
            },
            "13": {
                "N": "Accel",
                "V": "L"
            },
            "14": {
                "N": "RLDInformation"
            },
            "16": {
                "N": "Temperature",
                "V": "I"
            },
            "17": {
                "N": "vBat"
            },
            "18": {
                "N": "PatchId",
                "V": "A",
                "S": 1
            },
            "19": {
                "N": "SpO2",
                "V": "A",
                "S": 1
            },
            "20": {
                "N": "SpO2Raw",
                "V": "A",
                "S": 1
            },
            "21": {
                "N": "AmbientTemperature",
                "V": "L",
                "S": 4
            },
            "22": {
                "N": "SpO2WhiteData",
                "V": "A",
                "S": 1
            },
            "23": {
                "N": "SpO2SlowData",
                "V": "A",
                "S": 1
            },
            "25": {
                "N": "RespirationIQData",
                "V": "L",
                "S": 4
            },
            "32": {
                "N": "TsECG",
                "S": 6
            },
            "33": {
                "N": "TsResp",
                "S": 6
            },
            "34": {
                "N": "TsTemp",
                "S": 6
            },
            "35": {
                "N": "TsAccel",
                "S": 6
            },
            "36": {
                "N": "TsSpO2",
                "S": 6
            },
            "37": {
                "N": "TsECG0",
                "S": 6
            },
            "38": {
                "N": "TsECG1",
                "S": 6
            },
            "39": {
                "N": "TsECG2",
                "S": 6
            },
            "40": {
                "N": "TsECG3",
                "S": 6
            },
            "41": {
                "N": "TsECG4",
                "S": 6
            },
            "42": {
                "N": "TsECG5",
                "S": 6
            },
            "43": {
                "N": "TsECG6",
                "S": 6
            },
            "44": {
                "N": "TsECG7",
                "S": 6
            },
            "48": {
                "N": "FuelGauge",
                "V": "A"
            },
            "49": {
                "N": "ExtAmbientTemperature",
                "V": "L",
                "S": 4
            },
            "66": {
                "N": "CECG0",
                "V": "L"
            },
            "67": {
                "N": "CECG1",
                "V": "L"
            },
            "68": {
                "N": "CECG2",
                "V": "L"
            },
            "69": {
                "N": "CECG3",
                "V": "L"
            },
            "70": {
                "N": "CECG4",
                "V": "L"
            },
            "71": {
                "N": "CECG5",
                "V": "L"
            },
            "72": {
                "N": "CECG6",
                "V": "L"
            },
            "73": {
                "N": "CECG7",
                "V": "L"
            },
            "247": {
                "N": "Checksum",
                "S": 4
            }
        }
    },
    "2": {
        "N": "PatchInfo",
        "V": {
            "1": {
                "N": "PatchId",
                "V": "A",
                "S": 1
            },
            "2": {
                "N": "MacId",
                "V": "A",
                "S": 1
            },
            "3": {
                "N": "SerialNum",
                "V": "A",
                "S": 1
            },
            "4": {
                "N": "PartNumber",
                "V": "A",
                "S": 1
            },
            "5": {
                "N": "ProductPartNumber",
                "V": "A",
                "S": 1
            },
            "32": {
                "N": "VbatMinMax",
                "V": "A"
            },
            "33": {
                "N": "PatchDebug",
                "V": "A",
                "S": 1
            },
            "48": {
                "N": "RadioONStat",
                "S": 4
            },
            "64": {
                "N": "PatchIPAddress",
                "S": 4
            }
        }
    },
    "3": {
        "N": "SensorCalibration",
        "V": {
            "1": {
                "N": "IAConvLo1mv",
                "V": "A"
            },
            "2": {
                "N": "IAConvHi1mv",
                "V": "A"
            },
            "3": {
                "N": "Resp1Ohm",
                "V": "A"
            },
            "4": {
                "N": "TempCalib",
                "V": "L",
                "S": 4
            },
            "5": {
                "N": "AccelCalib",
                "V": "L"
            },
            "6": {
                "N": "SpO2Calib",
                "V": "A",
                "S": 1
            },
            "7": {
                "N": "IAGainLo",
                "V": "A",
                "S": 1
            },
            "8": {
                "N": "IAGainHi",
                "V": "A",
                "S": 1
            },
            "9": {
                "N": "SkinTempCalibExtRange",
                "V": "L",
                "S": 4
            }
        }
    },
    "4": {
        "N": "Capability",
        "V": {
            "1": {
                "N": "ECGSupportedCh"
            },
            "3": {
                "N": "ECGChSps",
                "V": "A",
                "S": 1
            },
            "5": {
                "N": "ClockDiv"
            },
            "6": {
                "N": "ECGWiring",
                "V": "A"
            },
            "7": {
                "N": "RespirationConfig"
            },
            "8": {
                "N": "StorageSize"
            },
            "9": {
                "N": "MaxPatchLife"
            },
            "10": {
                "N": "AccelInfo",
                "V": "A",
                "S": 1
            },
            "11": {
                "N": "TempSupported",
                "V": "A",
                "S": 1
            },
            "14": {
                "N": "SyncMode"
            },
            "15": {
                "N": "NetworkPorts",
                "V": "A"
            },
            "16": {
                "N": "MaxLatency",
                "V": "A"
            },
            "17": {
                "N": "PatchStatus"
            },
            "18": {
                "N": "TotalAvailSequence",
                "S": 4
            },
            "19": {
                "N": "StartTime",
                "S": 4
            },
            "21": {
                "N": "DestIP",
                "S": 4
            },
            "22": {
                "N": "SpO2Config",
                "V": "A"
            },
            "23": {
                "N": "SpO2SPS"
            },
            "24": {
                "N": "FeatureConfig"
            },
            "32": {
                "N": "PSConfig"
            },
            "33": {
                "N": "RequestBufferLen"
            },
            "48": {
                "N": "RangeEndSequenceNumber",
                "S": 4
            },
            "49": {
                "N": "StartAvailSequenceNumber",
                "S": 4
            },
            "50": {
                "N": "StartProcCommitTime",
                "S": 4
            },
            "51": {
                "N": "RespirationSampleSpacing",
                "S": 4
            },
            "52": {
                "N": "BroadcastInterval"
            }
        }
    },
    "5": {
        "N": "ConfigurePatch",
        "V": {
            "1": {
                "N": "ECGSupportedCh"
            },
            "3": {
                "N": "ECGChSps",
                "V": "A",
                "S": 1
            },
            "5": {
                "N": "ClockDiv"
            },
            "7": {
                "N": "RespirationConfig"
            },
            "9": {
                "N": "PatchLife"
            },
            "10": {
                "N": "AccelInfo",
                "V": "A",
                "S": 1
            },
            "11": {
                "N": "TempSupported",
                "V": "A",
                "S": 1
            },
            "14": {
                "N": "SyncMode"
            },
            "15": {
                "N": "NetworkPorts",
                "V": "A"
            },
            "16": {
                "N": "MaxLatency",
                "V": "A"
            },
            "22": {
                "N": "SpO2Config",
                "V": "A"
            },
            "23": {
                "N": "SpO2SPS"
            },
            "24": {
                "N": "FeatureConfig"
            },
            "32": {
                "N": "PSConfig"
            },
            "52": {
                "N": "BroadcastInterval"
            }
        }
    },
    "7": {
        "N": "VersionInfo",
        "V": {
            "1": {
                "N": "FWVersion",
                "S": 4
            },
            "2": {
                "N": "Hash",
                "S": 4
            },
            "3": {
                "N": "ProjVersion",
                "S": 4
            },
            "4": {
                "N": "SpO2Format"
            }
        }
    },
    "8": {
        "N": "Start",
        "V": {
            "19": {
                "N": "EpochTime",
                "S": 4
            },
            "21": {
                "N": "DestIP",
                "S": 4
            }
        }
    },
    "9": {
        "N": "StopAcq"
    },
    "10": {
        "N": "Commit",
        "V": {
            "14": {
                "N": "SyncMode"
            },
            "50": {
                "N": "StartProcCommitTime",
                "S": 4
            }
        }
    },
    "11": {
        "N": "TurnOff",
        "V": {
            "20": {
                "N": "EraseFlash"
            }
        }
    },
    "12": {
        "N": "Identify"
    },
    "13": {
        "N": "ConfigureSSID",
        "V": {
            "3": {
                "N": "SSID",
                "V": "A",
                "S": 1
            },
            "4": {
                "N": "Password",
                "V": "A",
                "S": 1
            }
        }
    },
    "14": {
        "N": "DataRequest",
        "V": {
            "1": {
                "N": "Seqs",
                "V": "A",
                "S": 4
            },
            "15": {
                "N": "RangeReq",
                "V": "A",
                "S": 4
            }
        }
    },
    "15": {
        "N": "Redirect",
        "V": {
            "21": {
                "N": "DestIP",
                "S": 4
            }
        }
    },
    "16": {
        "N": "Ok"
    },
    "17": {
        "N": "Fail"
    },
    "18": {
        "N": "ConfigSpO2",
        "V": {
            "23": {
                "N": "SpO2Config",
                "V": "A"
            }
        }
    },
    "19": {
        "N": "ProgramLED",
        "V": {
            "1": {
                "N": "Repeat"
            },
            "2": {
                "N": "RedGreenValue",
                "V": "A",
                "S": 1
            },
            "3": {
                "N": "Duration",
                "V": "A",
                "S": 1
            }
        }
    },
    "20": {
        "N": "SetEPochTime",
        "V": {
            "19": {
                "N": "EpochTime",
                "S": 4
            }
        }
    },
    "21": {
        "N": "GetSeqEpochTable"
    },
    "22": {
        "N": "SeqEpochTable",
        "V": "A",
        "S": 4
    },
    "23": {
        "N": "BiosensorLog",
        "V": {
            "1": {
                "N": "APPARMDebug",
                "V": "A",
                "S": 1
            },
            "2": {
                "N": "WLSMSGQueueInfo",
                "V": "A",
                "S": 1
            },
            "3": {
                "N": "IPCInfo",
                "V": "A",
                "S": 1
            },
            "4": {
                "N": "LIVEQInfo",
                "V": "A",
                "S": 1
            },
            "5": {
                "N": "BUFQInfo",
                "V": "A",
                "S": 1
            },
            "6": {
                "N": "WLANPRINT",
                "V": "A",
                "S": 1
            },
            "7": {
                "N": "WLANTRC",
                "V": "A",
                "S": 1
            },
            "8": {
                "N": "WLANPSSTA",
                "V": "A",
                "S": 1
            },
            "9": {
                "N": "WLANPSAP",
                "V": "A",
                "S": 1
            },
            "10": {
                "N": "WLANME",
                "V": "A",
                "S": 1
            },
            "11": {
                "N": "WLANPP",
                "V": "A",
                "S": 1
            },
            "12": {
                "N": "WLANMM",
                "V": "A",
                "S": 1
            },
            "13": {
                "N": "WLANMDS",
                "V": "A",
                "S": 1
            },
            "14": {
                "N": "WLANSEC",
                "V": "A",
                "S": 1
            },
            "15": {
                "N": "WLANPHY",
                "V": "A",
                "S": 1
            },
            "16": {
                "N": "WLANBUF",
                "V": "A",
                "S": 1
            },
            "17": {
                "N": "APPPRINT1",
                "V": "A",
                "S": 1
            },
            "18": {
                "N": "APPPRINT2",
                "V": "A",
                "S": 1
            },
            "19": {
                "N": "LOGPatchID",
                "V": "A",
                "S": 1
            }
        }
    },
    "24": {
        "N": "GetSeqEpochTsECGTable"
    },
    "25": {
        "N": "SeqEpochTsECGTable",
        "V": "A",
        "S": 4
    },
    "26": {
        "N": "ConfigurePassword",
        "V": {
            "17": {
                "N": "ProximityPassword",
                "V": "A",
                "S": 1
            },
            "10": {
                "N": "Password2",
                "V": "A",
                "S": 1
            }
        }
    },
    "30": {
        "N": "Reboot"
    },
    "31": {
        "N": "Reconnect"
    },
    "32": {
        "N": "OTAUpdateCommand",
        "V": {
            "1": {
                "N": "OTAConfigCommand"
            },
            "2": {
                "N": "OTAConfigCommandResponse"
            },
            "3": {
                "N": "OTACommandData",
                "V": "A",
                "S": 1
            },
            "4": {
                "N": "OTACommandDataResponse"
            },
            "5": {
                "N": "OTAFinalChecksumCommand",
                "V": "A",
                "S": 1
            },
            "6": {
                "N": "OTAFinalChecksumResponse"
            }
        }
    },
    "48": {
        "N": "StartPlacement"
    }
}`
