import axios from "axios"

const backendHost = process.env.BACKEND_HOST || "localhost"
const backendPort = process.env.BACKEND_PORT || "8085"

export const api = axios.create({
    baseURL: `http://${backendHost}:${backendPort}`,
    headers: { 'Content-Type': 'application/json' }
});
