package hweq

import (
	"errors"
	"strconv"
	"strings"

	"github.com/cassaram/ece1896/backend/ad5254"
	"github.com/cassaram/ece1896/backend/config"
)

type HardwareEQ struct {
	devices []*ad5254.AD5254
}

func NewHardwareEQ(devices []*ad5254.AD5254) *HardwareEQ {
	d := HardwareEQ{
		devices: devices,
	}
	return &d
}

// Endpoint for updating the all EQs from a config file
func (h *HardwareEQ) WriteFullConfig(config config.ShowConfig) error {
	errs := error(nil)
	pathsToUpdate := []string{
		"channel_cfgs.0.eq_cfg.high_pass_gain",
		"channel_cfgs.0.eq_cfg.mid_gain",
		"channel_cfgs.0.eq_cfg.low_pass_gain",
		"channel_cfgs.1.eq_cfg.high_pass_gain",
		"channel_cfgs.1.eq_cfg.mid_gain",
		"channel_cfgs.1.eq_cfg.low_pass_gain",
		"channel_cfgs.2.eq_cfg.high_pass_gain",
		"channel_cfgs.2.eq_cfg.mid_gain",
		"channel_cfgs.2.eq_cfg.low_pass_gain",
		"channel_cfgs.3.eq_cfg.high_pass_gain",
		"channel_cfgs.3.eq_cfg.mid_gain",
		"channel_cfgs.3.eq_cfg.low_pass_gain",
	}
	for _, path := range pathsToUpdate {
		err := h.UpdateFromPath(path, config)
		if err != nil {
			errs = errors.Join(errs, err)
		}
	}
	if errs != error(nil) {
		return errs
	}
	return nil
}

