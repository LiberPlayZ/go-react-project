
import { Badge, Box, Flex, Spinner, Text } from "@chakra-ui/react";
import { FaCheckCircle } from "react-icons/fa";
import { MdDelete } from "react-icons/md";
import { useColorModeValue } from "@/components/ui/color-mode"
import { TodoModel } from "@/models/TodoModel";
import { useState } from "react";
import { deleteTodo, updateTodoCompleted } from "@/services/todo_service";
const TodoItem = ({ todo, onUpdateCompleted, onDelete, onClick }:
    {
        todo: TodoModel;
        onUpdateCompleted: (todo: TodoModel) => void;
        onDelete: (id: string) => void;
        onClick?: () => void
    }) => {
    const [isUpdatingCompleted, setIsUpdatingCompleted] = useState(false);
    const [isDeleting, setIsDeleting] = useState(false);



    const updateTodo = async () => {
        if (todo.completed) {
            alert("todo already completed.")
            return;
        }

        setIsUpdatingCompleted(true)
        const res = await updateTodoCompleted(todo.id);
        if (res.error) {
            alert(JSON.stringify(res.error))
        }
        else {
            todo.completed = true;
            onUpdateCompleted(todo)
        }

        setIsUpdatingCompleted(false);
    };

    const deleteTodoOnClick = async () => {
        const confirmed = confirm("Are you sure you want to delete this todo?");
        if (!confirmed) return;
        setIsDeleting(true)
        const res = await deleteTodo(todo.id);
        if (res.error) {
            alert(JSON.stringify(res.error))
        }
        else {
            onDelete(todo.id)
        }

        setIsDeleting(false);
    };

    return (
        <Flex gap={2} alignItems={"center"} cursor={"pointer"} >
            <Flex style={{ border: '1px solid ', cursor: 'pointer' }} onClick={onClick}
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
            <Flex gap={2} alignItems={"center"} >
                <Box color={"green.500"} cursor={"pointer"} onClick={updateTodo}>
                    {!isUpdatingCompleted && <FaCheckCircle size={20} />}
                    {isUpdatingCompleted && <Spinner size={"sm"} />}

                </Box>
                <Box color={"red.500"} cursor={"pointer"} onClick={deleteTodoOnClick}>
                    {!isDeleting && <MdDelete size={25} />}
                    {isDeleting && <Spinner size={"sm"} />}

                </Box>
            </Flex>
        </Flex>
    );
};
export default TodoItem;