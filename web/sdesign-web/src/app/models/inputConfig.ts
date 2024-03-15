import { ShowConfig } from "./showConfig";

export interface InputConfig {
  invert_phase: boolean;
  stereo_group: boolean;
  gain: number;
}


export function SetInputConfigValue(cfg: ShowConfig, path: string[], value: string): ShowConfig {
  switch (path[3]) {
    case "invert_phase":
      cfg.channel_cfgs[+path[1]].input_cfg.invert_phase = value.toLowerCase() == 'true';
      break;
    case "stereo_group":
      cfg.channel_cfgs[+path[1]].input_cfg.stereo_group = value.toLowerCase() == 'true';
      break;
    case "gain":
      cfg.channel_cfgs[+path[1]].input_cfg.gain = +value;
      break;
  }
  return cfg;
}
