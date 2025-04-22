import { Box, Flex, Text, Container, IconButton } from "@chakra-ui/react";
import { ColorModeButton, useColorModeValue } from "@/components/ui/color-mode"
import { FaRegUser } from "react-icons/fa";
export default function Navbar({ onUserClick }: { onUserClick?: () => void }) {


    return (
        <Container maxW={"900px"}>
            <Box bg={useColorModeValue("gray.400", "gray.700")} px={4} my={4} borderRadius={"5"}>                <Flex h={16} alignItems={"center"} justifyContent={"space-between"}>
                {/* LEFT SIDE */}
                <Flex
                    justifyContent={"center"}
                    alignItems={"center"}
                    gap={3}
                    display={{ base: "none", sm: "flex" }}
                >

                    <IconButton
                        onClick={onUserClick}
                        variant="ghost"
                        aria-label="Toggle color mode"
                        size="sm"
                        css={{
                            _icon: {
                                width: "5",
                                height: "5",
                            },
                        }}
                    >
                        <FaRegUser />
                    </IconButton>


                </Flex>

                {/* RIGHT SIDE */}
                <Flex alignItems={"center"} gap={3}>
                    <Text fontSize={"lg"} fontWeight={500}>
                        Daily Tasks
                    </Text>
                    {/* Toggle Color Mode */}
                    <ColorModeButton></ColorModeButton>
                </Flex>
            </Flex>
            </Box>
        </Container>
    );
}