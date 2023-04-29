package parser

import (
	"encoding/json"
	"time"
)

type PngResult struct {
	BitDepth    int    `json:"BitDepth,omitempty"`
	ColorType   string `json:"ColorType,omitempty"`
	Compression string `json:"Compression,omitempty"`
	Filter      string `json:"Filter,omitempty"`
	ImageHeight int    `json:"ImageHeight,omitempty"`
	ImageWidth  int    `json:"ImageWidth,omitempty"`
	Interlace   string `json:"Interlace,omitempty"`
}

type ParseResult struct {
	ApertureValue            int             `json:"ApertureValue,omitempty"`
	Artist                   string          `json:"Artist,omitempty"`
	BrightnessValue          float64         `json:"BrightnessValue,omitempty"`
	ColorSpace               int             `json:"ColorSpace,omitempty"`
	ComponentsConfiguration  json.RawMessage `json:"ComponentsConfiguration,omitempty"`
	Contrast                 string          `json:"Contrast,omitempty"`
	Copyright                string          `json:"Copyright,omitempty"`
	CreateDate               time.Time       `json:"CreateDate,omitempty"`
	CustomRendered           string          `json:"CustomRendered,omitempty"`
	DateTimeOriginal         time.Time       `json:"DateTimeOriginal,omitempty"`
	DigitalZoomRatio         int             `json:"DigitalZoomRatio,omitempty"`
	ExifImageHeight          int             `json:"ExifImageHeight,omitempty"`
	ExifImageWidth           int             `json:"ExifImageWidth,omitempty"`
	ExifVersion              string          `json:"ExifVersion,omitempty"`
	ExposureCompensation     int             `json:"ExposureCompensation,omitempty"`
	ExposureMode             string          `json:"ExposureMode,omitempty"`
	ExposureProgram          string          `json:"ExposureProgram,omitempty"`
	ExposureTime             float64         `json:"ExposureTime,omitempty"`
	FNumber                  int             `json:"FNumber,omitempty"`
	FileSource               string          `json:"FileSource"`
	Flash                    string          `json:"Flash"`
	FlashpixVersion          string          `json:"FlashpixVersion"`
	FocalLength              int             `json:"FocalLength"`
	FocalLengthIn35MmFormat  int             `json:"FocalLengthIn35mmFormat"`
	FocalPlaneResolutionUnit string          `json:"FocalPlaneResolutionUnit"`
	FocalPlaneXResolution    float64         `json:"FocalPlaneXResolution"`
	FocalPlaneYResolution    float64         `json:"FocalPlaneYResolution"`
	GPSAltitude              float64         `json:"GPSAltitude"`
	GPSAltitudeRef           json.RawMessage `json:"GPSAltitudeRef"`
	GPSImgDirection          int             `json:"GPSImgDirection"`
	GPSImgDirectionRef       string          `json:"GPSImgDirectionRef"`
	GPSLatitude              []float64       `json:"GPSLatitude"`
	GPSLatitudeRef           string          `json:"GPSLatitudeRef"`
	GPSLongitude             []float64       `json:"GPSLongitude"`
	GPSLongitudeRef          string          `json:"GPSLongitudeRef"`
	GPSSpeed                 int             `json:"GPSSpeed"`
	GPSSpeedRef              string          `json:"GPSSpeedRef"`
	Iso                      int             `json:"ISO"`
	LensInfo                 []int           `json:"LensInfo"`
	LensModel                string          `json:"LensModel"`
	LightSource              string          `json:"LightSource"`
	Make                     string          `json:"Make"`
	MaxApertureValue         int             `json:"MaxApertureValue"`
	MeteringMode             string          `json:"MeteringMode"`
	Model                    string          `json:"Model"`
	ModifyDate               time.Time       `json:"ModifyDate"`
	OffsetTime               string          `json:"OffsetTime"`
	OffsetTimeDigitized      string          `json:"OffsetTimeDigitized"`
	OffsetTimeOriginal       string          `json:"OffsetTimeOriginal"`
	Orientation              string          `json:"Orientation"`
	RecommendedExposureIndex int             `json:"RecommendedExposureIndex"`
	ResolutionUnit           string          `json:"ResolutionUnit"`
	Saturation               string          `json:"Saturation"`
	SceneCaptureType         string          `json:"SceneCaptureType"`
	SceneType                string          `json:"SceneType"`
	SensitivityType          int             `json:"SensitivityType"`
	Sharpness                string          `json:"Sharpness"`
	ShutterSpeedValue        float64         `json:"ShutterSpeedValue"`
	Software                 string          `json:"Software"`
	SubSecTime               string          `json:"SubSecTime"`
	SubSecTimeDigitized      string          `json:"SubSecTimeDigitized"`
	SubSecTimeOriginal       string          `json:"SubSecTimeOriginal"`
	TileLength               int             `json:"TileLength"`
	TileWidth                int             `json:"TileWidth"`
	WhiteBalance             string          `json:"WhiteBalance"`
	XResolution              int             `json:"XResolution"`
	YCbCrPositioning         int             `json:"YCbCrPositioning"`
	YResolution              int             `json:"YResolution"`
	Latitude                 float64         `json:"latitude"`
	Longitude                float64         `json:"longitude"`
}

func (p *ParseResult) From(jsonStr string) (err error) {
	if err = json.Unmarshal([]byte(jsonStr), p); err != nil {
		return
	}
	return
}
