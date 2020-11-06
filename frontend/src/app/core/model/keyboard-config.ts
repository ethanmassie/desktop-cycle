
export interface KeyboardConfig {
  name: string;
  id: string;
  keys: KeyConfig[];
}

export interface KeyConfig {
  key: string;
  minSpeed: number;
  maxSpeed: number;
  action: 'HOLD' | 'TOGGLE';
}
