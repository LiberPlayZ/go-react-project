import { CreateTodoDto } from "@/types/dtos/todos/CreateTodoDto";

const API_URL = import.meta.env.VITE_API_URL + '/todos';


export const getTodos = async (userId: string) => {
    try {
        const response = await fetch(API_URL + '/' + userId);
        if (!response.ok) {
            const error = await response.json();
            return error;
        }
        return await response.json();
    } catch (error) {
        console.error("Error fetching todos:", error);
        return [];
    }
};


export const createTodo = async (todo: CreateTodoDto) => {
    try {

        const response = await fetch(API_URL, {
            method: "POST",
            headers: {
                "Content-Type": "application/json"
            },
            body: JSON.stringify(todo)
        });
        if (!response.ok) {
            const error = await response.json();
            return error;
        }
        return await response.json();
    } catch (error: any) {
        console.error("Error creating todo:", error);
        return { error: error.message || "Something went wrong" };
    }
};


export const updateTodoCompleted = async (todoId: string) => {
    try {
        const response = await fetch(API_URL + '/' + todoId, {
            method: 'PUT',
        });
        if (!response.ok) {
            const error = await response.json();
            return error;
        }
        return await response.json();
    } catch (error) {
        console.error("Error updating todo:", error);
        return { error: "Something went wrong" };
    }
};


export const deleteTodo = async (todoId: string) => {
    try {
        const response = await fetch(API_URL + '/' + todoId, {
            method: 'DELETE',
        });
        if (!response.ok) {
            const error = await response.json();
            return error;
        }
        return await response.json();
    } catch (error) {
        console.error("Error updating todo:", error);
        return null;
    }
};