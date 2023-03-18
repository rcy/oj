export function logout() {
  localStorage.clear();
  window.location.assign("/auth/logout");
}
