import { Injectable } from '@angular/core';
import { ShowConfig } from './models/showConfig';
import { Observable, Observer, Subject } from 'rxjs';
import { webSocket, WebSocketSubject } from 'rxjs/webSocket';
import { APICommandMethod, APIRequest } from './models/api';

@Injectable({
  providedIn: 'root'
})
export class BackendWsService {
  private socket$: WebSocketSubject<any>;
  private ShowConfig = new Subject<ShowConfig>;
  public ShowConfig$ = this.ShowConfig.asObservable();

  constructor() {
  }

  public Connect(url: string) {
    this.socket$ = webSocket(url);
    this.socket$.subscribe({
      next: msg => this.handleResponse(msg),
      error: err => console.error(err)
    })
  }

  private handleResponse(response: APIRequest) {
    if (response.method == APICommandMethod.SHOW_LOAD) {
      // Parse json
      this.ShowConfig.next(JSON.parse(response.data));
    }
  }

  public Subscribe(): void {
  }

  public Disconnect(): void {
  }
}
