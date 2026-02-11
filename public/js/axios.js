const api = axios.create({
    baseURL: "/api",
});

api.interceptors.request.use((config) => {
    const token = localStorage.getItem("token");
    if (token) config.headers.Authorization = `Bearer ${token}`;
    return config;
}, (error) => Promise.reject(error));

api.interceptors.response.use(response => {
    const token = response?.data?.data?.token;
    if (token) localStorage.setItem("token", token);
    return response;
}, error => Promise.reject(error));

window.api = api;