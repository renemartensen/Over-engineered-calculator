const API_BASE = window.location.hostname === "localhost"
  ? "http://localhost:8080"
  : "https://your-render-app.onrender.com";


document.getElementById("calculate").addEventListener("click", async () => {
  const expression = document.getElementById("expression").value;
  const resultDiv = document.getElementById("result");

  resultDiv.textContent = "Calculating...";

  try {
    const response = await fetch(`${API_BASE}/calculate`, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
        "Authorization": "Basic " + btoa("user@example.com:123456"),
      },
      body: JSON.stringify({ expression }),
    });

    if (!response.ok) {
      // Try to extract error message text from backend
      const errorText = await response.text();
      resultDiv.textContent = `Error ${response.status}: ${errorText || "Invalid input"}`;
      return;
    }

    const data = await response.json();
    resultDiv.textContent = `Result: ${data.result}`;

    loadHistory();

  } catch (err) {
    console.error("Fetch error:", err);
    resultDiv.textContent = "Network or server error.";
  }
});

async function loadHistory() {
  const response = await fetch(`${API_BASE}/history`, {
    headers: {
      "Authorization": "Basic " + btoa("user@example.com:123456"),
    },
  });
  const historyList = document.getElementById("history");
  historyList.innerHTML = "";

  if (!response.ok) {
    const text = await response.text();
    console.error("Failed to fetch history:", text);
    return;
  }

  const history = await response.json();
  history.forEach((item) => {
    const li = document.createElement("li");
    li.textContent = `${item.expression} = ${item.result}`;
    historyList.appendChild(li);
  });
}

loadHistory();
