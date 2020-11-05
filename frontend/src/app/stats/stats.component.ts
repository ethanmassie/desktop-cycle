import { Component, OnInit } from '@angular/core';
import { Observable } from 'rxjs';
import { Store } from '@ngrx/store';
import { AppState } from '../state';

@Component({
  selector: 'app-stats',
  templateUrl: './stats.component.html',
  styleUrls: ['./stats.component.css']
})
export class StatsComponent implements OnInit {

  speed$: Observable<number>;
  cadence$: Observable<number>;
  distanceTraveled$: Observable<number>;

  constructor(private store: Store<AppState>) { }

  ngOnInit() {
    this.speed$ = this.store.select(state => state.stats.speed);
    this.distanceTraveled$ = this.store.select(state => state.stats.distanceTraveled);
    this.cadence$ = this.store.select(state => state.stats.cadence);
  }

}
