import axios from "axios";

const BASE = "https://jsonplaceholder.typicode.com";

const api = axios.create({
  baseURL: BASE,
  timeout: 10000,
  headers: { "Content-Type": "application/json" },
});

export async function fetchPosts() {
  const resp = await api.get("/posts");
  return resp.data;
}

export async function fetchPost(id) {
  const resp = await api.get(`/posts/${id}`);
  return resp.data;
}

export async function createPost(post) {
  const resp = await api.post("/posts", post);
  return resp.data;
}

export async function updatePost(id, post) {
  const resp = await api.put(`/posts/${id}`, post);
  return resp.data;
}

export async function deletePost(id) {
  const resp = await api.delete(`/posts/${id}`);
  return resp.data;
}

export default api;
