import { InputConfig, SetInputConfigValue } from "./inputConfig";
import { EQConfig, SetEQConfigValue } from "./eqConfig";
import { CompressorConfig, SetCompressorConfigValue } from "./compressorConfig";
import { GateConfig, SetGateConfigValue } from "./gateConfig";
import { ShowConfig } from "./showConfig";

export interface ChannelConfig {
    name: string;
    id: number;
    color: string;
    input_cfg: InputConfig;
    eq_cfg: EQConfig;
    compressor_cfg: CompressorConfig;
    gate_cfg: GateConfig;
    pfl: boolean;
    afl: boolean;
}

export function SetChannelConfigValue(cfg: ShowConfig, path: string[], value: string): ShowConfig {
  switch (path[2]) {
    case "name":
      cfg.channel_cfgs[+path[1]].name = value;
      break;
    case "id":
      cfg.channel_cfgs[+path[1]].id = +value;
      break;
    case "color":
      cfg.channel_cfgs[+path[1]].color = value;
      break;
    case "input_cfg":
      cfg = SetInputConfigValue(cfg, path, value);
      break;
    case "eq_cfg":
      cfg = SetEQConfigValue(cfg, path, value);
      break;
    case "compressor_cfg":
      cfg = SetCompressorConfigValue(cfg, path, value);
      break;
    case "gate_cfg":
      cfg = SetGateConfigValue(cfg, path, value);
      break;
    case "pfl":
      cfg.channel_cfgs[+path[1]].pfl = value == 'true';
      break;
    case "afl":
      cfg.channel_cfgs[+path[1]].afl = value == 'true';
      break;
  }
  return cfg
}
