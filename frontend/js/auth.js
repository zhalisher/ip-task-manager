const loginBtn = document.getElementById("login-btn");
if (loginBtn) {
  loginBtn.addEventListener("click", async () => {
    const email = document.getElementById("email").value.trim();
    const password = document.getElementById("password").value.trim();
    const errorMsg = document.getElementById("error-msg");

    if (!email || !password) {
      errorMsg.textContent = "Email and password are required";
      return;
    }
    const emailRegex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;
    if (!emailRegex.test(email)) {
      errorMsg.textContent = "Invalid email format";
      return;
    }

    try {
      const data = await login(email, password);
      localStorage.setItem("token", data.access_token);
      window.location.href = "index.html";
    } catch (err) {
      errorMsg.textContent = err.message;
    }
  });
}

const registerBtn = document.getElementById("register-btn");
if (registerBtn) {
  registerBtn.addEventListener("click", async () => {
    const name = document.getElementById("name").value.trim();
    const email = document.getElementById("email").value.trim();
    const password = document.getElementById("password").value.trim();
    const errorMsg = document.getElementById("error-msg");

    if (!name || !email || !password) {
      errorMsg.textContent = "All fields are required";
      return;
    }
    const emailRegex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;
    if (!emailRegex.test(email)) {
      errorMsg.textContent = "Invalid email format";
      return;
    }
    if (password.length < 8) {
      errorMsg.textContent = "Password must be at least 8 characters";
      return;
    }

    try {
      await register(email, password, name);
      window.location.href = "login.html";
    } catch (err) {
      errorMsg.textContent = err.message;
    }
  });
}
