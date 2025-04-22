import { DataList } from "@chakra-ui/react";
import { UserDto } from "@/types/dtos/users/UserDto";
import GenericInfoDialog from "../ui/GenericDialog";


const UserInfoDialog = ({
    isOpen,
    onClose,
    user,
}: {
    isOpen: boolean;
    onClose: () => void;
    user: UserDto | null;
}) => {
    if (!user) return null;

    return (
        <GenericInfoDialog isOpen={isOpen} onClose={onClose} title={user.username}>
            <DataList.Root orientation="horizontal">
                <DataList.Item>
                    <DataList.ItemLabel>Email</DataList.ItemLabel>
                    <DataList.ItemValue>{user.email}</DataList.ItemValue>
                </DataList.Item>
                <DataList.Item>
                    <DataList.ItemLabel>Role</DataList.ItemLabel>
                    <DataList.ItemValue>{user.role}</DataList.ItemValue>
                </DataList.Item>
                <DataList.Item>
                    <DataList.ItemLabel>Created At</DataList.ItemLabel>
                    <DataList.ItemValue>{new Date(user.created_at).toLocaleString()}</DataList.ItemValue>
                </DataList.Item>
                <DataList.Item>
                    <DataList.ItemLabel>Last Updated</DataList.ItemLabel>
                    <DataList.ItemValue>{new Date(user.updated_at).toLocaleString()}</DataList.ItemValue>
                </DataList.Item>
            </DataList.Root>
        </GenericInfoDialog>
    );
};

export default UserInfoDialog;
