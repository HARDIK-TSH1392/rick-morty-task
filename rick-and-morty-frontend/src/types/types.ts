// src/types/types.ts

export interface Character {
    id: number;
    name: string;
    gender: string;
    species: string;
    status: string;
    type: string;
    episodes: string[]; // Make sure this is an array of strings (URLs)
  }
  