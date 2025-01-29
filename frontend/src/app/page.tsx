"use client"
import { useState } from 'react';

export default function Home() {
  const [email, setEmail] = useState('');
  const [favoriteTeam, setFavoriteTeam] = useState('');
  const [submitted, setSubmitted] = useState(false);

  const handleSubmit = (e: any) => {
    e.preventDefault();
    // Handle form submission (e.g., send data to an API)
    console.log('Email:', email);
    console.log('Favorite Team:', favoriteTeam);
    setSubmitted(true);
  };

  return (
    <div className="min-h-screen bg-gray-100 flex flex-col items-center justify-center p-4">
      <h1 className="text-3xl font-bold text-gray-800 mb-8">Fastball⚾; Didn't read!</h1>
      {submitted ? (
        <p className="text-green-600 text-lg">Thank you for submitting your favorite team!</p>
      ) : (
        <form onSubmit={handleSubmit} className="w-full max-w-md bg-white p-8 rounded-lg shadow-md">
          <div className="mb-6">
            <label htmlFor="email" className="block text-sm font-medium text-gray-700 mb-2">
              Email:
            </label>
            <input
              type="email"
              id="email"
              value={email}
              onChange={(e) => setEmail(e.target.value)}
              required
              className="w-full px-4 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500"
            />
          </div>
          <div className="mb-6">
            <label htmlFor="favoriteTeam" className="block text-sm font-medium text-gray-700 mb-2">
              Favorite MLB Team:
            </label>
            <select
              id="favoriteTeam"
              value={favoriteTeam}
              onChange={(e) => setFavoriteTeam(e.target.value)}
              required
              className="w-full px-4 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500"
            >
              <option value="">Select a team</option>
              <option value="Los Angeles Dodgers">Los Angeles Dodgers</option>
              <option value="Chicago Cubs">Chicago Cubs</option>
              <option value="San Francisco Giants">San Francisco Giants</option>
              <option value="Baltimore Orioles">Baltimore Orioles</option>
              <option value="Boston Red Sox">Boston Red Sox</option>
              <option value="New York Yankees">New York Yankees</option>
              <option value="Tampa Bay Rays">Tampa Bay Rays</option>
              <option value="Toronto Blue Jays">Toronto Blue Jays</option>
            </select>
          </div>
          <button
            type="submit"
            className="w-full bg-blue-500 text-white py-2 px-4 rounded-lg hover:bg-blue-600 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:ring-offset-2"
          >
            Submit
          </button>
        </form>
      )}
    </div>
  );
}