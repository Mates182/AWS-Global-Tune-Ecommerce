export const loginAuth = async (body) => {
  console.log(JSON.stringify(body));
  const response = await fetch(`http://18.208.109.124:81/login`, {
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
  const response = await fetch(`http://18.208.109.124:82/logout`, {
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