import { useEffect, useState } from "react";
import { Container, Flex, Spinner, Stack, Text, useDisclosure } from "@chakra-ui/react";
import TodoForm from "./TodoForm";
import TodoList from "./TodoList";
import { getTodos } from "@/services/todo_service";
import { TodoDto } from "@/types/dtos/todos/TodoDto";
import TodoInfoDialog from "./TodoInfoDialog";
import Navbar from "../ui/NavBar";
import { useColorModeValue } from "../ui/color-mode";
import { useAppSelector } from "@/store/hooks";
import { UserDto } from "@/types/dtos/users/UserDto";
import UserInfoDialog from "../auth-components/UserInfoDialog";
import { useNavigate } from "react-router-dom";

const TodoPage = () => {
    const [todos, setTodos] = useState<TodoDto[]>([]);
    const [isLoading, setIsLoading] = useState(true);
    const {
        open: isTodoDialogOpen,
        onOpen: onTodoDialogOpen,
        onClose: onTodoDialogClose,
    } = useDisclosure();
    const {
        open: isUserDialogOpen,
        onOpen: onUserDialogOpen,
        onClose: onUserDialogClose,
    } = useDisclosure();
    const [selectedTodo, setSelectedTodo] = useState<TodoDto | null>(null);
    const connectedUser: UserDto | null = useAppSelector((state) => state.userState.user);
    const navigate = useNavigate();

    useEffect(() => {
        const fetchTodos = async () => {
            if (connectedUser) {
                const data = await getTodos(connectedUser.id) || [];
                setTodos(data);

            }
        };
        fetchTodos();
        setIsLoading(false);
    }, []);

    const addTodo = (newTodo: TodoDto) => {
        setTodos((prevTodos) => [...prevTodos, newTodo]);
    };

    const handleUpdateTodo = (updatedTodo: TodoDto) => {

        todos.map((todo) => (todo.id === updatedTodo.id ? updatedTodo : todo))

    };

    const handleDeleteTodo = (id: string) => {
        setTodos((prev) => prev.filter((todo) => todo.id !== id));
    };

    const handleTodoClick = (todo: TodoDto) => {
        setSelectedTodo(todo);
        onTodoDialogOpen();
    };

    const handleUserInfoClick = () => {
        if (!connectedUser) {
            navigate('/login');
            return;
        }
        onUserDialogOpen();
    };

    return (
        <>
            <Stack h='100vh' bg={useColorModeValue("gray.50", "gray.900")}>
                <Navbar onUserClick={handleUserInfoClick} />
                <Container maxW={"2xl"}>

                    <TodoForm onAddTodo={addTodo} userId={connectedUser?.id || null} />
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
                        isOpen={isTodoDialogOpen}
                        onClose={onTodoDialogClose}
                        todo={selectedTodo} />

                    <UserInfoDialog
                        isOpen={isUserDialogOpen}
                        onClose={onUserDialogClose}
                        user={connectedUser}
                    ></UserInfoDialog>
                </Container>
            </Stack>
        </>
    );
};

export default TodoPage;
