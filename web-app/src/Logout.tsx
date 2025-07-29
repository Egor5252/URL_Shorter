import { useState } from 'react';
import { useNavigate } from 'react-router-dom';

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

function Logout() {
  const [successMessage, setSuccessMessage] = useState<string>('');
  const [errorMessage, setErrorMessage] = useState<string>('');
  const navigate = useNavigate();

  const doLogout = async () => {
    setSuccessMessage('');
    setErrorMessage('');

    try {
      const response = await fetch('http://localhost:8080/logout', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        credentials: 'include',
      });

      const result: ApiResponse = await response.json();

      if (result.success) {
        setSuccessMessage(`Успех: ${result.data.message}`);
        // перенаправляем пользователя после выхода
        setTimeout(() => navigate('/'), 1500);
      } else {
        setErrorMessage(`Ошибка: ${result.error.message}`);
      }
    } catch (error) {
      setErrorMessage('Ошибка запроса к серверу');
    }
  };

  return (
    <div className="container">
      <h1>Выход</h1>
      <button onClick={doLogout}>Выйти из аккаунта</button>
      {successMessage && <p style={{ color: 'green' }}>{successMessage}</p>}
      {errorMessage && <p style={{ color: 'red' }}>{errorMessage}</p>}
    </div>
  );
}

export default Logout;
