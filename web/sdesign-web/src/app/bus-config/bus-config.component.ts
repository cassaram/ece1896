import { Component, Input, OnInit } from '@angular/core';
import { MatTableModule } from '@angular/material/table';
import { BackendWsService } from '../backend-ws.service';
import { ShowConfig } from '../models/showConfig';
import { BusConfig } from '../models/busConfig';
import { APICommandMethod, APIRequest } from '../models/api';
import { MatSlideToggleModule } from '@angular/material/slide-toggle';
import { Console } from 'console';
import { MatInputModule } from '@angular/material/input';
import { FormsModule } from '@angular/forms';

@Component({
  selector: 'app-bus-config',
  standalone: true,
  imports: [
    FormsModule,
    MatTableModule,
    MatSlideToggleModule,
    MatInputModule
  ],
  templateUrl: './bus-config.component.html',
  styleUrl: './bus-config.component.scss'
})
export class BusConfigComponent implements OnInit {
  busConfigs: BusConfig[] = [];
  displayedColumns: string[] = ['id', 'name', 'master', 'pfl', 'afl'];

  constructor(private backendWs: BackendWsService) {}

  ngOnInit(): void {
    this.backendWs.ShowConfig$.subscribe({
      next: cfg => this.updateFromCfg(cfg)
    })
  }

  updateFromCfg(cfg: ShowConfig): void {
    this.busConfigs = cfg.bus_cfgs
  }

  setBusName(id: Number, name: string): void {
    let request: APIRequest = {
      method: APICommandMethod.SHOW_SET,
      path: "bus_cfgs.".concat(String(id).concat(".name")),
      data: name,
    };
    this.backendWs.SendRequest(request);
  }

  setBusMaster(id: Number, master: boolean): void {
    let request: APIRequest = {
      method: APICommandMethod.SHOW_SET,
      path: "bus_cfgs.".concat(String(id).concat(".master")),
      data: String(master),
    };
    this.backendWs.SendRequest(request);
  }

  setBusPFL(id: Number, pfl: boolean): void {
    let request: APIRequest = {
      method: APICommandMethod.SHOW_SET,
      path: "bus_cfgs.".concat(String(id).concat(".pfl")),
      data: String(pfl),
    };
    this.backendWs.SendRequest(request);
  }

  setBusAFL(id: Number, afl: boolean): void {
    let request: APIRequest = {
      method: APICommandMethod.SHOW_SET,
      path: "bus_cfgs.".concat(String(id).concat(".afl")),
      data: String(afl),
    };
    this.backendWs.SendRequest(request);
  }
}
