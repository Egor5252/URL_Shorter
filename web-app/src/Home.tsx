import { useState } from 'react'
import './Home.css'

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

function Home() {
  const [originalUrl, setOriginalUrl] = useState('')
  const [shortCode, setShortCode] = useState('')
  const [successMessage, setSuccessMessage] = useState<string>('')
  const [errorMessage, setErrorMessage] = useState<string>('')

  const doShortUrl = async () => {
    setSuccessMessage('')
    setErrorMessage('')

    try {
      const response = await fetch('http://localhost:8080/createshorturl', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        credentials: 'include',
        body: JSON.stringify({
          url: originalUrl,
          short_code: shortCode
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
      <h1>Сократить ссылку</h1>
      <input
        type="text"
        value={originalUrl}
        onChange={(e) => setOriginalUrl(e.target.value)}
        placeholder="Ссылка"
        className="input-field"
      />
      <input
        type="text"
        value={shortCode}
        onChange={(e) => setShortCode(e.target.value)}
        placeholder="Желаемое сокращение"
        className="input-field"
      />
      <button className="submit-button" onClick={doShortUrl}>
        Отправить
      </button>
      {successMessage && <p className="result-text">{successMessage}</p>}
      {errorMessage && <p className="result-text" style={{ color: 'red' }}>{errorMessage}</p>}
    </div>
  )
}

export default Home
