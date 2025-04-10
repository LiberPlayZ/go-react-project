import LoginPage from './components/auth-components/LoginPage'
import TodoPage from './components/todo-components/TodoPage'
import { Route, Routes } from 'react-router-dom'

function App() {

  return (

    <Routes>
      <Route path="/todos" element={<TodoPage />} />
      <Route path="/login" element={<LoginPage />} />
    </Routes>

  )

}

export default App
