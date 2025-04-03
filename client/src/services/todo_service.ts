const API_URL = import.meta.env.VITE_API_URL;


export const getTodos = async () => {
    const url = API_URL + '/todos';
    try {
        const response = await fetch(url);
        if (!response.ok) {
            throw new Error("Failed to fetch users");
        }
        return await response.json();
    } catch (error) {
        console.error("Error fetching users:", error);
        return [];
    }
};