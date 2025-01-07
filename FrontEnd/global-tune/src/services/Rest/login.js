export const loginAuth = async (body) => {
  console.log(JSON.stringify(body));
  const response = await fetch("http://localhost:8082/login", {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify(body),
  });
  if (!response.ok) {
    throw new Error(`Error on auth: ${response.statusText}`);
  }
  return response.json();
};
