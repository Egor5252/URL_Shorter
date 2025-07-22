import { useEffect, useState } from 'react'
import { Link } from 'react-router-dom';
import './Header.css';

type SuccessResponse = {
  data: { message: string };
  error: null;
  success: true;
};

type ErrorResponse = {
  data: null;
  error: { message: string };
  success: false;
};

type ApiResponse = SuccessResponse | ErrorResponse;

function Header() {
  const [username, setUsername] = useState<string | null>(null);

  const checkAuth = async () => {
    try {
      const response = await fetch('http://localhost:8080/whoami', {
        method: 'GET',
        headers: {
          'Content-Type': 'application/json',
        },
        credentials: 'include', // если куки используются для авторизации
      });

      const result: ApiResponse = await response.json();

      if (result.success) {
        setUsername(result.data.message); // имя пользователя
      } else {
        setUsername(null);
      }
    } catch (error) {
      setUsername(null);
    }
  };

  useEffect(() => {
    checkAuth();
  }, []);

  return (
    <header className="header">
      <Link to="/"><div className="logo">MyShortener</div></Link>
      <nav className="nav">
        {username ? (
          <>
            <span>{username}</span>
            <Link to="/logout">Выйти</Link>
          </>
        ) : (
          <>
            <Link to="/login">Вход</Link>
            <Link to="/register">Регистрация</Link>
          </>
        )}
      </nav>
    </header>
  );
}

export default Header;
