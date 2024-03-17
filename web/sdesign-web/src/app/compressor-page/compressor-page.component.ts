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
  selector: 'app-compressor-page',
  standalone: true,
  imports: [
    FormsModule,
    MatCardModule,
    MatListModule,
    MatSliderModule,
    MatSlideToggleModule
  ],
  templateUrl: './compressor-page.component.html',
  styleUrl: './compressor-page.component.scss'
})
export class CompressorPageComponent implements OnInit {
  @Input()
  backendWs: BackendWsService;

  public selectedChannel: number = 0;
  public bypass: Boolean = true;
  public preGain: number = 0;
  public postGain: number = 0;
  public attackTime: number = 0;
  public releaseTime: number = 0;
  public threshold: number = 0;
  public ratio: number = 0;

  ngOnInit(): void {
      this.backendWs.ShowConfig$.subscribe({
        next: cfg => this.updateFromCfg(cfg)
      })
  }

  private updateFromCfg(cfg: ShowConfig): void {
    this.selectedChannel = cfg.selected_channel;
    this.bypass = cfg.channel_cfgs[this.selectedChannel].compressor_cfg.bypass;
    this.preGain = cfg.channel_cfgs[this.selectedChannel].compressor_cfg.pre_gain;
    this.postGain = cfg.channel_cfgs[this.selectedChannel].compressor_cfg.post_gain;
    this.attackTime = cfg.channel_cfgs[this.selectedChannel].compressor_cfg.attack_time;
    this.releaseTime = cfg.channel_cfgs[this.selectedChannel].compressor_cfg.release_time;
    this.threshold = cfg.channel_cfgs[this.selectedChannel].compressor_cfg.threshold;
    this.ratio = cfg.channel_cfgs[this.selectedChannel].compressor_cfg.ratio;
  }

  public updateBypass(bypass: Boolean) {
    let request: APIRequest = {
      "method": APICommandMethod.SHOW_SET,
      "path": "channel_cfgs.".concat(String(this.selectedChannel).concat(".compressor_cfg.bypass")),
      "data": String(bypass)
    }
    this.backendWs.SendRequest(request);
  }

  public updatePreGain(preGain: number) {
    let request: APIRequest = {
      "method": APICommandMethod.SHOW_SET,
      "path": "channel_cfgs.".concat(String(this.selectedChannel).concat(".compressor_cfg.pre_gain")),
      "data": String(preGain)
    }
    this.backendWs.SendRequest(request);
  }

  public updatePostGain(postGain: number) {
    let request: APIRequest = {
      "method": APICommandMethod.SHOW_SET,
      "path": "channel_cfgs.".concat(String(this.selectedChannel).concat(".compressor_cfg.post_gain")),
      "data": String(postGain)
    }
    this.backendWs.SendRequest(request);
  }

  public updateAttackTime(time: number) {
    let request: APIRequest = {
      "method": APICommandMethod.SHOW_SET,
      "path": "channel_cfgs.".concat(String(this.selectedChannel).concat(".compressor_cfg.attack_time")),
      "data": String(time)
    }
    this.backendWs.SendRequest(request);
  }

  public updateReleaseTime(time: number) {
    let request: APIRequest = {
      "method": APICommandMethod.SHOW_SET,
      "path": "channel_cfgs.".concat(String(this.selectedChannel).concat(".compressor_cfg.release_time")),
      "data": String(time)
    }
    this.backendWs.SendRequest(request);
  }

  public updateThreshold(threshold: number) {
    let request: APIRequest = {
      "method": APICommandMethod.SHOW_SET,
      "path": "channel_cfgs.".concat(String(this.selectedChannel).concat(".compressor_cfg.threshold")),
      "data": String(threshold)
    }
    this.backendWs.SendRequest(request);
  }

  public updateRatio(ratio: number) {
    let request: APIRequest = {
      "method": APICommandMethod.SHOW_SET,
      "path": "channel_cfgs.".concat(String(this.selectedChannel).concat(".compressor_cfg.ratio")),
      "data": String(ratio)
    }
    this.backendWs.SendRequest(request);
  }
}
