import axios from 'axios';
import { API_URL } from '../config'; // Import the configuration file

const api = axios.create({
  baseURL: API_URL, // Use the imported variable
});

// Ensure no default base URL is set
axios.defaults.baseURL = '';

api.interceptors.request.use(config => {
  const token = localStorage.getItem('token');
  if (token) {
    config.headers.Authorization = `Bearer ${token}`; 
  }
  return config;
}, error => {
  return Promise.reject(error);
});

export default api;