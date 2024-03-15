import { ShowConfig } from "./showConfig";

export interface BusConfig {
  name: string;
  id: number;
}


export function SetBusConfigValue(cfg: ShowConfig, path: string[], value: string): ShowConfig {
  switch (path[2]) {
    case "name":
      cfg.bus_cfgs[+path[1]].name = value;
      break;
    case "id":
      cfg.bus_cfgs[+path[1]].id = +value;
      break;
  }
  return cfg;
}
