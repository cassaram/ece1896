import { ShowConfig } from "./showConfig";

export interface EQConfig {
  bypass: boolean;
  high_pass_enable: boolean;
  high_pass_gain: number;
  mid_enable: boolean;
  mid_gain: number;
  low_pass_enable: boolean;
  low_pass_gain: number;
}

export function SetEQConfigValue(cfg: ShowConfig, path: string[], value: string): ShowConfig {
  switch (path[3]) {
    case "bypass":
      cfg.channel_cfgs[+path[1]].eq_cfg.bypass = value.toLowerCase() == 'true';
      break;
    case "high_pass_enable":
      cfg.channel_cfgs[+path[1]].eq_cfg.high_pass_enable = value.toLowerCase() == 'true';
      break;
    case "high_pass_gain":
      cfg.channel_cfgs[+path[1]].eq_cfg.high_pass_gain = +value;
      break;
    case "mid_enable":
      cfg.channel_cfgs[+path[1]].eq_cfg.mid_enable = value.toLowerCase() == 'true';
      break;
    case "mid_gain":
      cfg.channel_cfgs[+path[1]].eq_cfg.mid_gain = +value;
      break;
    case "low_pass_enable":
      cfg.channel_cfgs[+path[1]].eq_cfg.low_pass_enable = value.toLowerCase() == 'true';
      break;
    case "low_pass_gain":
      cfg.channel_cfgs[+path[1]].eq_cfg.low_pass_gain = +value;
      break;
  }
  return cfg;
}
