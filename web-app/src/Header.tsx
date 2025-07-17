import React from 'react';
import { Link } from 'react-router-dom';
import './App.css';

const Header: React.FC = () => (
  <header className="app-header">
    <nav>
      <ul className="nav-list">
        <li><Link to="/">Главная</Link></li>
        <li><Link to="/login">Вход</Link></li>
        <li><Link to="/register">Регистрация</Link></li>
      </ul>
    </nav>
  </header>
);

export default Header;
