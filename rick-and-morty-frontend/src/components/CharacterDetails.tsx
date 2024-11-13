// src/components/CharacterDetails.tsx

import React, { useEffect, useState } from 'react';
import { Character } from '../types/types';
// import './CharacterDetails.css';

interface CharacterDetailsProps {
  character: Character;
}

const CharacterDetails: React.FC<CharacterDetailsProps> = ({ character }) => {
  const [episodes, setEpisodes] = useState<number[]>([]);

  useEffect(() => {
    // Extract episode numbers from URLs in the character's episodes array
    const episodeNumbers = character.episodes.map((episodeUrl) => {
      const parts = episodeUrl.split('/');
      return parseInt(parts[parts.length - 1]);
    });
    setEpisodes(episodeNumbers);
  }, [character]);

  return (
    <div>
      <h2>Character Details</h2>
      <p><strong>Name:</strong> {character.name}</p>
      <p><strong>Gender:</strong> {character.gender}</p>
      <p><strong>Species:</strong> {character.species}</p>
      <p><strong>Status:</strong> {character.status}</p>
      <p><strong>Type:</strong> {character.type || 'N/A'}</p>
      <h3>Episodes:</h3>
      <ul>
        {episodes.map((episodeNumber) => (
          <li key={episodeNumber}>Episode {episodeNumber}</li>
        ))}
      </ul>
    </div>
  );
};

export default CharacterDetails;
