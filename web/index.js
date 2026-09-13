const btnElement = document.querySelector("button");
const langHeading = document.getElementById("lang-heading");
const langRadios = document.querySelector(".language-radio-btns");
const outputEl = document.getElementById("output");

let isTranslated = false;

btnElement.addEventListener("click", async () => {
  if (isTranslated) {
    langHeading.textContent = "Select language 👇";
    langRadios.hidden = false;
    outputEl.hidden = true;
    btnElement.textContent = "Translate";
    isTranslated = false;
    return;
  }

  btnElement.disabled = true;
  btnElement.textContent = "...";
  const text = document.getElementById("text").value;
  const lang = document.querySelector('input[name="lang"]:checked').value;

  try {
    const res = await fetch("/api/translate", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({ text, lang }),
    });

    if (!res.ok) {
      throw new Error("Translation failed");
    }

    const data = await res.json();

    outputEl.textContent = data.output;
    langHeading.textContent = "Your Translation 👇";
    langRadios.hidden = true;
    outputEl.hidden = false;
    btnElement.textContent = "Start Over";
    isTranslated = true;
    btnElement.disabled = false;
  } catch (error) {
    console.error(error);
    alert("Something went wrong. Please try again.");
    btnElement.textContent = "Translate";
  } finally {
    btnElement.disabled = false;
  }
});
