import { Component, Input, OnInit, inject } from '@angular/core';
import { FormsModule } from '@angular/forms';
import { MatSlideToggle } from '@angular/material/slide-toggle';
import { MatSliderModule } from '@angular/material/slider';
import { MatGridListModule } from '@angular/material/grid-list';
import { MatButtonModule } from '@angular/material/button';
import { BackendWsService } from '../backend-ws.service';
import { ShowConfig } from '../models/showConfig';
import { APICommandMethod, APIRequest } from '../models/api';
import { MatCardModule } from '@angular/material/card';
import { MatListModule } from '@angular/material/list';

@Component({
  selector: 'app-input-page',
  templateUrl: './input-page.component.html',
  styleUrl: './input-page.component.scss',
  standalone: true,
  imports: [
    MatSlideToggle,
    MatSliderModule,
    MatGridListModule,
    FormsModule,
    MatButtonModule,
    MatCardModule,
    MatListModule
  ]
})
export class InputPageComponent implements OnInit {
  private selectedChannelIdx: number = 0;
  public invertPhase: Boolean = false;
  public stereoGroup: Boolean = false;
  public gain: number = 0;

  constructor(private backendWs: BackendWsService) {}

  ngOnInit(): void {
      this.backendWs.ShowConfig$.subscribe({
        next: cfg => this.updateFromCfg(cfg),
      })
  }

  private updateFromCfg(cfg: ShowConfig): void {
    this.selectedChannelIdx = cfg.selected_channel;
    this.invertPhase = cfg.channel_cfgs[this.selectedChannelIdx].input_cfg.invert_phase;
    this.stereoGroup = cfg.channel_cfgs[this.selectedChannelIdx].input_cfg.stereo_group;
    this.gain = cfg.channel_cfgs[this.selectedChannelIdx].input_cfg.gain;
  }

  public updateInvertPhase(): void {
    let request: APIRequest = {
      "method": APICommandMethod.SHOW_SET,
      "path": "channel_cfgs.".concat(String(this.selectedChannelIdx).concat(".input_cfg.invert_phase")),
      "data": String(!this.invertPhase)
    }
    this.backendWs.SendRequest(request);
  }

  public updateStereoGroup(): void {
    let request: APIRequest = {
      "method": APICommandMethod.SHOW_SET,
      "path": "channel_cfgs.".concat(String(this.selectedChannelIdx).concat(".input_cfg.stereo_group")),
      "data": String(!this.stereoGroup)
    }
    this.backendWs.SendRequest(request);
  }

  public updateGain(): void {
    let request: APIRequest = {
      "method": APICommandMethod.SHOW_SET,
      "path": "channel_cfgs.".concat(String(this.selectedChannelIdx).concat(".input_cfg.gain")),
      "data": String(this.gain)
    }
    this.backendWs.SendRequest(request);
  }
}
