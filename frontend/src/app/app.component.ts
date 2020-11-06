import { Component, OnInit } from '@angular/core';
import { interval } from 'rxjs';
import { Store } from '@ngrx/store';
import { AppState } from './state';
import * as StatsActions from './state/stats/stats.actions';

declare global {
  interface Window {
    backend: {
      getSpeed: () => Promise<number>;
      getCadence: () => Promise<number>;
      keyUp: (key: string) => Promise<void>;
      keyDown: (key: string) => Promise<void>;
      tapKey: (key: string) => Promise<void>;
    };
  }
}

const MILLISECONDS_IN_HOUR = 3600000;

@Component({
  selector: '[id="app"]',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent implements OnInit {
  title = 'Desktop Cycle';

  constructor(private store: Store<AppState>) {}

  ngOnInit(): void {
    let previousTime = new Date().getTime();
    interval(1500).subscribe(
      (_) => {
        window.backend.getSpeed().then(speed => {
          this.store.dispatch(StatsActions.setSpeed({speed}));
          const now = new Date().getTime();
          const deltaTime = now - previousTime;
          previousTime = now;
          const distance = (speed / MILLISECONDS_IN_HOUR) * deltaTime;
          this.store.dispatch(StatsActions.increaseDistanceTraveled({distance}));
        });
        window.backend.getCadence().then(cadence => this.store.dispatch(StatsActions.setCadence({cadence})));
      }
    );
  }
}
