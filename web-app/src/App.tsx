import { Routes, Route } from 'react-router-dom';
import Login from './Login';
import Register from './Register';
import Root from './Root';

export default function App() {
  return (
    <Routes>
      <Route path="/" element={<Root />} />
      <Route path="/login" element={<Login />} />
      <Route path="/register" element={<Register />} />
    </Routes>
  );
}