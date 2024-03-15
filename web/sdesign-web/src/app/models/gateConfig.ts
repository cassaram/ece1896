import { ShowConfig } from "./showConfig";

export interface GateConfig {
  bypass: boolean;
  threshold: number;
  depth: number;
  attack_time: number;
  hold_time: number;
  release_time: number;
}

export function SetGateConfigValue(cfg: ShowConfig, path: string[], value: string): ShowConfig {
  switch(path[3]) {
    case "bypass":
      cfg.channel_cfgs[+path[1]].gate_cfg.bypass = value.toLowerCase() == 'true';
      break;
    case "threshold":
      cfg.channel_cfgs[+path[1]].gate_cfg.threshold = +value;
      break;
    case "depth":
      cfg.channel_cfgs[+path[1]].gate_cfg.depth = +value;
      break;
    case "attack_time":
      cfg.channel_cfgs[+path[1]].gate_cfg.attack_time = +value;
      break;
    case "hold_time":
      cfg.channel_cfgs[+path[1]].gate_cfg.hold_time = +value;
      break;
    case "release_time":
      cfg.channel_cfgs[+path[1]].gate_cfg.release_time = +value;
      break;
  }
  return cfg;
}
