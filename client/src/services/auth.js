import api from "./axios.js";
import { clearToken } from "../utils/Token.js";

export const login = async (email, password) => {
    const response = await api.post("/login", { email, password });
    return response.data;
}

export const register = async (name, email, password) => {
    const response = await api.post("/signup", { name, email, password });
    return response.data;
}

export const verify = async () => {
    const response = await api.get("/verify");
    return response.data;
}

export const Logout = () => {
    return clearToken(); 
}
