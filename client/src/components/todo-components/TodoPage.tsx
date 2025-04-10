import { useEffect, useState } from "react";
import { Container, Flex, Spinner, Stack, Text, useDisclosure } from "@chakra-ui/react";
import TodoForm from "./TodoForm";
import TodoList from "./TodoList";
import { getTodos } from "@/services/todo_service";
import { TodoModel } from "@/models/TodoModel";
import TodoInfoDialog from "./TodoInfoDialog";
import Navbar from "../NavBar";
import { useColorModeValue } from "../ui/color-mode";

const TodoPage = () => {
    const [todos, setTodos] = useState<TodoModel[]>([]);
    const [isLoading, setIsLoading] = useState(true);
    const { open, onOpen, onClose } = useDisclosure();
    const [selectedTodo, setSelectedTodo] = useState<TodoModel | null>(null);

    useEffect(() => {
        const fetchTodos = async () => {
            const data = await getTodos() || [];
            setTodos(data);
            setIsLoading(false);
        };
        fetchTodos();
    }, []);

    const addTodo = (newTodo: TodoModel) => {
        setTodos((prevTodos) => [...prevTodos, newTodo]);
    };

    const handleUpdateTodo = (updatedTodo: TodoModel) => {

        todos.map((todo) => (todo.id === updatedTodo.id ? updatedTodo : todo))

    };

    const handleDeleteTodo = (id: string) => {
        setTodos((prev) => prev.filter((todo) => todo.id !== id));
    };

    const handleTodoClick = (todo: TodoModel) => {
        setSelectedTodo(todo);
        onOpen();
    };

    return (
        <>
            <Stack h='100vh' bg={useColorModeValue("gray.50", "gray.800")}>
                <Navbar />
                <Container maxW={"2xl"}>

                    <TodoForm onAddTodo={addTodo} />
                    <Text fontSize={"4xl"} textTransform={"uppercase"} fontWeight={"bold"} textAlign={"center"} my={2}>
                        Today's Tasks
                    </Text>
                    {isLoading ? (
                        <Flex justifyContent={"center"} my={4}>
                            <Spinner size={"xl"} />
                        </Flex>
                    ) : (
                        <TodoList
                            todos={todos}
                            onUpdateCompleted={handleUpdateTodo}
                            onDelete={handleDeleteTodo}
                            onTodoClick={handleTodoClick} />
                    )}
                    <TodoInfoDialog
                        isOpen={open}
                        onClose={onClose}
                        todo={selectedTodo} />
                </Container>
            </Stack>
        </>
    );
};

export default TodoPage;
