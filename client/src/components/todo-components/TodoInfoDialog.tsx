import { Badge, DataList } from "@chakra-ui/react";
import { TodoDto } from "@/dtos/todos/TodoDto";
import GenericInfoDialog from "../ui/GenericDialog";


const TodoInfoDialog = ({
    isOpen,
    onClose,
    todo,
}: {
    isOpen: boolean;
    onClose: () => void;
    todo: TodoDto | null;
}) => {
    if (!todo) return null;

    return (
        <GenericInfoDialog isOpen={isOpen} onClose={onClose} title={todo.title}>
            <DataList.Root orientation="horizontal">
                <DataList.Item>
                    <DataList.ItemLabel>Status</DataList.ItemLabel>
                    <DataList.ItemValue>
                        <Badge colorPalette={todo.completed ? "green" : "yellow"}>
                            {todo.completed ? "Completed" : "Pending"}
                        </Badge>
                    </DataList.ItemValue>
                </DataList.Item>

                {todo.description && (
                    <DataList.Item>
                        <DataList.ItemLabel>Description</DataList.ItemLabel>
                        <DataList.ItemValue>{todo.description}</DataList.ItemValue>
                    </DataList.Item>
                )}

                <DataList.Item>
                    <DataList.ItemLabel>Created At</DataList.ItemLabel>
                    <DataList.ItemValue>{new Date(todo.created_at).toLocaleString()}</DataList.ItemValue>
                </DataList.Item>

                <DataList.Item>
                    <DataList.ItemLabel>Last Updated</DataList.ItemLabel>
                    <DataList.ItemValue>{new Date(todo.updated_at).toLocaleString()}</DataList.ItemValue>
                </DataList.Item>
            </DataList.Root>
        </GenericInfoDialog>
    );
};

export default TodoInfoDialog;
