import { Component, OnInit, Input, Output } from '@angular/core';
import { MatGridListModule } from '@angular/material/grid-list';
import { ChannelConfig } from '../models/channelConfig';
import { BackendWsService } from '../backend-ws.service';
import { ShowConfig } from '../models/showConfig';
import { Subject } from 'rxjs';
import { APICommandMethod, APIRequest } from '../models/api';

@Component({
  selector: 'app-inputbar',
  templateUrl: './inputbar.component.html',
  styleUrl: './inputbar.component.scss',
  standalone: true,
  imports: [
    MatGridListModule,
  ]
})
export class InputbarComponent implements OnInit {
  @Input()
  backendWs: BackendWsService;

  channels: ChannelConfig[] = [];
  constructor(
  ) {}

  ngOnInit() {
    this.backendWs.ShowConfig$.subscribe({
      next: cfg => this.updateChannels(cfg)
    });
  }

  private updateChannels(showConfig: ShowConfig): void {
    this.channels = showConfig.channel_cfgs;
  }

  public setSelectedIdx(idx: number): void {
    var request: APIRequest = {
      "method": APICommandMethod.SHOW_SET,
      "path": "selected_channel",
      "data": String(idx)
    };
    this.backendWs.SendRequest(request);
  }
}
