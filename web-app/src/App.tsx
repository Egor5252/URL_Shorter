// import React from 'react';
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';
import Header from './components/Header';
import Home from './Home';
import Register from './Register';
import Login from './Login';
import Logout from './Logout';

// Компоненты страниц

function App() {
  return (
    <Router>
      <Header />
      <main style={{ padding: 200 }}>
        <Routes>
          <Route path="/" element={<Home />} />
          <Route path="/register" element={<Register />} />
          <Route path="/login" element={<Login />} />
          <Route path="/logout" element={<Logout />} />
        </Routes>
      </main>
    </Router>
  );
}

export default App;
