import { statsReducer } from './stats/stats.reducer';

export interface StatsState {
  speed: number;
  cadence: number;
  distanceTraveled: number;
}

export interface AppState {
  stats: StatsState;
}

export const reducerMap = {
  stats: statsReducer
};
