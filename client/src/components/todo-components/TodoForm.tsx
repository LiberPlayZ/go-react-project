import { TodoDto } from "@/dtos/TodoDto";
import { createTodo } from "@/services/todo_service";
import { Button, Flex, Input, Spinner } from "@chakra-ui/react";
import { useState } from "react";
import { IoMdAdd } from "react-icons/io";
import { TodoModel } from "@/models/TodoModel";

const TodoForm = ({ onAddTodo }: { onAddTodo: (todo: TodoModel) => void }) => {
    const [newTodoTitle, setNewTodoTitle] = useState("");
    const [isPending, setIsPending] = useState(false);

    const createTodoForm = async (e: React.FormEvent) => {
        e.preventDefault();
        if (!newTodoTitle.trim()) return;

        setIsPending(true);
        const todoDto: TodoDto = {
            title: newTodoTitle,
            description: "test",
        };
        const createdTodo = await createTodo(todoDto);
        setIsPending(false);

        onAddTodo(createdTodo); 
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
