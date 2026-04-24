import api from "./Axios.js";

export const createShortcode = async (url) => {
    const response = await api.post("/shorten", { original_url: url });
    return response.data;
}

