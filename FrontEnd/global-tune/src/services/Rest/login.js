export const loginAuth = async (body) => {
  console.log(JSON.stringify(body));
  const response = await fetch(`https://3.84.60.208/login`, {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify(body),
    credentials: "include",
  });
  if (!response.ok) {
    throw new Error(`Error on auth: ${response.statusText}`);
  }
  return response.json();
};

export const logout = async () => {
  const response = await fetch(`logout/logout`, {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    credentials: "include",
  });
  if (!response.ok) {
    throw new Error(`Error on logout: ${response.statusText}`);
  }
  return response.json();
};