// Endpoint for updating the EQs from path variables
func (h *HardwareEQ) UpdateFromPath(path string, config config.ShowConfig) error {
	// Determine which path we are using
	switch path {
	// ----- BYPASS VALUE -----
	case "channel_cfgs.0.eq_cfg.bypass":
		errs := h.UpdateFromPath("channel_cfgs.0.eq_cfg.high_pass_gain", config)
		err := h.UpdateFromPath("channel_cfgs.0.eq_cfg.low_pass_gain", config)
		if err != nil {
			errs = errors.Join(errs, err)
		}
		err = h.UpdateFromPath("channel_cfgs.0.eq_cfg.mid_gain", config)
		if err != nil {
			errs = errors.Join(errs, err)
		}
		return errs
	case "channel_cfgs.1.eq_cfg.bypass":
		errs := h.UpdateFromPath("channel_cfgs.1.eq_cfg.high_pass_gain", config)
		err := h.UpdateFromPath("channel_cfgs.1.eq_cfg.low_pass_gain", config)
		if err != nil {
			errs = errors.Join(errs, err)
		}
		err = h.UpdateFromPath("channel_cfgs.1.eq_cfg.mid_gain", config)
		if err != nil {
			errs = errors.Join(errs, err)
		}
		return errs
	case "channel_cfgs.2.eq_cfg.bypass":
		errs := h.UpdateFromPath("channel_cfgs.2.eq_cfg.high_pass_gain", config)
		err := h.UpdateFromPath("channel_cfgs.2.eq_cfg.low_pass_gain", config)
		if err != nil {
			errs = errors.Join(errs, err)
		}
		err = h.UpdateFromPath("channel_cfgs.2.eq_cfg.mid_gain", config)
		if err != nil {
			errs = errors.Join(errs, err)
		}
		return errs
	case "channel_cfgs.3.eq_cfg.bypass":
		errs := h.UpdateFromPath("channel_cfgs.3.eq_cfg.high_pass_gain", config)
		err := h.UpdateFromPath("channel_cfgs.3.eq_cfg.low_pass_gain", config)
		if err != nil {
			errs = errors.Join(errs, err)
		}
		err = h.UpdateFromPath("channel_cfgs.3.eq_cfg.mid_gain", config)
		if err != nil {
			errs = errors.Join(errs, err)
		}
		return errs
	// ----- GAIN VALUES -----
	case "channel_cfgs.0.eq_cfg.high_pass_gain":
		val, err := pathToGainVal(path, config)
		if err != nil {
			return err
		}
		return h.devices[0].WriteToRegister(0, val)
	case "channel_cfgs.0.eq_cfg.mid_gain":
		val, err := pathToGainVal(path, config)
		if err != nil {
			return err
		}
		return h.devices[0].WriteToRegister(1, val)
	case "channel_cfgs.0.eq_cfg.low_pass_gain":
		val, err := pathToGainVal(path, config)
		if err != nil {
			return err
		}
		return h.devices[0].WriteToRegister(2, val)
	case "channel_cfgs.1.eq_cfg.high_pass_gain":
		val, err := pathToGainVal(path, config)
		if err != nil {
			return err
		}
		return h.devices[0].WriteToRegister(3, val)
	case "channel_cfgs.1.eq_cfg.mid_gain":
		val, err := pathToGainVal(path, config)
		if err != nil {
			return err
		}
		return h.devices[1].WriteToRegister(0, val)
	case "channel_cfgs.1.eq_cfg.low_pass_gain":
		val, err := pathToGainVal(path, config)
		if err != nil {
			return err
		}
		return h.devices[1].WriteToRegister(1, val)
	case "channel_cfgs.2.eq_cfg.high_pass_gain":
		val, err := pathToGainVal(path, config)
		if err != nil {
			return err
		}
		return h.devices[1].WriteToRegister(2, val)
	case "channel_cfgs.2.eq_cfg.mid_gain":
		val, err := pathToGainVal(path, config)
		if err != nil {
			return err
		}
		return h.devices[1].WriteToRegister(3, val)
	case "channel_cfgs.2.eq_cfg.low_pass_gain":
		val, err := pathToGainVal(path, config)
		if err != nil {
			return err
		}
		return h.devices[2].WriteToRegister(0, val)
	case "channel_cfgs.3.eq_cfg.high_pass_gain":
		val, err := pathToGainVal(path, config)
		if err != nil {
			return err
		}
		return h.devices[2].WriteToRegister(1, val)
	case "channel_cfgs.3.eq_cfg.mid_gain":
		val, err := pathToGainVal(path, config)
		if err != nil {
			return err
		}
		return h.devices[2].WriteToRegister(2, val)
	case "channel_cfgs.3.eq_cfg.low_pass_gain":
		val, err := pathToGainVal(path, config)
		if err != nil {
			return err
		}
		return h.devices[2].WriteToRegister(3, val)
	// ----- ENABLE VALUES -----
	case "channel_cfgs.0.eq_cfg.high_pass_enable":
		val, err := pathToGainVal("channel_cfgs.0.eq_cfg.high_pass_gain", config)
		if err != nil {
			return err
		}
		return h.devices[0].WriteToRegister(0, val)
	case "channel_cfgs.0.eq_cfg.mid_enable":
		val, err := pathToGainVal("channel_cfgs.0.eq_cfg.mid_gain", config)
		if err != nil {
			return err
		}
		return h.devices[0].WriteToRegister(1, val)
	case "channel_cfgs.0.eq_cfg.low_pass_enable":
		val, err := pathToGainVal("channel_cfgs.0.eq_cfg.low_pass_gain", config)
		if err != nil {
			return err
		}
		return h.devices[0].WriteToRegister(2, val)
	case "channel_cfgs.1.eq_cfg.high_pass_enable":
		val, err := pathToGainVal("channel_cfgs.1.eq_cfg.high_pass_gain", config)
		if err != nil {
			return err
		}
		return h.devices[0].WriteToRegister(3, val)
	case "channel_cfgs.1.eq_cfg.mid_enable":
		val, err := pathToGainVal("channel_cfgs.1.eq_cfg.mid_gain", config)
		if err != nil {
			return err
		}
		return h.devices[1].WriteToRegister(0, val)
	case "channel_cfgs.1.eq_cfg.low_pass_enable":
		val, err := pathToGainVal("channel_cfgs.1.eq_cfg.low_pass_gain", config)
		if err != nil {
			return err
		}
		return h.devices[1].WriteToRegister(1, val)
	case "channel_cfgs.2.eq_cfg.high_pass_enable":
		val, err := pathToGainVal("channel_cfgs.2.eq_cfg.high_pass_gain", config)
		if err != nil {
			return err
		}
		return h.devices[1].WriteToRegister(2, val)
	case "channel_cfgs.2.eq_cfg.mid_enable":
		val, err := pathToGainVal("channel_cfgs.2.eq_cfg.mid_gain", config)
		if err != nil {
			return err
		}
		return h.devices[1].WriteToRegister(3, val)
	case "channel_cfgs.2.eq_cfg.low_pass_enable":
		val, err := pathToGainVal("channel_cfgs.2.eq_cfg.low_pass_gain", config)
		if err != nil {
			return err
		}
		return h.devices[2].WriteToRegister(0, val)
	case "channel_cfgs.3.eq_cfg.high_pass_enable":
		val, err := pathToGainVal("channel_cfgs.3.eq_cfg.high_pass_gain", config)
		if err != nil {
			return err
		}
		return h.devices[2].WriteToRegister(1, val)
	case "channel_cfgs.3.eq_cfg.mid_enable":
		val, err := pathToGainVal("channel_cfgs.3.eq_cfg.mid_gain", config)
		if err != nil {
			return err
		}
		return h.devices[2].WriteToRegister(2, val)
	case "channel_cfgs.3.eq_cfg.low_pass_enable":
		val, err := pathToGainVal("channel_cfgs.3.eq_cfg.low_pass_gain", config)
		if err != nil {
			return err
		}
		return h.devices[2].WriteToRegister(3, val)
	}

	// Not handled by this function. Return.
	return nil
}

