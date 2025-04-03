import { Flex, Spinner, Stack, Text } from "@chakra-ui/react";
import { useEffect, useState } from "react";
import TodoItem from "./TodoItem";
import { getTodos } from "@/services/todo_service";
import { TodoModel } from "@/models/TodoModel";

const TodoList = () => {
    const [isLoading, setIsLoading] = useState(true);
    const [todos, setTodos] = useState<TodoModel[]>([]);

    useEffect(() => {
        const fetchTodos = async () => {
            const data = await getTodos() || []; // Ensure it's never null
            setTodos(data);
            setIsLoading(false);
        };
        fetchTodos();
    }, []);

    useEffect(() => {
    }, [todos, isLoading]);

    return (
        <>
            <Text fontSize={"4xl"} textTransform={"uppercase"} fontWeight={"bold"} textAlign={"center"} my={2}>
                Today's Tasks
            </Text>

            {isLoading && (
                <Flex justifyContent={"center"} my={4}>
                    <Spinner size={"xl"} />
                </Flex>
            )}

            {!isLoading && (todos?.length ?? 0) === 0 && (
                <Stack alignItems={"center"} gap="3">
                    <Text fontSize={"xl"} textAlign={"center"} color={"gray.500"}>
                        All tasks completed! 🤞
                    </Text>
                </Stack>
            )}

            {!isLoading && todos.length > 0 && (
                <Stack gap={3} >
                    {todos.map((todo) => (
                        <TodoItem key={todo.id} todo={todo} />
                    ))}
                </Stack>
            )}
        </>
    );
};

export default TodoList;
