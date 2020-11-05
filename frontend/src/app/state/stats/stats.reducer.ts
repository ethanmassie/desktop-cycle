import { StatsState } from '../index';
import { Action, createReducer, on } from '@ngrx/store';
import * as StatsActions from './stats.actions';

export const initialState: StatsState = {
  speed: 0,
  cadence: 0,
  distanceTraveled: 0,
};

const STATS_REDUCER = createReducer(
  initialState,
  on(StatsActions.setSpeed, (state, {speed}) => ({...state, speed})),
  on(StatsActions.setCadence, (state, {cadence}) => ({...state, cadence})),
  on(StatsActions.increaseDistanceTraveled, (state, {distance}) => ({...state, distanceTraveled: state.distanceTraveled + distance}))
);


export function statsReducer(state: StatsState | undefined, action: Action) {
  return STATS_REDUCER(state, action);
}
