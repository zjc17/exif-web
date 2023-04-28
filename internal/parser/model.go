package parser

import (
	"encoding/json"
	"time"
)

type ParseResult struct {
	TileWidth               int             `json:"322,omitempty"`
	TileLength              int             `json:"323,omitempty"`
	ApertureValue           float64         `json:"ApertureValue,omitempty"`
	ColorSpace              int             `json:"ColorSpace,omitempty"`
	ComponentsConfiguration json.RawMessage `json:"ComponentsConfiguration,omitempty"`
	Contrast                string          `json:"Contrast,omitempty"`
	CreateDate              time.Time       `json:"CreateDate,omitempty"`
	DateTimeOriginal        time.Time       `json:"DateTimeOriginal,omitempty"`
	DigitalZoomRatio        int             `json:"DigitalZoomRatio,omitempty"`
	ExifImageHeight         int             `json:"ExifImageHeight,omitempty"`
	ExifImageWidth          int             `json:"ExifImageWidth,omitempty"`
	ExifVersion             string          `json:"ExifVersion,omitempty"`
	ExposureCompensation    int             `json:"ExposureCompensation,omitempty"`
	ExposureMode            string          `json:"ExposureMode,omitempty"`
	ExposureProgram         string          `json:"ExposureProgram,omitempty"`
	ExposureTime            float64         `json:"ExposureTime,omitempty"`
	FNumber                 float64         `json:"FNumber,omitempty"`
	FileSource              string          `json:"FileSource,omitempty"`
	Flash                   string          `json:"Flash,omitempty"`
	FlashpixVersion         string          `json:"FlashpixVersion,omitempty"`
	FocalLength             float64         `json:"FocalLength,omitempty"`
	FocalLengthIn35MmFormat int             `json:"FocalLengthIn35mmFormat,omitempty"`
	GPSAltitude             float64         `json:"GPSAltitude,omitempty"`
	GPSLatitude             []float64       `json:"GPSLatitude,omitempty"`
	GPSLatitudeRef          string          `json:"GPSLatitudeRef,omitempty"`
	GPSLongitude            []float64       `json:"GPSLongitude,omitempty"`
	GPSLongitudeRef         string          `json:"GPSLongitudeRef,omitempty"`
	GPSVersionID            string          `json:"GPSVersionID,omitempty"`
	GainControl             string          `json:"GainControl,omitempty"`
	Iso                     int             `json:"ISO,omitempty"`
	LensInfo                []float64       `json:"LensInfo,omitempty"`
	LensModel               string          `json:"LensModel,omitempty"`
	LightSource             string          `json:"LightSource,omitempty"`
	Make                    string          `json:"Make,omitempty"`
	MaxApertureValue        float64         `json:"MaxApertureValue,omitempty"`
	MeteringMode            string          `json:"MeteringMode,omitempty"`
	Model                   string          `json:"Model,omitempty"`
	ModifyDate              time.Time       `json:"ModifyDate,omitempty"`
	OffsetTime              string          `json:"OffsetTime,omitempty"`
	OffsetTimeDigitized     string          `json:"OffsetTimeDigitized,omitempty"`
	OffsetTimeOriginal      string          `json:"OffsetTimeOriginal,omitempty"`
	Orientation             string          `json:"Orientation,omitempty"`
	ResolutionUnit          string          `json:"ResolutionUnit,omitempty"`
	Saturation              string          `json:"Saturation,omitempty"`
	SceneCaptureType        string          `json:"SceneCaptureType,omitempty"`
	SceneType               string          `json:"SceneType,omitempty"`
	SerialNumber            string          `json:"SerialNumber,omitempty"`
	Sharpness               string          `json:"Sharpness,omitempty"`
	ShutterSpeedValue       float64         `json:"ShutterSpeedValue,omitempty"`
	Software                string          `json:"Software,omitempty"`
	SubSecTime              string          `json:"SubSecTime,omitempty"`
	SubSecTimeDigitized     string          `json:"SubSecTimeDigitized,omitempty"`
	SubSecTimeOriginal      string          `json:"SubSecTimeOriginal,omitempty"`
	WhiteBalance            string          `json:"WhiteBalance,omitempty"`
	XResolution             int             `json:"XResolution,omitempty"`
	YCbCrPositioning        int             `json:"YCbCrPositioning,omitempty"`
	YResolution             int             `json:"YResolution,omitempty"`
	Latitude                float64         `json:"latitude,omitempty"`
	Longitude               float64         `json:"longitude,omitempty"`
}
