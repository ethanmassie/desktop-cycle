import { createAction, props } from '@ngrx/store';

export const setSpeed = createAction('[StatsState] Set Speed', props<{speed: number}>());
export const setCadence = createAction('[StatsState] Set Cadence', props<{cadence: number}>());
export const increaseDistanceTraveled = createAction('[StatsState] Increase Distance Traveled',
  props<{distance: number}>());
