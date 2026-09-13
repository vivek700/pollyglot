const btnElement = document.querySelector("button");

btnElement.addEventListener("click", async () => {
  const text = document.getElementById("text").value;
  const lang = document.querySelector('input[name="lang"]:checked').value;

  const res = await fetch("/api/translate", {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify({text, lang})
  });

  const data = await res.json()
  console.log(data)


});
