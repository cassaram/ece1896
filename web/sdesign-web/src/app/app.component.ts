import { Component, OnInit } from '@angular/core';

import { BackendWsService } from './backend-ws.service';

import { MatSlideToggleModule } from '@angular/material/slide-toggle';
import { MainpageComponent } from './mainpage/mainpage.component';

@Component({
  selector: 'app-root',
  standalone: true,
  imports: [MatSlideToggleModule, MainpageComponent],
  templateUrl: './app.component.html',
  styleUrl: './app.component.scss'
})
export class AppComponent implements OnInit{
  title = 'sdesign-web';

  constructor(
    private backendWs: BackendWsService
  ) {}

  ngOnInit(): void {
    this.backendWs.Connect('ws://192.168.8.10:8080/api/v1/ws');
  }

}
