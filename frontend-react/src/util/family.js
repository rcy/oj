export function setFamilyMembershipId(familyMembershipId) {
  window.sessionStorage.setItem('familyMembershipId', familyMembershipId)
}

export function clearFamilyMembershipId() {
  window.sessionStorage.removeItem('familyMembershipId')
}

export function familyMembershipId() {
  return window.sessionStorage.getItem('familyMembershipId')
}
