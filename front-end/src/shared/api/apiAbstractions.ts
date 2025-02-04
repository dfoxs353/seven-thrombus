import axiosInstance from "../config/axios";
import { ApiResponse } from "../models/ApiResponse";

export const get = async <T>(url: string, params?: object): Promise<ApiResponse<T>> => {
  const response = await axiosInstance.get<T>(url, { params });
  return response;
};

export const post = async <T>(url: string, data?: object): Promise<ApiResponse<T>> => {
  const response = await axiosInstance.post<T>(url, data);
  return response;
};

export const put = async <T>(url: string, data?: object): Promise<ApiResponse<T>> => {
  const response = await axiosInstance.put<T>(url, data);
  return response;
};

export const deleteRequest = async <T>(url: string, params?: object): Promise<ApiResponse<T>> => {
  const response = await axiosInstance.delete<T>(url, { params });
  return response;
};