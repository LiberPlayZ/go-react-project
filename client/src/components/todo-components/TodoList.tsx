import { Stack, Text, Box } from "@chakra-ui/react";
import TodoItem from "./TodoItem";
import { TodoDto } from "@/dtos/todos/TodoDto";

const TodoList = ({
    todos,
    onUpdateCompleted,
    onDelete,
    onTodoClick
}: {
    todos: TodoDto[];
    onUpdateCompleted: (todo: TodoDto) => void;
    onDelete: (id: string) => void;
    onTodoClick: (todo: TodoDto) => void; // 👈 include this

}) => {
    return (
        <>
            {todos.length === 0 ? (
                <Stack alignItems={"center"} gap="3">
                    <Text fontSize={"xl"} textAlign={"center"} color={"gray.500"}>
                        All tasks completed! 🤞
                    </Text>
                </Stack>
            ) : (
                <Box
                    maxH="60vh" // or any height that fits your layout
                    overflowY="auto"
                    px={1}
                >
                    <Stack gap={3}>
                        {todos.map((todo) => (
                            <TodoItem
                                key={todo.id}
                                todo={todo}
                                onUpdateCompleted={onUpdateCompleted}
                                onDelete={onDelete}
                                onClick={() => onTodoClick(todo)}
                            />
                        ))}
                    </Stack>
                </Box>
            )}
        </>
    );
};

export default TodoList;
