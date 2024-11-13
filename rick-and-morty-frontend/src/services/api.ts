// src/services/api.ts
import axios from 'axios';
import { Character } from '../types/types';

const API_BASE_URL = 'http://localhost:8080';

export const fetchCharacter = async (name: string): Promise<Character> => {
  try {
    const response = await axios.get(`${API_BASE_URL}/search?name=${name}`);
    
    if (response.data) {
      const characterData = response.data; // Response directly contains the character data

      return {
        id: characterData.id,
        name: characterData.name,
        gender: characterData.gender,
        species: characterData.species,
        status: characterData.status,
        type: characterData.type || 'N/A', // If type is not available, fallback to 'N/A'
        episodes: characterData.episodes, // Assuming episodes is an array of episode URLs
      };
    } else {
      throw new Error("Character not found");
    }
  } catch (error) {
    console.error("Error fetching character data:", error);
    throw new Error("Character not found");
  }
};
