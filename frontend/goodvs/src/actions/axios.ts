import axios from 'axios';

export const axiosInstance = axios.create({
  baseURL: 'http://localhost:8080/api',
  headers: {
    'Content-Type': 'application/json',
  },
});

export const Ping = async () => {
    return axiosInstance.get('/ping');
}

export const PostSignIn = async (email: string, password: string) => {
    return axiosInstance.post('/user/login', {
        email: email,
        password: password,
    });
}

export const PostSignUp = async (email: string, password: string,name: string ) => {
    return axiosInstance.post('/user/register', {
        name: name,
        email: email,
        password: password,
    });
}

export const GetSearchResult = async (keyword: string) => {
    return axiosInstance.get(`/search?product=${keyword}`);
}