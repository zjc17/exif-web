package parser

import "time"

type ParseResult struct {
	ApertureValue            int       `json:"ApertureValue"`
	BrightnessValue          float64   `json:"BrightnessValue"`
	Copyright                any       `json:"Copyright"`
	CreateDate               time.Time `json:"CreateDate"`
	CustomRendered           string    `json:"CustomRendered"`
	DateTimeOriginal         time.Time `json:"DateTimeOriginal"`
	ExifImageHeight          int       `json:"ExifImageHeight"`
	ExifImageWidth           int       `json:"ExifImageWidth"`
	ExifVersion              string    `json:"ExifVersion"`
	ExposureCompensation     int       `json:"ExposureCompensation"`
	ExposureMode             string    `json:"ExposureMode"`
	ExposureProgram          string    `json:"ExposureProgram"`
	ExposureTime             float64   `json:"ExposureTime"`
	FNumber                  int       `json:"FNumber"`
	FileSource               string    `json:"FileSource"`
	Flash                    string    `json:"Flash"`
	FocalLength              float64   `json:"FocalLength"`
	FocalLengthIn35MmFormat  int       `json:"FocalLengthIn35mmFormat"`
	FocalPlaneResolutionUnit string    `json:"FocalPlaneResolutionUnit"`
	FocalPlaneXResolution    int       `json:"FocalPlaneXResolution"`
	FocalPlaneYResolution    int       `json:"FocalPlaneYResolution"`
	Iso                      int       `json:"ISO"`
	LensInfo                 []int     `json:"LensInfo"`
	LensMake                 string    `json:"LensMake"`
	LensModel                string    `json:"LensModel"`
	LensSerialNumber         string    `json:"LensSerialNumber"`
	LightSource              string    `json:"LightSource"`
	Make                     string    `json:"Make"`
	MaxApertureValue         int       `json:"MaxApertureValue"`
	MeteringMode             string    `json:"MeteringMode"`
	Model                    string    `json:"Model"`
	ResolutionUnit           string    `json:"ResolutionUnit"`
	SceneCaptureType         string    `json:"SceneCaptureType"`
	SceneType                string    `json:"SceneType"`
	SensingMethod            string    `json:"SensingMethod"`
	SensitivityType          int       `json:"SensitivityType"`
	SerialNumber             string    `json:"SerialNumber"`
	Sharpness                string    `json:"Sharpness"`
	ShutterSpeedValue        float64   `json:"ShutterSpeedValue"`
	Software                 string    `json:"Software"`
	SubjectDistanceRange     string    `json:"SubjectDistanceRange"`
	WhiteBalance             string    `json:"WhiteBalance"`
	XResolution              int       `json:"XResolution"`
	YResolution              int       `json:"YResolution"`
}
