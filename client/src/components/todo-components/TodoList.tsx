import { Stack, Text } from "@chakra-ui/react";
import TodoItem from "./TodoItem";
import { TodoModel } from "@/models/TodoModel";

const TodoList = ({ todos }: { todos: TodoModel[] }) => {

    const handleUpdateTodo = (updatedTodo: TodoModel) => {

        todos.map((todo) => (todo.id === updatedTodo.id ? updatedTodo : todo))

    };
    return (
        <>
            {todos.length === 0 ? (
                <Stack alignItems={"center"} gap="3">
                    <Text fontSize={"xl"} textAlign={"center"} color={"gray.500"}>
                        All tasks completed! 🤞
                    </Text>
                </Stack>
            ) : (
                <Stack gap={3}>
                    {todos.map((todo) => (
                        <TodoItem key={todo.id} todo={todo} onUpdateCompleted={handleUpdateTodo} />
                    ))}
                </Stack>
            )}
        </>
    );
};

export default TodoList;
