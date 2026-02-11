const form = document.getElementById("loginForm");

form.addEventListener("submit", async (e) => {
    e.preventDefault();

    const username = document.getElementById("username").value;
    const password = document.getElementById("password").value;

    try {
        const res = await window.api.post("/login", { username, password });
        alert(res.data.message);
        form.reset();
        window.location.href = "dashboard.html";

    } catch (err) {
        alert(err.response?.data?.message || "Something went wrong");
        form.reset();
        console.error(err);
    }
});
