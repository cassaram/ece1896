import { ShowConfig } from "./showConfig";

export interface CompressorConfig {
  bypass: boolean;
  pre_gain: number;
  post_gain: number;
  threshold: number;
  ratio: number;
  attack_time: number;
  release_time: number;
}

export function SetCompressorConfigValue(cfg: ShowConfig, path: string[], value: string): ShowConfig {
  switch (path[3]) {
    case "bypass":
      cfg.channel_cfgs[+path[1]].compressor_cfg.bypass = value.toLowerCase() == 'true';
      break;
    case "pre_gain":
      cfg.channel_cfgs[+path[1]].compressor_cfg.pre_gain = +value;
      break;
    case "post_gain":
      cfg.channel_cfgs[+path[1]].compressor_cfg.post_gain = +value;
      break;
    case "threshold":
      cfg.channel_cfgs[+path[1]].compressor_cfg.threshold = +value;
      break;
    case "ratio":
      cfg.channel_cfgs[+path[1]].compressor_cfg.ratio = +value;
      break;
    case "attack_time":
      cfg.channel_cfgs[+path[1]].compressor_cfg.attack_time = +value;
      break;
    case "release_time":
      cfg.channel_cfgs[+path[1]].compressor_cfg.release_time = +value;
      break;
  }
  return cfg;
}
