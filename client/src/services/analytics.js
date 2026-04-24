import api from "./Axios.js";


export const analytics = async (shortcode) => {
    const response = await api.get(`/analytics/${shortcode}`);
    return response.data;
}

export const ListAllUrls = async () => {
    const response = await api.get(`/urls`);
    return response.data;
}