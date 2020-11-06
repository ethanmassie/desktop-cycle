import { Component, OnInit, ViewChild } from '@angular/core';
import { MatTable, MatTableDataSource } from '@angular/material';
import { KeyboardService } from 'src/app/core/keyboard.service';
import { KeyboardConfig, KeyConfig } from 'src/app/core/model/keyboard-config';

const defaultKey: KeyConfig = {
  key: 'w',
  minSpeed: 0,
  maxSpeed: 0,
  action: 'HOLD'
};

@Component({
  selector: 'app-config-form',
  templateUrl: './config-form.component.html',
  styleUrls: ['./config-form.component.css']
})
export class ConfigFormComponent implements OnInit {

  @ViewChild('keysTable', {static: true}) keysTable: MatTable<KeyConfig>;

  displayedColumns = ['keyName', 'minSpeed', 'maxSpeed', 'action', 'delete'];
  dataSource = new MatTableDataSource<KeyConfig>([]);
  name = 'New Keyboard Config';

  constructor(private keyboardService: KeyboardService) { }

  ngOnInit() {
  }

  addKey() {
    this.dataSource.data.push(defaultKey);
    this.keysTable.renderRows();
  }

  deleteRow(index: number) {
    this.dataSource.data.splice(index, 1);
    this.keysTable.renderRows();
  }

  submit() {
    const keyboardConfig: KeyboardConfig = {
      name: this.name,
      id: 'new',
      keys: this.dataSource.data
    };
    this.keyboardService.startKeyboard(keyboardConfig);
  }

}
