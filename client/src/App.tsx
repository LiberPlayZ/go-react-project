import { Container, Stack } from '@chakra-ui/react'
import './App.css'
import Navbar from './components/NavBar'
import TodoForm from './components/todo-components/TodoForm'
import TodoList from './components/todo-components/TodoList'

function App() {

  return (

    <Stack h='100vh'>
      <Navbar />
      <Container>
        <TodoForm />
        <TodoList />
      </Container>
    </Stack>
  )

}

export default App
