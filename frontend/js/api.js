const API_URL = "http://localhost:8080";

function getToken() {
  return localStorage.getItem("token");
}

async function request(method, path, body = null) {
  const headers = { "Content-Type": "application/json" };
  const token = getToken();
  if (token) headers["Authorization"] = `Bearer ${token}`;

  const res = await fetch(API_URL + path, {
    method,
    headers,
    body: body ? JSON.stringify(body) : null,
  });

  const text = await res.text();
  let data = text;
  try {
    data = JSON.parse(text);
  } catch {}
  if (!res.ok) {
    throw new Error((data && data.message) || text || "request failed");
  }
  return data;
}

// auth
async function register(email, password, name) {
  return request("POST", "/auth/register", { email, password, name });
}
async function login(email, password) {
  return request("POST", "/auth/login", { email, password });
}

// categories
async function getCategories() {
  return request("GET", "/categories");
}
async function createCategory(name, color) {
  return request("POST", "/categories", { name, color });
}
async function updateCategory(id, name, color) {
  return request("PUT", `/categories/${id}`, { name, color });
}
async function deleteCategory(id) {
  return request("DELETE", `/categories/${id}`);
}

// tasks
async function getTasks(filters = {}) {
  const params = new URLSearchParams(filters).toString();
  return request("GET", `/tasks?${params}`);
}
async function createTask(task) {
  return request("POST", "/tasks", task);
}
async function updateTask(id, task) {
  return request("PUT", `/tasks/${id}`, task);
}
async function deleteTask(id) {
  return request("DELETE", `/tasks/${id}`);
}
