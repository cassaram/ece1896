import { Component, OnInit, Input } from '@angular/core';
import { Breakpoints, BreakpointObserver } from '@angular/cdk/layout';
import { map } from 'rxjs/operators';
import { AsyncPipe } from '@angular/common';
import { MatGridListModule } from '@angular/material/grid-list';
import { MatMenuModule } from '@angular/material/menu';
import { MatIconModule } from '@angular/material/icon';
import { MatButtonModule } from '@angular/material/button';
import { MatCardModule } from '@angular/material/card';
import { ChannelConfig } from '../models/channelConfig';
import { BackendWsService } from '../backend-ws.service';
import { ShowConfig } from '../models/showConfig';

@Component({
  selector: 'app-inputbar',
  templateUrl: './inputbar.component.html',
  styleUrl: './inputbar.component.scss',
  standalone: true,
  imports: [
    AsyncPipe,
    MatGridListModule,
    MatMenuModule,
    MatIconModule,
    MatButtonModule,
    MatCardModule
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
    console.log(this.channels);
  }
}
