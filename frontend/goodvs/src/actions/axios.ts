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

export const PostPlatformLogin = async (platform: string) => {
    return axiosInstance.post(`/platform/login`,{
        platform: platform,
    });
}

export const FollowProduct = async (userID: number, productID: string) => {
    return axiosInstance.post('/follow', {
        user_id: userID,
        product_id: productID,
    });
}

export const UnfollowProduct = async (userID: number, productID: string) => {
    return axiosInstance.post('/unfollow', {
        user_id: userID,
        product_id: productID,
    });
}

export const GetFollowingProducts = async (userID: number) => {
    return axiosInstance.get(`/follow?user_id=${userID}`);
}

export const GetProductPriceList = async (productID: string) => {
    return axiosInstance.get(`/pricelist?product_id=${productID}`);
}