import axios from 'axios';

// Вынести в env, пока в падлу
const BASE_URL = 'https://api.example.com'

const axiosInstance = axios.create({
  baseURL: BASE_URL,
  headers: {
    'Content-Type': 'application/json',
  },
});

export default axiosInstance;