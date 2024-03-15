import { ChannelConfig } from "./channelConfig";
import { BusConfig } from "./busConfig";
import { CrosspointConfig } from "./crosspointConfig";

export interface ShowConfig {
  name: string;
  filename: string;
  channel_cfgs: ChannelConfig[];
  bus_cfgs: BusConfig[];
  crosspoint_cfgs: CrosspointConfig[][];
}
