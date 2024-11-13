// src/App.tsx

import React from 'react';
import CharacterList from './components/CharacterList';
import './App.css';

const App: React.FC = () => (
  <div className="App">
    <h1>Rick and Morty Character Search</h1>
    <CharacterList />
  </div>
);

export default App;
