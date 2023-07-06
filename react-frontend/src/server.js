const express = require('express');
const cors = require('cors');

const app = express();
const PORT = process.env.PORT || 3000;

// Enable CORS
app.use(cors());

// Middleware for parsing JSON
app.use(express.json());

// Endpoint to handle transfer
app.post('/api/transfer', async (req, res) => {
  const { albanAccountID, irisAccountID, amount } = req.body;

  // Perform the transaction using provided account IDs and amount
  // ...
  // On success:
  res.json({ message: 'Transaction completed!' });
  
  // If there was an error:
  // res.status(500).json({ error: 'Transaction failed!' });
});

// Starting the server
app.listen(PORT, () => {
  console.log(`Server is running on port ${PORT}`);
});