// Helper to go from a path for a gain and config to a value 0-255
func pathToGainVal(path string, config config.ShowConfig) (uint8, error) {
	paths := strings.Split(path, ".")
	// Get value
	dBuStr, err := config.GetValue(paths)
	if err != nil {
		return 128, err
	}
	val, err := dBuStrToVal(dBuStr)
	if err != nil {
		return 128, err
	}
	// Get enable value
	pathsEnable := make([]string, len(paths))
	copy(pathsEnable, paths)
	enablePathStr := ""
	switch paths[len(paths)-1] {
	case "high_pass_gain":
		enablePathStr = "high_pass_enable"
	case "mid_gain":
		enablePathStr = "mid_enable"
	case "low_pass_gain":
		enablePathStr = "low_pass_enable"
	}
	pathsEnable[len(pathsEnable)-1] = enablePathStr
	enableStr, err := config.GetValue(pathsEnable)
	if err != nil {
		return 128, err
	}
	enable, err := strconv.ParseBool(enableStr)
	if err != nil {
		return 128, err
	}
	// Get bypass value
	pathsBypass := make([]string, len(paths))
	copy(pathsBypass, paths)
	pathsBypass[len(paths)-1] = "bypass"
	bypassStr, err := config.GetValue(pathsBypass)
	if err != nil {
		return 128, err
	}
	bypass, err := strconv.ParseBool(bypassStr)
	if err != nil {
		return 128, err
	}
	// Return value needed depending on the enable
	if enable || !bypass {
		return val, nil
	}
	return 128, nil
}

// Helper to go from dBu string to resistance value 0-255
func dBuStrToVal(dBuStr string) (uint8, error) {
	dBu, err := strconv.ParseFloat(dBuStr, 64)
	if err != nil {
		return 128, err
	}
	return dBuToVal(dBu), nil
}

// Helper to convert from dBu to resistance value 0-255
func dBuToVal(dBu float64) uint8 {
	const min float64 = -4.5
	const max float64 = 4.5
	const oldRange float64 = (max - min)
	const newRange float64 = (255 - 0)
	return 255 - uint8(((dBu-min)*newRange)/oldRange)
}
