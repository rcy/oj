export default function () {
  return (
    <div>
      <p>build out a kids view using bootstrap</p>
      <LogoutPersonButton />
    </div>
  )
}


function LogoutPersonButton() {
  function handleClick() {
    localStorage.clear()
    window.location.assign('/')
  }

  return (
    <button onClick={handleClick}>
      logout
    </button>
  )
}
