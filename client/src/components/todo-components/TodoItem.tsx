
import { Badge, Box, Flex, Spinner, Text } from "@chakra-ui/react";
import { FaCheckCircle } from "react-icons/fa";
import { MdDelete } from "react-icons/md";
import { useColorModeValue } from "@/components/ui/color-mode"
import { TodoModel } from "@/models/TodoModel";
import { useState } from "react";
import { updateTodoCompleted } from "@/services/todo_service";
const TodoItem = ({ todo, onUpdateCompleted }:
    { todo: TodoModel; onUpdateCompleted: (todo: TodoModel) => void }) => {
    const [isUpdatingCompleted, setIsUpdatingCompleted] = useState(false);
    const [isDeleting, setIsDeleting] = useState(false);



    const updateTodo = async () => {
        if (todo.completed) {
            alert("todo already completed.")
            return;
        }

        setIsUpdatingCompleted(true)
        await updateTodoCompleted(todo.id);
        todo.completed = true;
        onUpdateCompleted(todo)
        setIsUpdatingCompleted(false);
    };

    const deleteTodo = async () => {

        setIsUpdatingCompleted(true)
        await updateTodoCompleted(todo.id);
        todo.completed = true;
        onUpdateCompleted(todo)
        setIsUpdatingCompleted(false);
    };

    return (
        <Flex gap={2} alignItems={"center"}>
            <Flex style={{ border: '1px solid ' }}
                flex={1}
                alignItems={"center"}
                p={2}
                borderRadius={"lg"}
                justifyContent={"space-between"}
            >
                <Text
                    color={todo.completed ? useColorModeValue("green.400", "green.700")
                        : useColorModeValue("yellow.400", "yellow.700")}
                    textDecoration={todo.completed ? "line-through" : "none"}
                >
                    {todo.title}
                </Text>
                {todo.completed && (
                    <Badge ml='1' colorScheme='green'>
                        Done
                    </Badge>
                )}
                {!todo.completed && (
                    <Badge ml='1' colorScheme='yellow'>
                        In Progress
                    </Badge>
                )}
            </Flex>
            <Flex gap={2} alignItems={"center"} onClick={updateTodo}>
                <Box color={"green.500"} cursor={"pointer"}>
                    {!isUpdatingCompleted && <FaCheckCircle size={20} />}
                    {isUpdatingCompleted && <Spinner size={"sm"} />}

                </Box>
                <Box color={"red.500"} cursor={"pointer"} onClick={deleteTodo}>
                    {!isDeleting && <MdDelete size={25} />}
                    {isDeleting && <Spinner size={"sm"} />}

                </Box>
            </Flex>
        </Flex>
    );
};
export default TodoItem;