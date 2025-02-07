import axios from 'axios';
import { TTokenPair } from '../models/TokenPair';

// Вынести в env, пока в падлу
const BASE_URL = 'http://85.192.61.121:8085/api';

const axiosInstance = axios.create({
  baseURL: BASE_URL,
  headers: {
    'Content-Type': 'application/json'
  }
});

axiosInstance.interceptors.request.use((config) => {
  const token = localStorage.getItem("token");
  if (token) {
    config.headers.Authorization = `Bearer ${token}`;
  }
  return config;
});

axiosInstance.interceptors.response.use(
  (config) => {
    return config;
  },
  async (error) => {
    const originalRequest = error.config;
    if (
      error.response.status == 401 &&
      error.config &&
      !error.config._isRetry
    ) {
      originalRequest._isRetry = true;
      try {
        const response = await axios.post<TTokenPair>(
          `${BASE_URL}/refresh`,
          { withCredentials: true }
        );
        localStorage.setItem("token", response.data.accessToken);
        return axiosInstance.request(originalRequest);
      } catch (e) {
        console.log("Не авторизован");
      }
    }
    throw error;
  }
);

export default axiosInstance;
