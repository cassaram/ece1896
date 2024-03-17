import { Component, Inject, OnInit } from '@angular/core';
import { FormsModule } from '@angular/forms';
import { MatButtonModule } from '@angular/material/button';
import { MAT_DIALOG_DATA, MatDialogModule } from '@angular/material/dialog';
import { MatInputModule } from '@angular/material/input';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatListModule } from '@angular/material/list';
import { BackendWsService } from '../backend-ws.service';
import { ShowConfig } from '../models/showConfig';
import { NgxColorsModule } from 'ngx-colors';
import { APIRequest, APICommandMethod } from '../models/api';


export interface ChannelDialogData {
  backendWs: BackendWsService;
  selectedChannel: number;
  channelName: string;
  channelColor: string;
}

@Component({
  selector: 'app-channel-dialog',
  standalone: true,
  imports: [
    FormsModule,
    NgxColorsModule,
    MatDialogModule,
    MatButtonModule,
    MatListModule,
    MatInputModule,
    MatFormFieldModule
  ],
  templateUrl: './channel-dialog.component.html',
  styleUrl: './channel-dialog.component.scss'
})
export class ChannelDialogComponent implements OnInit{

  public selectedChannel: number = 0;
  public channelName: string = "";
  public channelColor: string = "#FFFFFF";

  constructor(@Inject(MAT_DIALOG_DATA) public data: ChannelDialogData) {

  }

  ngOnInit(): void {
    this.selectedChannel = this.data.selectedChannel;
    this.channelName = this.data.channelName;
    this.channelColor = this.data.channelColor;

    this.data.backendWs.ShowConfig$.subscribe({
      next: cfg => this.updateFromCfg(cfg)
    });
  }

  private updateFromCfg(cfg: ShowConfig): void {
    this.selectedChannel = cfg.selected_channel;
    this.channelName = cfg.channel_cfgs[this.selectedChannel].name;
    this.channelColor = cfg.channel_cfgs[this.selectedChannel].color;
  }

  public applySettings(): void {
    this.updateColor();
    this.updateName();
  }

  public updateColor(): void {
    let request: APIRequest = {
      "method": APICommandMethod.SHOW_SET,
      "path": "channel_cfgs.".concat(String(this.selectedChannel).concat(".color")),
      "data": this.channelColor
    }
    this.data.backendWs.SendRequest(request);
  }

  public updateName(): void {
    let request: APIRequest = {
      "method": APICommandMethod.SHOW_SET,
      "path": "channel_cfgs.".concat(String(this.selectedChannel).concat(".name")),
      "data": this.channelName
    }
    this.data.backendWs.SendRequest(request);
  }
}
