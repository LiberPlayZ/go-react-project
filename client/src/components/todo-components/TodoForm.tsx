import { CreateTodoDto } from "@/dtos/todos/CreateTodoDto";
import { createTodo } from "@/services/todo_service";
import { Button, Flex, Input, Spinner } from "@chakra-ui/react";
import { useState } from "react";
import { IoMdAdd } from "react-icons/io";
import { TodoDto } from "@/dtos/todos/TodoDto";
import { useAppDispatch } from "@/store/hooks";
import { addTodoToUser } from "@/store/userSlice";

const TodoForm = ({ onAddTodo, userId }:
    {
        onAddTodo: (todo: TodoDto) => void,
        userId: string | null
    }) => {
    const [newTodoTitle, setNewTodoTitle] = useState("");
    const [isPending, setIsPending] = useState(false);
    const dispatch = useAppDispatch();
    const createTodoForm = async (e: React.FormEvent) => {
        e.preventDefault();
        if (!newTodoTitle.trim()) return;
        if (!userId || !userId.trim()) {
            alert("pls log in to see todos.")
            return;
        }

        setIsPending(true);
        const todoDto: CreateTodoDto = {
            title: newTodoTitle,
            description: "test",
            userid: userId
        };
        const createdTodo = await createTodo(todoDto);
        if (createdTodo.error) {
            alert(JSON.stringify(createdTodo.error))
        }
        else {
            onAddTodo(createdTodo);
            dispatch(addTodoToUser(createdTodo.id))
        }
        setIsPending(false);
        setNewTodoTitle("");
    };

    return (
        <form onSubmit={createTodoForm}>
            <Flex gap={2}>
                <Input
                    type="text"
                    value={newTodoTitle}
                    onChange={(e) => setNewTodoTitle(e.target.value)}
                />
                <Button mx={2} type="submit" _active={{ transform: "scale(.97)" }}>
                    {isPending ? <Spinner size={"xs"} /> : <IoMdAdd size={30} />}
                </Button>
            </Flex>
        </form>
    );
};

export default TodoForm;
