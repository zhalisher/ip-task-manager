if (!localStorage.getItem("token")) {
  window.location.href = "login.html";
}

document.getElementById("logout-btn").addEventListener("click", (e) => {
  e.preventDefault();
  localStorage.removeItem("token");
  window.location.href = "login.html";
});

let categoriesCache = [];
let tasksCache = [];

function fillCategorySelect() {
  const select = document.getElementById("task-category");
  select.innerHTML = "";
  categoriesCache.forEach((cat) => {
    const opt = document.createElement("option");
    opt.value = cat.ID;
    opt.textContent = cat.Name;
    select.appendChild(opt);
  });
}

async function loadDashboard() {
  const search = document.getElementById("search").value;
  const status = document.getElementById("status-filter").value;
  const priority = document.getElementById("priority-filter").value;

  localStorage.setItem(
    "lastFilter",
    JSON.stringify({ search, status, priority }),
  );

  const filters = { page: 1, limit: 50 };
  if (search) filters.search = search;
  if (status) filters.status = status;
  if (priority) filters.priority = priority;

  try {
    const [categories, tasks] = await Promise.all([
      getCategories(),
      getTasks(filters),
    ]);
    categoriesCache = categories || [];
    tasksCache = tasks || [];
    fillCategorySelect();

    const container = document.getElementById("categories-container");
    container.innerHTML = "";

    categoriesCache.forEach((cat) => {
      const catTasks = tasksCache.filter((t) => t.CategoryID === cat.ID);
      const div = document.createElement("div");
      div.className = "categories";
      div.innerHTML = `
        <div class="cat-header">
          <h2 style="border-left:4px solid ${cat.Color}; padding-left:0.5rem">${cat.Name}</h2>
          <div class="actions">
            <button class="btn-small" onclick="editCategory('${cat.ID}', '${cat.Name}', '${cat.Color}')">Edit</button>
            <button class="btn-small danger" onclick="removeCategory('${cat.ID}')">Delete</button>
          </div>
        </div>
        ${catTasks.length === 0 ? '<p class="empty">No tasks</p>' : ""}
      `;
      catTasks.forEach((task) => {
        const taskDiv = document.createElement("div");
        taskDiv.className = "tasks";
        taskDiv.innerHTML = `
          <h2>${task.Title}</h2>
          <p>${task.Description || ""}</p>
          <span>${task.DueDate ? new Date(task.DueDate).toLocaleDateString() : ""}</span>
          <span class="badge">${task.Priority}</span>
          <div class="actions">
            <select onchange="changeStatus('${task.ID}', this.value)">
              <option value="todo" ${task.Status === "todo" ? "selected" : ""}>Todo</option>
              <option value="in_progress" ${task.Status === "in_progress" ? "selected" : ""}>In Progress</option>
              <option value="done" ${task.Status === "done" ? "selected" : ""}>Done</option>
            </select>
            <button class="btn-small danger" onclick="removeTask('${task.ID}')">Delete</button>
          </div>
        `;
        div.appendChild(taskDiv);
      });
      container.appendChild(div);
    });
  } catch (err) {
    alert("Failed to load: " + err.message);
  }
}

document
  .getElementById("add-category-btn")
  .addEventListener("click", async () => {
    const name = document.getElementById("cat-name").value.trim();
    const color = document.getElementById("cat-color").value;
    if (!name) return alert("Category name is required");
    try {
      await createCategory(name, color);
      document.getElementById("cat-name").value = "";
      loadDashboard();
    } catch (err) {
      alert(err.message);
    }
  });

document.getElementById("add-task-btn").addEventListener("click", async () => {
  const title = document.getElementById("task-title").value.trim();
  const description = document.getElementById("task-desc").value.trim();
  const category_id = document.getElementById("task-category").value;
  const status = document.getElementById("task-status").value;
  const priority = document.getElementById("task-priority").value;
  const due = document.getElementById("task-due").value;

  if (!title) return alert("Title is required");
  if (!category_id) return alert("Create a category first");

  const task = { category_id, title, description, status, priority };
  if (due) task.due_date = new Date(due).toISOString();

  try {
    await createTask(task);
    document.getElementById("task-title").value = "";
    document.getElementById("task-desc").value = "";
    document.getElementById("task-due").value = "";
    loadDashboard();
  } catch (err) {
    alert(err.message);
  }
});

async function removeTask(id) {
  if (!confirm("Delete this task?")) return;
  try {
    await deleteTask(id);
    loadDashboard();
  } catch (e) {
    alert(e.message);
  }
}

async function changeStatus(id, status) {
  const t = tasksCache.find((x) => x.ID === id);
  if (!t) return;
  try {
    await updateTask(id, {
      title: t.Title,
      description: t.Description,
      status: status,
      priority: t.Priority,
      due_date: t.DueDate,
    });
    loadDashboard();
  } catch (e) {
    alert(e.message);
  }
}

async function removeCategory(id) {
  if (!confirm("Delete this category?")) return;
  try {
    await deleteCategory(id);
    loadDashboard();
  } catch (e) {
    alert(e.message);
  }
}

async function editCategory(id, currentName, currentColor) {
  const name = prompt("New category name:", currentName);
  if (name === null) return;
  try {
    await updateCategory(id, name.trim() || currentName, currentColor);
    loadDashboard();
  } catch (e) {
    alert(e.message);
  }
}

const lastFilter = JSON.parse(localStorage.getItem("lastFilter") || "{}");
if (lastFilter.search)
  document.getElementById("search").value = lastFilter.search;
if (lastFilter.status)
  document.getElementById("status-filter").value = lastFilter.status;
if (lastFilter.priority)
  document.getElementById("priority-filter").value = lastFilter.priority;

document.getElementById("search").addEventListener("input", loadDashboard);
document
  .getElementById("status-filter")
  .addEventListener("change", loadDashboard);
document
  .getElementById("priority-filter")
  .addEventListener("change", loadDashboard);

loadDashboard();
