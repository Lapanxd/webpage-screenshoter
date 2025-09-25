const form = document.getElementById("form");
const input = document.getElementById("urlInput");
const result = document.getElementById("result");

form.addEventListener("submit", async (e) => {
    e.preventDefault();
    const url = input.value.trim();
    if (!url) return;

    const loader = document.createElement("span");
    loader.textContent = "Chargement...";
    result.innerHTML = "";
    result.appendChild(loader);

    const img = document.createElement("img");
    img.alt = "screenshot";

    img.onload = () => {
        result.innerHTML = "";
        result.appendChild(img);
    };

    img.onerror = () => {
        result.innerHTML = "";
        result.textContent = "Impossible de récupérer le screenshot.";
    };

    img.src = `/api/screenshot?url=${encodeURIComponent(url)}`;
});
