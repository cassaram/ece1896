import { Injectable } from '@angular/core';
import { GetBlankShowConfig, SetShowConfigValue, ShowConfig } from './models/showConfig';
import { BehaviorSubject, Observable, Observer, Subject } from 'rxjs';
import { webSocket, WebSocketSubject } from 'rxjs/webSocket';
import { APICommandMethod, APIRequest } from './models/api';
import { ConfigFile } from './models/configFile';

@Injectable({
  providedIn: 'root'
})
export class BackendWsService {
  private socket$: WebSocketSubject<any>;
  private ShowConfig = new BehaviorSubject<ShowConfig>(GetBlankShowConfig());
  private ShowConfig_Cache: ShowConfig;
  public ShowConfig$ = this.ShowConfig.asObservable();

  private ConfigFiles = new BehaviorSubject<ConfigFile[]>([]);
  public ConfigFiles$ = this.ConfigFiles.asObservable();

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
      this.ShowConfig_Cache = JSON.parse(response.data);
      this.ShowConfig.next(this.ShowConfig_Cache);
    }
    if (response.method == APICommandMethod.SHOW_SET) {
      let cfg = SetShowConfigValue(this.ShowConfig_Cache, response.path.split("."), response.data)
      this.ShowConfig_Cache = cfg;
      this.ShowConfig.next(this.ShowConfig_Cache);
    }
    if (response.method == APICommandMethod.SHOW_LIST) {
      let data = JSON.parse(response.data) as ConfigFile[];
      console.log(data);
      this.ConfigFiles.next(data);
    }
    if (response.method == APICommandMethod.ERROR) {
      console.error(response.data);
    }
  }

  public SendRequest(request: APIRequest): void {
    this.socket$.next(request);
  }
}
