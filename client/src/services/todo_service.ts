import { TodoDto } from "@/dtos/TodoDto";
import { TodoModel } from "@/models/TodoModel";

const API_URL = import.meta.env.VITE_API_URL + '/todos';


export const getTodos = async () => {
    try {
        const response = await fetch(API_URL);
        if (!response.ok) {
            throw new Error("Failed to fetch users");
        }
        return await response.json();
    } catch (error) {
        console.error("Error fetching users:", error);
        return [];
    }
};


export const createTodo = async (todo: TodoDto) => {
    try {
        const response = await fetch(API_URL, {
            method: "POST",
            headers: {
                "Content-Type": "application/json"
            },
            body: JSON.stringify(todo)
        });
        if (!response.ok) {
            throw new Error("Failed to fetch users");
        }
        return await response.json();
    } catch (error) {
        console.error("Error fetching users:", error);
        return [];
    }
};