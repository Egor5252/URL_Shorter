import { useState } from 'react';
import axios from 'axios';

export default function Root() {
  const [url, setUrl] = useState('');
  const [error, setError] = useState<string | null>(null);

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    setError(null);

    try {
      const res = await axios.post('/createshorturl', {
        url,
      });
      const message = res.data.message;
      alert(message);
      // перенаправление можно сделать через react-router-dom
    } catch (err) {
      setError('Ошибка');
    }
  };

  return (
    <div className="login-container">
      <form onSubmit={handleSubmit} className="login-form">
        <h1 className="form-group">Короткоссылка</h1>

        <div className="form-group">
          <input
            type="text"
            value={url}
            onChange={e => setUrl(e.target.value)}
            placeholder="Ссылка"
            required
          />
        </div>

        <button type="submit" className="bg-blue-500 text-white p-2 rounded">Короткоссыльнуть</button>
        {error && <div className="text-red-500">{error}</div>}
      </form>
    </div>
  );
}
