const modal = document.getElementById("modal");
const articlesEl = document.getElementById("articles");
const form = document.getElementById("articleForm");
const imageInput = document.getElementById("imageFile");
const preview = document.getElementById("preview");

window.openModal = function () {
    modal.classList.remove("hidden");
    modal.classList.add("flex");
};

window.closeModal = function () {
    modal.classList.add("hidden");
    modal.classList.remove("flex");
    form.reset();
    preview.classList.add("hidden");
};

imageInput.addEventListener("change", () => {
    const file = imageInput.files[0];
    if (!file) return;

    preview.src = URL.createObjectURL(file);
    preview.classList.remove("hidden");
});

form.addEventListener("submit", async (e) => {
    e.preventDefault();

    try {
        const imageForm = new FormData();
        imageForm.append("image", imageInput.files[0]);

        const uploadRes = await window.api.post("/uploads/image", imageForm, {
            headers: { "Content-Type": "multipart/form-data" },
        });

        const imageUrl = uploadRes.data.data.image_url;

        await window.api.post("/articles", {
            title: title.value,
            category: category.value,
            description: description.value,
            image: imageUrl,
        });

        closeModal();
        fetchMyArticles();

    } catch (err) {
        alert(err.response?.data?.message || "Failed to create article");
    }
});

async function fetchMyArticles() {
    articlesEl.innerHTML = "Loading...";

    try {
        const res = await window.api.get("/my-articles");

        articlesEl.innerHTML = "";
        res.data.data.forEach(a => {
            const card = document.createElement("div");
            card.className = "bg-white p-4 rounded-xl shadow";

            card.innerHTML = `
        <img src="${a.image}" class="w-full h-40 object-cover rounded mb-2">
        <h4 class="font-semibold">${a.title}</h4>
        <small class="text-gray-500">${a.category}</small>
        <p class="text-sm mt-2">${a.description}</p>
      `;

            articlesEl.appendChild(card);
        });

    } catch {
        articlesEl.innerHTML = "Failed to load articles";
    }
}

fetchMyArticles();