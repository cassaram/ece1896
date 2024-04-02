import { Component, Input, OnInit } from '@angular/core';
import { BackendWsService } from '../backend-ws.service';
import { ConfigFile } from '../models/configFile';
import { MatTableModule } from '@angular/material/table';
import { MatButtonModule } from '@angular/material/button';
import { APICommandMethod, APIRequest } from '../models/api';
import { MatSortModule } from '@angular/material/sort';
import { MatDialog, MatDialogModule } from '@angular/material/dialog';
import { LoadConfirmDialogComponent } from './load-confirm-dialog/load-confirm-dialog.component';
import path from 'path';
import { SaveDialogComponent } from './save-dialog/save-dialog.component';
import { ShowConfig } from '../models/showConfig';
import { RenameDialogComponent } from './rename-dialog/rename-dialog.component';

const sleep = (ms: number) => new Promise((resolve) => setTimeout(resolve, ms));

@Component({
  selector: 'app-load-config-page',
  standalone: true,
  imports: [
    MatTableModule,
    MatButtonModule,
    MatSortModule,
    MatDialogModule
  ],
  templateUrl: './load-config-page.component.html',
  styleUrl: './load-config-page.component.scss'
})
export class LoadConfigPageComponent implements OnInit {
  @Input()
  backendWs: BackendWsService;

  displayedColumns: string[] = ['name', 'filename', 'mod_time', 'size', 'actions'];
  configFiles: ConfigFile[] = [];
  currentFilename: string = "";
  currentName: string = "";

  constructor(public dialog: MatDialog) {

  }

  ngOnInit(): void {
    this.backendWs.ConfigFiles$.subscribe({
      next: cfg => this.updateFromCfg(cfg)
    });
    this.backendWs.ShowConfig$.subscribe({
      next: cfg => this.updateFromShowCfg(cfg)
    });
    this.refresh();
  }

  updateFromCfg(cfg: ConfigFile[]): void {
    this.configFiles = cfg;
  }

  updateFromShowCfg(cfg: ShowConfig): void {
    this.currentFilename = cfg.filename;
    this.currentName = cfg.name;
  }

  refresh(): void {
    let request: APIRequest = {
      method: APICommandMethod.SHOW_LIST,
      path: "",
      data: ""
    };
    this.backendWs.SendRequest(request);
  }

  loadConfig(name: string, filename: string): void {
    const dialogRef = this.dialog.open(LoadConfirmDialogComponent, {
      data:{
        name: name,
        filename: filename,
      },
    });

    dialogRef.afterClosed().subscribe(result => {
      if (result == true) {
        let request: APIRequest = {
          method: APICommandMethod.SHOW_LOAD,
          path: filename,
          data: "",
        };
        this.backendWs.SendRequest(request);
      }
    })
  }

  saveConfig(): void {
    const dialogRef = this.dialog.open(SaveDialogComponent, {
      data:{
        saveas: false,
        filename: this.currentFilename
      },
    })

    dialogRef.afterClosed().subscribe(result => {
      if (result) {
        let request: APIRequest = {
          method: APICommandMethod.SHOW_SAVE,
          path: "",
          data: "",
        };
        this.backendWs.SendRequest(request);
      }
    })

    this.delayRefresh();
  }

  saveAsConfig(): void {
    const dialogRef = this.dialog.open(SaveDialogComponent, {
      data:{
        saveas: true,
        filename: this.currentFilename
      },
    })

    dialogRef.afterClosed().subscribe(result => {
      if (result) {
        let request: APIRequest = {
          method: APICommandMethod.SHOW_SAVEAS,
          path: "",
          data: result,
        };
        this.backendWs.SendRequest(request);
      }
    })

    this.delayRefresh();
  }

  renameCurrentConfig(): void {
    const dialogRef = this.dialog.open(RenameDialogComponent, {
      data:{
        name: this.currentName
      },
    })

    dialogRef.afterClosed().subscribe(result => {
      if (result) {
        let request: APIRequest = {
          method: APICommandMethod.SHOW_SET,
          path: "name",
          data: result,
        };
        this.backendWs.SendRequest(request);
      }
    })

    this.delayRefresh();
  }

  async delayRefresh() {
    sleep(50);
    this.refresh();
  }
}