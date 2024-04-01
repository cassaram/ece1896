import { ShowConfig } from "./showConfig";

export interface BusConfig {
  name: string;
  id: number;
  master: boolean;
  pfl: boolean;
  afl: boolean;
}


export function SetBusConfigValue(cfg: ShowConfig, path: string[], value: string): ShowConfig {
  switch (path[2]) {
    case "name":
      cfg.bus_cfgs[+path[1]].name = value;
      break;
    case "id":
      cfg.bus_cfgs[+path[1]].id = +value;
      break;
    case "master":
      cfg.bus_cfgs[+path[1]].master = value.toLowerCase() == 'true';
      break;
    case "pfl":
      cfg.bus_cfgs[+path[1]].pfl = value.toLowerCase() == 'true';
      break;
    case "afl":
      cfg.bus_cfgs[+path[1]].afl = value.toLowerCase() == 'true';
      break;
  }
  return cfg;
}
