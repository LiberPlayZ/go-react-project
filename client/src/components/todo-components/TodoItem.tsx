
import { Badge, Box, Flex, Text } from "@chakra-ui/react";
import { FaCheckCircle } from "react-icons/fa";
import { MdDelete } from "react-icons/md";
import { useColorModeValue } from "@/components/ui/color-mode"
const TodoItem = ({ todo }: { todo: any }) => {
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
                    {todo.body}
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
            <Flex gap={2} alignItems={"center"}>
                <Box color={"green.500"} cursor={"pointer"}>
                    <FaCheckCircle size={20} />
                </Box>
                <Box color={"red.500"} cursor={"pointer"}>
                    <MdDelete size={25} />
                </Box>
            </Flex>
        </Flex>
    );
};
export default TodoItem;