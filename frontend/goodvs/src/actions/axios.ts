

import axios from 'axios';

export const axiosInstance = axios.create({
  baseURL: 'http://localhost:8000',
  headers: {
    'Content-Type': 'application/json',
  },
});

export const Ping = async () => {
    return axiosInstance.get('/ping');
}

export const SignIn = async (email: string, password: string) => {
    return axiosInstance.post('/login', {
        email: email,
        password: password,
    });
}