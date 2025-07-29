import { useState } from 'react'
import './Login.css'

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

function Login() {
  const [userName, setUserName] = useState('')
  const [password, setPassword] = useState('')
  const [successMessage, setSuccessMessage] = useState<string>('')
  const [errorMessage, setErrorMessage] = useState<string>('')

  const doLogin = async () => {
    setSuccessMessage('')
    setErrorMessage('')

    try {
      const response = await fetch('http://localhost:8080/login', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        credentials: 'include',
        body: JSON.stringify({
          username: userName,
          password: password
        }),
      })

      const result: ApiResponse = await response.json()

      if (result.success) {
        setSuccessMessage(`Успех: ${result.data.message}`)
      } else {
        setErrorMessage(`Ошибка: ${result.error.message}`)
      }
    } catch (error) {
      setErrorMessage('Ошибка запроса к серверу')
    }
  }

  return (
    <div className="container">
      <h1>Вход</h1>
      <input
        type="text"
        value={userName}
        onChange={(e) => setUserName(e.target.value)}
        placeholder="Логин"
        className="input-field"
      />
      <input
        type="password"
        value={password}
        onChange={(e) => setPassword(e.target.value)}
        placeholder="Пароль"
        className="input-field"
      />
      <button className="submit-button" onClick={doLogin}>
        Зарегистрироавться
      </button>
      {successMessage && <p className="result-text">{successMessage}</p>}
      {errorMessage && <p className="result-text" style={{ color: 'red' }}>{errorMessage}</p>}
    </div>
  )
}

export default Login
