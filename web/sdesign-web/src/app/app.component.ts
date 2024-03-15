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
  public backendWs = new BackendWsService();

  ngOnInit(): void {
    this.backendWs.Connect('ws://127.0.0.1:8080/api/v1/ws');
  }

}
