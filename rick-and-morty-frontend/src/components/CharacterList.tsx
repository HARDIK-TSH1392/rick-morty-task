// src/components/CharacterList.tsx

import React, { useState } from 'react';
import { fetchCharacter } from '../services/api';
import CharacterDetails from './CharacterDetails';
import { Character } from '../types/types'; // Import the Character type

const CharacterList: React.FC = () => {
  const [character, setCharacter] = useState<Character | null>(null);
  const [name, setName] = useState('');

  const handleSearch = async () => {
    try {
      const characterData = await fetchCharacter(name);
      setCharacter(characterData);
    } catch (error) {
      console.error("Character not found", error);
    }
  };

  return (
    <div>
      <input
        type="text"
        value={name}
        onChange={(e) => setName(e.target.value)}
        placeholder="Search for a character"
      />
      <button onClick={handleSearch}>Search</button>

      {character && <CharacterDetails character={character} />}
    </div>
  );
};

export default CharacterList;
