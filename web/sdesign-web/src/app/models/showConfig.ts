import { ChannelConfig, SetChannelConfigValue } from './channelConfig';
import { BusConfig, SetBusConfigValue } from './busConfig';
import { CrosspointConfig, SetCrosspointConfigValue } from './crosspointConfig';

export interface ShowConfig {
  name: string;
  filename: string;
  selected_channel: number;
  channel_cfgs: ChannelConfig[];
  bus_cfgs: BusConfig[];
  crosspoint_cfgs: CrosspointConfig[][];
}

export function SetShowConfigValue(
  cfg: ShowConfig,
  path: string[],
  value: string
): ShowConfig {
  switch (path[0]) {
    case 'name':
      cfg.name = value;
      break;
    case 'filename':
      cfg.filename = value;
      break;
    case 'selected_channel':
      cfg.selected_channel = +value;
      break;
    case 'channel_cfgs':
      cfg = SetChannelConfigValue(cfg, path, value);
      break;
    case 'bus_cfgs':
      cfg = SetBusConfigValue(cfg, path, value);
      break;
    case 'crosspoint_cfgs':
      cfg = SetCrosspointConfigValue(cfg, path, value);
      break;
  }

  return cfg;
}

export function GetBlankShowConfig(): ShowConfig {
  var cfg: ShowConfig = {
    name: 'NewShow',
    filename: 'NewShow.cfg',
    selected_channel: 0,
    channel_cfgs: [
      {
        name: 'Channel 1',
        id: 0,
        color: '#F6F',
        input_cfg: {
          invert_phase: false,
          stereo_group: false,
          gain: 0,
        },
        eq_cfg: {
          bypass: true,
          high_pass_enable: false,
          high_pass_gain: -24,
          mid_enable: false,
          mid_gain: 0,
          low_pass_enable: false,
          low_pass_gain: -24,
        },
        compressor_cfg: {
          bypass: true,
          pre_gain: 0,
          post_gain: 0,
          threshold: -18,
          ratio: 2,
          attack_time: 5,
          release_time: 75,
        },
        gate_cfg: {
          bypass: true,
          threshold: -42,
          depth: 0,
          attack_time: 16,
          hold_time: 0,
          release_time: 75,
        },
      },
      {
        name: 'Channel 2',
        id: 1,
        color: '#F6F',
        input_cfg: {
          invert_phase: false,
          stereo_group: false,
          gain: 0,
        },
        eq_cfg: {
          bypass: true,
          high_pass_enable: false,
          high_pass_gain: -24,
          mid_enable: false,
          mid_gain: 0,
          low_pass_enable: false,
          low_pass_gain: -24,
        },
        compressor_cfg: {
          bypass: true,
          pre_gain: 0,
          post_gain: 0,
          threshold: -18,
          ratio: 2,
          attack_time: 5,
          release_time: 75,
        },
        gate_cfg: {
          bypass: true,
          threshold: -42,
          depth: 0,
          attack_time: 16,
          hold_time: 0,
          release_time: 75,
        },
      },
      {
        name: 'Channel 3',
        id: 2,
        color: '#F6F',
        input_cfg: {
          invert_phase: false,
          stereo_group: false,
          gain: 0,
        },
        eq_cfg: {
          bypass: true,
          high_pass_enable: false,
          high_pass_gain: -24,
          mid_enable: false,
          mid_gain: 0,
          low_pass_enable: false,
          low_pass_gain: -24,
        },
        compressor_cfg: {
          bypass: true,
          pre_gain: 0,
          post_gain: 0,
          threshold: -18,
          ratio: 2,
          attack_time: 5,
          release_time: 75,
        },
        gate_cfg: {
          bypass: true,
          threshold: -42,
          depth: 0,
          attack_time: 16,
          hold_time: 0,
          release_time: 75,
        },
      },
      {
        name: 'Channel 4',
        id: 3,
        color: '#F6F',
        input_cfg: {
          invert_phase: false,
          stereo_group: false,
          gain: 0,
        },
        eq_cfg: {
          bypass: true,
          high_pass_enable: false,
          high_pass_gain: -24,
          mid_enable: false,
          mid_gain: 0,
          low_pass_enable: false,
          low_pass_gain: -24,
        },
        compressor_cfg: {
          bypass: true,
          pre_gain: 0,
          post_gain: 0,
          threshold: -18,
          ratio: 2,
          attack_time: 5,
          release_time: 75,
        },
        gate_cfg: {
          bypass: true,
          threshold: -42,
          depth: 0,
          attack_time: 16,
          hold_time: 0,
          release_time: 75,
        },
      },
    ],
    bus_cfgs: [
      {
        name: 'Bus 1',
        id: 0,
      },
      {
        name: 'Bus 2',
        id: 1,
      },
    ],
    crosspoint_cfgs: [
      [
        {
          bus_id: 0,
          channel_id: 0,
          enable: false,
          pan: 0,
          volume: -100,
        },
        {
          bus_id: 1,
          channel_id: 0,
          enable: false,
          pan: 0,
          volume: -100,
        },
      ],
      [
        {
          bus_id: 0,
          channel_id: 1,
          enable: false,
          pan: 0,
          volume: -100,
        },
        {
          bus_id: 1,
          channel_id: 1,
          enable: false,
          pan: 0,
          volume: -100,
        },
      ],
      [
        {
          bus_id: 0,
          channel_id: 2,
          enable: false,
          pan: 0,
          volume: -100,
        },
        {
          bus_id: 1,
          channel_id: 2,
          enable: false,
          pan: 0,
          volume: -100,
        },
      ],
      [
        {
          bus_id: 0,
          channel_id: 3,
          enable: false,
          pan: 0,
          volume: -100,
        },
        {
          bus_id: 1,
          channel_id: 3,
          enable: false,
          pan: 0,
          volume: -100,
        },
      ],
    ],
  };

  return cfg;
}
