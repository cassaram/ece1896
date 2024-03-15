import { ChannelConfig, SetChannelConfigValue } from "./channelConfig";
import { BusConfig, SetBusConfigValue } from "./busConfig";
import { CrosspointConfig, SetCrosspointConfigValue } from "./crosspointConfig";

export interface ShowConfig {
  name: string;
  filename: string;
  selected_channel: number;
  channel_cfgs: ChannelConfig[];
  bus_cfgs: BusConfig[];
  crosspoint_cfgs: CrosspointConfig[][];
}

export function SetShowConfigValue(cfg: ShowConfig, path: string[], value: string): ShowConfig {
  switch (path[0]) {
    case "name":
      cfg.name = value;
      break;
    case "filename":
      cfg.filename = value;
      break;
    case "selected_channel":
      cfg.selected_channel = +value;
      break;
    case "channel_cfgs":
      cfg = SetChannelConfigValue(cfg, path, value);
      break;
    case "bus_cfgs":
      cfg = SetBusConfigValue(cfg, path, value);
      break;
    case "crosspoint_cfgs":
      cfg = SetCrosspointConfigValue(cfg, path, value);
      break;
  }

  return cfg
}
