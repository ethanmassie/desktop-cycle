import { Injectable } from '@angular/core';
import { select, Store } from '@ngrx/store';
import { BehaviorSubject, Subject } from 'rxjs';
import { takeUntil } from 'rxjs/operators';
import { AppState } from '../state';
import { KeyboardConfig, KeyConfig } from './model/keyboard-config';

@Injectable({
  providedIn: 'root'
})
export class KeyboardService {

  private stop$: Subject<void>;
  private executing = false;
  private heldKeys: {[key: string]: boolean} = {};
  private toggledKeys: {[key: string]: boolean} = {};

  constructor(private store: Store<AppState>) { }

  startKeyboard(config: KeyboardConfig) {
    if (this.executing) {
      return;
    }
    console.log(config);
    this.stop$ = new Subject<void>();
    this.store.pipe(
      takeUntil(this.stop$),
      select(state => state.stats.speed)
    ).subscribe(
      {
        next: (speed) => {
          config.keys
            .forEach(key => {
              if (speed >= key.minSpeed && speed <= key.maxSpeed) {
                this.activateKey(key);
              } else {
                this.deactivateKey(key);
              }
            });
        },
        complete: () => {
          this.executing = false;
          config.keys
            .filter(key => !!this.heldKeys[key.key])
            .forEach(key => this.deactivateKey(key));
          this.heldKeys = {};
          this.toggledKeys = {};
        }
      }
    );
  }

  stop() {
    this.stop$.next();
  }

  private activateKey(key: KeyConfig) {
    if (key.action === 'HOLD' && !this.heldKeys[key.key]) {
      this.heldKeys[key.key] = true;
      window.backend.keyDown(key.key).then();
    } else if (key.action === 'TOGGLE' && !this.toggledKeys[key.key]) {
      this.toggledKeys[key.key] = true;
      window.backend.tapKey(key.key).then();
    }
  }

  private deactivateKey(key: KeyConfig) {
    if (key.action === 'HOLD' && this.heldKeys[key.key] === true) {
      this.heldKeys[key.key] = false;
      window.backend.keyUp(key.key).then();
    } else if (key.action === 'TOGGLE' && this.toggledKeys[key.key] === true) {
      this.toggledKeys[key.key] = false;
      window.backend.tapKey(key.key).then();
    }
  }
}
