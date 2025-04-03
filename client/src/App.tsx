import { Container, Stack } from '@chakra-ui/react'
import Navbar from './components/NavBar'
import TodoPage from './components/todo-components/TodoPage'

function App() {

  return (

    <Stack h='100vh'>
      <Navbar />
      <Container maxW={"2xl"}>
        <TodoPage></TodoPage>
      </Container>
    </Stack>
  )

}

export default App
