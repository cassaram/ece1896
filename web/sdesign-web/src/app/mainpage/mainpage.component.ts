import { Component, Input, OnInit, inject } from '@angular/core';
import { RouterOutlet } from '@angular/router';
import { BreakpointObserver, Breakpoints } from '@angular/cdk/layout';
import { AsyncPipe, CommonModule } from '@angular/common';
import { MatToolbarModule } from '@angular/material/toolbar';
import { MatButtonModule } from '@angular/material/button';
import { MatSidenavModule } from '@angular/material/sidenav';
import { MatListModule } from '@angular/material/list';
import { MatIconModule } from '@angular/material/icon';
import { MatTabsModule } from '@angular/material/tabs';
import { MatDialog, MatDialogModule } from '@angular/material/dialog';
import { Observable, Subject } from 'rxjs';
import { map, shareReplay } from 'rxjs/operators';
import { InputbarComponent } from '../inputbar/inputbar.component';
import { InputPageComponent } from '../input-page/input-page.component';
import { BackendWsService } from '../backend-ws.service';
import { ShowConfig } from '../models/showConfig';
import { EqPageComponent } from '../eq-page/eq-page.component';
import { ChannelDialogComponent } from '../channel-dialog/channel-dialog.component';
import { MatGridListModule } from '@angular/material/grid-list';
import { MatCardModule } from '@angular/material/card';
import { CompressorPageComponent } from '../compressor-page/compressor-page.component';
import { GatePageComponent } from '../gate-page/gate-page.component';
import { BusConfigComponent } from '../bus-config/bus-config.component';
import { LoadConfigPageComponent } from '../load-config-page/load-config-page.component';
import { RemoteControlComponent } from '../remote-control/remote-control.component';


@Component({
  selector: 'app-mainpage',
  templateUrl: './mainpage.component.html',
  styleUrl: './mainpage.component.scss',
  standalone: true,
  imports: [
    CommonModule,
    RouterOutlet,
    MatToolbarModule,
    MatButtonModule,
    MatSidenavModule,
    MatListModule,
    MatIconModule,
    MatTabsModule,
    MatDialogModule,
    MatCardModule,
    AsyncPipe,
    InputbarComponent,
    InputPageComponent,
    EqPageComponent,
    ChannelDialogComponent,
    CompressorPageComponent,
    GatePageComponent,
    BusConfigComponent,
    LoadConfigPageComponent,
    RemoteControlComponent
  ]
})
export class MainpageComponent implements OnInit {
  public showName: string = "";
  public selectedChannel: number = 0;
  public channelName: string = "";
  public channelColor: string = "";

  public channelConfigVisible: boolean = true;
  public busConfigVisible: boolean = false;
  public loadConfigVisible: boolean = false;
  public remoteControlVisible: boolean = false;

  constructor(public dialog: MatDialog, private backendWs: BackendWsService) {

  }

  ngOnInit(): void {
      this.backendWs.ShowConfig$.subscribe({
        next: cfg => this.updateFromCfg(cfg)
      });
  }

  public openDialog() {
    const dialogRef = this.dialog.open(ChannelDialogComponent, {
      data: {
        backendWs: this.backendWs,
        selectedChannel: this.selectedChannel,
        channelName: this.channelName,
        channelColor: this.channelColor,
      },
    });
  }

  private updateFromCfg(cfg: ShowConfig): void {
    this.showName = cfg.name;
    this.selectedChannel = cfg.selected_channel;
    this.channelName = cfg.channel_cfgs[this.selectedChannel].name;
    this.channelColor = cfg.channel_cfgs[this.selectedChannel].color;
  }

  public setPage(pageID: number): void {
    this.channelConfigVisible = false;
    this.busConfigVisible = false;
    this.loadConfigVisible = false;
    this.remoteControlVisible = false;

    switch (pageID) {
      case 0:
        this.channelConfigVisible = true;
        break;
      case 1:
        this.busConfigVisible = true;
        break;
      case 2:
        this.loadConfigVisible = true;
        break;
      case 3:
        this.remoteControlVisible = true;
        break;
    }
  }
}
