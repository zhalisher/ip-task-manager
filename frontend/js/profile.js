if (!localStorage.getItem("token")) {
  window.location.href = "login.html";
}

document.getElementById("logout-btn").addEventListener("click", (e) => {
  e.preventDefault();
  localStorage.removeItem("token");
  window.location.href = "login.html";
});

async function loadProfile() {
  try {
    const user = await request("GET", "/profile");
    document.getElementById("profile-name").textContent = user.Name;
    document.getElementById("profile-email").textContent = user.Email;
    document.getElementById("profile-created").textContent = new Date(
      user.CreatedAt,
    ).toLocaleDateString();
  } catch (err) {
    alert("Failed to load profile: " + err.message);
  }
}

loadProfile();
