import { TodoDto } from "@/dtos/TodoDto";

const API_URL = import.meta.env.VITE_API_URL + '/todos';


export const getTodos = async () => {
    try {
        const response = await fetch(API_URL);
        if (!response.ok) {
            throw new Error("Failed to fetch todos");
        }
        return await response.json();
    } catch (error) {
        console.error("Error fetching todos:", error);
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
            throw new Error("Failed to create todo");
        }
        return await response.json();
    } catch (error) {
        console.error("Error creating todo:", error);
        return [];
    }
};


export const updateTodoCompleted = async (todoId: string) => {
    try {
        const response = await fetch(API_URL + '/' + todoId, {
            method: 'PUT',
        });
        if (!response.ok) {
            throw new Error("Failed to update todo");
        }
        return await response.json();
    } catch (error) {
        console.error("Error updating todo:", error);
        return [];
    }
};