const form = document.getElementById("form");
const input = document.getElementById("url-input");
const result = document.getElementById("result");

form.addEventListener("submit", async (e) => {
    e.preventDefault();
    const url = input.value.trim();
    if (!url) return;

    const startTime = performance.now();

    const loader = document.createElement("div");
    loader.className = "loader";
    loader.textContent = "Capture en cours...";

    result.innerHTML = "";
    result.appendChild(loader);

    const container = document.createElement("div");
    container.className = "screenshot-container";

    const img = document.createElement("img");
    img.alt = "Screenshot de " + url;

    img.onload = () => {
        const endTime = performance.now();
        const loadTime = Math.round(endTime - startTime);

        result.innerHTML = "";

        const responseTimeElement = document.getElementById("response-time");
        responseTimeElement.textContent = `Chargé en ${loadTime} ms`;
        container.appendChild(img);
        result.appendChild(container);

        container.style.opacity = "0";
        container.style.transform = "translateY(20px)";

        setTimeout(() => {
            container.style.transition = "all 0.5s ease";
            container.style.opacity = "1";
            container.style.transform = "translateY(0)";
        }, 50);
    };

    img.onerror = () => {
        result.innerHTML = "";
        const errorDiv = document.createElement("div");
        errorDiv.className = "error-message";
        errorDiv.textContent = "Impossible de récupérer le screenshot de cette URL.";
        result.appendChild(errorDiv);
    };

    img.src = `/api/screenshot?url=${encodeURIComponent(url)}`;
});