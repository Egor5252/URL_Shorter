
import { Routes, Route } from 'react-router-dom';
import Login from './Login';
import Register from './Register';
import Root from './Root';
import Header from './Header';

export default function App() {
  return (
    <>
      <Header />
      <Routes>
        <Route path="/" element={<Root />} />
        <Route path="/login" element={<Login />} />
        <Route path="/register" element={<Register />} />
      </Routes>
    </>
  );
}