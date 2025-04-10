import { UserDto } from "@/dtos/UserDto";


const API_URL = import.meta.env.VITE_API_URL + '/users';



export const login = async (user: UserDto) => {
    try {

        const response = await fetch(API_URL, {
            method: "POST",
            headers: {
                "Content-Type": "application/json"
            },
            body: JSON.stringify(user)
        });
        if (!response.ok) {
            const error = await response.json();
            return { error: error };
        }
        return await response.json();
    } catch (error: any) {
        console.error("Error logging:", error);
        return { error: error.message || "Something went wrong" };
    }
};


