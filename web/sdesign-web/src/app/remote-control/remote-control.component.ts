import { Component, OnInit } from '@angular/core';
import { BackendWsService } from '../backend-ws.service';
import { MatCardModule } from '@angular/material/card';
import { MatButtonModule } from '@angular/material/button';
import { ShowConfig } from '../models/showConfig';
import { APICommandMethod, APIRequest } from '../models/api';
import { MatSliderModule } from '@angular/material/slider';
import { FormsModule } from '@angular/forms';
import { MatSlideToggleModule } from '@angular/material/slide-toggle';
import { MatListModule } from '@angular/material/list';
import { MatRadioModule } from '@angular/material/radio';

@Component({
  selector: 'app-remote-control',
  standalone: true,
  imports: [
    FormsModule,
    MatCardModule,
    MatButtonModule,
    MatSliderModule,
    MatSlideToggleModule,
    MatListModule,
    MatRadioModule
  ],
  templateUrl: './remote-control.component.html',
  styleUrl: './remote-control.component.scss'
})
export class RemoteControlComponent implements OnInit {
  public c1_volume: number = 0;
  public c1_mute: boolean = false;
  public c1_mon: number = 0;
  public c1_pan: number = 0;

  public c2_volume: number = 0;
  public c2_mute: boolean = false;
  public c2_mon: number = 0;
  public c2_pan: number = 0;

  public c3_volume: number = 0;
  public c3_mute: boolean = false;
  public c3_mon: number = 0;
  public c3_pan: number = 0;

  public c4_volume: number = 0;
  public c4_mute: boolean = false;
  public c4_mon: number = 0;
  public c4_pan: number = 0;

  constructor(private backendWs: BackendWsService) {

  }

  ngOnInit(): void {
    this.backendWs.ShowConfig$.subscribe({
      next: cfg => this.updateFromCfg(cfg)
    })
  }

  private updateFromCfg(cfg: ShowConfig): void {
    this.c1_volume = cfg.crosspoint_cfgs[0][0].volume;
    this.c1_mute = !cfg.crosspoint_cfgs[0][0].enable;
    this.c1_mon = cfg.channel_cfgs[0].monitor;
    this.c1_pan = cfg.crosspoint_cfgs[0][0].pan;
    this.c2_volume = cfg.crosspoint_cfgs[1][0].volume;
    this.c2_mute = !cfg.crosspoint_cfgs[1][0].enable;
    this.c2_mon = cfg.channel_cfgs[1].monitor;
    this.c2_pan = cfg.crosspoint_cfgs[1][0].pan;
    this.c3_volume = cfg.crosspoint_cfgs[2][0].volume;
    this.c3_mute = !cfg.crosspoint_cfgs[2][0].enable;
    this.c3_mon = cfg.channel_cfgs[2].monitor;
    this.c3_pan = cfg.crosspoint_cfgs[2][0].pan;
    this.c4_volume = cfg.crosspoint_cfgs[3][0].volume;
    this.c4_mute = !cfg.crosspoint_cfgs[3][0].enable;
    this.c4_mon = cfg.channel_cfgs[3].monitor;
    this.c4_pan = cfg.crosspoint_cfgs[3][0].pan;
  }

  setVolume(channel: number, volume: number): void {
    let request: APIRequest = {
      "method": APICommandMethod.SHOW_SET,
      "path": "crosspoint_cfgs.".concat(String(channel).concat(".0.volume")),
      "data": String(volume)
    }
    this.backendWs.SendRequest(request);
  }

  setMute(channel: number, muted: boolean): void {
    let request: APIRequest = {
      "method": APICommandMethod.SHOW_SET,
      "path": "crosspoint_cfgs.".concat(String(channel).concat(".0.enable")),
      "data": String(!muted)
    }
    this.backendWs.SendRequest(request);
  }

  setPflAfl(channel: number, mon: number): void {
    let reqData = mon;

    let request: APIRequest = {
      "method": APICommandMethod.SHOW_SET,
      "path": "channel_cfgs.".concat(String(channel).concat(".monitor")),
      "data": String(reqData)
    }
    this.backendWs.SendRequest(request);
  }

  setPan(channel: number, value: number): void {
    let request: APIRequest = {
      "method": APICommandMethod.SHOW_SET,
      "path": "crosspoint_cfgs.".concat(String(channel).concat(".0.pan")),
      "data": String(value)
    }
    this.backendWs.SendRequest(request);
  }
}
