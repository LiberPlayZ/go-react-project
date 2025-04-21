import {
    Badge,
    DataList,
    Dialog,
    DialogBody,
    DialogContent,
    DialogHeader,
    DialogTitle,
    Portal,
    DialogCloseTrigger,
    DialogPositioner,
    DialogBackdrop,
    CloseButton,
} from "@chakra-ui/react";
import { TodoDto } from "@/dtos/todos/TodoDto";

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
        <Dialog.Root open={isOpen}>
            <Portal>
                <DialogBackdrop />
                <DialogPositioner>
                    <DialogContent>
                        <DialogHeader>
                            <DialogTitle>{todo.title}</DialogTitle>
                        </DialogHeader>
                        <DialogBody pb="8">
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
                                    <DataList.ItemValue>
                                        {new Date(todo.created_at).toLocaleString()
                                        }
                                    </DataList.ItemValue>
                                </DataList.Item>

                                <DataList.Item>
                                    <DataList.ItemLabel>Last Updated</DataList.ItemLabel>
                                    <DataList.ItemValue>
                                        {new Date(todo.updated_at).toLocaleString()}
                                    </DataList.ItemValue>
                                </DataList.Item>
                            </DataList.Root>

                            <DialogCloseTrigger asChild>
                                <CloseButton size="sm" onClick={onClose} />
                            </DialogCloseTrigger>
                        </DialogBody>
                    </DialogContent>
                </DialogPositioner>
            </Portal>
        </Dialog.Root>
    );
};

export default TodoInfoDialog;
