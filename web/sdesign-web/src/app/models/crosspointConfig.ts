import { ShowConfig } from "./showConfig";

export interface CrosspointConfig {
  bus_id: number;
  channel_id: number;
  enable: boolean;
  pan: number;
  volume: number;
}

export function SetCrosspointConfigValue(cfg: ShowConfig, path: string[], value: string): ShowConfig {
  switch (path[3]) {
    case "bus_id":
      cfg.crosspoint_cfgs[+path[1]][+path[2]].bus_id = +value;
      break;
    case "channel_id":
      cfg.crosspoint_cfgs[+path[1]][+path[2]].channel_id = +value;
      break;
    case "enable":
      cfg.crosspoint_cfgs[+path[1]][+path[2]].enable = value.toLowerCase() == 'true';
      break;
    case "pan":
      cfg.crosspoint_cfgs[+path[1]][+path[2]].pan = +value;
      break;
    case "volume":
      cfg.crosspoint_cfgs[+path[1]][+path[2]].volume = +value;
      break;
  }
  return cfg;
}
