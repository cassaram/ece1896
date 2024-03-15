import { Component, Input, OnInit, inject } from '@angular/core';
import { RouterOutlet } from '@angular/router';
import { BreakpointObserver, Breakpoints } from '@angular/cdk/layout';
import { AsyncPipe } from '@angular/common';
import { MatToolbarModule } from '@angular/material/toolbar';
import { MatButtonModule } from '@angular/material/button';
import { MatSidenavModule } from '@angular/material/sidenav';
import { MatListModule } from '@angular/material/list';
import { MatIconModule } from '@angular/material/icon';
import { MatTabsModule } from '@angular/material/tabs';
import { Observable, Subject } from 'rxjs';
import { map, shareReplay } from 'rxjs/operators';
import { InputbarComponent } from '../inputbar/inputbar.component';
import { InputPageComponent } from '../input-page/input-page.component';
import { BackendWsService } from '../backend-ws.service';
import { ShowConfig } from '../models/showConfig';


@Component({
  selector: 'app-mainpage',
  templateUrl: './mainpage.component.html',
  styleUrl: './mainpage.component.scss',
  standalone: true,
  imports: [
    RouterOutlet,
    MatToolbarModule,
    MatButtonModule,
    MatSidenavModule,
    MatListModule,
    MatIconModule,
    MatTabsModule,
    AsyncPipe,
    InputbarComponent,
    InputPageComponent,
  ]
})
export class MainpageComponent implements OnInit {
  @Input()
  backendWs: BackendWsService;

  public showName: string = "";

  ngOnInit(): void {
      this.backendWs.ShowConfig$.subscribe({
        next: cfg => this.updateShowName(cfg)
      });
  }

  private updateShowName(cfg: ShowConfig): void {
    this.showName = cfg.name;
  }
}
