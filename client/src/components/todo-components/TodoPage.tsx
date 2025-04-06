import { useEffect, useState } from "react";
import { Flex, Spinner, Text } from "@chakra-ui/react";
import TodoForm from "./TodoForm";
import TodoList from "./TodoList";
import { getTodos } from "@/services/todo_service";
import { TodoModel } from "@/models/TodoModel";

const TodoPage = () => {
    const [todos, setTodos] = useState<TodoModel[]>([]);
    const [isLoading, setIsLoading] = useState(true);

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

    return (
        <>
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
                    onDelete={handleDeleteTodo} />
            )}
        </>
    );
};

export default TodoPage;
