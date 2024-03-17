import { Component, Input, OnInit } from '@angular/core';
import { FormsModule } from '@angular/forms';
import { MatCardModule } from '@angular/material/card';
import { MatListModule } from '@angular/material/list';
import { BackendWsService } from '../backend-ws.service';
import { ShowConfig } from '../models/showConfig';
import { APICommandMethod, APIRequest } from '../models/api';
import { MatSliderModule } from '@angular/material/slider';
import { MatSlideToggleModule } from '@angular/material/slide-toggle';

@Component({
  selector: 'app-gate-page',
  standalone: true,
  imports: [
    FormsModule,
    MatCardModule,
    MatListModule,
    MatSliderModule,
    MatSlideToggleModule
  ],
  templateUrl: './gate-page.component.html',
  styleUrl: './gate-page.component.scss'
})
export class GatePageComponent implements OnInit {
  @Input()
  backendWs: BackendWsService;

  public selectedChannel: number = 0;
  public bypass: Boolean = true;
  public attackTime: number = 0;
  public holdTime: number = 0;
  public releaseTime: number = 0;
  public threshold: number = 0;
  public depth: number = 0;

  ngOnInit(): void {
    this.backendWs.ShowConfig$.subscribe({
      next: cfg => this.updateFromCfg(cfg)
    })
  }

  private updateFromCfg(cfg: ShowConfig): void {
    this.selectedChannel = cfg.selected_channel;
    this.bypass = cfg.channel_cfgs[this.selectedChannel].gate_cfg.bypass;
    this.attackTime = cfg.channel_cfgs[this.selectedChannel].gate_cfg.attack_time;
    this.holdTime = cfg.channel_cfgs[this.selectedChannel].gate_cfg.hold_time;
    this.releaseTime = cfg.channel_cfgs[this.selectedChannel].gate_cfg.release_time;
    this.threshold = cfg.channel_cfgs[this.selectedChannel].gate_cfg.threshold;
    this.depth = cfg.channel_cfgs[this.selectedChannel].gate_cfg.depth;
  }

  public updateBypass(bypass: Boolean) {
    let request: APIRequest = {
      "method": APICommandMethod.SHOW_SET,
      "path": "channel_cfgs.".concat(String(this.selectedChannel).concat(".gate_cfg.bypass")),
      "data": String(bypass)
    }
    this.backendWs.SendRequest(request);
  }

  public updateAttackTime(time: number) {
    let request: APIRequest = {
      "method": APICommandMethod.SHOW_SET,
      "path": "channel_cfgs.".concat(String(this.selectedChannel).concat(".gate_cfg.attack_time")),
      "data": String(time)
    }
    this.backendWs.SendRequest(request);
  }

  public updateHoldTime(time: number) {
    let request: APIRequest = {
      "method": APICommandMethod.SHOW_SET,
      "path": "channel_cfgs.".concat(String(this.selectedChannel).concat(".gate_cfg.hold_time")),
      "data": String(time)
    }
    this.backendWs.SendRequest(request);
  }

  public updateReleaseTime(time: number) {
    let request: APIRequest = {
      "method": APICommandMethod.SHOW_SET,
      "path": "channel_cfgs.".concat(String(this.selectedChannel).concat(".gate_cfg.release_time")),
      "data": String(time)
    }
    this.backendWs.SendRequest(request);
  }

  public updateThreshold(value: number) {
    let request: APIRequest = {
      "method": APICommandMethod.SHOW_SET,
      "path": "channel_cfgs.".concat(String(this.selectedChannel).concat(".gate_cfg.threshold")),
      "data": String(value)
    }
    this.backendWs.SendRequest(request);
  }

  public updateDepth(value: number) {
    let request: APIRequest = {
      "method": APICommandMethod.SHOW_SET,
      "path": "channel_cfgs.".concat(String(this.selectedChannel).concat(".gate_cfg.depth")),
      "data": String(value)
    }
    this.backendWs.SendRequest(request);
  }
}
