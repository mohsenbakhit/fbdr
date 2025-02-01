"use client"
import { useState, useEffect } from 'react';
import axios from 'axios';

export default function Home() {
  const [email, setEmail] = useState('');
  const [favoriteTeams, setFavoriteTeams] = useState<string[]>([]);
  const [favoritePlayers, setFavoritePlayers] = useState<string[]>([]);
  const [submitted, setSubmitted] = useState(false);
  const [playerSearch, setPlayerSearch] = useState('');
  const [playerSuggestions, setPlayerSuggestions] = useState<string[]>([]);
  const [allPlayers, setAllPlayers] = useState<string[]>([]);

  useEffect(() => {
    // Fetch the list of MLB players from the backend API
    axios.get('/api/mlb-players')
      .then(response => {
        setAllPlayers(response.data.players);
      })
      .catch(error => {
        console.error('Error fetching MLB players:', error);
      });
  }, []);

  const handleCheckboxChange = (team: string) => {
    setFavoriteTeams((prev) =>
      prev.includes(team) ? prev.filter((t) => t !== team) : [...prev, team]
    );
  };

  const handlePlayerSearchChange = (e: any) => {
    const searchValue = e.target.value;
    setPlayerSearch(searchValue);
    if (searchValue) {
      const suggestions = allPlayers.filter(player =>
        player.toLowerCase().includes(searchValue.toLowerCase())
      );
      setPlayerSuggestions(suggestions);
    } else {
      setPlayerSuggestions([]);
    }
  };

  const handlePlayerSelect = (player: string) => {
    if (!favoritePlayers.includes(player)) {
      setFavoritePlayers([...favoritePlayers, player]);
    }
    setPlayerSearch('');
    setPlayerSuggestions([]);
  };

  const handleSubmit = (e: any) => {
    e.preventDefault();
    axios.post('/api/submit', {
      "email": email,
      "favoriteTeams": favoriteTeams,
      "favoritePlayers": favoritePlayers
    })
    console.log('Email:', email);
    console.log('Favorite Teams:', favoriteTeams);
    console.log('Favorite Players:', favoritePlayers);
    setSubmitted(true);
  };

  const teams = {
    "AL East": ["Baltimore Orioles", "Boston Red Sox", "New York Yankees", "Tampa Bay Rays", "Toronto Blue Jays"],
    "AL Central": ["Chicago White Sox", "Cleveland Guardians", "Detroit Tigers", "Kansas City Royals", "Minnesota Twins"],
    "AL West": ["Houston Astros", "Los Angeles Angels", "Oakland Athletics", "Seattle Mariners", "Texas Rangers"],
    "NL East": ["Atlanta Braves", "Miami Marlins", "New York Mets", "Philadelphia Phillies", "Washington Nationals"],
    "NL Central": ["Chicago Cubs", "Cincinnati Reds", "Milwaukee Brewers", "Pittsburgh Pirates", "St. Louis Cardinals"],
    "NL West": ["Arizona Diamondbacks", "Colorado Rockies", "Los Angeles Dodgers", "San Diego Padres", "San Francisco Giants"]
  };

  return (
    <div className="min-h-screen bg-gray-100 flex flex-col items-center justify-center p-4">
      <h1 className="text-3xl font-bold text-gray-900 mb-8">Fastball⚾; Didn't read!</h1>
      {submitted ? (
        <p className="text-green-700 text-lg">Thank you for submitting your favorite team!</p>
      ) : (
        <form onSubmit={handleSubmit} className="w-full max-w-2xl bg-white p-8 rounded-lg shadow-md">
          <div className="mb-6 flex flex-col items-center">
            <label htmlFor="email" className="block text-sm font-medium text-gray-900 mb-2">
              Email:
            </label>
            <input
              type="email"
              id="email"
              value={email}
              onChange={(e) => setEmail(e.target.value)}
              required
              className="w-64 px-4 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500"
            />
          </div>
          <div className="mb-6">
            <label className="block text-sm font-medium text-gray-900 mb-2">
              Pick the MLB Teams you want to follow:
            </label>
            <div className="grid grid-cols-1 md:grid-cols-3 gap-4">
              {Object.entries(teams).map(([division, teams]) => (
                <div key={division} className="mb-4">
                  <h3 className="font-semibold text-gray-900">{division}</h3>
                    <div className="grid grid-cols-2 md:grid-cols-2 gap-4">
                      {teams.map((team) => (
                        <div key={team} className="flex items-center mb-2">
                          <label htmlFor={team} className="text-sm text-gray-900">{team}</label>
                          <input
                            type="checkbox"
                            id={team}
                            value={team}
                            checked={favoriteTeams.includes(team)}
                            onChange={() => handleCheckboxChange(team)}
                            className="mr-2"
                          />
                        </div>
                      ))}
                    </div>
                  </div>
              ))}
            </div>
          <div className="mb-6 flex flex-col items-center">
            <label htmlFor="playerSearch" className="block text-sm font-medium text-gray-900 mb-2">
              Search for MLB Players:
            </label>
            <input
              type="text"
              id="playerSearch"
              value={playerSearch}
              onChange={handlePlayerSearchChange}
              className="w-64 px-4 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500"
            />
            {playerSuggestions.length > 0 && (
              <ul className="w-64 bg-white border border-gray-300 rounded-lg mt-2">
                {playerSuggestions.map((player) => (
                  <li
                    key={player}
                    onClick={() => handlePlayerSelect(player)}
                    className="px-4 py-2 cursor-pointer hover:bg-gray-200"
                  >
                    {player}
                  </li>
                ))}
              </ul>
            )}
            </div>
          </div>
          <div className="mb-6">
            <label className="block text-sm font-medium text-gray-900 mb-2">
              Favorite Players:
            </label>
            <ul className="list-disc list-inside">
              {favoritePlayers.map((player) => (
                <li key={player} className="text-sm text-gray-900">{player}</li>
              ))}
            </ul>
          </div>
          <div className="flex justify-center">
            <button
              type="submit"
              className="w-64 bg-blue-500 text-white py-2 px-4 rounded-lg hover:bg-blue-600 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:ring-offset-2"
            >Submit</button>
          </div>
        </form>
      )}
    </div>
  );
}