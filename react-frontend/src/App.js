import React, { useState } from 'react';
import axios from 'axios';
import './App.css';

function App() {
  const [albanAccountID, setAlbanAccountID] = useState('');
  const [irisAccountID, setIrisAccountID] = useState('');
  const [amount, setAmount] = useState('');
  const [message, setMessage] = useState('');

  const handleSubmit = async () => {
    if (!albanAccountID || !irisAccountID || !amount) {
      setMessage('Please fill in all fields.');
      return;
    }

    try {
      const response = await axios.post('http://localhost:8000/transfer', {
        albanAccountID,
        irisAccountID,
        amount: Number(amount),
      });
      setMessage('Transaction successful! ' + response.data);
    } catch (error) {
      setMessage('An error occurred: ' + error.toString());
    }
  };

  return (
    <div className="App">
      <header className="App-header">
        <h1>USDC Transfer</h1>
        <input
          type="text"
          placeholder="Alban's Account ID"
          value={albanAccountID}
          onChange={(e) => setAlbanAccountID(e.target.value)}
        />
        <input
          type="text"
          placeholder="Iris's Account ID"
          value={irisAccountID}
          onChange={(e) => setIrisAccountID(e.target.value)}
        />
        <input
          type="number"
          placeholder="Amount"
          value={amount}
          onChange={(e) => setAmount(e.target.value)}
        />
        <button onClick={handleSubmit}>Transfer USDC</button>
        {message && <p>{message}</p>}
      </header>
    </div>
  );
}

export default App;
