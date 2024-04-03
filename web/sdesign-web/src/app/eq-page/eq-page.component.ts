import { Component, Input, OnInit } from '@angular/core';
import { MatCardModule } from '@angular/material/card';
import { BackendWsService } from '../backend-ws.service';
import { ShowConfig } from '../models/showConfig';
import { MatListModule } from '@angular/material/list';
import { MatSlideToggle } from '@angular/material/slide-toggle';
import { MatSliderModule } from '@angular/material/slider';
import { FormsModule } from '@angular/forms';
import { APIRequest, APICommandMethod } from '../models/api';

@Component({
  selector: 'app-eq-page',
  standalone: true,
  imports: [
    FormsModule,
    MatCardModule,
    MatListModule,
    MatSlideToggle,
    MatSliderModule
  ],
  templateUrl: './eq-page.component.html',
  styleUrl: './eq-page.component.scss'
})
export class EqPageComponent implements OnInit{
  public selectedChannel: number = 0;
  public eqBypass: Boolean = true;
  public highPassEnable: Boolean = false;
  public highPassGain: number = 0;
  public midEnable: Boolean = false;
  public midGain: number = 0;
  public lowPassEnable: Boolean = false;
  public lowPassGain: number = 0;

  constructor(private backendWs: BackendWsService) {}

  ngOnInit(): void {
    this.backendWs.ShowConfig$.subscribe({
      next: cfg => this.updateFromConfig(cfg)
    })
  }

  private updateFromConfig(cfg: ShowConfig): void {
    this.selectedChannel = cfg.selected_channel;
    this.eqBypass = cfg.channel_cfgs[this.selectedChannel].eq_cfg.bypass;
    this.highPassEnable = cfg.channel_cfgs[this.selectedChannel].eq_cfg.high_pass_enable;
    this.highPassGain = cfg.channel_cfgs[this.selectedChannel].eq_cfg.high_pass_gain;
    this.midEnable = cfg.channel_cfgs[this.selectedChannel].eq_cfg.mid_enable;
    this.midGain = cfg.channel_cfgs[this.selectedChannel].eq_cfg.mid_gain;
    this.lowPassEnable = cfg.channel_cfgs[this.selectedChannel].eq_cfg.low_pass_enable;
    this.lowPassGain = cfg.channel_cfgs[this.selectedChannel].eq_cfg.low_pass_gain;
  }

  public updateEQBypass(): void {
    let request: APIRequest = {
      "method": APICommandMethod.SHOW_SET,
      "path": "channel_cfgs.".concat(String(this.selectedChannel).concat(".eq_cfg.bypass")),
      "data": String(!this.eqBypass)
    }
    this.backendWs.SendRequest(request);
  }

  public updateHighPassEnable(): void {
    let request: APIRequest = {
      "method": APICommandMethod.SHOW_SET,
      "path": "channel_cfgs.".concat(String(this.selectedChannel).concat(".eq_cfg.high_pass_enable")),
      "data": String(!this.highPassEnable)
    }
    this.backendWs.SendRequest(request);
  }

  public updateHighPassGain(): void {
    let request: APIRequest = {
      "method": APICommandMethod.SHOW_SET,
      "path": "channel_cfgs.".concat(String(this.selectedChannel).concat(".eq_cfg.high_pass_gain")),
      "data": String(this.highPassGain)
    }
    this.backendWs.SendRequest(request);
  }

  public updateMidEnable(): void {
    let request: APIRequest = {
      "method": APICommandMethod.SHOW_SET,
      "path": "channel_cfgs.".concat(String(this.selectedChannel).concat(".eq_cfg.mid_enable")),
      "data": String(!this.midEnable)
    }
    this.backendWs.SendRequest(request);
  }

  public updateMidGain(): void {
    let request: APIRequest = {
      "method": APICommandMethod.SHOW_SET,
      "path": "channel_cfgs.".concat(String(this.selectedChannel).concat(".eq_cfg.mid_gain")),
      "data": String(this.midGain)
    }
    this.backendWs.SendRequest(request);
  }

  public updateLowPassEnable(): void {
    let request: APIRequest = {
      "method": APICommandMethod.SHOW_SET,
      "path": "channel_cfgs.".concat(String(this.selectedChannel).concat(".eq_cfg.low_pass_enable")),
      "data": String(!this.lowPassEnable)
    }
    this.backendWs.SendRequest(request);
  }

  public updateLowPassGain(): void {
    let request: APIRequest = {
      "method": APICommandMethod.SHOW_SET,
      "path": "channel_cfgs.".concat(String(this.selectedChannel).concat(".eq_cfg.low_pass_gain")),
      "data": String(this.lowPassGain)
    }
    this.backendWs.SendRequest(request);
  }
}
