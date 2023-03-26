export function logout() {
  localStorage.removeItem("sessionKey");
  window.location.assign("/auth/logout");
}
