import {
    Box,
    Button,
    Field,
    Heading,
    Input,
    Stack,
    Flex,
} from "@chakra-ui/react";
import Navbar from "../NavBar";
import { useColorModeValue } from "../ui/color-mode";
import { useForm } from "react-hook-form";
import { useState } from "react";

const LoginPage = () => {
    const { register, handleSubmit } = useForm();
    const [loading, setLoading] = useState(false);
    const onLogin = async (data: any) => {
        const formValues: { email: string, password: string } = data;
        console.log(formValues);
        setLoading(true);



        // if (!newTodoTitle.trim()) return;

        // setIsPending(true);
        // const todoDto: TodoDto = {
        //     title: newTodoTitle,
        //     description: "test",
        // };
        // const createdTodo = await createTodo(todoDto);
        // if (createdTodo.error) {
        //     alert(JSON.stringify(createdTodo.error))
        // }
        // else {
        //     onAddTodo(createdTodo);

        // }
        // setIsPending(false);
        // setNewTodoTitle("");
    };

    return (
        <Flex direction="column" minH="100vh" bg={useColorModeValue("gray.50", "gray.800")}>
            {/* Navbar at the top */}
            <Navbar />

            {/* Remaining area: center the login form */}
            <Flex
                flex="1"
                align="center"
                justify="center"
                bg={useColorModeValue("gray.50", "gray.800")}
            >
                <Box
                    p={8}
                    width="full"
                    maxW="md"
                    borderWidth={1}
                    borderRadius="xl"
                    boxShadow="lg"
                    bg={useColorModeValue("white", "gray.700")}
                    css={{ "--field-label-width": "96px" }}
                >     <form onSubmit={handleSubmit(onLogin)}>
                        <Stack gap={6}>
                            <Heading size="lg" textAlign="center">
                                Login
                            </Heading>

                            <Field.Root orientation="horizontal">
                                <Field.Label>Email</Field.Label>
                                <Input {...register("email")} type="email" placeholder="me@example.com" flex="1" />
                            </Field.Root>

                            <Field.Root orientation="horizontal">
                                <Field.Label>Password</Field.Label>
                                <Input {...register("password")} type="password" flex="1" />
                            </Field.Root>

                            <Button
                                loading={loading}
                                loadingText="Loading"
                                spinnerPlacement="end"
                                colorScheme="blue"
                                w="full"
                                type="submit"
                                _active={{ transform: "scale(.97)" }}>
                                Log In
                            </Button>

                        </Stack>
                    </form>
                </Box>
            </Flex>
        </Flex >
    );
};

export default LoginPage;
