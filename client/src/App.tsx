import LoginPage from './components/auth-components/LoginPage';
import TodoPage from './components/todo-components/TodoPage';
import { Route, Routes, Navigate } from 'react-router-dom';


function App() {

  return (


    <Routes>
      <Route path="/" element={<Navigate to="/login" replace />} />

      <Route path="/todos" element={<TodoPage />} />
      <Route path="/login" element={<LoginPage />} />
    </Routes>

  )

}

export default App
