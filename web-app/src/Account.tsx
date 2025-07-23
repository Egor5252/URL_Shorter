import { useState } from 'react'
import './Home.css'

type Message = {
  ID: number;
  CreatedAt: string;
  UpdatedAt: string;
  DeletedAt: string | null;
  UserID: number;
  OriginalURL: string;
  ShortCode: string;
};

type SuccessResponse = {
  data: {
    message: Message[];
  };
  error: null;
  success: true;
};

type ErrorResponse = {
  data: null;
  error: { message: string };
  success: false;
};

type ApiResponse = SuccessResponse | ErrorResponse;

function Account() {
  const [successMessage, setSuccessMessage] = useState<string>('')
  const [errorMessage, setErrorMessage] = useState<string>('')

  const doShortUrl = async () => {
    setSuccessMessage('')
    setErrorMessage('')

    try {
      const response = await fetch('http://localhost:8080/account', {
        method: 'GET',
        headers: {
          'Content-Type': 'application/json',
        },
        credentials: 'include',
      })

      const result: ApiResponse = await response.json()

      if (result.success) {
         const messages = result.data.message
        .map(m => `🔗 ${m.ShortCode} → ${m.OriginalURL}`)
        .join('\n');
  setSuccessMessage(`Успех:\n${messages}`);
      } else {
        setErrorMessage(`Ошибка: ${result.error.message}`)
      }
    } catch (error) {
      setErrorMessage('Ошибка запроса к серверу')
    }
  }

  return (
    <div className="container">
      <h1>Аккаунт</h1>
      <button className="submit-button" onClick={doShortUrl}>
        Получить информацию
      </button>
      {successMessage && <pre className="result-text">{successMessage}</pre>}
      {errorMessage && <p className="result-text" style={{ color: 'red' }}>{errorMessage}</p>}
    </div>
  )
}

export default Account
