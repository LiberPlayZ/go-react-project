import {
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
import { ReactNode } from "react";

type GenericInfoDialogProps = {
    isOpen: boolean;
    onClose: () => void;
    title: string;
    children: ReactNode;
};

const GenericInfoDialog = ({ isOpen, onClose, title, children }: GenericInfoDialogProps) => {
    return (
        <Dialog.Root open={isOpen}>
            <Portal>
                <DialogBackdrop />
                <DialogPositioner>
                    <DialogContent>
                        <DialogHeader>
                            <DialogTitle>{title}</DialogTitle>
                        </DialogHeader>
                        <DialogBody pb="8">
                            {children}
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

export default GenericInfoDialog;
