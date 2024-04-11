package fpga

import (
	"math"
	"strconv"
)

type fpga_writer struct {
	address   uint8
	formatter func(string) []byte
}

var fpga_memMap = map[string]fpga_writer{
	"channel_cfgs.0.compressor_cfg.threshold":    {0x01, formatThreshold},
	"channel_cfgs.0.compressor_cfg.ratio":        {0x02, formatRatio},
	"channel_cfgs.0.compressor_cfg.post_gain":    {0x03, formatGain},
	"channel_cfgs.0.compressor_cfg.attack_time":  {0x04, formatTime},
	"channel_cfgs.0.compressor_cfg.release_time": {0x05, formatTime},
	"channel_cfgs.1.compressor_cfg.threshold":    {0x08, formatThreshold},
	"channel_cfgs.1.compressor_cfg.ratio":        {0x09, formatRatio},
	"channel_cfgs.1.compressor_cfg.post_gain":    {0x0A, formatGain},
	"channel_cfgs.1.compressor_cfg.attack_time":  {0x0B, formatTime},
	"channel_cfgs.1.compressor_cfg.release_time": {0x0C, formatTime},
	"channel_cfgs.2.compressor_cfg.threshold":    {0x10, formatThreshold},
	"channel_cfgs.2.compressor_cfg.ratio":        {0x11, formatRatio},
	"channel_cfgs.2.compressor_cfg.post_gain":    {0x12, formatGain},
	"channel_cfgs.2.compressor_cfg.attack_time":  {0x13, formatTime},
	"channel_cfgs.2.compressor_cfg.release_time": {0x14, formatTime},
	"channel_cfgs.3.compressor_cfg.threshold":    {0x18, formatThreshold},
	"channel_cfgs.3.compressor_cfg.ratio":        {0x19, formatRatio},
	"channel_cfgs.3.compressor_cfg.post_gain":    {0x1A, formatGain},
	"channel_cfgs.3.compressor_cfg.attack_time":  {0x1B, formatTime},
	"channel_cfgs.3.compressor_cfg.release_time": {0x1C, formatTime},
	"channel_cfgs.0.gate_cfg.threshold":          {0x20, formatThreshold},
	"channel_cfgs.0.gate_cfg.attack_time":        {0x21, formatTime},
	"channel_cfgs.0.gate_cfg.hold_time":          {0x22, formatTime},
	"channel_cfgs.0.gate_cfg.release_time":       {0x23, formatTime},
	"channel_cfgs.1.gate_cfg.threshold":          {0x28, formatThreshold},
	"channel_cfgs.1.gate_cfg.attack_time":        {0x29, formatTime},
	"channel_cfgs.1.gate_cfg.hold_time":          {0x2A, formatTime},
	"channel_cfgs.1.gate_cfg.release_time":       {0x2B, formatTime},
	"channel_cfgs.2.gate_cfg.threshold":          {0x30, formatThreshold},
	"channel_cfgs.2.gate_cfg.attack_time":        {0x31, formatTime},
	"channel_cfgs.2.gate_cfg.hold_time":          {0x32, formatTime},
	"channel_cfgs.2.gate_cfg.release_time":       {0x33, formatTime},
	"channel_cfgs.3.gate_cfg.threshold":          {0x38, formatThreshold},
	"channel_cfgs.3.gate_cfg.attack_time":        {0x39, formatTime},
	"channel_cfgs.3.gate_cfg.hold_time":          {0x3A, formatTime},
	"channel_cfgs.3.gate_cfg.release_time":       {0x3B, formatTime},
	"crosspoint_cfgs.0.0.volume":                 {0x40, formatVolume},
	"crosspoint_cfgs.0.0.pan":                    {0x41, formatPan},
	"crosspoint_cfgs.1.0.volume":                 {0x48, formatVolume},
	"crosspoint_cfgs.1.0.pan":                    {0x49, formatPan},
	"crosspoint_cfgs.2.0.volume":                 {0x50, formatVolume},
	"crosspoint_cfgs.2.0.pan":                    {0x51, formatPan},
	"crosspoint_cfgs.3.0.volume":                 {0x58, formatVolume},
	"crosspoint_cfgs.3.0.pan":                    {0x59, formatPan},
	"bus_cfgs.0.volume":                          {0x68, formatVolume},
	"bus_cfgs.0.pan":                             {0x69, formatPan},
}

func voltsToVal(volts float64) int64 {
	oldRange := float64(1.736 - 0)
	newRange := float64(0x7FFFFF - 0)
	return int64(((volts-0)*newRange)/oldRange) + 0
}

func dBuToVal(dbu float64) int64 {
	volts := math.Pow(10, dbu/20) * 0.775
	return voltsToVal(volts)
}

func int32ToInt24(val int32) []byte {
	b := make([]byte, 3)
	b[0] = byte(val >> 0)
	b[1] = byte(val >> 8)
	b[2] = byte(val >> 16)
	return b
}

func formatThreshold(str string) []byte {
	valstr, _ := strconv.ParseFloat(str, 64) // Ignore error
	val := int32(dBuToVal(valstr))
	return int32ToInt24(val)
}

func formatRatio(str string) []byte {
	valstr, _ := strconv.ParseFloat(str, 64) // Ignore error
	val := int32(dBuToVal(valstr))
	return int32ToInt24(val)
}

func formatGain(str string) []byte {
	valstr, _ := strconv.ParseFloat(str, 64) // Ignore error
	val := int32(dBuToVal(valstr))
	return int32ToInt24(val)
}

func formatTime(str string) []byte {
	valstr, _ := strconv.ParseFloat(str, 64) // Ignore error
	val := int32(valstr)
	return int32ToInt24(val)
}

func formatPan(str string) []byte {
	valstr, _ := strconv.ParseFloat(str, 64) // Ignore error
	val := int32(valstr)
	return int32ToInt24(val)
}

func formatVolume(str string) []byte {
	valstr, _ := strconv.ParseFloat(str, 64) // Ignore error
	val := int32(dBuToVal(valstr))
	return int32ToInt24(val)
}
