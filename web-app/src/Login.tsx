import { useState } from 'react';
import axios from 'axios';

export default function Login() {
  const [user, setUser] = useState('');
  const [password, setPassword] = useState('');
  const [error, setError] = useState<string | null>(null);

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    setError(null);

    try {
      await axios.post('/login', {
        user,
        password,
      });
      // localStorage.setItem('token', token);
      alert('Вы вошли!');
      // перенаправление можно сделать через react-router-dom
    } catch (err) {
      setError('Неверный логин или пароль');
    }
  };

  return (
    <div className="login-container">
      <form onSubmit={handleSubmit} className="login-form">
        <h1 className="form-group">Вход</h1>

        <div className="form-group">
          <input
            type="text"
            value={user}
            onChange={e => setUser(e.target.value)}
            placeholder="Логин"
            required
          />
        </div>

        <div className="form-group">
          <input
            type="password"
            value={password}
            onChange={e => setPassword(e.target.value)}
            placeholder="Пароль"
            required
          />
        </div>

        <button type="submit" className="bg-blue-500 text-white p-2 rounded">Войти</button>
        {error && <div className="text-red-500">{error}</div>}
      </form>
    </div>
  );
}
