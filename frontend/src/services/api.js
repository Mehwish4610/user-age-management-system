import axios from "axios";

const API = axios.create({
  baseURL: import.meta.env.VITE_API_URL || "http://localhost:8081",
});

export const getUsers = (page = 1, limit = 5) => {
  return API.get(`/users?page=${page}&limit=${limit}`);
};

export const getUserById = (id) => {
  return API.get(`/users/${id}`);
};

export const createUser = (userData) => {
  return API.post("/users", userData);
};

export const updateUser = (id, userData) => {
  return API.put(`/users/${id}`, userData);
};

export const deleteUser = (id) => {
  return API.delete(`/users/${id}`);
};