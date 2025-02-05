import axios from 'axios';
import { TTokenPair } from '../models/TokenPair';
import { get } from '../api/apiAbstractions';

// Вынести в env, пока в падлу
const BASE_URL = 'http://85.192.61.121:8085/api';

const axiosInstance = axios.create({
  baseURL: BASE_URL,
  withCredentials: true,
  headers: {
    'Content-Type': 'application/json'
  }
});

export const refreshAccessTokenFn = async () => {
  const response = await get<TTokenPair>('/refresh');
  return response.data;
};

axiosInstance.interceptors.response.use(
  (response) => {
    return response;
  },
  async (error) => {
    const originalRequest = error.config;
    const errMessage = error.response.data.message as string;
    if (errMessage.includes('not logged in') && !originalRequest._retry) {
      originalRequest._retry = true;
      await refreshAccessTokenFn();
      return axiosInstance(originalRequest);
    }
    return Promise.reject(error);
  }
);

export default axiosInstance;
